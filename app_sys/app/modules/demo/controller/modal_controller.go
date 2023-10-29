package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/response"
)

func (w DemoController) Dialog(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/modal/dialog").WriteTpl()
}

func (w DemoController) Form(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/modal/form").WriteTpl()
}

func (w DemoController) Layer(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/modal/layer").WriteTpl()
}

func (w DemoController) Table(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/modal/table").WriteTpl()
}

func (w DemoController) Check(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/modal/table/check").WriteTpl()
}

func (w DemoController) Parent(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/modal/table/parent").WriteTpl()
}

func (w DemoController) Radio(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/modal/table/radio").WriteTpl()
}
