// ==========================================================================
// LV自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-11-26 16:27:18 +0800 CST
// 生成人：lv
// ==========================================================================
package robvi

import (
        "lostvip.com/web/router"
        "robvi/app/middleware/auth"
        "robvi/app/modules/biz/controller"
)

//加载路由
func init() {
	// 参数路由
	g1 := router.New( "/biz/task", auth.Auth)

	web := controller.DpcTaskController{}
	g1.GET("/", "biz:task:view", web.List)
	g1.POST("/list", "biz:task:list", web.ListAjax)
	g1.GET("/add", "biz:task:add", web.Add)
	g1.POST("/add", "biz:task:add", web.AddSave)
	g1.POST("/remove", "biz:task:remove", web.Remove)
	g1.GET("/edit", "biz:task:edit", web.Edit)
	g1.POST("/edit", "biz:task:edit",web.EditSave)
	g1.POST("/export", "biz:task:export", web.Export)
}
