package main

import "fmt"

func main() {
	app := App{}
	env := getEnv()
	app.Initialize(env)
	app.Run(fmt.Sprintf(":%d", env.port))
}
