package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	"net/http"
	"robvi/app/common/model_cmn"
	"robvi/app/system/model/system/menu"
	"robvi/app/system/service"
	menuService "robvi/app/system/service/system/menu"
)

type MenuController struct {
}

// 列表页
func (w *MenuController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "system/menu/list").WriteTpl()
}

// 列表分页数据
func (w *MenuController) ListAjax(c *gin.Context) {
	var req = new(menu.SelectPageReq)
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
		return
	}
	rows := make([]menu.SysMenu, 0)
	result, err := menuService.SelectListAll(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	c.JSON(http.StatusOK, rows)
}

// 新增页面
func (w *MenuController) Add(c *gin.Context) {
	pid := lv_conv.Int64(c.Query("pid"))
	var pmenu menu.EntityExtend
	pmenu.MenuId = 0
	pmenu.MenuName = "主目录"

	tmp, err := menuService.SelectRecordById(pid)
	if err == nil && tmp != nil && tmp.MenuId > 0 {
		pmenu.MenuId = tmp.MenuId
		pmenu.MenuName = tmp.MenuName
	}
	lv_web.BuildTpl(c, "system/menu/add").WriteTpl(gin.H{"menu": pmenu})
}

// 新增页面保存
func (w *MenuController) AddSave(c *gin.Context) {
	var req = new(menu.AddReq)
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Add).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
		return
	}

	if menuService.CheckMenuNameUniqueAll(req.MenuName, req.ParentId) == "1" {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Add).SetMsg("菜单名称已存在").Log("菜单管理", req).WriteJsonExit()
		return
	}

	id, err := menuService.AddSave(req, c)

	if err != nil || id <= 0 {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Add).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetBtype(model_cmn.Buniss_Add).SetData(id).Log("菜单管理", req).WriteJsonExit()
}

// 修改页面
func (w *MenuController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, model_cmn.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	menu, err := menuService.SelectRecordById(id)

	if err != nil || menu == nil {
		lv_web.BuildTpl(c, model_cmn.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "菜单不存在",
		})
		return
	}

	lv_web.BuildTpl(c, "system/menu/edit").WriteTpl(gin.H{
		"menu": menu,
	})
}

// 修改页面保存
func (w *MenuController) EditSave(c *gin.Context) {
	var req = new(menu.EditReq)
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
		return
	}

	if menuService.CheckMenuNameUnique(req.MenuName, req.MenuId, req.ParentId) == "1" {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg("菜单名称已存在").Log("菜单管理", req).WriteJsonExit()
		return
	}

	rs, err := menuService.EditSave(req, c)

	if err != nil || rs <= 0 {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).Log("菜单管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetBtype(model_cmn.Buniss_Edit).SetData(rs).Log("菜单管理", req).WriteJsonExit()
}

// 删除数据
func (w *MenuController) Remove(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	rs := menuService.DeleteRecordById(id)

	if rs {
		lv_web.SucessResp(c).SetBtype(model_cmn.Buniss_Del).Log("菜单管理", gin.H{"id": id}).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Del).Log("菜单管理", gin.H{"id": id}).WriteJsonExit()
	}
}

// 选择菜单树
func (w *MenuController) SelectMenuTree(c *gin.Context) {
	menuId := lv_conv.Int64(c.Query("menuId"))
	menu, err := menuService.SelectRecordById(menuId)
	if err != nil {
		lv_web.BuildTpl(c, model_cmn.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "菜单不存在",
		})
		return
	}
	lv_web.BuildTpl(c, "system/menu/tree").WriteTpl(gin.H{
		"menu": menu,
	})
}

// 加载所有菜单列表树
func (w *MenuController) MenuTreeData(c *gin.Context) {
	var userService service.UserService
	user := userService.GetProfile(c)
	if user == nil {
		lv_web.ErrorResp(c).SetMsg("登录超时").Log("菜单管理", gin.H{"userId": user.UserId}).WriteJsonExit()
		return
	}
	ztrees, err := menuService.MenuTreeData(user.UserId)
	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("菜单管理", gin.H{"userId": user.UserId}).WriteJsonExit()
		return
	}
	c.JSON(http.StatusOK, ztrees)
}

// 选择图标
func (w *MenuController) Icon(c *gin.Context) {
	lv_web.BuildTpl(c, "system/menu/icon").WriteTpl()
}

// 加载角色菜单列表树
func (w *MenuController) RoleMenuTreeData(c *gin.Context) {
	var userService service.UserService
	roleId := lv_conv.Int64(c.Query("roleId"))
	user := userService.GetProfile(c)
	if user == nil || user.UserId <= 0 {
		lv_web.ErrorResp(c).SetMsg("登录超时").Log("菜单管理", gin.H{"roleId": roleId}).WriteJsonExit()
		return
	}

	result, err := menuService.RoleMenuTreeData(roleId, user.UserId)

	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("菜单管理", gin.H{"roleId": roleId}).WriteJsonExit()
		return
	}

	c.JSON(http.StatusOK, result)
}

// 检查菜单名是否已经存在不包括自身
func (w *MenuController) CheckMenuNameUnique(c *gin.Context) {
	var req *menu.CheckMenuNameReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := menuService.CheckMenuNameUnique(req.MenuName, req.MenuId, req.ParentId)

	c.Writer.WriteString(result)
}

// 检查菜单名是否已经存在
func (w *MenuController) CheckMenuNameUniqueAll(c *gin.Context) {
	var req *menu.CheckMenuNameALLReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := menuService.CheckMenuNameUniqueAll(req.MenuName, req.ParentId)

	c.Writer.WriteString(result)
}
