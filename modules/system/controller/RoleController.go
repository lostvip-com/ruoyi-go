package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"system/service"
)

// 列表页
func (w *RoleController) List(c *gin.Context) {
	util.BuildTpl(c, "system/role/list").WriteTpl()
}

// 新增页面
func (w *RoleController) Add(c *gin.Context) {
	util.BuildTpl(c, "system/role/add").WriteTpl()
}

// 修改页面
func (w *RoleController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		util.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	roleService := service.RoleService{}
	role, err := roleService.SelectRecordById(id)

	if err != nil || role == nil {
		util.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
		return
	}

	util.BuildTpl(c, "system/role/edit").WriteTpl(gin.H{
		"role": role,
	})
}

// 分配用户添加
func (w *RoleController) SelectUser(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		util.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	roleService := service.RoleService{}
	role, err := roleService.SelectRecordById(id)

	if err != nil {
		util.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		util.BuildTpl(c, "system/role/selectUser").WriteTpl(gin.H{
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
		util.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		util.BuildTpl(c, "system/role/dataScope").WriteTpl(gin.H{
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
		util.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		util.BuildTpl(c, "system/role/authUser").WriteTpl(gin.H{
			"role": role,
		})
	}
}
