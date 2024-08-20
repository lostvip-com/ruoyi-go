package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"html/template"
	"main/internal/system/service"
	"main/internal/system/vo"
)

type OperlogController struct {
}

// List  用户列表页
func (w *OperlogController) List(c *gin.Context) {
	util.WriteTpl(c, "monitor/list_oper_log")
}

// ListAjax 用户列表分页数据
func (w *OperlogController) ListAjax(c *gin.Context) {
	var req vo.OperLogPageReq
	//获取参数
	err := c.ShouldBind(&req)
	lv_err.HasErrAndPanic(err)
	var operlogService service.OperLogService
	rows, total, err := operlogService.SelectPageList(&req)
	lv_err.HasErrAndPanic(err)
	util.SucessPage(c, rows, total)
}

// 清空记录
func (w *OperlogController) Clean(c *gin.Context) {
	var operlogService service.OperLogService
	err := operlogService.DeleteRecordAll()
	lv_err.HasErrAndPanic(err)
	util.Success(c, nil, "success")
}

// 删除数据
func (w *OperlogController) Remove(c *gin.Context) {
	var req *lv_dto.IdsReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util.ErrorResp(c).SetBtype(lv_dto.Buniss_Del).SetMsg(err.Error()).Log("操作日志管理", req).WriteJsonExit()
		return
	}
	var operlogService service.OperLogService
	err := operlogService.DeleteRecordByIds(req.Ids)
	lv_err.HasErrAndPanic(err)
	util.Success(c, nil, "删除成功")
}

// 记录详情
func (w *OperlogController) Detail(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	var operlogService service.OperLogService
	operLog, err := operlogService.SelectRecordById(id)
	lv_err.HasErrAndPanic(err)
	jsonResult := template.HTML(operLog.JsonResult)
	operParam := template.HTML(operLog.OperParam)
	util.WriteTpl(c, "monitor/detail_oper_log", gin.H{
		"operLog":    operLog,
		"jsonResult": jsonResult,
		"operParam":  operParam,
	})
}
