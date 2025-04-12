package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (w *HomeController) Home(c *gin.Context) {
	util.BuildTpl(c, "home").WriteTpl()
}
