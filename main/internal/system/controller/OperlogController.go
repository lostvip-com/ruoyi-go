package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	"github.com/lostvip-com/lv_framework/web/dto"
	"html/template"
	"main/internal/system/model/monitor/oper_log"
	operlogService "main/internal/system/service/monitor/operlog"
	"net/http"
)

type OperlogController struct {
}

// 用户列表页
func (w *OperlogController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "monitor/operlog/list").WriteTpl()
}

// 用户列表分页数据
func (w *OperlogController) ListAjax(c *gin.Context) {
	var req *oper_log.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, dto.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
	}

	rows := make([]oper_log.Entity, 0)

	result, page, err := operlogService.SelectPageList(req)

	if err == nil && len(*result) > 0 {
		rows = *result
	}

	lv_web.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 清空记录
func (w *OperlogController) Clean(c *gin.Context) {

	rs, _ := operlogService.DeleteRecordAll()

	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(dto.Buniss_Del).SetData(rs).Log("操作日志管理", "all").WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).Log("操作日志管理", "all").WriteJsonExit()
	}
}

// 删除数据
func (w *OperlogController) Remove(c *gin.Context) {
	var req *dto.IdsReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).SetMsg(err.Error()).Log("操作日志管理", req).WriteJsonExit()
		return
	}

	rs := operlogService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(dto.Buniss_Del).SetData(rs).Log("操作日志管理", req).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).Log("操作日志管理", req).WriteJsonExit()
	}
}

// 记录详情
func (w *OperlogController) Detail(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	operLog, err := operlogService.SelectRecordById(id)

	if err != nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "数据不存在",
		})
		return
	}

	jsonResult := template.HTML(operLog.JsonResult)
	operParam := template.HTML(operLog.OperParam)
	lv_web.BuildTpl(c, "monitor/operlog/detail").WriteTpl(gin.H{
		"operLog":    operLog,
		"jsonResult": jsonResult,
		"operParam":  operParam,
	})
}

// 导出
func (w *OperlogController) Export(c *gin.Context) {
	var req *oper_log.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("导出操作日志", req).WriteJsonExit()
		return
	}
	url, err := operlogService.Export(req)

	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("导出操作日志", req).WriteJsonExit()
	} else {
		lv_web.SucessResp(c).SetMsg(url).Log("导出操作日志", req).WriteJsonExit()
	}
}
