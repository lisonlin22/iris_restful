package main

import (
	"gdopo.com/app"
	"github.com/kataras/iris/v12"
)

func main() {
	app := app.App
	app.Run(iris.Addr(":80"))
}
