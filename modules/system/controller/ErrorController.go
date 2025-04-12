package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
)

type ErrorController struct{}

func (w *ErrorController) Unauth(c *gin.Context) {
	util.BuildTpl(c, "error/unauth").WriteTpl()
}

func (w *ErrorController) Error(c *gin.Context) {
	util.BuildTpl(c, "error/500").WriteTpl()
}

func (w *ErrorController) NotFound(c *gin.Context) {
	util.BuildTpl(c, "error/404").WriteTpl()
}
