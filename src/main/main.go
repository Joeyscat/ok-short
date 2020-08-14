package main

import (
	"fmt"
	"github.com/joeyscat/ok-short/store"
)

func main() {
	app := App{}
	env := getEnv()
	app.Initialize(env)
	app.Run(fmt.Sprintf(":%d", env.port))

	defer store.MyDB.Close() // TODO 集中关闭资源
}
