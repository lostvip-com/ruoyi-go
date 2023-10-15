package logme

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"lostvip.com/conf"
	"os"
)

type MyLogrus struct {
	Logger *logrus.Logger
}

var Log *MyLogrus

func InitLog(logfile string) {
	if Log == nil {
		Log = new(MyLogrus)
		Log.Logger = logrus.New()
	}
	// 输出到标准输出，而不是默认的标准错误
	// Log 为JSON而不是默认的ASCII格式。
	if conf.Config().IsDebug() {
		fmt.Println("============ debug模式，输出日志到控制台 ============")
		Log.Logger.SetOutput(os.Stdout)
		Log.Logger.SetLevel(logrus.DebugLevel)
	} else {
		Log.Logger.SetLevel(logrus.InfoLevel)
		fmt.Println("============ release模式，输出日志json到文件============")
		file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			Log.Logger.SetOutput(file)
		} else {
			Log.Logger.Info("Failed to log to file, using default stderr")
		}
		Log.Logger.SetFormatter(&logrus.JSONFormatter{})
	}

	Log.Logger.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("初始化测试! Log.WithFields(logrus.Field{}).Info : ！A group of walrus emerges from the ocean")

	Log.Logger.WithFields(logrus.Fields{
		"omg":    true,
		"number": 123456,
	}).Warn("初始化测试！Log.WithFields(logrus.Field{}).Warn:  The group's number increased tremendously!")

	//Log.WithFields(logrus.Fields{
	//	"number": 100,
	//}).Fatal("The ice breaks!")
	// 一种常见的模式是通过重用
	//从WithFields返回的logrus.Entry 来重用日志记录语句之间的字段
	Log.Logger.Info("初始化测试Log.Info: I'll be logged with common and other field")
	Log.Logger.Info("初始化测试Log.Info 完成 ！！")
}

func (log *MyLogrus) Error(args ...interface{}) {
	log.Logger.Error(args)
}

func (log *MyLogrus) ErrorTraceId(c *gin.Context, args ...interface{}) {
	ip := c.ClientIP()
	traceId := c.GetHeader("traceId")
	log.Logger.WithFields(logrus.Fields{
		"ip":      ip,
		"traceId": traceId,
	}).Error(args)
}
func (log *MyLogrus) Fatal(args ...interface{}) {
	log.Logger.Fatal(args)
}
func (log *MyLogrus) FatalWithTrace(c *gin.Context, args ...interface{}) {
	ip := c.ClientIP()
	traceId := c.GetHeader("traceId")
	log.Logger.WithFields(logrus.Fields{
		"ip":      ip,
		"traceId": traceId,
	}).Fatal(args)
}
func (log *MyLogrus) Warn(args ...interface{}) {
	ip := ""
	traceId := ""
	log.Logger.WithFields(logrus.Fields{
		"ip":      ip,
		"traceId": traceId,
	}).Warn(args)
}

func (log *MyLogrus) Info(args ...interface{}) {
	log.Logger.Info(args)
}

func (log *MyLogrus) InfoWithTraceId(c *gin.Context, args ...interface{}) {
	ip := c.ClientIP()
	traceId := c.GetHeader(conf.TraceId)
	log.Logger.WithFields(logrus.Fields{
		"ip":      ip,
		"traceId": traceId,
	}).Info(args)
}

func (log *MyLogrus) Debug(args ...interface{}) {
	fmt.Println(args)
}
