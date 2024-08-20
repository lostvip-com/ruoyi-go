// ==========================================================================
// LV自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-07-19 17:16:13 +0800 CST
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
	g1 := router.New("/iot_dev/product", auth.TokenCheck())
	//产品
	web := controller.IotProductController{}
	g1.GET("/prePrdDetailTabs", "iot_dev:product:view", web.PrePrdDetailTabs)
	g1.GET("/", "iot_dev:product:view", web.PreList)
	g1.GET("/edit", "iot_dev:product:edit", web.PreEdit)
	g1.GET("/add", "iot_dev:product:add", web.PreAdd)

	// api
	g1.GET("/list", "iot_dev:product:list", web.ListAjax)
	g1.POST("/list", "iot_dev:product:list", web.ListAjax)
	g1.POST("/add", "iot_dev:product:add", web.AddSave)
	g1.POST("/remove", "iot_dev:product:remove", web.Remove)
	g1.POST("/edit", "iot_dev:product:edit", web.EditSave)
	g1.POST("/export", "iot_dev:product:export", web.Export)
	// 产品属性 页面
	g11 := router.New("/iot_dev/property", auth.TokenCheck())
	webProp := controller.IotPrdPropertyController{}
	g11.GET("/", "iot_dev:product:view", webProp.List)
	g11.GET("/add", "iot_dev:product:add", webProp.Add)
	g11.GET("/edit", "iot_dev:product:edit", webProp.Edit)
	// 产品属性 api
	g11.POST("/list", "iot_dev:product:list", webProp.ListAjax)
	g11.POST("/add", "iot_dev:product:add", webProp.AddSave)
	g11.POST("/remove", "iot_dev:product:remove", webProp.Remove)
	g11.POST("/edit", "iot_dev:product:edit", webProp.EditSave)
	g11.POST("/export", "iot_dev:product:export", webProp.Export)
	//
	g12 := router.New("/iot_dev/action", auth.TokenCheck())
	webAction := controller.IotPrdActionController{}
	g12.GET("/", "iot_dev:product:view", webAction.List)
	g12.GET("/add", "iot_dev:product:add", webAction.Add)
	g12.GET("/edit", "iot_dev:product:edit", webAction.Edit)
	// api
	g12.POST("/list", "iot_dev:product:list", webAction.ListAjax)
	g12.POST("/add", "iot_dev:product:add", webAction.AddSave)
	g12.POST("/remove", "iot_dev:product:remove", webAction.Remove)
	g12.POST("/edit", "iot_dev:product:edit", webAction.EditSave)
	g12.POST("/export", "iot_dev:product:export", webAction.Export)
	// event
	g13 := router.New("/iot_dev/event", auth.TokenCheck())
	webEvent := controller.IotPrdEventController{}
	g13.GET("/", "iot_dev:product:view", webEvent.List)
	g13.GET("/add", "iot_dev:product:add", webEvent.Add)
	g13.GET("/edit", "iot_dev:product:edit", webEvent.Edit)
	// api
	g13.POST("/list", "iot_dev:product:list", webEvent.ListAjax)
	g13.POST("/add", "iot_dev:product:add", webEvent.AddSave)
	g13.POST("/remove", "iot_dev:product:remove", webEvent.Remove)
	g13.POST("/edit", "iot_dev:product:edit", webEvent.EditSave)
	g13.POST("/export", "iot_dev:product:export", webEvent.Export)
}
