package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/db"
	"github.com/lv_framework/logme"
	"github.com/lv_framework/utils/lv_err"
	"github.com/lv_framework/web/server"
	"github.com/spf13/cast"
	_ "main/app"
	"main/app/common/global"
	my "main/app/mywork/model"
	"main/app/system/model/monitor/online"
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
	cfg := global.GetConfigInstance()
	logme.InitLog("logru.log")
	if cfg.IsDebug() {
		gin.SetMode("debug")
	}
	//自动建表
	err := db.GetMasterGorm().AutoMigrate(my.DpcTask{}, online.UserOnline{})
	lv_err.HasErrAndPanic(err)
	httpSvr = server.New("0.0.0.0:" + cast.ToString(cfg.GetServerPort()))
	go httpSvr.Start()
	//监听信号
	catchSignal()
}

// 捕捉信号
func catchSignal() {
	logme.Info("⛲ ⛲ ⛲ ⛲ ⛲ ⛲  running!  ⛲ ⛲ ⛲ ⛲ ⛲ ⛲ ")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	for {
		s := <-c
		logme.Info("⛳  收到信号:", s)
		switch s {
		case syscall.SIGHUP:
			logme.Info("收到终端断开信号, 忽略")
		case syscall.SIGINT, syscall.SIGTERM:
			shutdown()
		}
	}
}

// 应用退出
func shutdown() {
	defer func() {
		logme.Info("⛔️ 已经退出应用!")
		os.Exit(0)
	}()
	logme.Info("⌛  即将停止服务!")
	httpSvr.ShutDown()
	logme.Info("❌  已经停止服务!")
}
