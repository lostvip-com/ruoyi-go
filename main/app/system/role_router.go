package system

import (
	"lostvip.com/web/router"
	"robvi/app/common/middleware/auth"
	"robvi/app/common/middleware/token"
	"robvi/app/system/controller/system/role"
)

// 加载路由
func init() {
	// 角色路由
	g1 := router.New("/system/role", token.TokenMiddleware(), auth.Auth)

	g1.GET("/", "system:post:view", role.List)
	g1.POST("/list", "system:post:list", role.ListAjax)
	g1.GET("/add", "system:post:add", role.Add)
	g1.POST("/add", "system:post:add", role.AddSave)
	g1.POST("/remove", "system:post:remove", role.Remove)
	g1.GET("/edit", "system:post:edit", role.Edit)
	g1.POST("/edit", "system:post:edit", role.EditSave)
	g1.POST("/export", "system:post:export", role.Export)
	g1.POST("/checkRoleKeyUnique", "system:post:view", role.CheckRoleKeyUnique)
	g1.POST("/checkRoleNameUniqueAll", "system:post:view", role.CheckRoleNameUniqueAll)
	g1.POST("/checkRoleNameUnique", "system:post:view", role.CheckRoleNameUnique)
	g1.POST("/checkRoleKeyUniqueAll", "system:post:view", role.CheckRoleKeyUniqueAll)
	g1.GET("/authDataScope", "system:post:view", role.AuthDataScope)
	g1.POST("/authDataScope", "system:post:view", role.AuthDataScopeSave)
	g1.GET("/authUser", "system:post:view", role.AuthUser)
	g1.POST("/allocatedList", "system:post:view", role.AllocatedList)
	g1.GET("/selectUser", "system:post:view", role.SelectUser)
	g1.POST("/unallocatedList", "system:post:view", role.UnallocatedList)
	g1.POST("/selectAll", "system:post:view", role.SelectAll)
	g1.POST("/cancel", "system:post:view", role.Cancel)
	g1.POST("/cancelAll", "system:post:view", role.CancelAll)
}
