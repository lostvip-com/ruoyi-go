package main

import (
	"common/models"
	"common/myconf"
	_ "demo"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/lv_log/lv_log_impl"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/web/server"
	"os"
	"os/signal"
	"syscall"
	_ "system"
)

var httpSvr *server.MyHttpServer

// @title LV 自动生成API文档
// @version 1.0
// @description 生成文档请在调试模式下进行<a href="/tool/swagger?a=r">重新生成文档</a>
// @host localhost
// @BasePath /
func main() {
	fmt.Println(lv_conv.ToJsonStr(new(models.BaseModel)))

	cfg := myconf.GetConfigInstance()
	log := lv_log_impl.InitLog(cfg.GetAppName() + ".log")
	lv_log.Log = log
	if lv_global.IsDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	go server.NewHttpServer().ListenAndServe()
	//监听信号
	catchSignal()
}

// 捕捉信号
func catchSignal() {
	lv_log.Info("⛲ ⛲ ⛲ ⛲ ⛲ ⛲  running!  ⛲ ⛲ ⛲ ⛲ ⛲ ⛲ ")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	for {
		s := <-c
		lv_log.Info("⛳  收到信号:", s)
		switch s {
		case syscall.SIGHUP:
			lv_log.Info("收到终端断开信号, 忽略")
		case syscall.SIGINT, syscall.SIGTERM:
			shutdown()
		}
	}
}

// 应用退出
func shutdown() {
	defer func() {
		lv_log.Info("⛔️ 已经退出应用!")
		os.Exit(0)
	}()
	lv_log.Info("⌛  即将停止服务!")
	httpSvr.ShutDown()
	lv_log.Info("❌  已经停止服务!")
}
