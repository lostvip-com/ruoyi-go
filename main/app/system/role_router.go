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
	g1 := router.New("/system/role", token.TokenMiddleware(), auth.Auth)
	controller := controller.RoleController{}
	g1.GET("/", "system:post:view", controller.List)
	g1.POST("/list", "system:post:list", controller.ListAjax)
	g1.GET("/add", "system:post:add", controller.Add)
	g1.POST("/add", "system:post:add", controller.AddSave)
	g1.POST("/remove", "system:post:remove", controller.Remove)
	g1.GET("/edit", "system:post:edit", controller.Edit)
	g1.POST("/edit", "system:post:edit", controller.EditSave)
	g1.POST("/export", "system:post:export", controller.Export)
	g1.POST("/checkRoleKeyUnique", "system:post:view", controller.CheckRoleKeyUnique)
	g1.POST("/checkRoleNameUniqueAll", "system:post:view", controller.CheckRoleNameUniqueAll)
	g1.POST("/checkRoleNameUnique", "system:post:view", controller.CheckRoleNameUnique)
	g1.POST("/checkRoleKeyUniqueAll", "system:post:view", controller.CheckRoleKeyUniqueAll)
	g1.GET("/authDataScope", "system:post:view", controller.AuthDataScope)
	g1.POST("/authDataScope", "system:post:view", controller.AuthDataScopeSave)
	g1.GET("/authUser", "system:post:view", controller.AuthUser)
	g1.POST("/allocatedList", "system:post:view", controller.AllocatedList)
	g1.GET("/selectUser", "system:post:view", controller.SelectUser)
	g1.POST("/unallocatedList", "system:post:view", controller.UnallocatedList)
	g1.POST("/selectAll", "system:post:view", controller.SelectAll)
	g1.POST("/cancel", "system:post:view", controller.Cancel)
	g1.POST("/cancelAll", "system:post:view", controller.CancelAll)
}
