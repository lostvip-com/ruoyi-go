// ==========================================================================
// LV自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-06-29 22:21:21 +0800 CST
// 生成路径: app/controller/modules/tenant/TenantController.go
// 生成人：robnote
// ==========================================================================
package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/gconv"
	response2 "lostvip.com/utils/response"
	"robvi/app/modules/sys/model"
	tenant2 "robvi/app/modules/sys/model/module/tenant"
	"robvi/app/modules/sys/service"
)

type TenantController struct {
}

// 列表页
func (w TenantController) List(c *gin.Context) {
	response2.BuildTpl(c, "modules/tenant/list.html").WriteTpl()
}

// 列表分页数据
func (w TenantController) ListAjax(c *gin.Context) {
	req := new(tenant2.SelectPageReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response2.ErrorResp(c).SetMsg(err.Error()).Log("tenant管理", req).WriteJsonExit()
		return
	}
	rows := make([]tenant2.SysTenant, 0)
	tenantService := service.TenantService{}
	result, page, err := tenantService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response2.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 新增页面
func (w TenantController) Add(c *gin.Context) {
	response2.BuildTpl(c, "modules/tenant/add.html").WriteTpl()
}

// 新增页面保存
func (w TenantController) AddSave(c *gin.Context) {
	req := new(tenant2.AddReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("商户信息新增数据", req).WriteJsonExit()
		return
	}
	tenantService := service.TenantService{}
	id, err := tenantService.AddSave(req, c)

	if err != nil || id <= 0 {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).Log("商户信息新增数据", req).WriteJsonExit()
		return
	}
	response2.SucessResp(c).SetData(id).Log("商户信息新增数据", req).WriteJsonExit()
}

// 修改页面
func (w TenantController) Edit(c *gin.Context) {
	id := gconv.Int64(c.Query("id"))

	if id <= 0 {
		response2.ErrorTpl(c).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	tenantService := service.TenantService{}
	entity, err := tenantService.SelectRecordById(id)

	if err != nil || entity == nil {
		response2.ErrorTpl(c).WriteTpl(gin.H{
			"desc": "数据不存在",
		})
		return
	}

	response2.BuildTpl(c, "modules/tenant/edit.html").WriteTpl(gin.H{
		"tenant": entity,
	})
}

// 修改页面保存
func (w TenantController) EditSave(c *gin.Context) {
	var req = new(tenant2.EditReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("商户信息修改数据", req).WriteJsonExit()
		return
	}
	tenantService := service.TenantService{}
	rs, err := tenantService.EditSave(req, c)
	if err != nil || rs <= 0 {
		response2.ErrorResp(c).SetBtype(model.Buniss_Edit).Log("商户信息修改数据", req).WriteJsonExit()
		return
	}
	response2.SucessResp(c).SetBtype(model.Buniss_Edit).Log("商户信息修改数据", req).WriteJsonExit()
}

// 删除数据
func (w TenantController) Remove(c *gin.Context) {
	req := new(model.RemoveReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("商户信息删除数据", req).WriteJsonExit()
		return
	}
	tenantService := service.TenantService{}
	rs := tenantService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response2.SucessResp(c).SetBtype(model.Buniss_Del).Log("商户信息删除数据", req).WriteJsonExit()
	} else {
		response2.ErrorResp(c).SetBtype(model.Buniss_Del).Log("商户信息删除数据", req).WriteJsonExit()
	}
}

// 导出
func (w TenantController) Export(c *gin.Context) {
	req := new(tenant2.SelectPageReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response2.ErrorResp(c).Log("商户信息导出数据", req).WriteJsonExit()
		return
	}

	tenantService := service.TenantService{}
	url, err := tenantService.Export(req)

	if err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Other).Log("商户信息导出数据", req).WriteJsonExit()
		return
	}
	response2.SucessResp(c).SetBtype(model.Buniss_Other).SetMsg(url).WriteJsonExit()
}
