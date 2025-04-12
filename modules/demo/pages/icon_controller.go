package pages

import (
	"common/util"
	"github.com/gin-gonic/gin"
)

type DemoController struct{}

func (w DemoController) Fontawesome(c *gin.Context) {
	util.BuildTpl(c, "demo/icon/fontawesome").WriteTpl()
}

func (w DemoController) Glyphicons(c *gin.Context) {
	util.BuildTpl(c, "demo/icon/glyphicons").WriteTpl()
}
