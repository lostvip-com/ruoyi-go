package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/utils/lv_conv"
	"github.com/lv_framework/utils/lv_err"
	"github.com/lv_framework/utils/lv_web"
	"github.com/lv_framework/web/dto"
	service2 "main/internal/system/service"
	userModel "main/internal/system/vo"
)

// 用户列表页
func (w *UserApi) List(c *gin.Context) {
	lv_web.BuildTpl(c, "system/user/list").WriteTpl()
}

// 用户列表分页数据
func (w *UserApi) ListAjax(c *gin.Context) {
	var req *userModel.SelectUserPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
		return
	}
	var userService service2.UserService
	result, total, err := userService.SelectRecordList(req)
	lv_err.HasErrAndPanic(err)
	lv_web.BuildTable(c, total, result).WriteJsonExit()
}

// 用户新增页面
func (w *UserApi) Add(c *gin.Context) {
	var roleService service2.RoleService
	rolesP, _ := roleService.SelectRecordAll(nil)
	var postService service2.SysPostService
	postP, _ := postService.SelectListAll(nil)
	lv_web.BuildTpl(c, "system/user/add").WriteTpl(gin.H{
		"roles": rolesP,
		"posts": postP,
	})
}

// 保存新增用户数据
func (w *UserApi) AddSave(c *gin.Context) {
	var req *userModel.AddUserReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg(err.Error()).Log("新增用户", req).WriteJsonExit()
		return
	}
	var userService service2.UserService
	//判断登录名是否已注册
	isHadName := userService.CheckLoginName(req.LoginName)
	if isHadName {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg("登录名已经存在").Log("新增用户", req).WriteJsonExit()
		return
	}
	//判断手机号码是否已注册
	isHadPhone := userService.CheckPhoneUniqueAll(req.Phonenumber)
	if isHadPhone {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg("手机号码已经存在").Log("新增用户", req).WriteJsonExit()
		return
	}
	//判断邮箱是否已注册
	isHadEmail := userService.CheckEmailUniqueAll(req.Email)
	if isHadEmail {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg("邮箱已经存在").Log("新增用户", req).WriteJsonExit()
		return
	}

	uid, err := userService.AddSave(req, c)

	if err != nil || uid <= 0 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).Log("新增用户", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(uid).SetBtype(dto.Buniss_Add).Log("新增用户", req).WriteJsonExit()
}

// 用户修改页面
func (w *UserApi) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	var userService service2.UserService
	user, err := userService.SelectRecordById(id)
	if err != nil || user == nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "用户不存在",
		})
		return
	}
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
	lv_web.BuildTpl(c, "system/user/edit").WriteTpl(gin.H{
		"user":     user,
		"deptName": deptName,
		"roles":    rolesP,
		"posts":    postP,
	})
}

// 重置密码
func (w *UserApi) ResetPwd(c *gin.Context) {
	id := lv_conv.Int64(c.Query("userId"))
	if id <= 0 {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	var userService service2.UserService
	user, err := userService.SelectRecordById(id)

	if err != nil || user == nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "用户不存在",
		})
		return
	}
	lv_web.BuildTpl(c, "system/user/resetPwd").WriteTpl(gin.H{
		"user": user,
	})
}
