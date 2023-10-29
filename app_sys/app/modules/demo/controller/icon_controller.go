package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/response"
)

type DemoController struct{}

func (w DemoController) Fontawesome(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/icon/fontawesome").WriteTpl()
}

func (w DemoController) Glyphicons(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/icon/glyphicons").WriteTpl()
}
