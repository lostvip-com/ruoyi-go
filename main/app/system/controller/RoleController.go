package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	"lostvip.com/web/dto"
	"main/app/system/service"
)

// 列表页
func (w *RoleController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "system/role/list").WriteTpl()
}

// 新增页面
func (w *RoleController) Add(c *gin.Context) {
	lv_web.BuildTpl(c, "system/role/add").WriteTpl()
}

// 修改页面
func (w *RoleController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	roleService := service.RoleService{}
	role, err := roleService.SelectRecordById(id)

	if err != nil || role == nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
		return
	}

	lv_web.BuildTpl(c, "system/role/edit").WriteTpl(gin.H{
		"role": role,
	})
}

// 分配用户添加
func (w *RoleController) SelectUser(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	roleService := service.RoleService{}
	role, err := roleService.SelectRecordById(id)

	if err != nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		lv_web.BuildTpl(c, "system/role/selectUser").WriteTpl(gin.H{
			"role": role,
		})
	}
}

// 数据权限
func (w *RoleController) AuthDataScope(c *gin.Context) {
	roleId := lv_conv.Int64(c.Query("id"))
	roleService := service.RoleService{}
	role, err := roleService.SelectRecordById(roleId)
	if err != nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		lv_web.BuildTpl(c, "system/role/dataScope").WriteTpl(gin.H{
			"role": role,
		})
	}
}

// 分配用户
func (w *RoleController) AuthUser(c *gin.Context) {
	roleId := lv_conv.Int64(c.Query("id"))
	roleService := service.RoleService{}
	role, err := roleService.SelectRecordById(roleId)
	if err != nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		lv_web.BuildTpl(c, "system/role/authUser").WriteTpl(gin.H{
			"role": role,
		})
	}
}
