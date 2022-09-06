package route

import (
	"gdopo.com/controllers"
	"gdopo.com/err"
	"github.com/kataras/iris/v12"
)

// 注册路由
func RegisterRouter(app *iris.Application) {
	// 处理跨域
	app.Use(Cors)
	// 处理错误
	app.OnErrorCode(iris.StatusNotFound, err.StatusNotFound)
	app.OnErrorCode(iris.StatusInternalServerError, err.StatusInternalServerError)
	// 全局API: /v1/api
	api := app.Party("/v1").Party("/api")

	// User : /v1/api/user
	api.PartyFunc("/user", controllers.User)
}

// 处理跨域
func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
