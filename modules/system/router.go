package system

import (
	"common/middleware/auth"
	"github.com/lostvip-com/lv_framework/web/router"
	"system/controller"
)

func init() {
	// 加载登录路由
	g0 := router.New("/")
	login := controller.LoginController{}
	g0.GET("/login", "", login.Login)
	g0.POST("/login", "", login.CheckLogin)
	g0.GET("/captchaImage", "", login.CaptchaImage)
	errorc := controller.ErrorController{}
	g0.GET("/500", "", errorc.Error)
	g0.GET("/404", "", errorc.NotFound)
	g0.GET("/403", "", errorc.Unauth)
	//下在要检测是否登录
	g1 := router.New("/", auth.TokenCheck())
	index := controller.IndexController{}
	g1.GET("/", "", index.Index)
	g1.GET("/index", "", index.Index)
	g1.GET("/index_left", "", index.Index)
	g1.GET("/logout", "", index.Logout)

}
