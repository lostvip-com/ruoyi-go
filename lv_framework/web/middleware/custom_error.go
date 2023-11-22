package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lostvip.com/logme"
	"lostvip.com/utils/response"
	"lostvip.com/web/dto"
	"net/http"
	"strings"
)

func RecoverError(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			switch errTypeObj := err.(type) {
			case string:
				if strings.HasPrefix(errTypeObj, "{") {
					fmt.Println("============>" + errTypeObj)
					c.Header("Content-Type", "application/json; charset=utf-8")
					c.String(http.StatusOK, errTypeObj)
					c.Abort()
				} else {
					response.Err(c, errTypeObj)
				}
			case dto.Resp: //封装过的
				c.AbortWithStatusJSON(http.StatusOK, errTypeObj)
				response.ErrResp(c, errTypeObj)
			case error: // 原始的错误
				logme.Log.Error(c, "CustomError XXXXXXXXXX: ", errTypeObj)
				response.Error(c, errTypeObj)
			default:
				logme.Log.Error(c, "default CustomErrorXXXXXXXXXX: ", errTypeObj)
				response.Err(c, "未知错误!")
			}
		} else {
			logme.Log.Info(c, "-----------request over----------")
		}
	}()
	c.Next()
}
