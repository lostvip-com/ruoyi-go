// ==========================================================================
// LV自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-11-26 16:09:17 +0800 CST
// 生成人：lv
// ==========================================================================
package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	model2 "robvi/app/biz/model"
	"robvi/app/biz/service"
	sysmodel "robvi/app/system/model"
)

type HisPatientController struct{}

// /////////////////////////////////////////////////////////////////////////////////////
// 页面部分
// 列表分页数据
// /////////////////////////////////////////////////////////////////////////////////////
// 查询页
func (w HisPatientController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "biz/patient/list.html").WriteTpl()
}

// 新增页
func (w HisPatientController) Add(c *gin.Context) {
	lv_web.BuildTpl(c, "biz/patient/add.html").WriteTpl()
}

// 修改页
func (w HisPatientController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	patientService := service.HisPatientService{}
	entity, err := patientService.SelectRecordById(id)
	if err != nil || entity == nil {
		lv_web.ErrorTpl(c).WriteTpl(gin.H{"desc": "数据不存在"})
		return
	}
	lv_web.BuildTpl(c, "biz/patient/edit.html").WriteTpl(gin.H{
		"patient": entity,
	})
}

// /////////////////////////////////////////////////////////////////////////////////////
// api接口
// 列表分页数据
// /////////////////////////////////////////////////////////////////////////////////////
func (w HisPatientController) ListAjax(c *gin.Context) {
	req := new(model2.PageHisPatientReq)
	if err := c.ShouldBind(req); err != nil { //获取参数
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("patient管理", req).WriteJsonExit()
		return
	}
	rows := make([]model2.HisPatient, 0)
	patientService := service.HisPatientService{}
	result, page, err := patientService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	lv_web.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 新增页面保存
func (w HisPatientController) AddSave(c *gin.Context) {
	req := new(model2.AddHisPatientReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("患者基本信息新增数据", req).WriteJsonExit()
		return
	}
	patientService := service.HisPatientService{}
	id, err := patientService.AddSave(req, c)

	if err != nil || id <= 0 {
		lv_web.ErrorResp(c).Log("患者基本信息新增数据", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(id).Log("患者基本信息新增数据", req).WriteJsonExit()
}

// 修改页面保存
func (w HisPatientController) EditSave(c *gin.Context) {
	req := new(model2.EditHisPatientReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("患者基本信息修改数据", req).WriteJsonExit()
		return
	}
	patientService := service.HisPatientService{}
	rs, err := patientService.EditSave(req, c)

	if err != nil || rs <= 0 {
		lv_web.ErrorResp(c).Log("患者基本信息修改数据", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).Log("患者基本信息修改数据", req).WriteJsonExit()
}

// 删除数据
func (w HisPatientController) Remove(c *gin.Context) {
	req := new(sysmodel.RemoveReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("患者基本信息删除数据", req).WriteJsonExit()
		return
	}
	patientService := service.HisPatientService{}
	rs := patientService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		lv_web.SucessResp(c).Log("患者基本信息删除数据", req).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).Log("患者基本信息删除数据", req).WriteJsonExit()
	}
}

// 导出
func (w HisPatientController) Export(c *gin.Context) {
	req := new(model2.PageHisPatientReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).Log("患者基本信息导出数据", req).WriteJsonExit()
		return
	}

	patientService := service.HisPatientService{}
	url, err := patientService.Export(req)

	if err != nil {
		lv_web.ErrorResp(c).Log("患者基本信息导出数据", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetMsg(url).WriteJsonExit()
}
