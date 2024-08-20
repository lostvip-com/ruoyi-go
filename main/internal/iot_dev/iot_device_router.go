// ==========================================================================
// LV自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-07-19 17:09:35 +0800 CST
// 生成人：lv
// ==========================================================================
package iot_dev

import (
	"github.com/lostvip-com/lv_framework/web/router"
	"main/internal/common/middleware/auth"
	"main/internal/iot_dev/controller"
)

// 加载路由
func init() {
	// 参数路由
	g1 := router.New("/iot_dev/device", auth.TokenCheck())

	web := controller.IotDeviceController{}
	g1.GET("/preListTree", "iot_dev:device:view", web.PreListTree) //设备
	// preTabs 二级页，设备信息
	g1.GET("/preListDevice", "iot_dev:device:view", web.PreListDevice)   //设备
	g1.GET("/preListGateway", "iot_dev:device:view", web.PreListGateway) //Tabs网关列表
	// preTabs3 三级页
	g1.GET("/edit", "iot_dev:device:edit", web.PreEdit)                 //设备编辑
	g1.GET("/add", "iot_dev:device:add", web.PreAdd)                    //设备添加
	g1.GET("/popSubDevice", "iot_dev:device:view", web.PrePopSubDevice) //选择子设备

	g1.GET("/preDeviceDetailTabs", "iot_dev:product:view", web.PreDeviceDetailTabs)
	// 监控数据
	//g1.GET("/preTabListRealtime", "iot_dev:device:view", web.PreTabListRealtime) //Tabs设备列表

	//api
	g1.GET("/list", "iot_dev:device:list", web.ListAjax)
	g1.POST("/list", "iot_dev:device:list", web.ListAjax)
	g1.POST("/add", "iot_dev:device:add", web.AddSave)
	g1.POST("/remove", "iot_dev:device:remove", web.Remove)
	g1.POST("/edit", "iot_dev:device:edit", web.EditSave)
	g1.POST("/export", "iot_dev:device:export", web.Export)
	//
	g1.POST("/bindGateway", "iot_dev:device:edit", web.BindGateway)

}
