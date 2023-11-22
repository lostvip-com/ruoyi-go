package sys

import (
	"lostvip.com/web/router"
	"robvi/app/middleware/auth"
	"robvi/app/middleware/token"
	"robvi/app/modules/sys/controller/monitor/logininfor"
	"robvi/app/modules/sys/controller/monitor/online"
	"robvi/app/modules/sys/controller/monitor/operlog"
	"robvi/app/modules/sys/controller/monitor/server"
)

// 加载路由
func init() {
	// 服务监控
	g1 := router.New("/monitor/server", token.TokenMiddleware(), auth.Auth)
	g1.GET("/", "monitor:server:view", server.Server)

	//登陆日志
	g2 := router.New("/monitor/logininfor", token.TokenMiddleware(), auth.Auth)
	g2.GET("/", "monitor:logininfor:view", logininfor.List)
	g2.POST("/list", "monitor:logininfor:list", logininfor.ListAjax)
	g2.POST("/export", "monitor:logininfor:export", logininfor.Export)
	g2.POST("/clean", "monitor:logininfor:remove", logininfor.Clean)
	g2.POST("/remove", "monitor:logininfor:remove", logininfor.Remove)
	g2.POST("/unlock", "monitor:logininfor:unlock", logininfor.Unlock)

	//操作日志
	g3 := router.New("/monitor/operlog", token.TokenMiddleware(), auth.Auth)
	g3.GET("/", "monitor:operlog:view", operlog.List)
	g3.POST("/list", "monitor:operlog:list", operlog.ListAjax)
	g3.POST("/export", "monitor:operlog:export", operlog.Export)
	g3.POST("/remove", "monitor:operlog:export", operlog.Remove)
	g3.POST("/clean", "monitor:operlog:export", operlog.Clean)
	g3.GET("/detail", "monitor:operlog:detail", operlog.Detail)

	//在线用户
	g4 := router.New("/monitor/online", token.TokenMiddleware(), auth.Auth)
	g4.GET("/", "monitor:online:view", online.List)
	g4.POST("/list", "monitor:online:list", online.ListAjax)
	g4.POST("/forceLogout", "monitor:online:forceLogout", online.ForceLogout)
	g4.POST("/batchForceLogout", "monitor:online:batchForceLogout", online.BatchForceLogout)
}
