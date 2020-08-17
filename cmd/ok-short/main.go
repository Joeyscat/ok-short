package main

import (
	"fmt"
	. "github.com/joeyscat/ok-short/internel/app"
	. "github.com/joeyscat/ok-short/internel/pkg"
)

func main() {
	app := App{}
	context := GetContext()
	app.Initialize(context)
	app.Run(fmt.Sprintf(":%d", context.Port))

	defer MyDB.Close() // TODO 集中关闭资源
	defer ReCli.Close()
}
