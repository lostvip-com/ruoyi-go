package controller

import (
	"common/common_vo"
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"github.com/spf13/cast"
	dao2 "system/dao"
	"system/service"
)

type ConfigController struct{}

// 列表页
func (w *ConfigController) List(c *gin.Context) {
	util.BuildTpl(c, "system/config/list").WriteTpl()
}

// 列表分页数据
func (w *ConfigController) ListAjax(c *gin.Context) {
	req := new(common_vo.SelectConfigPageReq)
	var configService service.ConfigService
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		util.ErrorResp(c).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}
	result, total, err := configService.SelectListByPage(req)
	lv_err.HasErrAndPanic(err)
	util.BuildTable(c, total, result).WriteJsonExit()
}

// 新增页面
func (w *ConfigController) Add(c *gin.Context) {
	util.BuildTpl(c, "system/config/add").WriteTpl()
}

// 新增页面保存
func (w *ConfigController) AddSave(c *gin.Context) {
	req := new(common_vo.AddConfigReq)
	//获取参数
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var config dao2.ConfigDao
	count, err := config.Count(req.ConfigKey)
	lv_err.HasErrAndPanic(err)
	if count > 0 {
		util.Err(c, "此编号已经存在！请更换编号")
	} else {
		util.Success(c, "", "操作成功！")
	}
}

// 修改页面
func (w *ConfigController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	var configService service.ConfigService
	entity, err := configService.SelectRecordById(id)
	lv_err.HasErrAndPanic(err)
	util.WriteTpl(c, "system/config/edit", gin.H{
		"config": entity,
	})
}

// 修改页面保存
func (w *ConfigController) EditSave(c *gin.Context) {
	req := new(common_vo.EditConfigReq)
	//获取参数
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var configService service.ConfigService
	configService.EditSave(req, c)
	util.Success(c, "", "操作成功！")
}

// 删除数据
func (w *ConfigController) Remove(c *gin.Context) {
	req := new(lv_dto.IdsReq)
	//获取参数
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var configService service.ConfigService
	configService.DeleteRecordByIds(req.Ids)
	util.Success(c, "", "删除成功！")
}

// 导出
func (w *ConfigController) Export(c *gin.Context) {
	req := new(common_vo.SelectConfigPageReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var configService service.ConfigService
	url, err := configService.Export(req)
	lv_err.HasErrAndPanic(err)

	util.Success(c, url, "success!")
}

// CheckConfigKeyUnique 检查参数键名是否已经存在不包括本参数,使用jquery validator 默认1 存在，0 不存在
func (w *ConfigController) CheckConfigKeyUnique(c *gin.Context) {
	var req *common_vo.CheckConfigKeyReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}
	var configService service.ConfigService
	num, err := configService.CountKey(req.ConfigKey)
	lv_err.HasErrAndPanic(err)
	c.Writer.WriteString(cast.ToString(num))
}
