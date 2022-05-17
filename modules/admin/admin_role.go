package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris_web/models"
	"strconv"
	"time"
)

type AdminRoleController struct {
	BaseController
}

func (this *AdminRoleController) Get() {
	this.render("admin/admin_role/index.html", nil)
}

func (this *AdminRoleController) GetList() mvc.Result {
	limit, offset := this.initPage()
	var model = models.AdminRole{}
	total, _ := this.Db.Count(&model)

	var list []map[string]interface{}
	var all []models.AdminRole
	this.Db.Limit(limit, offset).Find(&all)
	for _, one := range all {
		item := make(map[string]interface{})
		item["id"] = one.Id
		item["name"] = one.Name
		item["time"] = one.Time.Format(DatetimeFormat)
		list = append(list, item)
	}
	return this.respList(list, total)
}

func (this *AdminRoleController) GetAdd() {
	this.render("admin/admin_role/add.html", nil)
}

func (this *AdminRoleController) PostAdd() mvc.Result {
	name := this.Ctx.PostValue("name")
	var role models.AdminRole
	exists, _ := this.Db.Where("name = ?", name).Get(&role)
	if exists {
		return this.error("存在相同名称")
	}
	role.Name = name
	role.Time = time.Now()
	_, err := this.Db.Insert(&role)
	if err != nil {
		return this.ok("添加失败")
	}
	return this.ok("添加成功")
}

func (this *AdminRoleController) GetEdit() {
	id := this.Ctx.URLParam("id")
	var model models.AdminRole
	exists, _ := this.Db.Id(id).Get(&model)
	if !exists {
		this.App.Logger().Errorf("数据不存在: id=%v", id)
		// todo throw exception
	}
	tv := iris.Map{}
	tv["model"] = model
	this.render("admin/admin_role/edit.html", tv)
}

func (this *AdminRoleController) PostEdit() mvc.Result {
	id := this.Ctx.PostValue("id")
	name := this.Ctx.PostValue("name")
	var model models.AdminRole
	exists, _ := this.Db.Where("name = ? and id != ?", name, id).Exist(&model)
	if exists {
		return this.error("存在相同名称")
	}
	exists, _ = this.Db.Id(id).Get(&model)
	if !exists {
		return this.error("数据不存在")
	}
	var role models.AdminRole
	role.Name = name
	this.Db.Id(id).Update(&role)
	return this.ok("修改成功")
}

func (this *AdminRoleController) GetDelete() mvc.Result {
	id := this.Ctx.URLParam("id")
	var model models.AdminRole
	this.Db.Id(id).Delete(&model)
	return this.ok("删除成功")
}

func (this *AdminRoleController) GetAuth() {
	role := this.get("id")
	var moduleList []map[string]interface{}
	var modules []models.AdminModule
	err := this.Db.OrderBy("sort asc").Find(&modules)
	if err != nil {
		this.App.Logger().Errorf("查询module列表失败: %v", err)
		// todo throw exception
		return
	}
	for _, module := range modules {
		checkAll := true
		moduleItem := make(map[string]interface{})
		moduleItem["id"] = module.Id
		moduleItem["name"] = module.Name
		moduleItem["sort"] = module.Sort
		var actionList []map[string]interface{}
		var actions []models.AdminAction
		err = this.Db.Where("mid = ?", module.Id).OrderBy("sort asc").Find(&actions)
		if err != nil {
			this.App.Logger().Errorf("查询action列表失败: %v", err)
			// todo throw exception
			return
		}
		for _, action := range actions {
			actionItem := make(map[string]interface{})
			actionItem["id"] = action.Id
			actionItem["mid"] = action.Mid
			actionItem["name"] = action.Name
			actionItem["sort"] = action.Sort
			var auth models.AdminAuth
			exists, _ := this.Db.Where("role = ? and action = ?", role, action.Id).Exist(&auth)
			if exists {
				actionItem["checked"] = true
			} else {
				actionItem["checked"] = false
				checkAll = checkAll && exists
			}
			actionList = append(actionList, actionItem)
		}
		moduleItem["actions"] = actionList
		moduleItem["checked_all"] = checkAll
		moduleList = append(moduleList, moduleItem)
	}
	tv := iris.Map{}
	tv["id"] = role
	tv["module_list"] = moduleList
	this.render("admin/admin_role/auth.html", tv)
}

func (this *AdminRoleController) PostAuth() mvc.Result {
	role, _ := this.Ctx.PostValueInt("id")
	actions := this.Ctx.PostValues("actions[]")
	var auth models.AdminAuth
	auth.Role = role
	this.Db.Delete(&auth)
	for _, action := range actions {
		var model models.AdminAuth
		model.Role = role
		model.Action, _ = strconv.Atoi(action)
		model.Time = time.Now()
		this.Db.Insert(&model)
	}
	return this.ok("设置成功")
}
