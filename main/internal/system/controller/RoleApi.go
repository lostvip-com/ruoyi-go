package controller

import (
	"common/cm_vo"
	"github.com/gin-gonic/gin"
	db2 "github.com/lostvip-com/lv_framework/db"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	"github.com/lostvip-com/lv_framework/web/dto"
	"main/internal/system/dao"
	"main/internal/system/model"
	"main/internal/system/service"
	"net/http"
)

type RoleController struct {
}

// 列表分页数据
func (w *RoleController) ListAjax(c *gin.Context) {
	var req *cm_vo.RolePageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	rows := make([]model.SysRole, 0)
	roleService := service.RoleService{}
	result, total, err := roleService.SelectRecordPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	lv_web.BuildTable(c, total, rows).WriteJsonExit()
}

// 新增页面保存
func (w *RoleController) AddSave(c *gin.Context) {
	var req *cm_vo.AddRoleReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	roleDao := dao.SysRoleDao{}

	count, err := roleDao.FindCount("", req.RoleName)
	if count >= 1 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg("角色名称已存在").Log("角色管理", req).WriteJsonExit()
		return
	}
	count, err = roleDao.FindCount(req.RoleKey, "")
	if count >= 1 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg("角色权限已存在").Log("角色管理", req).WriteJsonExit()
		return
	}
	roleService := service.RoleService{}
	rid, err := roleService.AddSave(req, c)

	if err != nil || rid <= 0 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).Log("角色管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(rid).SetBtype(dto.Buniss_Add).Log("角色管理", req).WriteJsonExit()
}

// 修改页面保存
func (w *RoleController) EditSave(c *gin.Context) {
	var req *cm_vo.EditRoleReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	roleService := service.RoleService{}
	rs, err := roleService.EditSave(req, c)
	if err != nil || rs <= 0 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).Log("角色管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetBtype(dto.Buniss_Edit).SetData(rs).Log("角色管理", req).WriteJsonExit()
}

// 修改页面保存
func (w *RoleController) ChangeStatus(c *gin.Context) {
	roleId := c.PostForm("roleId")
	status := c.PostForm("status")
	sql := " update sys_role set status=? where role_id = ? "
	rows := db2.GetMasterGorm().Exec(sql, status, roleId).RowsAffected
	lv_web.SuccessData(c, rows)
}

// 获取用户列表
func (w *RoleController) UnallocatedList(c *gin.Context) {
	roleId := lv_conv.Int64(c.PostForm("roleId"))
	loginName := c.PostForm("loginName")
	phonenumber := c.PostForm("phonenumber")
	var rows []map[string]string
	var userService service.UserService
	userList, err := userService.SelectUnallocatedList(roleId, loginName, phonenumber)

	if err == nil && userList != nil {
		rows = *userList
	}

	c.JSON(http.StatusOK, dto.TableDataInfo{
		Code:  200,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}

// 删除数据
func (w *RoleController) Remove(c *gin.Context) {
	var req *dto.IdsReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	roleService := service.RoleService{}
	rs, _ := roleService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(dto.Buniss_Del).SetData(rs).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).WriteJsonExit()
	}
}

// 数据权限保存
func (w *RoleController) AuthDataScopeSave(c *gin.Context) {
	var req *cm_vo.DataScopeReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	roleService := service.RoleService{}
	if !roleService.CheckRoleAllowed(req.RoleId) {
		lv_web.ErrorResp(c).SetMsg("不允许操作超级管理员角色").WriteJsonExit()
		return
	}

	rs, err := roleService.AuthDataScope(req, c)
	if err != nil || rs <= 0 {
		lv_web.ErrorResp(c).SetMsg("保存数据失败").SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
	} else {
		lv_web.SucessResp(c).Log("角色管理", req).WriteJsonExit()
	}
}

// 查询已分配用户角色列表
func (w *RoleController) AllocatedList(c *gin.Context) {
	roleId := lv_conv.Int64(c.PostForm("roleId"))
	loginName := c.PostForm("loginName")
	phonenumber := c.PostForm("phonenumber")
	var rows []map[string]string

	var userService service.UserService
	userList, err := userService.SelectAllocatedList(roleId, loginName, phonenumber)

	if err == nil && userList != nil {
		rows = *userList
	}

	c.JSON(http.StatusOK, dto.TableDataInfo{
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
		lv_web.ErrorResp(c).SetMsg("参数错误1").SetBtype(dto.Buniss_Add).WriteJsonExit()
		return
	}
	if userIds == "" {
		lv_web.ErrorResp(c).SetMsg("参数错误2").SetBtype(dto.Buniss_Add).WriteJsonExit()
		return
	}
	roleService := service.RoleService{}
	err := roleService.InsertAuthUsers(roleId, userIds)
	if err == nil {
		lv_web.SucessResp(c).SetBtype(dto.Buniss_Add).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).WriteJsonExit()
	}
}

// 取消用户角色授权
func (w *RoleController) CancelAll(c *gin.Context) {
	roleId := lv_conv.Int64(c.PostForm("roleId"))
	userIds := c.PostForm("userIds")
	roleService := service.RoleService{}
	if roleId > 0 && userIds != "" {
		roleService.DeleteUserRoleInfos(roleId, userIds)
		lv_web.SucessResp(c).SetBtype(dto.Buniss_Del).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).SetMsg("参数错误").WriteJsonExit()
	}
}

// 批量取消用户角色授权
func (w *RoleController) Cancel(c *gin.Context) {
	roleId := lv_conv.Int64(c.PostForm("roleId"))
	userId := lv_conv.Int64(c.PostForm("userId"))
	roleService := service.RoleService{}
	if roleId > 0 && userId > 0 {
		roleService.DeleteUserRoleInfo(userId, roleId)
		lv_web.SucessResp(c).SetBtype(dto.Buniss_Del).Log("角色管理", gin.H{
			"roleId": roleId,
			"userId": userId,
		}).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).SetMsg("参数错误").WriteJsonExit()
	}
}
