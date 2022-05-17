package main

import (
	"github.com/kataras/iris/v12"
	"iris_web/routes"
)

func main() {
	app := iris.New()
	//file, _ := os.Create("main.log")
	//app.Logger().SetOutput(file)
	tmpl := iris.Django("./views", ".html").Reload(true)
	app.RegisterView(tmpl)

	routes.InitRoute(app)
	app.Listen(":80")
}
