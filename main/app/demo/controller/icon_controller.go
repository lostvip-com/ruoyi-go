package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_web"
)

type DemoController struct{}

func (w DemoController) Fontawesome(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/icon/fontawesome").WriteTpl()
}

func (w DemoController) Glyphicons(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/icon/glyphicons").WriteTpl()
}
