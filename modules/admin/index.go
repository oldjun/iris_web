package admin

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.render("admin/index/index.html", nil)
}
