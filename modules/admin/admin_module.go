package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris_web/models"
	"time"
)

type AdminModuleController struct {
	BaseController
}

func (this *AdminModuleController) Get() {
	this.render("admin/admin_module/index.html", nil)
}

func (this *AdminModuleController) GetList() mvc.Result {
	limit, offset := this.initPage()
	var model models.AdminModule
	total, _ := this.Db.Count(&model)
	var all []models.AdminModule
	this.Db.Limit(limit, offset).Find(&all)
	var list []map[string]interface{}
	for _, one := range all {
		item := make(map[string]interface{})
		item["id"] = one.Id
		item["name"] = one.Name
		item["sort"] = one.Sort
		item["time"] = one.Time.Format(DatetimeFormat)
		list = append(list, item)
	}
	return this.respList(list, total)
}

func (this *AdminModuleController) GetAdd() {
	this.render("admin/admin_module/add.html", nil)
}

func (this *AdminModuleController) PostAdd() mvc.Result {
	name := this.post("name")
	var model models.AdminModule
	model.Name = name
	exists, _ := this.Db.Exist(&model)
	if exists {
		return this.error("存在相同名称")
	}
	model = models.AdminModule{}
	model.Name = name
	model.Time = time.Now()
	_, err := this.Db.Insert(&model)
	if err != nil {
		this.App.Logger().Errorf("添加失败: %v", err)
		return this.error("添加失败")
	}
	return this.ok("添加成功")
}

func (this *AdminModuleController) GetEdit() {
	id := this.get("id")
	var model models.AdminModule
	exists, err := this.Db.Id(id).Get(&model)
	if !exists {
		this.App.Logger().Errorf("数据不存在: %v", err)
		// todo throw exception
	}
	tv := iris.Map{}
	tv["model"] = model
	this.render("admin/admin_module/edit.html", tv)
}

func (this *AdminModuleController) PostEdit() mvc.Result {
	id := this.post("id")
	name := this.post("name")
	var model models.AdminModule
	exists, err := this.Db.Id(id).Get(&model)
	if !exists {
		this.App.Logger().Errorf("数据不存在: %v", err)
		return this.error("数据不存在")
	}
	model = models.AdminModule{}
	model.Name = name
	_, err = this.Db.Id(id).Update(&model)
	if err != nil {
		this.App.Logger().Errorf("修改失败: %v", err)
		return this.error("修改失败")
	}
	return this.ok("修改成功")
}

func (this *AdminModuleController) GetDelete() mvc.Result {
	id := this.get("id")
	var model models.AdminModule
	this.Db.Id(id).Delete(&model)
	return this.ok("删除成功")
}
