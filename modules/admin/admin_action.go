package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris_web/models"
	"strconv"
	"time"
)

type AdminActionController struct {
	BaseController
}

func (this *AdminActionController) Get() {
	var moduleList []models.AdminModule
	this.Db.Find(&moduleList)
	tv := iris.Map{}
	tv["module_list"] = moduleList
	this.render("admin/admin_action/index.html", tv)
}

func (this *AdminActionController) GetList() mvc.Result {
	limit, offset := this.initPage()
	mid := this.get("mid")
	menu := this.get("menu")
	var cond models.AdminAction
	if mid != "" {
		cond.Mid, _ = strconv.Atoi(mid)
	}
	if menu != "" {
		cond.Menu, _ = strconv.Atoi(menu)
	}
	total, _ := this.Db.Count(&cond)
	var all []models.AdminAction
	this.Db.Limit(limit, offset).Find(&all, &cond)
	var list []iris.Map
	for _, one := range all {
		item := make(iris.Map)
		item["id"] = one.Id
		item["mid"] = one.Mid
		item["name"] = one.Name
		if one.Menu == 1 {
			item["menu"] = "是"
		} else {
			item["menu"] = "否"
		}
		item["action"] = one.Action
		item["time"] = one.Time.Format(DatetimeFormat)
		var module models.AdminModule
		exists, _ := this.Db.Id(one.Mid).Get(&module)
		if exists {
			item["module"] = module.Name
		} else {
			item["module"] = ""
		}
		list = append(list, item)
	}
	return this.respList(list, total)
}

func (this *AdminActionController) GetAdd() {
	var moduleList []models.AdminModule
	this.Db.Find(&moduleList)
	tv := iris.Map{}
	tv["module_list"] = moduleList
	this.render("admin/admin_action/add.html", tv)
}

func (this *AdminActionController) PostAdd() mvc.Result {
	mid := this.post("mid")
	name := this.post("name")
	action := this.post("action")
	menu := this.post("menu")
	if name == "" {
		return this.error("名称不能为空")
	}
	var model models.AdminAction
	model.Name = name
	exists, _ := this.Db.Where("name = ?", name).Exist(&model)
	if exists {
		return this.error("存在相同名称")
	}
	model = models.AdminAction{}
	model.Mid, _ = strconv.Atoi(mid)
	model.Name = name
	model.Action = action
	model.Menu, _ = strconv.Atoi(menu)
	model.Time = time.Now()
	_, err := this.Db.Insert(&model)
	if err != nil {
		this.App.Logger().Errorf("添加失败: %v", err)
		return this.error("添加失败")
	}
	return this.ok("添加成功")
}

func (this *AdminActionController) GetEdit() {
	id := this.get("id")
	var model models.AdminAction
	exists, _ := this.Db.Id(id).Get(&model)
	if !exists {
		this.App.Logger().Errorf("数据不存在: id=%v", id)
		// todo throw exception
	}
	var moduleList []models.AdminModule
	this.Db.Find(&moduleList)

	tv := iris.Map{}
	tv["model"] = model
	tv["module_list"] = moduleList
	this.render("admin/admin_action/edit.html", tv)
}

func (this *AdminActionController) PostEdit() mvc.Result {
	id := this.post("id")
	mid := this.post("mid")
	name := this.post("name")
	action := this.post("action")
	menu := this.post("menu")

	var model models.AdminAction
	exists, _ := this.Db.Id(id).Get(&model)
	if !exists {
		return this.error("数据不存在")
	}
	model = models.AdminAction{}
	model.Mid, _ = strconv.Atoi(mid)
	model.Name = name
	model.Action = action
	model.Menu, _ = strconv.Atoi(menu)
	_, err := this.Db.Id(id).Update(&model)
	if err != nil {
		this.App.Logger().Errorf("修改失败: %v", err)
		return this.error("修改失败")
	}
	return this.ok("修改成功")
}

func (this *AdminActionController) GetDelete() mvc.Result {
	id := this.get("id")
	var model models.AdminAction
	this.Db.Id(id).Delete(&model)
	return this.ok("删除成功")
}
