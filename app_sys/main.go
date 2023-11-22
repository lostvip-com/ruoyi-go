package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	db "lostvip.com/db"
	"lostvip.com/web/server"
	"robvi/app/global"
	_ "robvi/app/modules"
	"robvi/app/modules/sys/model/module/tenant"
	"robvi/app/modules/sys/model/system/dept"
	"robvi/app/modules/sys/model/system/post"
	"robvi/app/modules/sys/model/system/user"
)

// @title LV 自动生成API文档
// @version 1.0
// @description 生成文档请在调试模式下进行<a href="/tool/swagger?a=r">重新生成文档</a>

// @host localhost
// @BasePath /api
func main() {
	//str, _ := gmd5.Encrypt("123456")
	//println(">>>>>>>>>>>>>: " + str)
	cfg := global.GetConfigInstance()
	if cfg.IsDebug() {
		gin.SetMode("debug")
		db.Instance().Engine().ShowSQL(true)
	}
	db.Instance().Engine().Sync2(dept.SysDept{}, tenant.SysTenant{}, user.SysUser{}, post.SysPost{})
	//后台服务状态
	httpSvr := server.New("0.0.0.0:" + cast.ToString(cfg.GetServerPort()))
	httpSvr.Start()
	fmt.Println("------------exit----------------------")

}
