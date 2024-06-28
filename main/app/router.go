package app

import (
	"github.com/lv_framework/web/router"
	"main/app/common/middleware/auth"
	_ "main/app/demo"
	_ "main/app/mywork"
	_ "main/app/system"
	sysController "main/app/system/controller"
)

func init() {
	// 加载登录路由
	g0 := router.New("/")
	login := sysController.LoginController{}
	g0.GET("/login", "", login.Login)
	g0.POST("/login", "", login.CheckLogin)
	g0.GET("/captchaImage", "", login.CaptchaImage)
	errorc := sysController.ErrorController{}
	g0.GET("/500", "", errorc.Error)
	g0.GET("/404", "", errorc.NotFound)
	g0.GET("/403", "", errorc.Unauth)
	//下在要检测是否登录
	g1 := router.New("/", auth.TokenCheck())
	main := sysController.MainController{}
	g1.GET("/", "", main.Index)
	g1.GET("/index", "", main.Index)
	g1.GET("/index_left", "", main.Index)
	g1.GET("/logout", "", main.Logout)

}
