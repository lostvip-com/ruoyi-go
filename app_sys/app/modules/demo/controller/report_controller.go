package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/response"
)

func (w DemoController) Echarts(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/report/echarts").WriteTpl()
}

func (w DemoController) Metrics(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/report/metrics").WriteTpl()
}

func (w DemoController) Peity(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/report/peity").WriteTpl()
}

func (w DemoController) Sparkline(c *gin.Context) {
	response.BuildTpl(c, "modules/demo/report/sparkline").WriteTpl()
}
