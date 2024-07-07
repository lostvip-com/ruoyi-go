package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/utils/lv_web"
)

func (w DemoController) Dialog(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/modal/dialog").WriteTpl()
}

func (w DemoController) Form(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/modal/form").WriteTpl()
}

func (w DemoController) Layer(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/modal/layer").WriteTpl()
}

func (w DemoController) Table(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/modal/table").WriteTpl()
}

func (w DemoController) Check(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/modal/table/check").WriteTpl()
}

func (w DemoController) Parent(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/modal/table/parent").WriteTpl()
}

func (w DemoController) Radio(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/modal/table/radio").WriteTpl()
}
