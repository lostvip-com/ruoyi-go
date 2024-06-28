package logme

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/utils/lv_file"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"lostvip.com/lv_global"
	"os"
)

var logRu *logrus.Logger

func GetLog() *logrus.Logger {
	if logRu == nil {
		logRu = logrus.New()
	}
	return logRu
}

func InitLog(fileName string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pwd = pwd + "/logs"
	err = lv_file.PathCreateIfNotExist(pwd)
	if err != nil {
		panic(err)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	if lv_global.IsDebug {
		fmt.Println("============ debug模式，输出日志到控制台 ============")
		GetLog().SetLevel(logrus.DebugLevel)
	} else {
		GetLog().SetLevel(logrus.InfoLevel)
		fmt.Println("============ release模式，输出日志json到文件============")
	}
	fileLog := &lumberjack.Logger{
		Filename:   pwd + "/" + fileName,
		MaxSize:    500,  // 日志文件最大 size, 单位是 MB
		MaxBackups: 3,    // 最大过期日志保留的个数
		MaxAge:     28,   //保留过期文件的最大时间间隔,单位是天
		Compress:   true, // disabled by default,是否需要压缩滚动日志, 使用的 gzip 压缩
	}
	var writers []io.Writer
	output := lv_global.LogOutputType
	if output == "stdout" { //只写控制台
		writers = append(writers, os.Stdout)
	} else if output == "file" { // 只写文件
		writers = append(writers, fileLog)
	} else { //同时写文件和屏幕
		writers = append(writers, os.Stdout, fileLog)
	}
	multiWriter := io.MultiWriter(writers...)
	log.SetOutput(multiWriter)
	GetLog().Info("初始化测试Log.Info 完成 ！！stdout:", output)
}

func Error(args ...interface{}) {
	GetLog().Error(args)
}

func Errorf(format string, args ...interface{}) {
	GetLog().Errorf(format, args)
}
func ErrorTraceId(c *gin.Context, args ...interface{}) {
	ip := c.ClientIP()
	traceId := c.GetHeader("traceId")
	GetLog().WithFields(logrus.Fields{
		"ip":      ip,
		"traceId": traceId,
	}).Error(args)
}
func Fatal(args ...interface{}) {
	GetLog().Fatal(args)
}
func FatalWithTrace(c *gin.Context, args ...interface{}) {
	ip := c.ClientIP()
	traceId := c.GetHeader("traceId")
	GetLog().WithFields(logrus.Fields{
		"ip":      ip,
		"traceId": traceId,
	}).Fatal(args)
}
func Warn(args ...interface{}) {
	GetLog().Warn(args)
}

func Info(args ...interface{}) {
	GetLog().Info(args)
}

func InfoWithTraceId(c *gin.Context, args ...interface{}) {
	ip := c.ClientIP()
	traceId := c.GetHeader(lv_global.TraceId)
	GetLog().WithFields(logrus.Fields{
		"ip":      ip,
		"traceId": traceId,
	}).Info(args)
}

func Debug(args ...interface{}) {
	fmt.Println(args)
}
