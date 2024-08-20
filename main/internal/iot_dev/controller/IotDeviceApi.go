//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2024-07-19 17:09:35 +0800 CST
// author：lv
// ==========================================================================
package controller

import (
	service2 "common/service"
	"common/session"
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"github.com/spf13/cast"
	"main/internal/iot_dev/service"
	"main/internal/iot_dev/vo"
	sysService "main/internal/system/service"
	"strings"
)

type IotDeviceController struct{}

//	========================================================================
//
// api
// =========================================================================

// ListAjax 新增页面保存
func (w IotDeviceController) ListAjax(c *gin.Context) {
	req := new(vo.PageIotDeviceReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	if req.NodeType != "" {
		req.NodeTypeArr = strings.Split(req.NodeType, ",")
	}
	if req.DeptId == 0 {
		req.DeptId = session.GetLoginInfo(c).DeptId
	}
	var svc service.IotDeviceService
	result, total, _ := svc.ListByPage(req)
	util.SucessPage(c, result, total)
}

// AddSave 新增页面保存
func (w IotDeviceController) AddSave(c *gin.Context) {
	req := new(vo.AddIotDeviceReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	//
	strPid := c.Query("productId")
	if strPid != "" {
		deviceService := service.IotProductService{}
		entity, err := deviceService.FindById(lv_conv.Int64(strPid))
		lv_err.HasErrAndPanic(err)
		req.NodeType = entity.NodeType
	}

	var svc service.IotDeviceService
	var userService sysService.UserService
	user := userService.GetProfile(c)
	req.CreateBy = user.LoginName
	id, err := svc.AddSave(req)
	lv_err.HasErrAndPanic(err)
	util.SucessData(c, id)
}

// EditSave 修改页面保存
func (w IotDeviceController) EditSave(c *gin.Context) {
	req := new(vo.EditIotDeviceReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.IotDeviceService
	var userService sysService.UserService
	user := userService.GetProfile(c)
	req.UpdateBy = user.LoginName
	err = svc.EditSave(req)
	lv_err.HasErrAndPanic(err)
	util.Success(c, nil, "success")
}

// Remove 删除数据
func (w IotDeviceController) Remove(c *gin.Context) {
	req := new(lv_dto.IdsReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.IotDeviceService
	rs := svc.DeleteByIds(req.Ids)
	util.SuccessData(c, rs)
}

// 导出
func (w IotDeviceController) Export(c *gin.Context) {
	req := new(vo.PageIotDeviceReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.IotDeviceService
	url, err := svc.ExportAll(req)
	lv_err.HasErrAndPanic(err)
	util.SucessDataMsg(c, url, url)
}

// EditSave 修改页面保存
func (w IotDeviceController) BindGateway(c *gin.Context) {
	//gatewayId+"&deviceId="+deviceId
	gatewayId := c.Query("gatewayId")
	deviceId := c.Query("deviceId")
	bind := c.Query("bind")
	var svc service.IotDeviceService
	err := svc.BindGateway(bind, cast.ToInt64(gatewayId), cast.ToInt64(deviceId))
	lv_err.HasErrAndPanic(err)
	util.Success(c, nil, "success")
}

func (w IotDeviceController) PreDeviceDetailTabs(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	svc := service.IotDeviceService{}
	entity, err := svc.FindById(id)
	lv_err.HasErrAndPanic(err)
	svcPrd := service.IotProductService{}
	prd, err := svcPrd.FindById(entity.ProductId)
	if err == nil {
		entity.ProductName = prd.Name
	}
	cacheSvc := service2.CacheService{}
	label := cacheSvc.GetDictLabel("iot_node_type", entity.NodeType)
	entity.NodeTypeName = label
	util.WriteTpl(c, "iot_dev/device/device_detail_tabs.html", gin.H{
		"entity": entity,
	})
}
