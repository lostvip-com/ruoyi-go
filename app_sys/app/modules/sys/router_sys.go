package sys

import (
	"lostvip.com/router"
	"robvi/app/middleware/auth"
	"robvi/app/middleware/jwt"
	"robvi/app/modules/sys/controller/system/index"
)

func init() {
	// 加载框架路由
	g2 := router.New("/system", jwt.JWTAuthMiddleware(), auth.Auth)
	g2.GET("/main", "", index.Main)
	g2.GET("/switchSkin", "", index.SwitchSkin)
	g2.GET("/download", "", index.Download)
}
