package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	db "lostvip.com/db"
	"lostvip.com/logme"
	"lostvip.com/web/server"
	_ "robvi/app"
	"robvi/app/common/global"
	"robvi/app/mywork/model"
)

// @title LV 自动生成API文档
// @version 1.0
// @description 生成文档请在调试模式下进行<a href="/tool/swagger?a=r">重新生成文档</a>

// @host localhost
// @BasePath /api
func main() {
	cfg := global.GetConfigInstance()
	logme.InitLog("logru.log")
	if cfg.IsDebug() {
		// Only log the warning severity or above.
		gin.SetMode("debug")
		db.Instance().Engine().ShowSQL(true)
	}
	db.Instance().Engine().Sync2(model.DpcTask{})
	//后台服务状态
	httpSvr := server.New("0.0.0.0:" + cast.ToString(cfg.GetServerPort()))
	httpSvr.Start()
	fmt.Println("------------exit----------------------")

}
