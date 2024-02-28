// ==========================================================================
// LV自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-02-28 14:21:50 +0800 CST
// 生成人：lv
// ==========================================================================
package mywork

import (
        "lostvip.com/web/router"
        "main/app/common/middleware/auth"
        "main/app/mywork/controller"
)

//加载路由
func init() {
	// 参数路由
	g1 := router.New( "/mywork/params", auth.TokenCheck())

	web := controller.AppGenParamsController{}
	g1.GET("/", "mywork:params:view", web.List)
	g1.POST("/list", "mywork:params:list", web.ListAjax)
	g1.GET("/add", "mywork:params:add", web.Add)
	g1.POST("/add", "mywork:params:add", web.AddSave)
	g1.POST("/remove", "mywork:params:remove", web.Remove)
	g1.GET("/edit", "mywork:params:edit", web.Edit)
	g1.POST("/edit", "mywork:params:edit",web.EditSave)
	g1.POST("/export", "mywork:params:export", web.Export)
}
