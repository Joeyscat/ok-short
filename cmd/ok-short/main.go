package main

import (
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/routers/api/v2"
)

// @title 短链接服务
// @version 1.0
// @description GoGo
// @termsOfService mm
func main() {
	global.InitEnv()
	defer global.Redis.Close()

	e := v2.NewRouter()

	//e.Debug = true
	e.Logger.Fatal(e.Start(":" + global.ServerSetting.HttpPort))
}
