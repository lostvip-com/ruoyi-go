package config

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/gconv"
	response2 "lostvip.com/utils/response"
	"robvi/app/modules/sys/model"
	config2 "robvi/app/modules/sys/model/system/config"
	configService "robvi/app/modules/sys/service/system/config"
)

// 列表页
func List(c *gin.Context) {
	response2.BuildTpl(c, "system/config/list").WriteTpl()
}

// 列表分页数据
func ListAjax(c *gin.Context) {
	req := new(config2.SelectPageReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response2.ErrorResp(c).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}
	rows := make([]config2.Entity, 0)
	result, page, err := configService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	response2.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 新增页面
func Add(c *gin.Context) {
	response2.BuildTpl(c, "system/config/add").WriteTpl()
}

// 新增页面保存
func AddSave(c *gin.Context) {
	req := new(config2.AddReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}

	if configService.CheckConfigKeyUniqueAll(req.ConfigKey) == "1" {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("参数键名已存在").Log("参数管理", req).WriteJsonExit()
		return
	}

	rid, err := configService.AddSave(req, c)

	if err != nil || rid <= 0 {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).Log("参数管理", req).WriteJsonExit()
		return
	}
	response2.SucessResp(c).SetData(rid).Log("参数管理", req).WriteJsonExit()
}

// 修改页面
func Edit(c *gin.Context) {
	id := gconv.Int64(c.Query("id"))
	if id <= 0 {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	entity, err := configService.SelectRecordById(id)

	if err != nil || entity == nil {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "数据不存在",
		})
		return
	}

	response2.BuildTpl(c, "system/config/edit").WriteTpl(gin.H{
		"config": entity,
	})
}

// 修改页面保存
func EditSave(c *gin.Context) {
	req := new(config2.EditReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}

	if configService.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId) == "1" {
		response2.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg("参数键名已存在").Log("参数管理", req).WriteJsonExit()
		return
	}

	rs, err := configService.EditSave(req, c)

	if err != nil || rs <= 0 {
		response2.ErrorResp(c).SetBtype(model.Buniss_Edit).Log("参数管理", req).WriteJsonExit()
		return
	}
	response2.SucessResp(c).SetBtype(model.Buniss_Edit).Log("参数管理", req).WriteJsonExit()
}

// 删除数据
func Remove(c *gin.Context) {
	req := new(model.RemoveReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}

	rs := configService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response2.SucessResp(c).SetBtype(model.Buniss_Del).Log("参数管理", req).WriteJsonExit()
	} else {
		response2.ErrorResp(c).SetBtype(model.Buniss_Del).Log("参数管理", req).WriteJsonExit()
	}
}

// 导出
func Export(c *gin.Context) {
	req := new(config2.SelectPageReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response2.ErrorResp(c).Log("参数管理", req).WriteJsonExit()
		return
	}
	url, err := configService.Export(req)

	if err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Other).Log("参数管理", req).WriteJsonExit()
		return
	}

	response2.SucessResp(c).SetBtype(model.Buniss_Other).SetMsg(url).WriteJsonExit()
}

// 检查参数键名是否已经存在不包括本参数
func CheckConfigKeyUnique(c *gin.Context) {
	var req *config2.CheckConfigKeyReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := configService.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId)

	c.Writer.WriteString(result)
}

// 检查参数键名是否已经存在
func CheckConfigKeyUniqueAll(c *gin.Context) {
	var req *config2.CheckPostCodeALLReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := configService.CheckConfigKeyUniqueAll(req.ConfigKey)

	c.Writer.WriteString(result)
}
