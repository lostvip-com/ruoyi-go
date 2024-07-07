package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/utils/lv_web"
)

type ErrorController struct{}

func (w *ErrorController) Unauth(c *gin.Context) {
	lv_web.BuildTpl(c, "error/unauth").WriteTpl()
}

func (w *ErrorController) Error(c *gin.Context) {
	lv_web.BuildTpl(c, "error/500").WriteTpl()
}

func (w *ErrorController) NotFound(c *gin.Context) {
	lv_web.BuildTpl(c, "error/404").WriteTpl()
}
