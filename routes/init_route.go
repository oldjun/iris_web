package routes

import "github.com/kataras/iris/v12"

func InitRoute(app *iris.Application) {
	InitAdminRoute(app)
}
