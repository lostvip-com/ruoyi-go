package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	"robvi/app/common/model_cmn"
	service2 "robvi/app/system/service"
	roleService "robvi/app/system/service/system/role"
	userModel "robvi/app/system/vo"
)

type UserController struct {
}

// 用户列表页
func (w *UserController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "system/user/list").WriteTpl()
}

// 用户列表分页数据
func (w *UserController) ListAjax(c *gin.Context) {
	var req *userModel.SelectUserPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
		return
	}
	var rows []map[string]string
	var userService service2.UserService
	result, total, err := userService.SelectRecordList(req)

	if err == nil && len(*result) > 0 {
		rows = *result
	}
	lv_web.BuildTable(c, total, rows).WriteJsonExit()
}

// 用户新增页面
func (w *UserController) Add(c *gin.Context) {
	rolesP, _ := roleService.SelectRecordAll(nil)
	var postService service2.SysPostService
	postP, _ := postService.SelectListAll(nil)
	lv_web.BuildTpl(c, "system/user/add").WriteTpl(gin.H{
		"roles": rolesP,
		"posts": postP,
	})
}

// 保存新增用户数据
func (w *UserController) AddSave(c *gin.Context) {
	var req *userModel.AddReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Add).SetMsg(err.Error()).Log("新增用户", req).WriteJsonExit()
		return
	}
	var userService service2.UserService
	//判断登录名是否已注册
	isHadName := userService.CheckLoginName(req.LoginName)
	if isHadName {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Add).SetMsg("登录名已经存在").Log("新增用户", req).WriteJsonExit()
		return
	}
	//判断手机号码是否已注册
	isHadPhone := userService.CheckPhoneUniqueAll(req.Phonenumber)
	if isHadPhone {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Add).SetMsg("手机号码已经存在").Log("新增用户", req).WriteJsonExit()
		return
	}
	//判断邮箱是否已注册
	isHadEmail := userService.CheckEmailUniqueAll(req.Email)
	if isHadEmail {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Add).SetMsg("邮箱已经存在").Log("新增用户", req).WriteJsonExit()
		return
	}

	uid, err := userService.AddSave(req, c)

	if err != nil || uid <= 0 {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Add).Log("新增用户", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(uid).SetBtype(model_cmn.Buniss_Add).Log("新增用户", req).WriteJsonExit()
}

// 用户修改页面
func (w *UserController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	var userService service2.UserService
	user, err := userService.SelectRecordById(id)
	if err != nil || user == nil {
		lv_web.BuildTpl(c, model_cmn.ERROR_PAGE).WriteTpl(gin.H{
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
func (w *UserController) ResetPwd(c *gin.Context) {
	id := lv_conv.Int64(c.Query("userId"))
	if id <= 0 {
		lv_web.BuildTpl(c, model_cmn.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	var userService service2.UserService
	user, err := userService.SelectRecordById(id)

	if err != nil || user == nil {
		lv_web.BuildTpl(c, model_cmn.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "用户不存在",
		})
		return
	}
	lv_web.BuildTpl(c, "system/user/resetPwd").WriteTpl(gin.H{
		"user": user,
	})
}

// 重置密码保存
func (w *UserController) ResetPwdSave(c *gin.Context) {
	var req *userModel.ResetPwdReq
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg(err.Error()).Log("重置密码", req).WriteJsonExit()
	}
	var userService service2.UserService
	result, err := userService.ResetPassword(req)

	if err != nil || !result {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg(err.Error()).Log("重置密码", req).WriteJsonExit()
	} else {
		lv_web.SucessResp(c).SetBtype(model_cmn.Buniss_Edit).Log("重置密码", req).WriteJsonExit()
	}
}

// 保存修改用户数据
func (w *UserController) EditSave(c *gin.Context) {
	var req *userModel.EditReq
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	var userService service2.UserService
	err := userService.EditSave(req, c)
	if err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(req.UserId).SetBtype(model_cmn.Buniss_Edit).WriteJsonExit()
}

// 删除数据
func (w *UserController) Remove(c *gin.Context) {
	var req *model_cmn.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Del).SetMsg(err.Error()).WriteJsonExit()
	}
	var userService service2.UserService
	err := userService.DeleteRecordByIds(req.Ids)
	if err == nil {
		lv_web.SucessResp(c).SetBtype(model_cmn.Buniss_Del).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Del).WriteJsonExit()
	}
}

// 导出
func (w *UserController) Export(c *gin.Context) {
	var req *userModel.SelectUserPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("导出Excel", req).WriteJsonExit()
	}
	var userService service2.UserService
	url, err := userService.Export(req)

	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("导出Excel", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetMsg(url).Log("导出Excel", req).WriteJsonExit()
}
