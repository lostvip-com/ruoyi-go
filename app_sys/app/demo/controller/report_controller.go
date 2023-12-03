package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_web"
)

func (w DemoController) Echarts(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/report/echarts").WriteTpl()
}

func (w DemoController) Metrics(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/report/metrics").WriteTpl()
}

func (w DemoController) Peity(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/report/peity").WriteTpl()
}

func (w DemoController) Sparkline(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/report/sparkline").WriteTpl()
}
