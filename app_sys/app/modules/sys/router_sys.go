package sys

import (
	"lostvip.com/web/router"
	"robvi/app/middleware/auth"
	"robvi/app/middleware/token"
	"robvi/app/modules/sys/controller/system/index"
)

func init() {
	// 加载框架路由
	g2 := router.New("/system", token.TokenMiddleware(), auth.Auth)
	g2.GET("/main", "", index.Main)
	g2.GET("/switchSkin", "", index.SwitchSkin)
	g2.GET("/download", "", index.Download)
}
