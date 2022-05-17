package routes

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"iris_web/config"
	"iris_web/modules/admin"
	"os"
	"time"
)

const ISLOGIN = "is_login"
const USERNAME = "username"

func InitAdminRoute(app *iris.Application) {
	sessionManager := sessions.New(sessions.Config{
		Cookie:  "IRIS_SESSION",
		Expires: 24 * time.Hour,
	})
	app.Use(sessionManager.Handler())
	adminParty := app.Party("admin.")
	adminParty.HandleDir("/js", "./js")
	adminParty.HandleDir("/css", "./css")
	adminParty.HandleDir("/images", "./images")

	database := config.Database
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s",
		database["user"], database["password"], database["host"], database["port"], database["database"], database["charset"])
	db, err := xorm.NewEngine(database["type"], dataSourceName)
	if err != nil {
		app.Logger().Fatalf("failed to connect mysql: %v", err)
		os.Exit(1)
	}
	adminApp := mvc.New(adminParty)
	adminApp.Register(app, sessionManager.Start, db)
	adminApp.Party("/test").Handle(new(admin.TestController))
	adminApp.Party("/login").Handle(new(admin.LoginController))
	adminApp.Party("/", checkLogin).Handle(new(admin.IndexController))
	adminApp.Party("/logout", checkLogin).Handle(new(admin.LogoutController))
	adminApp.Party("/admin", checkLogin, checkAuth).Handle(new(admin.AdminController))
	adminApp.Party("/admin-role", checkLogin, checkAuth).Handle(new(admin.AdminRoleController))
	adminApp.Party("/admin-module", checkLogin, checkAuth).Handle(new(admin.AdminModuleController))
	adminApp.Party("/admin-action", checkLogin, checkAuth).Handle(new(admin.AdminActionController))
}

// 登录验证
func checkLogin(ctx iris.Context) {
	session := sessions.Get(ctx)
	isLogin, _ := session.GetBoolean(ISLOGIN)
	if !isLogin {
		ctx.Redirect("/login")
		return
	}
	fmt.Println("check login")
	ctx.Next()
}

// 权限验证
func checkAuth(ctx iris.Context) {
	session := sessions.Get(ctx)
	username := session.Get(USERNAME)
	//user := models.Admin{}
	fmt.Println("check auth: " + username.(string))
	ctx.Next()
}
