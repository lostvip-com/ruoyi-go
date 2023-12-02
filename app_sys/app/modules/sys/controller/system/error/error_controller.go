package error

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_web"
)

func Unauth(c *gin.Context) {
	lv_web.BuildTpl(c, "error/unauth").WriteTpl()
}

func Error(c *gin.Context) {
	lv_web.BuildTpl(c, "error/500").WriteTpl()
}

func NotFound(c *gin.Context) {
	lv_web.BuildTpl(c, "error/404").WriteTpl()
}
