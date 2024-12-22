package middleware

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
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
					c.Header("Content-Type", "application/json; charset=utf-8")
					c.String(http.StatusOK, errTypeObj)
					c.Abort()
				} else {
					util.Err(c, errTypeObj)
				}
			case lv_dto.Resp: //封装过的
				c.AbortWithStatusJSON(http.StatusOK, errTypeObj)
				util.ErrResp(c, errTypeObj)
			case error: // 原始的错误
				if gin.IsDebugging() {
					lv_err.PrintStackTrace(errTypeObj)
				}
				lv_log.Error(c, "CustomError XXXXXXXXXX: ", errTypeObj)
				util.Error(c, errTypeObj)
			default:
				lv_log.Error(c, "default CustomErrorXXXXXXXXXX: ", errTypeObj)
				util.Err(c, "未知错误!")
			}
		} else {
			lv_log.Info(c, "-----------request over----------")
		}
	}()
	c.Next()
}
