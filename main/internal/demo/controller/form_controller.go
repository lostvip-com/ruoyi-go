package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
)

func (w DemoController) Autocomplete(c *gin.Context) {

	lv_web.BuildTpl(c, "demo/form/autocomplete").WriteTpl()
}

func (w DemoController) Basic(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/basic").WriteTpl()
}

func (w DemoController) DemoControllerButton(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/button").WriteTpl()
}

func (w DemoController) Cards(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/cards").WriteTpl()
}

func (w DemoController) Datetime(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/datetime").WriteTpl()
}

func (w DemoController) Duallistbox(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/duallistbox").WriteTpl()
}

func (w DemoController) Grid(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/grid").WriteTpl()
}

func (w DemoController) Jasny(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/jasny").WriteTpl()
}

func (w DemoController) Select(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/select").WriteTpl()
}

func (w DemoController) Sortable(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/sortable").WriteTpl()
}

func (w DemoController) Summernote(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/summernote").WriteTpl()
}

func (w DemoController) Tabs_panels(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/tabs_panels").WriteTpl()
}

func (w DemoController) Timeline(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/timeline").WriteTpl()
}

func (w DemoController) Upload(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/upload").WriteTpl()
}

func (w DemoController) Validate(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/validate").WriteTpl()
}

func (w DemoController) Wizard(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/form/wizard").WriteTpl()
}

func (w DemoController) UserModel(c *gin.Context) {
	data := make([]map[string]any, 6)
	data[0] = map[string]any{"userId": 1, "userName": "abcdefg 1 "}
	data[1] = map[string]any{"userId": 2, "userName": "ruoyi 2 "}
	data[2] = map[string]any{"userId": 3, "userName": "ruoyi 3 "}
	data[3] = map[string]any{"userId": 4, "userName": "test 4 "}
	data[4] = map[string]any{"userId": 5, "userName": "lostvip 5 "}
	data[5] = map[string]any{"userId": 6, "userName": "gogo 6 "}
	lv_web.Success(c, data, "success")
}
