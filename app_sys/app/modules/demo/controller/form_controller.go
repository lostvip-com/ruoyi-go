package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/response"
)

func (w DemoController) Autocomplete(c *gin.Context) {

	response.BuildTpl(c, "modules/demo/form/autocomplete").WriteTpl()
}

func (w DemoController) Basic(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/basic").WriteTpl()
}

func (w DemoController) DemoControllerButton(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/button").WriteTpl()
}

func (w DemoController) Cards(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/cards").WriteTpl()
}

func (w DemoController) Datetime(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/datetime").WriteTpl()
}

func (w DemoController) Duallistbox(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/duallistbox").WriteTpl()
}

func (w DemoController) Grid(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/grid").WriteTpl()
}

func (w DemoController) Jasny(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/jasny").WriteTpl()
}

func (w DemoController) Select(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/select").WriteTpl()
}

func (w DemoController) Sortable(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/sortable").WriteTpl()
}

func (w DemoController) Summernote(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/summernote").WriteTpl()
}

func (w DemoController) Tabs_panels(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/tabs_panels").WriteTpl()
}

func (w DemoController) Timeline(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/timeline").WriteTpl()
}

func (w DemoController) Upload(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/upload").WriteTpl()
}

func (w DemoController) Validate(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/validate").WriteTpl()
}

func (w DemoController) Wizard(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/form/wizard").WriteTpl()
}
