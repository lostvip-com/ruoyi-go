package system

import (
	"lostvip.com/web/router"
	"robvi/app/common/middleware/auth"
	"robvi/app/common/middleware/token"
	"robvi/app/system/controller"
)

// 加载路由
func init() {
	// 岗位路由
	g1 := router.New("/system/post", token.TokenMiddleware(), auth.Auth)
	controller := controller.PostController{}
	g1.GET("/", "system:post:view", controller.List)
	g1.POST("/list", "system:post:list", controller.ListAjax)
	g1.GET("/add", "system:post:add", controller.Add)
	g1.POST("/add", "system:post:add", controller.AddSave)
	g1.POST("/remove", "system:post:remove", controller.Remove)
	g1.GET("/edit", "system:post:edit", controller.Edit)
	g1.POST("/edit", "system:post:edit", controller.EditSave)
	g1.POST("/export", "system:post:export", controller.Export)
	g1.POST("/checkPostCodeUniqueAll", "system:post:list", controller.CheckPostCodeUniqueAll)
	g1.POST("/checkPostCodeUnique", "system:post:list", controller.CheckPostCodeUnique)
	g1.POST("/checkPostNameUniqueAll", "system:post:list", controller.CheckPostNameUniqueAll)
	g1.POST("/checkPostNameUnique", "system:post:list", controller.CheckPostNameUnique)
}
