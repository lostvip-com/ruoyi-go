package lv_log

import (
	"fmt"
	"github.com/lostvip-com/lv_framework/lv_conf"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/utils/lv_file"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var log *logrus.Logger

func GetLog() *logrus.Logger {
	if log == nil {
		InitLog("boot.log")
	}
	return log
}

func InitLog(fileName string) {
	if lv_global.Config() == nil {
		cfg := new(lv_conf.ConfigDefault)
		lv_global.RegisterCfg(cfg)
	}
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pwd = pwd + "/logs"
	err = lv_file.PathCreateIfNotExist(pwd)
	if err != nil {
		panic(err)
	}
	if log == nil {
		log = logrus.New()
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	if lv_global.Config().IsDebug() == "true" {
		fmt.Println("============ debug模式，输出日志到控制台 ============")
		lv_global.IsDebug = true
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
		lv_global.IsDebug = false
		fmt.Println("============ release模式，输出日志json到文件============")
	}
	maxSize := lv_global.Config().GetValueStr("go.log.max-size")
	maxBackups := lv_global.Config().GetValueStr("go.log.max-backups")
	maxAge := lv_global.Config().GetValueStr("go.log.max-age")
	if maxSize == "" {
		maxSize = "10"
	}
	if maxBackups == "" {
		maxBackups = "5"
	}
	if maxAge == "" {
		maxAge = "30"
	}
	fileLog := &lumberjack.Logger{
		Filename:   pwd + "/" + fileName,
		MaxSize:    cast.ToInt(maxSize),    // 日志文件最大 size, 单位是 MB
		MaxBackups: cast.ToInt(maxBackups), // 最大过期日志保留的个数
		MaxAge:     cast.ToInt(maxAge),     //保留过期文件的最大时间间隔,单位是天
		Compress:   true,                   // disabled by default,是否需要压缩滚动日志, 使用的 gzip 压缩
	}
	var writers []io.Writer
	output := lv_global.Config().GetValueStr("go.log.output")
	if output == "stdout" { //只写控制台
		writers = append(writers, os.Stdout)
	} else if output == "file" { // 只写文件
		writers = append(writers, fileLog)
	} else { //同时写文件和屏幕
		writers = append(writers, os.Stdout, fileLog)
	}
	multiWriter := io.MultiWriter(writers...)
	log.SetOutput(multiWriter)
	log.Info("初始化测试Log.Info 完成 ！！stdout:", output)
}

func Error(args ...interface{}) {
	GetLog().Error(args)
}

func ErrorTraceId(traceId any, args ...interface{}) {
	GetLog().WithFields(logrus.Fields{
		"traceId": traceId,
	}).Error(args)
}
func Fatal(args ...interface{}) {
	GetLog().Fatal(args)
}
func FatalTraceId(traceId any, args ...interface{}) {
	GetLog().WithFields(logrus.Fields{
		"traceId": traceId,
	}).Fatal(args)
}
func Warn(args ...interface{}) {
	GetLog().Warn(args)
}
func WarnTraceId(traceId any, args ...interface{}) {
	GetLog().WithFields(logrus.Fields{
		"traceId": traceId,
	}).Warn(args)
}
func Info(args ...interface{}) {
	GetLog().Info(args)
}
func InfoTraceId(traceId any, args ...interface{}) {
	GetLog().WithFields(logrus.Fields{
		"traceId": traceId,
	}).Info(args)
}
func Debug(args ...interface{}) {
	GetLog().Debug(args...)
}
func DebugTraceId(traceId any, args ...interface{}) {
	GetLog().WithFields(logrus.Fields{
		"traceId": traceId,
	}).Debug(args)
}
func Errorf(format string, args ...interface{}) {
	GetLog().Errorf(format, args...)
}
func Warnf(format string, args ...interface{}) {
	GetLog().Warnf(format, args...)
}
func Infof(format string, args ...interface{}) {
	GetLog().Infof(format, args...)
}

func Debugf(format string, args ...interface{}) {
	GetLog().Debugf(format, args...)
}
