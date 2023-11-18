package sys

import (
	"lostvip.com/router"
	"robvi/app/middleware/auth"
	"robvi/app/middleware/token"
	"robvi/app/modules/sys/controller/system/user"
)

// 加载路由
func init() {
	// 用户管理路由
	g1 := router.New("/system/user", token.TokenMiddleware(), auth.Auth)
	g1.GET("/", "system:user:view", user.List)
	g1.POST("/list", "system:user:list", user.ListAjax)
	g1.GET("/add", "system:user:add", user.Add)
	g1.POST("/add", "system:user:add", user.AddSave)
	g1.POST("/remove", "system:user:remove", user.Remove)
	g1.GET("/edit", "system:user:edit", user.Edit)
	g1.POST("/edit", "system:user:edit", user.EditSave)
	g1.POST("/export", "system:user:export", user.Export)
	g1.GET("/resetPwd", "system:user:resetPwd", user.ResetPwd)
	g1.POST("/resetPwd", "system:user:resetPwd", user.ResetPwdSave)

	// 个人中心路由
	g2 := router.New("/system/user/profile", token.TokenMiddleware(), auth.Auth)
	g2.GET("/", "", user.Profile)
	g2.GET("/avatar", "", user.Avatar)
	g2.GET("/resetPwd", "", user.EditPwd)
	g2.POST("/update", "", user.Update)
	g2.POST("/resetSavePwd", "", user.UpdatePassword)
	g2.POST("/checkEmailUnique", "", user.CheckEmailUnique)
	g2.POST("/checkPhoneUnique", "", user.CheckPhoneUnique)
	g2.POST("/checkLoginNameUnique", "", user.CheckLoginNameUnique)
	g2.POST("/checkEmailUniqueAll", "", user.CheckEmailUniqueAll)
	g2.POST("/checkPhoneUniqueAll", "", user.CheckPhoneUniqueAll)
	g2.POST("/checkPassword", "", user.CheckPassword)
	g2.POST("/updateAvatar", "", user.UpdateAvatar)
}
