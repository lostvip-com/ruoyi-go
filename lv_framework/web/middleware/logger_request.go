package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/logme"
	"time"
)

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logData := map[string]interface{}{
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIP":    clientIP,
			"method":      reqMethod,
			"uri":         reqUri,
		}
		logme.Info(c, logData)
		// 禁用日志写入数据库的功能 ssz20210702
		//if c.Request.Method != "GET" && c.Request.Method != "OPTIONS" && conf.LoggerConfig.EnabledDB {
		//	SetDBOperLog(c, clientIP, statusCode, reqUri, reqMethod, latencyTime)
		//}
	}
}
