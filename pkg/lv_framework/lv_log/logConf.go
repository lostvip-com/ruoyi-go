package lv_log

import (
	"fmt"
	"io"
)

type ILog interface {
	Error(args ...interface{})
	ErrorTraceId(traceId any, args ...interface{})
	Fatal(args ...interface{})
	FatalTraceId(traceId any, args ...interface{})
	Warn(args ...interface{})
	WarnTraceId(traceId any, args ...interface{})
	Info(args ...interface{})
	InfoTraceId(traceId any, args ...interface{})
	Debug(args ...interface{})
	DebugTraceId(traceId any, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	GetLogWriter() io.Writer
}

var Log ILog //主数据库

func GetLog() ILog {
	if Log == nil {
		fmt.Println("log is nil!!!!!!!")
	}
	return Log
}

func Error(args ...interface{}) {
	if Log != nil {
		Log.Error(args)
	} else {
		fmt.Println(args)
	}
}

func ErrorTraceId(traceId any, args ...interface{}) {
	if Log != nil {
		Log.Error(args)
	} else {
		fmt.Println(args)
	}
}
func Fatal(args ...interface{}) {
	Log.Fatal(args)
}
func FatalTraceId(traceId any, args ...interface{}) {
	if Log != nil {
		Log.Fatal(args)
	} else {
		fmt.Println(args)
	}
}
func Warn(args ...interface{}) {
	if Log != nil {
		Log.Warn(args)
	} else {
		fmt.Println(args)
	}
}
func WarnTraceId(traceId any, args ...interface{}) {
	if Log != nil {
		Log.Warn(args)
	} else {
		fmt.Println(args)
	}
}
func Info(args ...interface{}) {
	if Log != nil {
		Log.Info(args)
	} else {
		fmt.Println(args)
	}
}
func InfoTraceId(traceId any, args ...interface{}) {
	if Log != nil {
		Log.Info(args)
	} else {
		fmt.Println(args)
	}
}
func Debug(args ...interface{}) {
	if Log != nil {
		Log.Debug(args)
	} else {
		fmt.Println(args)
	}
}
func DebugTraceId(traceId any, args ...interface{}) {
	if Log != nil {
		Log.Debug(args)
	} else {
		fmt.Println(args)
	}
}
func Errorf(format string, args ...interface{}) {
	if Log != nil {
		Log.Errorf(format, args)
	} else {
		fmt.Printf("\n"+format, args)
	}
}
func Warnf(format string, args ...interface{}) {
	if Log != nil {
		Log.Warnf(format, args)
	} else {
		fmt.Printf("\n"+format, args)
	}
}
func Infof(format string, args ...interface{}) {
	if Log != nil {
		Log.Infof(format, args)
	} else {
		fmt.Printf("\n"+format, args)
	}
}

func Debugf(format string, args ...interface{}) {
	if Log != nil {
		Log.Debugf(format, args)
	} else {
		fmt.Printf("\n"+format, args)
	}
}
