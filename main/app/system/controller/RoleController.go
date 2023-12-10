package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	"net/http"
	"robvi/app/system/model"
	"robvi/app/system/model/system/role"
	userModel "robvi/app/system/model/system/user"
	"robvi/app/system/service"
	roleService "robvi/app/system/service/system/role"
)

type RoleController struct {
}

// 列表页
func (w *RoleController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "system/role/list").WriteTpl()
}

// 列表分页数据
func (w *RoleController) ListAjax(c *gin.Context) {
	var req *role.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	rows := make([]role.Entity, 0)
	result, page, err := roleService.SelectRecordPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	lv_web.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 新增页面
func (w *RoleController) Add(c *gin.Context) {
	lv_web.BuildTpl(c, "system/role/add").WriteTpl()
}

// 新增页面保存
func (w *RoleController) AddSave(c *gin.Context) {
	var req *role.AddReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}

	if roleService.CheckRoleNameUniqueAll(req.RoleName) == "1" {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("角色名称已存在").Log("角色管理", req).WriteJsonExit()
		return
	}

	if roleService.CheckRoleKeyUniqueAll(req.RoleKey) == "1" {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("角色权限已存在").Log("角色管理", req).WriteJsonExit()
		return
	}

	rid, err := roleService.AddSave(req, c)

	if err != nil || rid <= 0 {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).Log("角色管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(rid).SetBtype(model.Buniss_Add).Log("角色管理", req).WriteJsonExit()
}

// 修改页面
func (w *RoleController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	role, err := roleService.SelectRecordById(id)

	if err != nil || role == nil {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
		return
	}

	lv_web.BuildTpl(c, "system/role/edit").WriteTpl(gin.H{
		"role": role,
	})
}

// 修改页面保存
func (w *RoleController) EditSave(c *gin.Context) {
	var req *role.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}

	if roleService.CheckRoleNameUnique(req.RoleName, req.RoleId) == "1" {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg("角色名称已存在").Log("角色管理", req).WriteJsonExit()
		return
	}

	if roleService.CheckRoleKeyUnique(req.RoleKey, req.RoleId) == "1" {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg("角色权限已存在").Log("角色管理", req).WriteJsonExit()
		return
	}

	rs, err := roleService.EditSave(req, c)

	if err != nil || rs <= 0 {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).Log("角色管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetBtype(model.Buniss_Edit).SetData(rs).Log("角色管理", req).WriteJsonExit()
}

// 分配用户添加
func (w *RoleController) SelectUser(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	role, err := roleService.SelectRecordById(id)

	if err != nil {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		lv_web.BuildTpl(c, "system/role/selectUser").WriteTpl(gin.H{
			"role": role,
		})
	}
}

// 获取用户列表
func (w *RoleController) UnallocatedList(c *gin.Context) {
	roleId := lv_conv.Int64(c.PostForm("roleId"))
	loginName := c.PostForm("loginName")
	phonenumber := c.PostForm("phonenumber")
	var rows []userModel.SysUser
	var userService service.UserService
	userList, err := userService.SelectUnallocatedList(roleId, loginName, phonenumber)

	if err == nil && userList != nil {
		rows = userList
	}

	c.JSON(http.StatusOK, model.TableDataInfo{
		Code:  200,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}

// 删除数据
func (w *RoleController) Remove(c *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}

	rs := roleService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(model.Buniss_Del).SetData(rs).Log("角色管理", req).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Del).Log("角色管理", req).WriteJsonExit()
	}
}

// 导出
func (w *RoleController) Export(c *gin.Context) {
	var req *role.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	url, err := roleService.Export(req)

	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetMsg(url).Log("角色管理", req).WriteJsonExit()
}

// 数据权限
func (w *RoleController) AuthDataScope(c *gin.Context) {
	roleId := lv_conv.Int64(c.Query("id"))
	role, err := roleService.SelectRecordById(roleId)
	if err != nil {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		lv_web.BuildTpl(c, "system/role/dataScope").WriteTpl(gin.H{
			"role": role,
		})
	}
}

// 数据权限保存
func (w *RoleController) AuthDataScopeSave(c *gin.Context) {
	var req *role.DataScopeReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	if !roleService.CheckRoleAllowed(req.RoleId) {
		lv_web.ErrorResp(c).SetMsg("不允许操作超级管理员角色").Log("角色管理", req).WriteJsonExit()
		return
	}

	rs, err := roleService.AuthDataScope(req, c)
	if err != nil || rs <= 0 {
		lv_web.ErrorResp(c).SetMsg("保存数据失败").SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
	} else {
		lv_web.SucessResp(c).Log("角色管理", req).WriteJsonExit()
	}
}

// 分配用户
func (w *RoleController) AuthUser(c *gin.Context) {
	roleId := lv_conv.Int64(c.Query("id"))
	role, err := roleService.SelectRecordById(roleId)
	if err != nil {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		lv_web.BuildTpl(c, "system/role/authUser").WriteTpl(gin.H{
			"role": role,
		})
	}
}

// 查询已分配用户角色列表
func (w *RoleController) AllocatedList(c *gin.Context) {
	roleId := lv_conv.Int64(c.PostForm("roleId"))
	loginName := c.PostForm("loginName")
	phonenumber := c.PostForm("phonenumber")
	var rows []userModel.SysUser

	var userService service.UserService
	userList, err := userService.SelectAllocatedList(roleId, loginName, phonenumber)

	if err == nil && userList != nil {
		rows = userList
	}

	c.JSON(http.StatusOK, model.TableDataInfo{
		Code:  200,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}

// 保存角色选择
func (w *RoleController) SelectAll(c *gin.Context) {
	roleId := lv_conv.Int64(c.PostForm("roleId"))
	userIds := c.PostForm("userIds")

	if roleId <= 0 {
		lv_web.ErrorResp(c).SetMsg("参数错误1").SetBtype(model.Buniss_Add).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
		return
	}
	if userIds == "" {
		lv_web.ErrorResp(c).SetMsg("参数错误2").SetBtype(model.Buniss_Add).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
		return
	}

	rs := roleService.InsertAuthUsers(roleId, userIds)
	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(model.Buniss_Add).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	}

}

// 取消用户角色授权
func (w *RoleController) CancelAll(c *gin.Context) {
	roleId := lv_conv.Int64(c.PostForm("roleId"))
	userIds := c.PostForm("userIds")
	if roleId > 0 && userIds != "" {
		roleService.DeleteUserRoleInfos(roleId, userIds)
		lv_web.SucessResp(c).SetBtype(model.Buniss_Del).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Del).SetMsg("参数错误").Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	}
}

// 批量取消用户角色授权
func (w *RoleController) Cancel(c *gin.Context) {
	roleId := lv_conv.Int64(c.PostForm("roleId"))
	userId := lv_conv.Int64(c.PostForm("userId"))
	if roleId > 0 && userId > 0 {
		roleService.DeleteUserRoleInfo(userId, roleId)
		lv_web.SucessResp(c).SetBtype(model.Buniss_Del).Log("角色管理", gin.H{
			"roleId": roleId,
			"userId": userId,
		}).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Del).SetMsg("参数错误").Log("角色管理", gin.H{
			"roleId": roleId,
			"userId": userId,
		}).WriteJsonExit()
	}
}

// 检查角色是否已经存在不包括本角色
func (w *RoleController) CheckRoleNameUnique(c *gin.Context) {
	var req *role.CheckRoleNameReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := roleService.CheckRoleNameUnique(req.RoleName, req.RoleId)

	c.Writer.WriteString(result)
}

// 检查角色是否已经存在
func (w *RoleController) CheckRoleNameUniqueAll(c *gin.Context) {
	var req *role.CheckRoleNameALLReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := roleService.CheckRoleNameUniqueAll(req.RoleName)

	c.Writer.WriteString(result)
}

// 检查角色是否已经存在不包括本角色
func (w *RoleController) CheckRoleKeyUnique(c *gin.Context) {
	var req *role.CheckRoleKeyReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := roleService.CheckRoleKeyUnique(req.RoleKey, req.RoleId)

	c.Writer.WriteString(result)
}

// 检查角色是否已经存在
func (w *RoleController) CheckRoleKeyUniqueAll(c *gin.Context) {
	var req *role.CheckRoleKeyALLReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := roleService.CheckRoleKeyUniqueAll(req.RoleKey)

	c.Writer.WriteString(result)
}
