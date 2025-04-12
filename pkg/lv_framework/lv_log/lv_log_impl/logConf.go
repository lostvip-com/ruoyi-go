package lv_log_impl

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
	"strings"
)

type LvLogImpl struct {
	Log *logrus.Logger
}

func InitLog(fileName string) *LvLogImpl {
	e := new(LvLogImpl)
	if lv_global.Config() == nil {
		cfg := new(lv_conf.ConfigDefault)
		lv_global.RegisterCfg(cfg)
	}
	if e.Log == nil {
		e.Log = logrus.New()
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	level := lv_global.Config().GetLogLevel()
	switch level {
	case "":
		e.Log.SetLevel(logrus.ErrorLevel)
	case "debug":
		lv_global.IsDebug = true
		e.Log.SetLevel(logrus.DebugLevel)
		fmt.Println("============ debug模式，输出日志到控制台 ============")
	case "info":
		e.Log.SetLevel(logrus.InfoLevel)
	case "warn":
		e.Log.SetLevel(logrus.WarnLevel)
	case "fatal":
		e.Log.SetLevel(logrus.FatalLevel)
		break
	case "error":
		e.Log.SetLevel(logrus.ErrorLevel)
	default:
		panic("Log level is not support: " + level)
	}
	maxSize := lv_global.Config().GetValueStr("go.log.max-size")
	maxBackups := lv_global.Config().GetValueStr("go.log.max-backups")
	maxAge := lv_global.Config().GetValueStr("go.log.max-age")
	if maxSize == "" {
		maxSize = "200"
	}
	if maxBackups == "" {
		maxBackups = "7"
	}
	if maxAge == "" {
		maxAge = "7"
	}
	logPath := lv_file.GetCurrentPath() + "/" + lv_global.Config().GetValueStr("go.log.path")
	err := lv_file.PathCreateIfNotExist(logPath)
	if err != nil {
		panic(err)
	}
	fileLog := &lumberjack.Logger{
		Filename:   logPath + "/" + fileName,
		MaxSize:    cast.ToInt(maxSize),    // 日志文件最大 size, 单位是 MB
		MaxBackups: cast.ToInt(maxBackups), // 最大过期日志保留的个数
		MaxAge:     cast.ToInt(maxAge),     //保留过期文件的最大时间间隔,单位是天
		Compress:   true,                   // disabled by default,是否需要压缩滚动日志, 使用的 gzip 压缩
	}
	var writers []io.Writer
	output := lv_global.Config().GetValueStr("go.log.output")
	if strings.Contains(output, "stdout") { //只写控制台
		writers = append(writers, os.Stdout)
	}
	if strings.Contains(output, "file") { // 只写文件
		writers = append(writers, fileLog)
	}
	if len(writers) == 0 { //同时写文件和控制台
		writers = append(writers, os.Stdout)
	}
	multiWriter := io.MultiWriter(writers...)
	e.Log.SetOutput(multiWriter)
	e.Log.Info("初始化测试Log.Info 完成 ！！stdout:", output)
	return e
}

func (e *LvLogImpl) GetLogWriter() io.Writer {
	return e.Log.Out
}
func (e *LvLogImpl) Error(args ...interface{}) {
	e.Log.Error(args)
}

func (e *LvLogImpl) ErrorTraceId(traceId any, args ...interface{}) {
	e.Log.WithFields(logrus.Fields{
		"traceId": traceId,
	}).Error(args)
}
func (e *LvLogImpl) Fatal(args ...interface{}) {
	e.Log.Fatal(args)
}
func (e *LvLogImpl) FatalTraceId(traceId any, args ...interface{}) {
	e.Log.WithFields(logrus.Fields{
		"traceId": traceId,
	}).Fatal(args)
}
func (e *LvLogImpl) Warn(args ...interface{}) {
	e.Log.Warn(args)
}
func (e *LvLogImpl) WarnTraceId(traceId any, args ...interface{}) {
	e.Log.WithFields(logrus.Fields{
		"traceId": traceId,
	}).Warn(args)
}
func (e *LvLogImpl) Info(args ...interface{}) {
	e.Log.Info(args)
}
func (e *LvLogImpl) InfoTraceId(traceId any, args ...interface{}) {
	e.Log.WithFields(logrus.Fields{
		"traceId": traceId,
	}).Info(args)
}
func (e *LvLogImpl) Debug(args ...interface{}) {
	e.Log.Debug(args...)
}
func (e *LvLogImpl) DebugTraceId(traceId any, args ...interface{}) {
	e.Log.WithFields(logrus.Fields{
		"traceId": traceId,
	}).Debug(args)
}
func (e *LvLogImpl) Errorf(format string, args ...interface{}) {
	e.Log.Errorf(format, args...)
}
func (e *LvLogImpl) Warnf(format string, args ...interface{}) {
	e.Log.Warnf(format, args...)
}
func (e *LvLogImpl) Infof(format string, args ...interface{}) {
	e.Log.Infof(format, args...)
}

func (e *LvLogImpl) Debugf(format string, args ...interface{}) {
	e.Log.Debugf(format, args...)
}
