package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	"robvi/app/common/model_cmn"
	"robvi/app/system/model/system/config"
	configService "robvi/app/system/service/system/config"
)

type ConfigController struct{}

// 列表页
func (w *ConfigController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "system/config/list").WriteTpl()
}

// 列表分页数据
func (w *ConfigController) ListAjax(c *gin.Context) {
	req := new(config.SelectPageReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}
	rows := make([]config.Entity, 0)
	result, page, err := configService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	lv_web.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 新增页面
func (w *ConfigController) Add(c *gin.Context) {
	lv_web.BuildTpl(c, "system/config/add").WriteTpl()
}

// 新增页面保存
func (w *ConfigController) AddSave(c *gin.Context) {
	req := new(config.AddReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Add).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}

	if configService.CheckConfigKeyUniqueAll(req.ConfigKey) == "1" {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Add).SetMsg("参数键名已存在").Log("参数管理", req).WriteJsonExit()
		return
	}

	rid, err := configService.AddSave(req, c)

	if err != nil || rid <= 0 {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Add).Log("参数管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(rid).Log("参数管理", req).WriteJsonExit()
}

// 修改页面
func (w *ConfigController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, model_cmn.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	entity, err := configService.SelectRecordById(id)

	if err != nil || entity == nil {
		lv_web.BuildTpl(c, model_cmn.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "数据不存在",
		})
		return
	}

	lv_web.BuildTpl(c, "system/config/edit").WriteTpl(gin.H{
		"config": entity,
	})
}

// 修改页面保存
func (w *ConfigController) EditSave(c *gin.Context) {
	req := new(config.EditReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}

	if configService.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId) == "1" {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).SetMsg("参数键名已存在").Log("参数管理", req).WriteJsonExit()
		return
	}

	rs, err := configService.EditSave(req, c)

	if err != nil || rs <= 0 {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Edit).Log("参数管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetBtype(model_cmn.Buniss_Edit).Log("参数管理", req).WriteJsonExit()
}

// 删除数据
func (w *ConfigController) Remove(c *gin.Context) {
	req := new(model_cmn.RemoveReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Del).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}

	rs := configService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(model_cmn.Buniss_Del).Log("参数管理", req).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Del).Log("参数管理", req).WriteJsonExit()
	}
}

// 导出
func (w *ConfigController) Export(c *gin.Context) {
	req := new(config.SelectPageReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).Log("参数管理", req).WriteJsonExit()
		return
	}
	url, err := configService.Export(req)

	if err != nil {
		lv_web.ErrorResp(c).SetBtype(model_cmn.Buniss_Other).Log("参数管理", req).WriteJsonExit()
		return
	}

	lv_web.SucessResp(c).SetBtype(model_cmn.Buniss_Other).SetMsg(url).WriteJsonExit()
}

// 检查参数键名是否已经存在不包括本参数
func (w *ConfigController) CheckConfigKeyUnique(c *gin.Context) {
	var req *config.CheckConfigKeyReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := configService.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId)

	c.Writer.WriteString(result)
}

// 检查参数键名是否已经存在
func (w *ConfigController) CheckConfigKeyUniqueAll(c *gin.Context) {
	var req *config.CheckPostCodeALLReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := configService.CheckConfigKeyUniqueAll(req.ConfigKey)

	c.Writer.WriteString(result)
}
