package controller

import (
	"common/common_vo"
	"common/model"
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"main/internal/system/service"
)

type DictDataController struct {
}

// 列表分页数据
func (w *DictDataController) ListAjax(c *gin.Context) {
	var req *common_vo.SelectDictDataPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util.ErrorResp(c).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
		return
	}
	rows := make([]model.SysDictData, 0)
	var dictService service.DictDataService
	result, total, err := dictService.SelectListByPage(req)

	if err == nil && len(*result) > 0 {
		rows = *result
	}

	util.BuildTable(c, total, rows).WriteJsonExit()
}

// 新增页面
func (w *DictDataController) Add(c *gin.Context) {
	dictType := c.Query("dictType")
	util.BuildTpl(c, "system/dict/data/add").WriteTpl(gin.H{"dictType": dictType})
}

// 新增页面保存
func (w *DictDataController) AddSave(c *gin.Context) {
	var req *common_vo.AddDictDataReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util.ErrorResp(c).SetBtype(lv_dto.Buniss_Add).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
		return
	}
	var dictService service.DictDataService
	rid, err := dictService.AddSave(req, c)

	if err != nil || rid <= 0 {
		util.ErrorResp(c).SetBtype(lv_dto.Buniss_Add).Log("字典数据管理", req).WriteJsonExit()
		return
	}
	util.SucessResp(c).SetData(rid).SetBtype(lv_dto.Buniss_Add).Log("字典数据管理", req).WriteJsonExit()
}

// 修改页面
func (w *DictDataController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		util.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典数据错误",
		})
		return
	}
	var dictService service.DictDataService
	entity, err := dictService.SelectRecordById(id)

	if err != nil || entity == nil {
		util.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典数据不存在",
		})
		return
	}

	util.BuildTpl(c, "system/dict/data/edit").WriteTpl(gin.H{
		"dict": entity,
	})
}

// 修改页面保存
func (w *DictDataController) EditSave(c *gin.Context) {
	var req *common_vo.EditDictDataReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util.ErrorResp(c).SetBtype(lv_dto.Buniss_Edit).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
		return
	}
	var dictService service.DictDataService
	err := dictService.EditSave(req, c)

	if err != nil {
		util.ErrorResp(c).SetBtype(lv_dto.Buniss_Edit).WriteJsonExit()
		return
	}
	util.SucessResp(c).SetBtype(lv_dto.Buniss_Edit).WriteJsonExit()
}

// 删除数据
func (w *DictDataController) Remove(c *gin.Context) {
	var req *lv_dto.IdsReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util.ErrorResp(c).SetBtype(lv_dto.Buniss_Del).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
		return
	}
	var dictService service.DictDataService
	err := dictService.DeleteRecordByIds(req.Ids)

	if err == nil {
		util.SucessResp(c).SetBtype(lv_dto.Buniss_Del).WriteJsonExit()
	} else {
		util.ErrorResp(c).SetBtype(lv_dto.Buniss_Del).WriteJsonExit()
	}
}

// 导出
func (w *DictDataController) Export(c *gin.Context) {
	var req *common_vo.SelectDictDataPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util.ErrorResp(c).SetMsg(err.Error()).Log("字典数据导出", req).WriteJsonExit()
		return
	}
	var dictService service.DictDataService
	url, err := dictService.Export(req)

	if err != nil {
		util.ErrorResp(c).SetMsg(err.Error()).Log("字典数据导出", req).WriteJsonExit()
		return
	}
	util.SucessResp(c).SetMsg(url).Log("导出Excel", req).WriteJsonExit()
}
