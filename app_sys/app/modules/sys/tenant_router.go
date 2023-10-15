// ==========================================================================
// LV自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-06-29 22:21:21 +0800 CST
// 生成路径: app/controller/modules/tenant_router.go
// 生成人：robnote
// ==========================================================================
package sys

import (
	"lostvip.com/router"
	"robvi/app/middleware/auth"
	"robvi/app/middleware/jwt"
	"robvi/app/modules/sys/controller"
)

// 加载路由
func init() {
	// 参数路由
	g1 := router.New("modules/tenant", jwt.JWTAuthMiddleware(), auth.Auth)
	controller := controller.TenantController{}
	g1.GET("/", "modules:tenant:view", controller.List)
	g1.POST("/list", "modules:tenant:list", controller.ListAjax)
	g1.GET("/add", "modules:tenant:add", controller.Add)
	g1.POST("/add", "modules:tenant:add", controller.AddSave)
	g1.POST("/remove", "modules:tenant:remove", controller.Remove)
	g1.GET("/edit", "modules:tenant:edit", controller.Edit)
	g1.POST("/edit", "modules:tenant:edit", controller.EditSave)
	g1.POST("/export", "modules:tenant:export", controller.Export)
}
