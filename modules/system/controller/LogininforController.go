package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"system/model"
	"system/service"
	"system/vo"
)

type LoginInforController struct {
}

// 用户列表页
func (w *LoginInforController) List(c *gin.Context) {
	util.BuildTpl(c, "monitor/list_login_info").WriteTpl()
}

// 用户列表分页数据
func (w *LoginInforController) ListAjax(c *gin.Context) {
	var req *vo.LoginInfoPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}

	rows := make([]model.SysLoginInfo, 0)
	var logininforService service.LoginInforService
	result, total, err := logininforService.SelectPageList(req)

	if err == nil && len(*result) > 0 {
		rows = *result
	}
	util.BuildTable(c, total, rows).WriteJsonExit()
}

// 删除数据
func (w *LoginInforController) Remove(c *gin.Context) {
	var req *lv_dto.IdsReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util.ErrorResp(c).Log("登录日志管理", req).WriteJsonExit()
		return
	}
	var logininforService service.LoginInforService
	err := logininforService.DeleteRecordByIds(req.Ids)

	if err == nil {
		util.SucessResp(c).Log("登录日志管理", req).WriteJsonExit()
	} else {
		util.ErrorResp(c).Log("登录日志管理", req).WriteJsonExit()
	}
}

// 清空记录
func (w *LoginInforController) Clean(c *gin.Context) {
	var logininforService service.LoginInforService
	err := logininforService.DeleteRecordAll()

	if err == nil {
		util.SucessResp(c).Log("登录日志管理", "all").WriteJsonExit()
	} else {
		util.ErrorResp(c).Log("登录日志管理", "all").WriteJsonExit()
	}
}

// 导出
func (w *LoginInforController) Export(c *gin.Context) {
	var req *vo.LoginInfoPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util.ErrorResp(c).SetMsg(err.Error()).Log("导出登录日志", req).WriteJsonExit()
		return
	}
	var logininforService service.LoginInforService
	url, err := logininforService.Export(req)
	if err != nil {
		util.ErrorResp(c).SetMsg(err.Error()).Log("导出登录日志", req).WriteJsonExit()
	} else {
		util.SucessResp(c).SetMsg(url).Log("导出登录日志", req).WriteJsonExit()
	}
}

// 解锁账号
func (w *LoginInforController) Unlock(c *gin.Context) {
	loginName := c.Query("loginName")
	if loginName == "" {
		util.ErrorResp(c).SetMsg("参数错误").Log("解锁账号", "loginName="+loginName).WriteJsonExit()
	} else {
		var logininforService service.LoginInforService
		logininforService.RemovePasswordCounts(loginName)
		logininforService.Unlock(loginName)
		util.SucessResp(c).Log("解锁账号", "loginName="+loginName).WriteJsonExit()
	}

}
