package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	"github.com/lostvip-com/lv_framework/web/dto"
	"main/internal/system/dao"
	"main/internal/system/model"
	"main/internal/system/service"
	"main/internal/system/vo"
	"net/http"
)

type MenuController struct {
}

// 列表页
func (w *MenuController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "system/menu/list").WriteTpl()
}

// 列表分页数据
func (w *MenuController) ListAjax(c *gin.Context) {
	var req = new(vo.SelectMenuPageReq)
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
		return
	}
	rows := make([]model.SysMenu, 0)
	var dao dao.MenuDao
	result, err := dao.SelectListAll(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	c.JSON(http.StatusOK, rows)
}

// 新增页面
func (w *MenuController) Add(c *gin.Context) {
	pid := lv_conv.Int64(c.Query("pid"))
	var pmenu model.SysMenu
	pmenu.MenuId = 0
	pmenu.MenuName = "主目录"
	var dao dao.MenuDao
	tmp, err := dao.SelectRecordById(pid)
	if err == nil && tmp != nil && tmp.MenuId > 0 {
		pmenu.MenuId = tmp.MenuId
		pmenu.MenuName = tmp.MenuName
	}
	lv_web.BuildTpl(c, "system/menu/add").WriteTpl(gin.H{"menu": pmenu})
}

// 新增页面保存
func (w *MenuController) AddSave(c *gin.Context) {
	var req = new(vo.AddMenuReq)
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
		return
	}
	var service service.MenuService
	id, err := service.AddSave(req, c)

	if err != nil || id <= 0 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetBtype(dto.Buniss_Add).SetData(id).Log("菜单管理", req).WriteJsonExit()
}

// 修改页面
func (w *MenuController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	var dao dao.MenuDao
	menu, err := dao.SelectRecordById(id)
	if err != nil || menu == nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "菜单不存在",
		})
		return
	}
	parent, err := dao.SelectRecordById(menu.ParentId)
	if err == nil {
		menu.ParentName = parent.MenuName
	}
	lv_web.BuildTpl(c, "system/menu/edit").WriteTpl(gin.H{
		"menu": menu,
	})
}

// 修改页面保存
func (w *MenuController) EditSave(c *gin.Context) {
	var req = new(vo.EditMenuReq)
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
		return
	}
	var service service.MenuService
	rs, err := service.EditSave(req, c)

	if err != nil || rs <= 0 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).Log("菜单管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetBtype(dto.Buniss_Edit).SetData(rs).Log("菜单管理", req).WriteJsonExit()
}

// 删除数据
func (w *MenuController) Remove(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	var service service.MenuService
	rs := service.DeleteRecordById(id)
	if rs {
		lv_web.SucessResp(c).SetBtype(dto.Buniss_Del).Log("菜单管理", gin.H{"id": id}).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).Log("菜单管理", gin.H{"id": id}).WriteJsonExit()
	}
}

// 选择菜单树
func (w *MenuController) SelectMenuTree(c *gin.Context) {
	menuId := lv_conv.Int64(c.Query("menuId"))
	var dao dao.MenuDao
	menu, err := dao.SelectRecordById(menuId)
	if err != nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
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
	var service service.MenuService
	ztrees, err := service.MenuTreeData(user.UserId)
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
	var service service.MenuService
	result, err := service.RoleMenuTreeData(roleId, user.UserId)

	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("菜单管理", gin.H{"roleId": roleId}).WriteJsonExit()
		return
	}

	c.JSON(http.StatusOK, result)
}

// 检查菜单名是否已经存在不包括自身
func (w *MenuController) CheckMenuNameUnique(c *gin.Context) {
	var req *vo.CheckMenuNameReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}
	var service service.MenuService
	result := service.CheckMenuNameUnique(req.MenuName, req.MenuId, req.ParentId)

	c.Writer.WriteString(result)
}

// 检查菜单名是否已经存在
func (w *MenuController) CheckMenuNameUniqueAll(c *gin.Context) {
	var req *vo.CheckMenuNameALLReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}
	var dao dao.MenuDao
	exists := dao.CheckMenuNameExists(req.MenuName, req.ParentId)
	if exists {
		c.Writer.WriteString(req.MenuName)
	} else {
		c.Writer.WriteString("")
	}

}
