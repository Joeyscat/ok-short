package main

import (
	"github.com/joeyscat/ok-short/internal"
	"github.com/joeyscat/ok-short/internal/global"
)

// @title 短链接服务
// @version 1.0
// @description GoGo
// @termsOfService mm
func main() {
	global.InitEnv()

	defer global.Redis.Close()

	server := internal.NewServer()

	err := server.Start()

	if err != nil {
		panic(err)
	}
}
