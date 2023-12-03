package user

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	"robvi/app/common/session"
	"robvi/app/system/model"
	postModel "robvi/app/system/model/system/post"
	roleModel "robvi/app/system/model/system/role"
	userModel "robvi/app/system/model/system/user"
	service2 "robvi/app/system/service"
	postService "robvi/app/system/service/system/post"
	roleService "robvi/app/system/service/system/role"
)

// 用户列表页
func List(c *gin.Context) {
	lv_web.BuildTpl(c, "system/user/list").WriteTpl()
}

// 用户列表分页数据
func ListAjax(c *gin.Context) {
	var req *userModel.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
		return
	}
	tenantId := session.GetProfile(c).TenantId
	req.TenantId = tenantId
	rows := make([]userModel.UserListEntity, 0)
	var userService service2.UserService
	result, page, err := userService.SelectRecordList(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	lv_web.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 用户新增页面
func Add(c *gin.Context) {
	var paramsRole *roleModel.SelectPageReq
	var paramsPost *postModel.SelectPageReq

	roles := make([]roleModel.EntityFlag, 0)
	posts := make([]postModel.EntityFlag, 0)

	rolesP, _ := roleService.SelectRecordAll(paramsRole)

	if rolesP != nil {
		roles = rolesP
	}

	postP, _ := postService.SelectListAll(paramsPost)

	if postP != nil {
		posts = postP
	}
	lv_web.BuildTpl(c, "system/user/add").WriteTpl(gin.H{
		"roles": roles,
		"posts": posts,
	})
}

// 保存新增用户数据
func AddSave(c *gin.Context) {
	var req *userModel.AddReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("新增用户", req).WriteJsonExit()
		return
	}
	var userService service2.UserService
	//判断登陆名是否已注册
	isHadName := userService.CheckLoginName(req.LoginName)
	if isHadName {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("登陆名已经存在").Log("新增用户", req).WriteJsonExit()
		return
	}

	//判断手机号码是否已注册
	isHadPhone := userService.CheckPhoneUniqueAll(req.Phonenumber)
	if isHadPhone {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("手机号码已经存在").Log("新增用户", req).WriteJsonExit()
		return
	}

	//判断邮箱是否已注册
	isHadEmail := userService.CheckEmailUniqueAll(req.Email)
	if isHadEmail {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("邮箱已经存在").Log("新增用户", req).WriteJsonExit()
		return
	}

	uid, err := userService.AddSave(req, c)

	if err != nil || uid <= 0 {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).Log("新增用户", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(uid).SetBtype(model.Buniss_Add).Log("新增用户", req).WriteJsonExit()
}

// 用户修改页面
func Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))

	if id <= 0 {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	var userService service2.UserService
	user, err := userService.SelectRecordById(id)

	if err != nil || user == nil {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
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

	roles := make([]roleModel.EntityFlag, 0)
	posts := make([]postModel.EntityFlag, 0)

	rolesP, _ := roleService.SelectRoleContactVo(id)

	if rolesP != nil {
		roles = rolesP
	}

	postP, _ := postService.SelectPostsByUserId(id)

	if postP != nil {
		posts = postP
	}

	lv_web.BuildTpl(c, "system/user/edit").WriteTpl(gin.H{
		"user":     user,
		"deptName": deptName,
		"roles":    roles,
		"posts":    posts,
	})
}

// 重置密码
func ResetPwd(c *gin.Context) {
	id := lv_conv.Int64(c.Query("userId"))
	if id <= 0 {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	var userService service2.UserService
	user, err := userService.SelectRecordById(id)

	if err != nil || user == nil {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "用户不存在",
		})
		return
	}
	lv_web.BuildTpl(c, "system/user/resetPwd").WriteTpl(gin.H{
		"user": user,
	})
}

// 重置密码保存
func ResetPwdSave(c *gin.Context) {
	var req *userModel.ResetPwdReq
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("重置密码", req).WriteJsonExit()
	}
	var userService service2.UserService
	result, err := userService.ResetPassword(req)

	if err != nil || !result {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("重置密码", req).WriteJsonExit()
	} else {
		lv_web.SucessResp(c).SetBtype(model.Buniss_Edit).Log("重置密码", req).WriteJsonExit()
	}
}

// 保存修改用户数据
func EditSave(c *gin.Context) {
	var req *userModel.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("修改用户", req).WriteJsonExit()
		return
	}
	var userService service2.UserService
	//判断手机号码是否已注册
	isHadPhone := userService.CheckPhoneUnique(req.UserId, req.Phonenumber)
	if isHadPhone {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg("手机号码已经存在").Log("修改用户", req).WriteJsonExit()
		return
	}

	//判断邮箱是否已注册
	isHadEmail := userService.CheckEmailUnique(req.UserId, req.Email)
	if isHadEmail {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg("邮箱已经存在").Log("修改用户", req).WriteJsonExit()
		return
	}

	uid, err := userService.EditSave(req, c)

	if err != nil || uid <= 0 {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).Log("修改用户", req).WriteJsonExit()
		return
	}

	lv_web.SucessResp(c).SetData(uid).SetBtype(model.Buniss_Edit).Log("修改用户", req).WriteJsonExit()
}

// 删除数据
func Remove(c *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("删除用户", req).WriteJsonExit()
	}
	var userService service2.UserService
	rs := userService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		lv_web.SucessResp(c).SetData(rs).SetBtype(model.Buniss_Del).Log("删除用户", req).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Del).Log("删除用户", req).WriteJsonExit()
	}
}

// 导出
func Export(c *gin.Context) {
	var req *userModel.SelectPageReq
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
