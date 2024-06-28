package controller

import (
	"github.com/gin-gonic/gin"
	db2 "github.com/lv_framework/db"
	"github.com/lv_framework/utils/lv_web"
	"github.com/lv_framework/web/dto"
	service2 "main/app/system/service"
	userModel "main/app/system/vo"
)

type UserApi struct {
}

// 修改页面保存
func (w *UserApi) ChangeStatus(c *gin.Context) {
	userId := c.PostForm("userId")
	status := c.PostForm("status")
	sql := " update sys_user set status=? where user_id = ? "
	rows := db2.GetMasterGorm().Exec(sql, status, userId).RowsAffected
	lv_web.SuccessData(c, rows)
}

// 重置密码保存
func (w *UserApi) ResetPwdSave(c *gin.Context) {
	var req *userModel.ResetPwdReq
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).SetMsg(err.Error()).Log("重置密码", req).WriteJsonExit()
	}
	var userService service2.UserService
	result, err := userService.ResetPassword(req)

	if err != nil || !result {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).SetMsg(err.Error()).Log("重置密码", req).WriteJsonExit()
	} else {
		lv_web.SucessResp(c).SetBtype(dto.Buniss_Edit).Log("重置密码", req).WriteJsonExit()
	}
}

// 保存修改用户数据
func (w *UserApi) EditSave(c *gin.Context) {
	var req *userModel.EditUserReq
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	var userService service2.UserService
	err := userService.EditSave(req, c)
	if err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(req.UserId).SetBtype(dto.Buniss_Edit).WriteJsonExit()
}

// 删除数据
func (w *UserApi) Remove(c *gin.Context) {
	var req *dto.IdsReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).SetMsg(err.Error()).WriteJsonExit()
	}
	var userService service2.UserService
	err := userService.DeleteRecordByIds(req.Ids)
	if err == nil {
		lv_web.SucessResp(c).SetBtype(dto.Buniss_Del).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).WriteJsonExit()
	}
}

// 导出
func (w *UserApi) Export(c *gin.Context) {
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

func (w *UserApi) GetUserInfo(c *gin.Context) {
	var req *dto.IdsReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).SetMsg(err.Error()).WriteJsonExit()
	}
	var userService service2.UserService
	user := userService.GetProfile(c)

	data := make(map[string]any)
	data["roles"] = ""
	data["permissions"] = ""
	data["user"] = user
	data["dept"] = ""
	lv_web.Success(c, data, "success")
}
