package admin

import (
	"github.com/kataras/iris/v12/mvc"
	"iris_web/models"
)

const USERNAME = "username"
const PASSWORD = "password"
const ISLOGIN = "is_login"

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	isLogin, _ := this.Session.GetBoolean(ISLOGIN)
	if isLogin {
		this.Ctx.Redirect("/")
	}
	this.Ctx.View("admin/login/index.html")
}

func (this *LoginController) Post() mvc.Result {
	username := this.Ctx.PostValue("username")
	password := this.Ctx.PostValue("password")
	if username == "" {
		return this.error("用户名不能为空")
	}
	if password == "" {
		return this.error("密码不能为空")
	}

	//var user models.Admin
	//exists, _ := this.Db.Where("username = ?", username).Get(&user)
	//if !exists {
	//	return this.error("用户不存在")
	//}
	var user models.Admin
	user.Username = username
	exists, err := this.Db.Get(&user)
	if err != nil {
		return this.error(err.Error())
	}
	if !exists {
		return this.error("用户不存在")
	}
	if user.IsLock() {
		return this.error("用户已被锁定")
	}
	if !user.CheckPassword(password) {
		return this.error("密码错误")
	}
	this.Session.Set(USERNAME, username)
	this.Session.Set(PASSWORD, password)
	this.Session.Set(ISLOGIN, true)
	return this.ok("登录成功")
}
