package sys

import (
	"lostvip.com/web/router"
	"robvi/app/common/middleware/auth"
	"robvi/app/common/middleware/token"
	user2 "robvi/app/sys/controller/system/user"
)

// 加载路由
func init() {
	// 用户管理路由
	g1 := router.New("/system/user", token.TokenMiddleware(), auth.Auth)
	g1.GET("/", "system:user:view", user2.List)
	g1.POST("/list", "system:user:list", user2.ListAjax)
	g1.GET("/add", "system:user:add", user2.Add)
	g1.POST("/add", "system:user:add", user2.AddSave)
	g1.POST("/remove", "system:user:remove", user2.Remove)
	g1.GET("/edit", "system:user:edit", user2.Edit)
	g1.POST("/edit", "system:user:edit", user2.EditSave)
	g1.POST("/export", "system:user:export", user2.Export)
	g1.GET("/resetPwd", "system:user:resetPwd", user2.ResetPwd)
	g1.POST("/resetPwd", "system:user:resetPwd", user2.ResetPwdSave)

	// 个人中心路由
	g2 := router.New("/system/user/profile", token.TokenMiddleware(), auth.Auth)
	g2.GET("/", "", user2.Profile)
	g2.GET("/avatar", "", user2.Avatar)
	g2.GET("/resetPwd", "", user2.EditPwd)
	g2.POST("/update", "", user2.Update)
	g2.POST("/resetSavePwd", "", user2.UpdatePassword)
	g2.POST("/checkEmailUnique", "", user2.CheckEmailUnique)
	g2.POST("/checkPhoneUnique", "", user2.CheckPhoneUnique)
	g2.POST("/checkLoginNameUnique", "", user2.CheckLoginNameUnique)
	g2.POST("/checkEmailUniqueAll", "", user2.CheckEmailUniqueAll)
	g2.POST("/checkPhoneUniqueAll", "", user2.CheckPhoneUniqueAll)
	g2.POST("/checkPassword", "", user2.CheckPassword)
	g2.POST("/updateAvatar", "", user2.UpdateAvatar)
}
