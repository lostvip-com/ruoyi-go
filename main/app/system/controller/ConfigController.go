package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/utils/lv_conv"
	"github.com/lv_framework/utils/lv_err"
	"github.com/lv_framework/utils/lv_web"
	"github.com/lv_framework/web/dto"
	"main/app/system/service"
	"main/app/system/vo"
)

type ConfigController struct{}

// 列表页
func (w *ConfigController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "system/config/list").WriteTpl()
}

// 列表分页数据
func (w *ConfigController) ListAjax(c *gin.Context) {
	req := new(vo.SelectConfigPageReq)
	var configService service.ConfigService
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}
	result, total, err := configService.SelectListByPage(req)
	lv_err.HasErrAndPanic(err)
	lv_web.BuildTable(c, total, result).WriteJsonExit()
}

// 新增页面
func (w *ConfigController) Add(c *gin.Context) {
	lv_web.BuildTpl(c, "system/config/add").WriteTpl()
}

// 新增页面保存
func (w *ConfigController) AddSave(c *gin.Context) {
	req := new(vo.AddConfigReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}
	var configService service.ConfigService
	if configService.CheckConfigKeyUniqueAll(req.ConfigKey) == "1" {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg("参数键名已存在").Log("参数管理", req).WriteJsonExit()
		return
	}

	rid, err := configService.AddSave(req, c)

	if err != nil || rid <= 0 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).Log("参数管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(rid).Log("参数管理", req).WriteJsonExit()
}

// 修改页面
func (w *ConfigController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	var configService service.ConfigService
	entity, err := configService.SelectRecordById(id)

	if err != nil || entity == nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
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
	req := new(vo.EditConfigReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}
	var configService service.ConfigService
	if configService.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId) == "1" {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).SetMsg("参数键名已存在").Log("参数管理", req).WriteJsonExit()
		return
	}
	configService.EditSave(req, c)
	lv_web.SucessResp(c).SetBtype(dto.Buniss_Edit).Log("参数管理", req).WriteJsonExit()
}

// 删除数据
func (w *ConfigController) Remove(c *gin.Context) {
	req := new(dto.IdsReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}
	var configService service.ConfigService
	configService.DeleteRecordByIds(req.Ids)
	lv_web.SucessResp(c).SetBtype(dto.Buniss_Del).Log("参数管理", req).WriteJsonExit()
}

// 导出
func (w *ConfigController) Export(c *gin.Context) {
	req := new(vo.SelectConfigPageReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		lv_web.ErrorResp(c).Log("参数管理", req).WriteJsonExit()
		return
	}
	var configService service.ConfigService
	url, err := configService.Export(req)

	if err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Other).Log("参数管理", req).WriteJsonExit()
		return
	}

	lv_web.SucessResp(c).SetBtype(dto.Buniss_Other).SetMsg(url).WriteJsonExit()
}

// 检查参数键名是否已经存在不包括本参数
func (w *ConfigController) CheckConfigKeyUnique(c *gin.Context) {
	var req *vo.CheckConfigKeyReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}
	var configService service.ConfigService
	result := configService.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId)

	c.Writer.WriteString(result)
}

// 检查参数键名是否已经存在
func (w *ConfigController) CheckConfigKeyUniqueAll(c *gin.Context) {
	var req *vo.AddConfigReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}
	var configService service.ConfigService
	result := configService.CheckConfigKeyUniqueAll(req.ConfigKey)

	c.Writer.WriteString(result)
}
