package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_web"
)

func (w DemoController) Echarts(c *gin.Context) {
	lv_web.BuildTpl(c, "modules/demo/report/echarts").WriteTpl()
}

func (w DemoController) Metrics(c *gin.Context) {
	lv_web.BuildTpl(c, "modules/demo/report/metrics").WriteTpl()
}

func (w DemoController) Peity(c *gin.Context) {
	lv_web.BuildTpl(c, "modules/demo/report/peity").WriteTpl()
}

func (w DemoController) Sparkline(c *gin.Context) {
	lv_web.BuildTpl(c, "modules/demo/report/sparkline").WriteTpl()
}
