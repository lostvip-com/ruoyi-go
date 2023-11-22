package modules

import (
	"lostvip.com/web/router"
	"robvi/app/middleware/token"
	_ "robvi/app/modules/biz"
	_ "robvi/app/modules/demo"
	_ "robvi/app/modules/sys"
	errorc "robvi/app/modules/sys/controller/system/error"
	"robvi/app/modules/sys/controller/system/index"
)

func init() {
	// 加载登陆路由
	g0 := router.New("/")
	g0.GET("/login", "", index.Login)
	g0.POST("/login", "", index.CheckLogin)
	g0.GET("/captchaImage", "", index.CaptchaImage)
	g0.GET("/500", "", errorc.Error)
	g0.GET("/404", "", errorc.NotFound)
	g0.GET("/403", "", errorc.Unauth)
	//下在要检测是否登陆
	g1 := router.New("/", token.TokenMiddleware())
	g1.GET("/", "", index.Index)
	g1.GET("/index", "", index.Index)
	g1.GET("/index_left", "", index.Index)
	g1.GET("/logout", "", index.Logout)

}
