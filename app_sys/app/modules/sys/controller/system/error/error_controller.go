package error

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/response"
)

func Unauth(c *gin.Context) {
	response.BuildTpl(c, "error/unauth").WriteTpl()
}

func Error(c *gin.Context) {
	response.BuildTpl(c, "error/500").WriteTpl()
}

func NotFound(c *gin.Context) {
	response.BuildTpl(c, "error/404").WriteTpl()
}
