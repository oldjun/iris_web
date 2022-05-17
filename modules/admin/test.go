package admin

import (
	"github.com/kataras/iris/v12/mvc"
	"iris_web/models"
	"time"
)

type TestController struct {
	BaseController
}

func (this *TestController) Get() mvc.Result {
	isLogin, _ := this.Session.GetBoolean(ISLOGIN)
	if isLogin {
		return this.ok("用户已登录")
	} else {
		return this.error("用户未登录")
	}
}

func (this *TestController) GetAdmin() mvc.Result {
	var user models.Admin
	user.Username = "admin"
	user.Phone = "18976641111"
	user.Type = 1
	user.SetPassword("123456")
	user.Time = time.Now()
	id, err := this.Db.Insert(&user)
	if err != nil {
		this.App.Logger().Errorf("failed to insert: %v", err)
		return this.error("创建管理员失败")
	}
	resp := make(map[string]interface{})
	resp["id"] = id
	resp["username"] = "admin"
	return this.resp(resp)
}
