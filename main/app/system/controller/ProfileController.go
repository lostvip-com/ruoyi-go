package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_web"
	"net/http"
	"os"
	"robvi/app/common/model_cmn"
	userModel "robvi/app/system/model/system"
	"robvi/app/system/service"
	"strconv"
	"time"
)

type ProfileController struct {
}

// 用户资料页面
func (w *ProfileController) Profile(c *gin.Context) {
	var userService service.UserService
	user := userService.GetProfile(c)
	lv_web.BuildTpl(c, "system/user/profile/profile").WriteTpl(gin.H{
		"user": user,
	})
}

// 修改用户信息
func (w *ProfileController) Update(c *gin.Context) {
	var req *userModel.ProfileReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
		return
	}
	var userService service.UserService
	err := userService.UpdateProfile(req, c)

	if err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
	} else {
		lv_web.SucessResp(c).SetBtype(model_cmn.Buniss_Edit).Log("用户管理", req).WriteJsonExit()
	}
}

// 修改用户密码
func (w *ProfileController) UpdatePassword(c *gin.Context) {
	var req *userModel.PasswordReq
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
	}
	var userService service.UserService
	err := userService.UpdatePassword(req, c)

	if err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
	} else {
		lv_web.SucessResp(c).SetBtype(model_cmn.Buniss_Edit).Log("修改用户密码", req).WriteJsonExit()
	}
}

// 修改头像页面
func (w *ProfileController) Avatar(c *gin.Context) {
	var userService service.UserService
	user := userService.GetProfile(c)
	lv_web.BuildTpl(c, "system/user/profile/avatar").WriteTpl(gin.H{
		"user": user,
	})
}

// 修改密码页面
func (w *ProfileController) EditPwd(c *gin.Context) {
	var userService service.UserService
	user := userService.GetProfile(c)
	lv_web.BuildTpl(c, "system/user/profile/resetPwd").WriteTpl(gin.H{
		"user": user,
	})
}

// 检查登录名是否存在
func (w *ProfileController) CheckLoginNameUnique(c *gin.Context) {
	var req *userModel.CheckLoginNameReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}
	var userService service.UserService
	result := userService.CheckLoginName(req.LoginName)

	if result {
		c.Writer.WriteString("1")
	} else {
		c.Writer.WriteString("0")
	}
}

// 检查邮箱是否存在
func (w *ProfileController) CheckEmailUnique(c *gin.Context) {
	var req *userModel.CheckEmailReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}
	var userService service.UserService
	result := userService.CheckEmailUnique(req.UserId, req.Email)

	if result {
		c.Writer.WriteString("1")
	} else {
		c.Writer.WriteString("0")
	}
}

// 检查邮箱是否存在
func (w *ProfileController) CheckEmailUniqueAll(c *gin.Context) {
	var req *userModel.CheckEmailAllReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}
	var userService service.UserService
	result := userService.CheckEmailUniqueAll(req.Email)

	if result {
		c.Writer.WriteString("1")
	} else {
		c.Writer.WriteString("0")
	}
}

// 检查手机号是否存在
func (w *ProfileController) CheckPhoneUnique(c *gin.Context) {
	var req *userModel.CheckPhoneReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}
	var userService service.UserService
	result := userService.CheckPhoneUnique(req.UserId, req.Phonenumber)

	if result {
		c.Writer.WriteString("1")
	} else {
		c.Writer.WriteString("0")
	}

}

// 检查手机号是否存在
func (w *ProfileController) CheckPhoneUniqueAll(c *gin.Context) {
	var req *userModel.CheckPhoneAllReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, model_cmn.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
	}
	var userService service.UserService
	result := userService.CheckPhoneUniqueAll(req.Phonenumber)

	if result {
		c.Writer.WriteString("1")
	} else {
		c.Writer.WriteString("0")
	}

}

// 校验密码是否正确
func (w *ProfileController) CheckPassword(c *gin.Context) {
	var req *userModel.CheckPasswordReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, model_cmn.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
	}
	var userService service.UserService
	user := userService.GetProfile(c)

	result := userService.CheckPassword(user, req.Password)

	if result {
		c.Writer.WriteString("true")
	} else {
		c.Writer.WriteString("false")
	}
}

// 保存头像
func (w *ProfileController) UpdateAvatar(c *gin.Context) {
	var userService service.UserService
	user := userService.GetProfile(c)

	curDir, err := os.Getwd()

	if err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg(err.Error()).Log("保存头像", gin.H{"userid": user.UserId}).WriteJsonExit()
	}

	saveDir := curDir + "/static/upload/"

	fileHead, err := c.FormFile("avatarfile")

	if err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg("没有获取到上传文件").Log("保存头像", gin.H{"userid": user.UserId}).WriteJsonExit()
	}

	curdate := time.Now().UnixNano()
	filename := user.LoginName + strconv.FormatInt(curdate, 10) + ".png"
	dts := saveDir + filename

	if err := c.SaveUploadedFile(fileHead, dts); err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg(err.Error()).Log("保存头像", gin.H{"userid": user.UserId}).WriteJsonExit()
	}

	avatar := "/upload/" + filename

	err = userService.UpdateAvatar(avatar, c)

	if err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg(err.Error()).Log("保存头像", gin.H{"userid": user.UserId}).WriteJsonExit()
	} else {
		lv_web.SucessResp(c).SetBtype(model_cmn.Buniss_Edit).Log("保存头像", gin.H{"userid": user.UserId}).WriteJsonExit()
	}
}
