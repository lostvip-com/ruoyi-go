package logme

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"lostvip.com/conf"
	"lostvip.com/utils/lv_file"
	"os"
)

type MyLogrus struct {
	Logger *logrus.Logger
}

var Log *MyLogrus

func InitLog(logfile string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pwd = pwd + "/logs"
	lv_file.PathCreateIfNotExist(pwd)
	if Log == nil {
		Log = new(MyLogrus)
		Log.Logger = logrus.New()
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	if conf.Config().IsDebug() {
		fmt.Println("============ debug模式，输出日志到控制台 ============")
		Log.Logger.SetLevel(logrus.DebugLevel)
	} else {
		Log.Logger.SetLevel(logrus.InfoLevel)
		fmt.Println("============ release模式，输出日志json到文件============")
		//Log.Logger.SetFormatter(&logrus.JSONFormatter{})
	}
	if conf.Config().GetValueStr("go.log.output") == "stdout" {
		fmt.Println("============ debug模式，输出日志到控制台 ============")
		log.SetOutput(os.Stdout)
		Log.Logger.SetOutput(os.Stdout) //调用 logrus 的 SetOutput()函数
	} else {
		logger := &lumberjack.Logger{
			Filename:   logfile,
			MaxSize:    500,  // 日志文件最大 size, 单位是 MB
			MaxBackups: 3,    // 最大过期日志保留的个数
			MaxAge:     28,   //保留过期文件的最大时间间隔,单位是天
			Compress:   true, // disabled by default,是否需要压缩滚动日志, 使用的 gzip 压缩
		}
		log.SetOutput(logger)
		Log.Logger.SetOutput(logger) //调用 logrus 的 SetOutput()函数
	}

	Log.Logger.WithFields(logrus.Fields{
		"omg":    true,
		"number": 123456,
	}).Warn("初始化测试！Log.WithFields(logrus.Field{}).Warn:  The group's number increased tremendously!")

	//logrus.SetOutput(&logrus.RotateWriter{
	//	Filename:   file.Name(),
	//	MaxSize:    10 * 1024 * 1024, // 每个日志文件最大10MB
	//	MaxBackups: 5,                // 最多保存5个旧日志文件
	//	MaxAge:     30,               // 日志文件最多保存30天
	//	LocalTime:  true,
	//})

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
