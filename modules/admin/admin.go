package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris_web/models"
	"strconv"
	"time"
)

type AdminController struct {
	BaseController
}

func (this *AdminController) Get() {
	var roleList []models.AdminRole
	this.Db.Find(&roleList)

	tv := iris.Map{}
	tv["role_list"] = roleList
	this.render("admin/admin/index.html", tv)
}

func (this *AdminController) GetList() mvc.Result {
	limit, offset := this.initPage()
	username := this.Ctx.URLParam("username")
	phone := this.Ctx.URLParam("phone")
	role := this.Ctx.URLParam("role")
	var model models.Admin
	total, _ := this.Db.Where("type = ?", 0).Count(&model)
	var cond models.Admin
	if username != "" {
		cond.Username = username
	}
	if phone != "" {
		cond.Phone = phone
	}
	if role != "" {
		cond.Role, _ = strconv.Atoi(role)
	}
	cond.Type = 0
	var all []models.Admin
	this.Db.Where("type = ?", 0).Limit(limit, offset).Find(&all, &cond)
	var list []map[string]interface{}
	for _, one := range all {
		item := make(map[string]interface{})
		item["id"] = one.Id
		item["username"] = one.Username
		item["phone"] = one.Phone
		item["lock"] = one.Lock
		item["time"] = one.Time.Format(DatetimeFormat)
		var role models.AdminRole
		role.Id = one.Role
		exists, _ := this.Db.Get(&role)
		if exists {
			item["role"] = role.Name
		} else {
			item["role"] = ""
		}
		list = append(list, item)
	}
	return this.respList(list, total)
}

func (this *AdminController) GetAdd() {
	var roleList []models.AdminRole
	this.Db.Find(&roleList)
	tv := iris.Map{}
	tv["role_list"] = roleList
	this.render("admin/admin/add.html", tv)
}

func (this *AdminController) PostAdd() mvc.Result {
	role := this.Ctx.PostValue("role")
	username := this.Ctx.PostValue("username")
	phone := this.Ctx.PostValue("phone")
	var user models.Admin
	exists, _ := this.Db.Where("username = ?", username).Exist(&user)
	if exists {
		return this.error("存在相同名称")
	}
	var model models.Admin
	model.Username = username
	model.Phone = phone
	model.Role, _ = strconv.Atoi(role)
	model.SetPassword("123456")
	model.Time = time.Now()
	_, err := this.Db.Insert(&model)
	if err != nil {
		this.App.Logger().Errorf("添加失败: %v", err)
		return this.error("添加失败")
	}
	return this.ok("添加成功")
}

func (this *AdminController) GetEdit() {
	id := this.Ctx.URLParam("id")
	var model models.Admin
	exists, _ := this.Db.ID(id).Get(&model)
	if !exists {
		this.App.Logger().Errorf("数据不存在: %v", id)
		// todo throw exception
		return
	}
	var roleList []models.AdminRole
	this.Db.Find(&roleList)

	tv := iris.Map{}
	tv["model"] = model
	tv["role_list"] = roleList
	this.render("admin/admin/edit.html", tv)
}

func (this *AdminController) PostEdit() mvc.Result {
	id := this.Ctx.PostValue("id")
	role := this.Ctx.PostValue("role")
	username := this.Ctx.PostValue("username")
	phone := this.Ctx.PostValue("phone")
	var user models.Admin
	exists, _ := this.Db.Where("username = ? and id != ?", username, id).Exist(&user)
	if exists {
		return this.error("存在相同名称")
	}
	user = models.Admin{}
	exists, _ = this.Db.Id(id).Get(&user)
	if !exists {
		return this.error("数据不存在")
	}
	var model = models.Admin{}
	model.Username = username
	model.Phone = phone
	model.Role, _ = strconv.Atoi(role)
	_, err := this.Db.Id(id).Update(&model)
	if err != nil {
		this.App.Logger().Errorf("修改失败: %v", err)
		return this.error("修改失败")
	}
	return this.ok("修改成功")
}

func (this *AdminController) GetDelete() mvc.Result {
	id := this.Ctx.URLParam("id")
	var model models.Admin
	this.Db.ID(id).Delete(&model)
	return this.ok("删除成功")
}

func (this *AdminController) GetLock() mvc.Result {
	id := this.Ctx.URLParam("id")
	lock := this.Ctx.URLParam("lock")
	var model models.Admin
	exists, _ := this.Db.ID(id).Get(&model)
	if !exists {
		return this.error("数据不存在")
	}
	model = models.Admin{}
	model.Lock, _ = strconv.Atoi(lock)
	this.Db.ID(id).Cols("lock").Update(&model)
	return this.ok("成功")
}
