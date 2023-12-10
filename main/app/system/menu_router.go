package system

import (
	"lostvip.com/web/router"
	"robvi/app/common/middleware/auth"
	"robvi/app/common/middleware/token"
	"robvi/app/system/controller"
)

// 加载路由
func init() {
	// 角色路由
	g1 := router.New("/system/menu", token.TokenMiddleware(), auth.Auth)
	controller := controller.MenuController{}
	g1.GET("/", "system:menu:view", controller.List)
	g1.POST("/list", "system:menu:list", controller.ListAjax)
	g1.GET("/add", "system:menu:add", controller.Add)
	g1.POST("/add", "system:menu:add", controller.AddSave)
	g1.GET("/remove", "system:menu:remove", controller.Remove)
	g1.POST("/remove", "system:menu:remove", controller.Remove)
	g1.GET("/edit", "system:menu:edit", controller.Edit)
	g1.POST("/edit", "system:menu:edit", controller.EditSave)
	g1.GET("/icon", "system:menu:view", controller.Icon)
	g1.GET("/selectMenuTree", "system:menu:view", controller.SelectMenuTree)
	g1.GET("/roleMenuTreeData", "system:menu:view", controller.RoleMenuTreeData)
	g1.GET("/menuTreeData", "system:menu:view", controller.MenuTreeData)
	g1.POST("/checkMenuNameUnique", "system:menu:view", controller.CheckMenuNameUnique)
	g1.POST("/checkMenuNameUniqueAll", "system:menu:view", controller.CheckMenuNameUniqueAll)
}
