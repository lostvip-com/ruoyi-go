package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lostvip.com/dto"
	"lostvip.com/logme"
	"lostvip.com/utils/lib_net"
	"net/http"
	"strings"
)

func CustomError(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			//if c.IsAborted() {
			//	c.Status(200)
			//}
			switch errTypeObj := err.(type) {
			case string:
				if strings.HasPrefix(errTypeObj, "{") {
					fmt.Println("============>" + errTypeObj)
					c.Header("Content-Type", "application/json; charset=utf-8")
					c.String(http.StatusOK, errTypeObj)
					c.Abort()
				} else {
					lib_net.Err(c, errTypeObj)
				}
			case dto.Resp: //封装过的
				c.AbortWithStatusJSON(http.StatusOK, errTypeObj)
				lib_net.ErrResp(c, errTypeObj)
			case error: // 原始的错误
				logme.Log.Error(c, "CustomError XXXXXXXXXX: ", errTypeObj)
				lib_net.Error(c, errTypeObj)
			default:
				logme.Log.Error(c, "default CustomErrorXXXXXXXXXX: ", errTypeObj)
				lib_net.Err(c, "未知错误!")
			}
		} else {
			logme.Log.Info(c, "-----------request over----------")
		}
	}()
	c.Next()
}
