package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
)

func (w DemoController) Dialog(c *gin.Context) {
	util.BuildTpl(c, "demo/modal/dialog").WriteTpl()
}

func (w DemoController) Form(c *gin.Context) {
	util.BuildTpl(c, "demo/modal/form").WriteTpl()
}

func (w DemoController) Layer(c *gin.Context) {
	util.BuildTpl(c, "demo/modal/layer").WriteTpl()
}

func (w DemoController) Table(c *gin.Context) {
	util.BuildTpl(c, "demo/modal/table").WriteTpl()
}

func (w DemoController) Check(c *gin.Context) {
	util.BuildTpl(c, "demo/modal/table/check").WriteTpl()
}

func (w DemoController) Parent(c *gin.Context) {
	//s := time.Now().UnixMilli()
	util.BuildTpl(c, "demo/modal/table/parent").WriteTpl()

	//fmt.Println(time.Now().UnixMilli() - s)
}

func (w DemoController) Radio(c *gin.Context) {
	util.BuildTpl(c, "demo/modal/table/radio").WriteTpl()
}
