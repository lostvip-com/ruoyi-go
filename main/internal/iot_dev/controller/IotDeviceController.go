//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2024-07-19 17:09:35 +0800 CST
// author：lv
// ==========================================================================
package controller

import (
	"common/session"
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_logic"
	"main/internal/iot_dev/service"
)

// PreListTree 设备菜单入口，1级页面，共3级
func (w IotDeviceController) PreListTree(c *gin.Context) {
	util.BuildTpl(c, "iot_dev/device/index_tree_tabs.html").WriteTpl()
}

// PreListDevice 设备
func (w IotDeviceController) PreListDevice(c *gin.Context) {
	strPid := c.Query("productId")
	var productId int64
	if strPid != "" {
		productId = lv_conv.Int64(strPid)
	}
	productName := ""
	//产品信息
	deviceService := service.IotProductService{}
	prd, _ := deviceService.FindById(productId)
	productName = prd.Name
	//部门信息
	deptId := c.Query("deptId")
	nodeType := c.Query("nodeType")
	data := gin.H{
		"productId":   productId,
		"productName": productName,
		"nodeType":    nodeType,
	}
	if lv_logic.IsEmpty(deptId) {
		data["deptId"] = session.GetLoginInfo(c).DeptId
	} else {
		data["deptId"] = deptId
	}
	util.BuildTpl(c, "iot_dev/device/list_device.html").WriteTpl(data)
}

// PreListGateway 网关
func (w IotDeviceController) PreListGateway(c *gin.Context) {
	deptId := c.Query("deptId")
	nodeType := c.Query("nodeType")
	util.BuildTpl(c, "iot_dev/device/list_gateway.html").WriteTpl(gin.H{
		"deptId":   deptId,
		"nodeType": nodeType,
	})
}

// PrePopSubDevice 选择网关子设备
func (w IotDeviceController) PrePopSubDevice(c *gin.Context) {
	deptId := c.Query("deptId")
	gatewayId := c.Query("gatewayId")
	util.BuildTpl(c, "iot_dev/device/pop_sub_device.html").WriteTpl(gin.H{
		"deptId":    deptId,
		"gatewayId": gatewayId,
	})
}

// Add 新增页
func (w IotDeviceController) PreAdd(c *gin.Context) {
	strPid := c.Query("productId")
	deptId := c.Query("deptId")
	var productId int64
	productName := ""
	nodeType := ""
	if strPid != "" {
		productId = lv_conv.Int64(strPid)
		deviceService := service.IotProductService{}
		prd, _ := deviceService.FindById(productId)
		productName = prd.Name
		nodeType = prd.NodeType
	}
	util.BuildTpl(c, "iot_dev/device/add_device.html").WriteTpl(gin.H{
		"productId":   productId,
		"productName": productName,
		"nodeType":    nodeType,
		"deptId":      deptId,
	})
}

// Edit 修改页
func (w IotDeviceController) PreEdit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	deviceService := service.IotDeviceService{}
	entity, err := deviceService.FindById(id)
	if err != nil || entity == nil {
		util.ErrorTpl(c).WriteTpl(gin.H{"desc": "数据不存在"})
		return
	}
	util.BuildTpl(c, "iot_dev/device/edit_device.html").WriteTpl(gin.H{
		"device": entity,
	})
}
