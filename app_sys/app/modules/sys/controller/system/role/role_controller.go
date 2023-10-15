package role

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/gconv"
	response2 "lostvip.com/utils/response"
	"net/http"
	"robvi/app/modules/sys/model"
	role2 "robvi/app/modules/sys/model/system/role"
	userModel "robvi/app/modules/sys/model/system/user"
	"robvi/app/modules/sys/service"
	roleService "robvi/app/modules/sys/service/system/role"
)

// 列表页
func List(c *gin.Context) {
	response2.BuildTpl(c, "system/role/list").WriteTpl()
}

// 列表分页数据
func ListAjax(c *gin.Context) {
	var req *role2.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	rows := make([]role2.Entity, 0)
	result, page, err := roleService.SelectRecordPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	response2.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 新增页面
func Add(c *gin.Context) {
	response2.BuildTpl(c, "system/role/add").WriteTpl()
}

// 新增页面保存
func AddSave(c *gin.Context) {
	var req *role2.AddReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}

	if roleService.CheckRoleNameUniqueAll(req.RoleName) == "1" {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("角色名称已存在").Log("角色管理", req).WriteJsonExit()
		return
	}

	if roleService.CheckRoleKeyUniqueAll(req.RoleKey) == "1" {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("角色权限已存在").Log("角色管理", req).WriteJsonExit()
		return
	}

	rid, err := roleService.AddSave(req, c)

	if err != nil || rid <= 0 {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).Log("角色管理", req).WriteJsonExit()
		return
	}
	response2.SucessResp(c).SetData(rid).SetBtype(model.Buniss_Add).Log("角色管理", req).WriteJsonExit()
}

// 修改页面
func Edit(c *gin.Context) {
	id := gconv.Int64(c.Query("id"))
	if id <= 0 {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	role, err := roleService.SelectRecordById(id)

	if err != nil || role == nil {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
		return
	}

	response2.BuildTpl(c, "system/role/edit").WriteTpl(gin.H{
		"role": role,
	})
}

// 修改页面保存
func EditSave(c *gin.Context) {
	var req *role2.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}

	if roleService.CheckRoleNameUnique(req.RoleName, req.RoleId) == "1" {
		response2.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg("角色名称已存在").Log("角色管理", req).WriteJsonExit()
		return
	}

	if roleService.CheckRoleKeyUnique(req.RoleKey, req.RoleId) == "1" {
		response2.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg("角色权限已存在").Log("角色管理", req).WriteJsonExit()
		return
	}

	rs, err := roleService.EditSave(req, c)

	if err != nil || rs <= 0 {
		response2.ErrorResp(c).SetBtype(model.Buniss_Edit).Log("角色管理", req).WriteJsonExit()
		return
	}
	response2.SucessResp(c).SetBtype(model.Buniss_Edit).SetData(rs).Log("角色管理", req).WriteJsonExit()
}

// 分配用户添加
func SelectUser(c *gin.Context) {
	id := gconv.Int64(c.Query("id"))
	if id <= 0 {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	role, err := roleService.SelectRecordById(id)

	if err != nil {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		response2.BuildTpl(c, "system/role/selectUser").WriteTpl(gin.H{
			"role": role,
		})
	}
}

// 获取用户列表
func UnallocatedList(c *gin.Context) {
	roleId := gconv.Int64(c.PostForm("roleId"))
	loginName := c.PostForm("loginName")
	phonenumber := c.PostForm("phonenumber")
	var rows []userModel.SysUser
	var userService service.UserService
	userList, err := userService.SelectUnallocatedList(roleId, loginName, phonenumber)

	if err == nil && userList != nil {
		rows = userList
	}

	c.JSON(http.StatusOK, model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}

// 删除数据
func Remove(c *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}

	rs := roleService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response2.SucessResp(c).SetBtype(model.Buniss_Del).SetData(rs).Log("角色管理", req).WriteJsonExit()
	} else {
		response2.ErrorResp(c).SetBtype(model.Buniss_Del).Log("角色管理", req).WriteJsonExit()
	}
}

// 导出
func Export(c *gin.Context) {
	var req *role2.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	url, err := roleService.Export(req)

	if err != nil {
		response2.ErrorResp(c).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	response2.SucessResp(c).SetMsg(url).Log("角色管理", req).WriteJsonExit()
}

// 数据权限
func AuthDataScope(c *gin.Context) {
	roleId := gconv.Int64(c.Query("id"))
	role, err := roleService.SelectRecordById(roleId)
	if err != nil {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		response2.BuildTpl(c, "system/role/dataScope").WriteTpl(gin.H{
			"role": role,
		})
	}
}

// 数据权限保存
func AuthDataScopeSave(c *gin.Context) {
	var req *role2.DataScopeReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	if !roleService.CheckRoleAllowed(req.RoleId) {
		response2.ErrorResp(c).SetMsg("不允许操作超级管理员角色").Log("角色管理", req).WriteJsonExit()
		return
	}

	rs, err := roleService.AuthDataScope(req, c)
	if err != nil || rs <= 0 {
		response2.ErrorResp(c).SetMsg("保存数据失败").SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
	} else {
		response2.SucessResp(c).Log("角色管理", req).WriteJsonExit()
	}
}

// 分配用户
func AuthUser(c *gin.Context) {
	roleId := gconv.Int64(c.Query("id"))
	role, err := roleService.SelectRecordById(roleId)
	if err != nil {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		response2.BuildTpl(c, "system/role/authUser").WriteTpl(gin.H{
			"role": role,
		})
	}
}

// 查询已分配用户角色列表
func AllocatedList(c *gin.Context) {
	roleId := gconv.Int64(c.PostForm("roleId"))
	loginName := c.PostForm("loginName")
	phonenumber := c.PostForm("phonenumber")
	var rows []userModel.SysUser

	var userService service.UserService
	userList, err := userService.SelectAllocatedList(roleId, loginName, phonenumber)

	if err == nil && userList != nil {
		rows = userList
	}

	c.JSON(http.StatusOK, model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}

// 保存角色选择
func SelectAll(c *gin.Context) {
	roleId := gconv.Int64(c.PostForm("roleId"))
	userIds := c.PostForm("userIds")

	if roleId <= 0 {
		response2.ErrorResp(c).SetMsg("参数错误1").SetBtype(model.Buniss_Add).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
		return
	}
	if userIds == "" {
		response2.ErrorResp(c).SetMsg("参数错误2").SetBtype(model.Buniss_Add).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
		return
	}

	rs := roleService.InsertAuthUsers(roleId, userIds)
	if rs > 0 {
		response2.SucessResp(c).SetBtype(model.Buniss_Add).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	} else {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	}

}

// 取消用户角色授权
func CancelAll(c *gin.Context) {
	roleId := gconv.Int64(c.PostForm("roleId"))
	userIds := c.PostForm("userIds")
	if roleId > 0 && userIds != "" {
		roleService.DeleteUserRoleInfos(roleId, userIds)
		response2.SucessResp(c).SetBtype(model.Buniss_Del).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	} else {
		response2.ErrorResp(c).SetBtype(model.Buniss_Del).SetMsg("参数错误").Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	}
}

// 批量取消用户角色授权
func Cancel(c *gin.Context) {
	roleId := gconv.Int64(c.PostForm("roleId"))
	userId := gconv.Int64(c.PostForm("userId"))
	if roleId > 0 && userId > 0 {
		roleService.DeleteUserRoleInfo(userId, roleId)
		response2.SucessResp(c).SetBtype(model.Buniss_Del).Log("角色管理", gin.H{
			"roleId": roleId,
			"userId": userId,
		}).WriteJsonExit()
	} else {
		response2.ErrorResp(c).SetBtype(model.Buniss_Del).SetMsg("参数错误").Log("角色管理", gin.H{
			"roleId": roleId,
			"userId": userId,
		}).WriteJsonExit()
	}
}

// 检查角色是否已经存在不包括本角色
func CheckRoleNameUnique(c *gin.Context) {
	var req *role2.CheckRoleNameReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := roleService.CheckRoleNameUnique(req.RoleName, req.RoleId)

	c.Writer.WriteString(result)
}

// 检查角色是否已经存在
func CheckRoleNameUniqueAll(c *gin.Context) {
	var req *role2.CheckRoleNameALLReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := roleService.CheckRoleNameUniqueAll(req.RoleName)

	c.Writer.WriteString(result)
}

// 检查角色是否已经存在不包括本角色
func CheckRoleKeyUnique(c *gin.Context) {
	var req *role2.CheckRoleKeyReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := roleService.CheckRoleKeyUnique(req.RoleKey, req.RoleId)

	c.Writer.WriteString(result)
}

// 检查角色是否已经存在
func CheckRoleKeyUniqueAll(c *gin.Context) {
	var req *role2.CheckRoleKeyALLReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := roleService.CheckRoleKeyUniqueAll(req.RoleKey)

	c.Writer.WriteString(result)
}
