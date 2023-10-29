package sys

import (
	"lostvip.com/router"
	"robvi/app/middleware/auth"
	"robvi/app/middleware/jwt"
	"robvi/app/modules/sys/controller/system/menu"
)

// 加载路由
func init() {
	// 角色路由
	g1 := router.New("/system/menu", jwt.JWTAuthMiddleware(), auth.Auth)
	g1.GET("/", "system:menu:view", menu.List)
	g1.POST("/list", "system:menu:list", menu.ListAjax)
	g1.GET("/add", "system:menu:add", menu.Add)
	g1.POST("/add", "system:menu:add", menu.AddSave)
	g1.GET("/remove", "system:menu:remove", menu.Remove)
	g1.POST("/remove", "system:menu:remove", menu.Remove)
	g1.GET("/edit", "system:menu:edit", menu.Edit)
	g1.POST("/edit", "system:menu:edit", menu.EditSave)
	g1.GET("/icon", "system:menu:view", menu.Icon)
	g1.GET("/selectMenuTree", "system:menu:view", menu.SelectMenuTree)
	g1.GET("/roleMenuTreeData", "system:menu:view", menu.RoleMenuTreeData)
	g1.GET("/menuTreeData", "system:menu:view", menu.MenuTreeData)
	g1.POST("/checkMenuNameUnique", "system:menu:view", menu.CheckMenuNameUnique)
	g1.POST("/checkMenuNameUniqueAll", "system:menu:view", menu.CheckMenuNameUniqueAll)
}
