package admin

type LogoutController struct {
	BaseController
}

func (this *LogoutController) Get() {
	this.Session.Delete(ISLOGIN)
	this.Session.Delete(USERNAME)
	this.Session.Delete(PASSWORD)
	this.Ctx.Redirect("/login")
}
