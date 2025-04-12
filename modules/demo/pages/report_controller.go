package pages

import (
	"common/util"
	"github.com/gin-gonic/gin"
)

func (w DemoController) Echarts(c *gin.Context) {
	util.BuildTpl(c, "demo/report/echarts").WriteTpl()
}

func (w DemoController) Metrics(c *gin.Context) {
	util.BuildTpl(c, "demo/report/metrics").WriteTpl()
}

func (w DemoController) Peity(c *gin.Context) {
	util.BuildTpl(c, "demo/report/peity").WriteTpl()
}

func (w DemoController) Sparkline(c *gin.Context) {
	util.BuildTpl(c, "demo/report/sparkline").WriteTpl()
}
