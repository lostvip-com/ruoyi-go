package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lostvip.com/logme"
	"lostvip.com/utils/lv_web"
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
					lv_web.Err(c, errTypeObj)
				}
			case dto.Resp: //封装过的
				c.AbortWithStatusJSON(http.StatusOK, errTypeObj)
				lv_web.ErrResp(c, errTypeObj)
			case error: // 原始的错误
				logme.Log.Error(c, "CustomError XXXXXXXXXX: ", errTypeObj)
				lv_web.Error(c, errTypeObj)
			default:
				logme.Log.Error(c, "default CustomErrorXXXXXXXXXX: ", errTypeObj)
				lv_web.Err(c, "未知错误!")
			}
		} else {
			logme.Log.Info(c, "-----------request over----------")
		}
	}()
	c.Next()
}
