// ==========================================================================
// LV自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-12-24 16:29:05 +0800 CST
// 生成人：lv
// ==========================================================================
package robvi

import (
        "lostvip.com/web/router"
        "robvi/app/common/middleware/auth"
        "robvi/app/mywork/controller"
)

//加载路由
func init() {
	// 参数路由
	g1 := router.New( "/mywork/item", auth.Auth)

	web := controller.DpcTaskItemController{}
	g1.GET("/", "mywork:item:view", web.List)
	g1.POST("/list", "mywork:item:list", web.ListAjax)
	g1.GET("/add", "mywork:item:add", web.Add)
	g1.POST("/add", "mywork:item:add", web.AddSave)
	g1.POST("/remove", "mywork:item:remove", web.Remove)
	g1.GET("/edit", "mywork:item:edit", web.Edit)
	g1.POST("/edit", "mywork:item:edit",web.EditSave)
	g1.POST("/export", "mywork:item:export", web.Export)
}
