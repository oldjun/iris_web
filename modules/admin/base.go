package admin

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"iris_web/models"
	"strconv"
)

const DatetimeFormat = "2006-01-02 15:04:05"

type BaseController struct {
	App     *iris.Application
	Ctx     iris.Context
	Session *sessions.Session
	Db      *xorm.Engine
}

func (this *BaseController) get(name string) string {
	return this.Ctx.URLParam(name)
}

func (this *BaseController) post(name string) string {
	return this.Ctx.PostValue(name)
}

func (this *BaseController) ok(data string) mvc.Result {
	return mvc.Response{
		Object: map[string]interface{}{
			"code": 0,
			"data": data,
		},
	}
}

func (this *BaseController) error(data string) mvc.Result {
	return mvc.Response{
		Object: map[string]interface{}{
			"code": 1,
			"data": data,
		},
	}
}

func (this *BaseController) resp(data map[string]interface{}) mvc.Result {
	return mvc.Response{
		Object: map[string]interface{}{
			"code": 0,
			"data": data,
		},
	}
}

func (this *BaseController) respList(list []map[string]interface{}, total int64) mvc.Result {
	return mvc.Response{
		Object: map[string]interface{}{
			"code":  0,
			"data":  "",
			"total": total,
			"list":  list,
		},
	}
}

func (this *BaseController) initPage() (int, int) {
	pageStr := this.get("page")
	page, _ := strconv.Atoi(pageStr)
	if page == 0 {
		page = 1
	}
	limitStr := this.get("limit")
	limit, _ := strconv.Atoi(limitStr)
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	return limit, offset
}

func (this *BaseController) render(filename string, tv iris.Map) {
	for key, value := range tv {
		this.Ctx.ViewData(key, value)
	}

	username := this.Session.Get(USERNAME)
	var user models.Admin
	this.Db.Where("username = ?", username).Get(&user)

	// 定义子菜单
	type MenuAction struct {
		Name   string
		Action string
		Active bool
	}

	// 定义主菜单
	type MenuModule struct {
		Name           string
		Active         bool
		MenuActionList []MenuAction
	}

	// 当前用户菜单生成
	menuModuleList := make([]MenuModule, 0)
	var midList []int
	var actionList []models.AdminAction
	this.Db.Where("menu = ?", 1).GroupBy("mid").Find(&actionList)
	for _, action := range actionList {
		midList = append(midList, action.Mid)
	}
	var moduleList []models.AdminModule
	this.Db.In("id", midList).OrderBy("sort asc").Find(&moduleList)
	for _, module := range moduleList {
		var actions []models.AdminAction
		this.Db.Where("mid = ? and menu = ?", module.Id, 1).OrderBy("sort asc").Find(&actions)
		if actions == nil {
			continue
		}
		var hasMenuModuleAuth = false
		var isMenuModuleActive = false
		menuActionList := make([]MenuAction, 0)
		for _, action := range actions {
			hasMenuActionAuth := false
			if user.IsSuper() {
				hasMenuActionAuth = true
			} else {
				var auth models.AdminAuth
				exists, _ := this.Db.Where("role = ? and action = ?", user.Role, action.Id).Exist(&auth)
				if exists {
					hasMenuActionAuth = true
				}
			}
			if hasMenuActionAuth {
				hasMenuModuleAuth = true
				menuAction := MenuAction{}
				menuAction.Name = action.Name
				menuAction.Action = action.Action
				if action.Action == this.Ctx.Path() {
					menuAction.Active = true
					isMenuModuleActive = true
				} else {
					menuAction.Active = false
				}
				menuActionList = append(menuActionList, menuAction)
			}
		}
		if hasMenuModuleAuth {
			menuModule := MenuModule{}
			menuModule.Name = module.Name
			menuModule.Active = isMenuModuleActive
			menuModule.MenuActionList = menuActionList
			menuModuleList = append(menuModuleList, menuModule)
		}
	}

	this.Ctx.ViewData("user", user)
	this.Ctx.ViewData("menu_module_list", menuModuleList)
	this.Ctx.View(filename)
}
