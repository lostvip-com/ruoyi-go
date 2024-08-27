package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/web/server"
	"github.com/spf13/cast"
	_ "main/internal"
	"main/internal/common/myconf"
	"os"
	"os/signal"
	"syscall"
)

var httpSvr *server.MyServer

// @title LV 自动生成API文档
// @version 1.0
// @description 生成文档请在调试模式下进行<a href="/tool/swagger?a=r">重新生成文档</a>
// @host localhost
// @BasePath /
func main() {
	cfg := myconf.GetConfigInstance()
	lv_log.InitLog("logru.log")
	if lv_global.IsDebug {
		gin.SetMode("debug")
	}
	httpSvr = server.New("0.0.0.0:" + cast.ToString(cfg.GetServerPort()))
	go httpSvr.Start()
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
