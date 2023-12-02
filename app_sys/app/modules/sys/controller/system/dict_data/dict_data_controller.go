package dict_data

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	"robvi/app/modules/sys/model"
	dict_data2 "robvi/app/modules/sys/model/system/dict_data"
	dictService "robvi/app/modules/sys/service/system/dict_data"
)

// 列表分页数据
func ListAjax(c *gin.Context) {
	var req *dict_data2.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
		return
	}
	rows := make([]dict_data2.Entity, 0)
	result, page, err := dictService.SelectListByPage(req)

	if err == nil && len(*result) > 0 {
		rows = *result
	}

	lv_web.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 新增页面
func Add(c *gin.Context) {
	dictType := c.Query("dictType")
	lv_web.BuildTpl(c, "system/dict/data/add").WriteTpl(gin.H{"dictType": dictType})
}

// 新增页面保存
func AddSave(c *gin.Context) {
	var req *dict_data2.AddReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
		return
	}

	rid, err := dictService.AddSave(req, c)

	if err != nil || rid <= 0 {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).Log("字典数据管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(rid).SetBtype(model.Buniss_Add).Log("字典数据管理", req).WriteJsonExit()
}

// 修改页面
func Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典数据错误",
		})
		return
	}

	entity, err := dictService.SelectRecordById(id)

	if err != nil || entity == nil {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典数据不存在",
		})
		return
	}

	lv_web.BuildTpl(c, "system/dict/data/edit").WriteTpl(gin.H{
		"dict": entity,
	})
}

// 修改页面保存
func EditSave(c *gin.Context) {
	var req *dict_data2.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
		return
	}

	rs, err := dictService.EditSave(req, c)

	if err != nil || rs <= 0 {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).Log("字典数据管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetBtype(model.Buniss_Edit).SetData(rs).Log("字典数据管理", req).WriteJsonExit()
}

// 删除数据
func Remove(c *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
		return
	}

	rs := dictService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(model.Buniss_Del).SetData(rs).Log("字典数据管理", req).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Del).Log("字典数据管理", req).WriteJsonExit()
	}
}

// 导出
func Export(c *gin.Context) {
	var req *dict_data2.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("字典数据导出", req).WriteJsonExit()
		return
	}
	url, err := dictService.Export(req)

	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("字典数据导出", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetMsg(url).Log("导出Excel", req).WriteJsonExit()
}
