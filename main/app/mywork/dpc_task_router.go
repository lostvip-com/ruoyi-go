// ==========================================================================
// LV自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-12-24 14:57:01 +0800 CST
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
	g1 := router.New( "/mywork/task", auth.Auth)

	web := controller.DpcTaskController{}
	g1.GET("/", "mywork:task:view", web.List)
	g1.POST("/list", "mywork:task:list", web.ListAjax)
	g1.GET("/add", "mywork:task:add", web.Add)
	g1.POST("/add", "mywork:task:add", web.AddSave)
	g1.POST("/remove", "mywork:task:remove", web.Remove)
	g1.GET("/edit", "mywork:task:edit", web.Edit)
	g1.POST("/edit", "mywork:task:edit",web.EditSave)
	g1.POST("/export", "mywork:task:export", web.Export)
}
