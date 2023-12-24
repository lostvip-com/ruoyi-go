package logme

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"lostvip.com/conf"
	"lostvip.com/utils/lv_file"
	"os"
)

var log *logrus.Logger

func GetLog() *logrus.Logger {
	return log
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
	if log == nil {
		log = logrus.New()
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	if conf.Config().IsDebug() {
		fmt.Println("============ debug模式，输出日志到控制台 ============")
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
		fmt.Println("============ release模式，输出日志json到文件============")
		//log.SetFormatter(&logrus.JSONFormatter{})
	}
	if conf.Config().GetValueStr("go.log.output") == "stdout" {
		fmt.Println("============ debug模式，输出日志到控制台 ============")
		log.SetOutput(os.Stdout)
		log.SetOutput(os.Stdout) //调用 logrus 的 SetOutput()函数
	} else {
		logger := &lumberjack.Logger{
			Filename:   pwd + "/" + fileName,
			MaxSize:    500,  // 日志文件最大 size, 单位是 MB
			MaxBackups: 3,    // 最大过期日志保留的个数
			MaxAge:     28,   //保留过期文件的最大时间间隔,单位是天
			Compress:   true, // disabled by default,是否需要压缩滚动日志, 使用的 gzip 压缩
		}
		log.SetOutput(logger)
		log.SetOutput(logger) //调用 logrus 的 SetOutput()函数
	}

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 123456,
	}).Warn("初始化测试！Log.WithFields(logrus.Field{}).Warn:  The group's number increased tremendously!")

	log.Info("初始化测试Log.Info 完成 ！！")
}

func Error(args ...interface{}) {
	log.Info(args)
}

func ErrorTraceId(c *gin.Context, args ...interface{}) {
	ip := c.ClientIP()
	traceId := c.GetHeader("traceId")
	log.WithFields(logrus.Fields{
		"ip":      ip,
		"traceId": traceId,
	}).Error(args)
}
func Fatal(args ...interface{}) {
	log.Fatal(args)
}
func FatalWithTrace(c *gin.Context, args ...interface{}) {
	ip := c.ClientIP()
	traceId := c.GetHeader("traceId")
	log.WithFields(logrus.Fields{
		"ip":      ip,
		"traceId": traceId,
	}).Fatal(args)
}
func Warn(args ...interface{}) {
	log.Warn(args)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func InfoWithTraceId(c *gin.Context, args ...interface{}) {
	ip := c.ClientIP()
	traceId := c.GetHeader(conf.TraceId)
	log.WithFields(logrus.Fields{
		"ip":      ip,
		"traceId": traceId,
	}).Info(args)
}

func Debug(args ...interface{}) {
	fmt.Println(args)
}
