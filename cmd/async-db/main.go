package main

import (
	"fmt"
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/dao"
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/internal/msg"
	"github.com/joeyscat/ok-short/pkg/codec"
	"github.com/joeyscat/ok-short/pkg/setting"
	stan "github.com/nats-io/go-nats-streaming"
	"io"
	"log"
	"sync"
)

var (
	d             *dao.Dao
	linkDetailSub stan.Subscription
	linkTraceSub  stan.Subscription
)

func main() {
	initEnv()
	defer global.DBEngine.Close()
	defer global.StanConn.Close()

	d = dao.New(global.DBEngine)

	wg := sync.WaitGroup{}
	wg.Add(1)
	var err error
	linkDetailSub, err = global.StanConn.Subscribe(global.NatsSetting.Subj.LinkDetail,
		linkDetailMsgHandler,
		stan.DurableName("i-will-remember"))
	if err != nil {
		wg.Done()
		fmt.Println(err)
	}

	linkTraceSub, err = global.StanConn.Subscribe(global.NatsSetting.Subj.LinkTrace,
		linkTraceMsgHandler,
		stan.DurableName("i-will-remember"))
	if err != nil {
		wg.Done()
		fmt.Println(err)
	}
	wg.Wait()
}

func linkDetailMsgHandler(m *stan.Msg) {
	var linkMsg msg.LinkMsg
	err := codec.Decoder(m.Data, &linkMsg)
	if err != nil {
		closeNatsSub(linkDetailSub)
	}

	fmt.Printf("subDbLink received: %+v\n", &linkMsg)
	_, err = d.CreateLink(linkMsg.Sc, linkMsg.URL, linkMsg.Exp)
	if err != nil {
		panic(err)
	}
}

func linkTraceMsgHandler(m *stan.Msg) {
	var linkTraceMsg msg.LinkTraceMsg
	err := codec.Decoder(m.Data, &linkTraceMsg)
	if err != nil {
		closeNatsSub(linkTraceSub)
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

func closeNatsSub(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatalf("close error: %s", err)
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
	global.StanConn, err = stan.Connect("test-cluster", "async-db-01", stan.NatsURL(global.NatsSetting.Url))
	if err != nil {
		panic(err)
	}

	return nil
}
