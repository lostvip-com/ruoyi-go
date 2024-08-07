package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	"github.com/lostvip-com/lv_framework/web/dto"
	"main/internal/system/model/monitor/logininfor"
	logininforService "main/internal/system/service/monitor/logininfor"
)

type LoginInforController struct {
}

// 用户列表页
func (w *LoginInforController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "monitor/logininfor/list").WriteTpl()
}

// 用户列表分页数据
func (w *LoginInforController) ListAjax(c *gin.Context) {
	var req *logininfor.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}

	rows := make([]logininfor.Entity, 0)

	result, page, err := logininforService.SelectPageList(req)

	if err == nil && len(*result) > 0 {
		rows = *result
	}
	lv_web.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 删除数据
func (w *LoginInforController) Remove(c *gin.Context) {
	var req *dto.IdsReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).SetMsg(err.Error()).Log("登录日志管理", req).WriteJsonExit()
		return
	}

	rs := logininforService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(dto.Buniss_Del).SetData(rs).Log("登录日志管理", req).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).Log("登录日志管理", req).WriteJsonExit()
	}
}

// 清空记录
func (w *LoginInforController) Clean(c *gin.Context) {

	rs, _ := logininforService.DeleteRecordAll()

	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(dto.Buniss_Del).SetData(rs).Log("登录日志管理", "all").WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).Log("登录日志管理", "all").WriteJsonExit()
	}
}

// 导出
func (w *LoginInforController) Export(c *gin.Context) {
	var req *logininfor.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("导出登录日志", req).WriteJsonExit()
		return
	}

	url, err := logininforService.Export(req)

	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("导出登录日志", req).WriteJsonExit()
	} else {
		lv_web.SucessResp(c).SetMsg(url).Log("导出登录日志", req).WriteJsonExit()
	}
}

// 解锁账号
func (w *LoginInforController) Unlock(c *gin.Context) {
	loginName := c.Query("loginName")
	if loginName == "" {
		lv_web.ErrorResp(c).SetMsg("参数错误").Log("解锁账号", "loginName="+loginName).WriteJsonExit()
	} else {
		logininforService.RemovePasswordCounts(loginName)
		logininforService.Unlock(loginName)
		lv_web.SucessResp(c).Log("解锁账号", "loginName="+loginName).WriteJsonExit()
	}

}
