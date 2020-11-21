package main

import (
	"fmt"
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/dao"
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/internal/msg"
	"github.com/joeyscat/ok-short/pkg/codec"
	"github.com/joeyscat/ok-short/pkg/setting"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
)

var (
	d *dao.Dao
)

func main() {
	initEnv()
	defer global.DBEngine.Close()
	defer global.Nats.Close()

	d = dao.New(global.DBEngine)

	wg := sync.WaitGroup{}
	wg.Add(1)
	_, err := global.Nats.Subscribe(global.NatsSetting.Subj.LinkDetail, linkDetailMsgHandler)
	if err != nil {
		wg.Done()
		fmt.Println(err)
	}

	_, err = global.Nats.Subscribe(global.NatsSetting.Subj.LinkTrace, linkTraceMsgHandler)
	if err != nil {
		wg.Done()
		fmt.Println(err)
	}
	wg.Wait()
}

func linkDetailMsgHandler(nMsg *nats.Msg) {
	var linkMsg msg.LinkMsg
	err := codec.Decoder(nMsg.Data, &linkMsg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("subDbLink received: %+v\n", &linkMsg)
	_, err = d.CreateLink(linkMsg.Sc, linkMsg.URL, linkMsg.Exp)
	if err != nil {
		panic(err)
	}
}

func linkTraceMsgHandler(nMsg *nats.Msg) {
	var linkTraceMsg msg.LinkTraceMsg
	err := codec.Decoder(nMsg.Data, &linkTraceMsg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("linkTrace received: %+v\n", &linkTraceMsg)
	l := &model.LinkTrace{
		Sc:     linkTraceMsg.Sc,
		URL:    linkTraceMsg.URL,
		Ip:     linkTraceMsg.Ip,
		UA:     linkTraceMsg.UA,
		Cookie: linkTraceMsg.Cookie,
	}
	_, err = d.CreateLinkTrace(l)
	if err != nil {
		panic(err)
	}
}

func initEnv() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("setupDBEngine err: %v", err)
	}

	err = setupNats()
	if err != nil {
		log.Fatalf("setupNats err: %v", err)
	}
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Nats", &global.NatsSetting)
	if err != nil {
		return err
	}

	log.Printf("DatabaseSetting: %v", global.DatabaseSetting)
	log.Printf("Nats: %v", global.NatsSetting)

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupNats() error {
	var err error
	global.Nats, err = nats.Connect(global.NatsSetting.Url)
	if err != nil {
		panic(err)
	}

	return nil
}
