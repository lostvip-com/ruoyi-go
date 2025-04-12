package controller

import (
	userModel "common/common_vo"
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	service2 "system/service"
)

// 用户列表页
func (w *UserApi) List(c *gin.Context) {
	util.BuildTpl(c, "system/user/list").WriteTpl()
}

// 用户列表分页数据
func (w *UserApi) ListAjax(c *gin.Context) {
	var req *userModel.SelectUserPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util.ErrorResp(c).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
		return
	}
	var userService service2.UserService
	result, total, err := userService.SelectRecordList(req)
	lv_err.HasErrAndPanic(err)
	util.BuildTable(c, total, result).WriteJsonExit()
}

// 用户新增页面
func (w *UserApi) Add(c *gin.Context) {
	var roleService service2.RoleService
	rolesP, _ := roleService.SelectRecordAll(nil)
	var postService service2.SysPostService
	postP, _ := postService.SelectListAll(nil)
	util.BuildTpl(c, "system/user/add").WriteTpl(gin.H{
		"roles": rolesP,
		"posts": postP,
	})
}

// 保存新增用户数据
func (w *UserApi) AddSave(c *gin.Context) {
	var req *userModel.AddUserReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util.ErrorResp(c).SetBtype(lv_dto.Buniss_Add).SetMsg(err.Error()).Log("新增用户", req).WriteJsonExit()
		return
	}
	var userService service2.UserService
	//判断登录名是否已注册
	count, err := userService.CountCol("username", req.LoginName)
	if count > 0 {
		util.Err(c, "登录名已经存在")
		return
	}
	//判断手机号码是否已注册
	count, _ = userService.CountCol("phonenumber", req.Phonenumber)
	if count > 0 {
		util.Fail(c, "手机号码已经存在")
		return
	}
	uid, err := userService.AddSave(req, c)
	lv_err.HasErrAndPanic(err)
	util.Success(c, uid, "新增用户成功")
}

// 用户修改页面
func (w *UserApi) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	var userService service2.UserService
	user, err := userService.SelectRecordById(id)
	lv_err.HasErrAndPanic(err)
	//获取部门信息
	deptName := ""
	if user.DeptId > 0 {
		var deptServic service2.DeptService
		dept := deptServic.SelectDeptById(user.DeptId)
		if dept != nil {
			deptName = dept.DeptName
		}
	}
	var roleService service2.RoleService
	rolesP, _ := roleService.SelectRoleContactVo(id)
	var postService service2.SysPostService
	postP, _ := postService.SelectPostsByUserId(id)
	util.WriteTpl(c, "system/user/edit", gin.H{
		"user":     user,
		"deptName": deptName,
		"roles":    rolesP,
		"posts":    postP,
	})
}

// 重置密码
func (w *UserApi) ResetPwd(c *gin.Context) {
	id := lv_conv.Int64(c.Query("userId"))
	var userService service2.UserService
	user, err := userService.SelectRecordById(id)
	if err != nil || user == nil {
		util.WriteErrorTPL(c, gin.H{
			"desc": "用户不存在",
		})
		return
	}
	util.WriteTpl(c, "system/user/resetPwd", gin.H{
		"user": user,
	})
}
