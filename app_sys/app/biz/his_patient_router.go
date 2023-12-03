// ==========================================================================
// LV自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-11-26 16:09:17 +0800 CST
// 生成人：lv
// ==========================================================================
package robvi

import (
	"lostvip.com/web/router"
	"robvi/app/biz/controller"
	"robvi/app/common/middleware/auth"
)

// 加载路由
func init() {
	// 参数路由
	g1 := router.New("/biz/patient", auth.Auth)

	web := controller.HisPatientController{}
	g1.GET("/", "biz:patient:view", web.List)
	g1.POST("/list", "biz:patient:list", web.ListAjax)
	g1.GET("/add", "biz:patient:add", web.Add)
	g1.POST("/add", "biz:patient:add", web.AddSave)
	g1.POST("/remove", "biz:patient:remove", web.Remove)
	g1.GET("/edit", "biz:patient:edit", web.Edit)
	g1.POST("/edit", "biz:patient:edit", web.EditSave)
	g1.POST("/export", "biz:patient:export", web.Export)
}
