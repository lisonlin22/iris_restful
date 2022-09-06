package app

import (
	"fmt"

	"gdopo.com/database"
	"gdopo.com/models"
	"gdopo.com/route"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

var App = iris.New()

func init() {
	// app init
	fmt.Println("app init")
	app := App
	app.Use(logger.New())

	// auto migrate model
	fmt.Println("auto migrate model")
	db := database.DB
	db.AutoMigrate(models.AutoMigrateModel...)

	// register routers
	fmt.Println("register routers")
	route.RegisterRouter(app)
}
