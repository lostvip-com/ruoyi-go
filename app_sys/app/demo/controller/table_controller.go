package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_web"
	"net/http"
	"robvi/app/sys/model"
)

func (w DemoController) Button(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/button").WriteTpl()
}

func (w DemoController) Child(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/child").WriteTpl()
}

func (w DemoController) Curd(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/curd").WriteTpl()
}

func (w DemoController) Detail(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/detail").WriteTpl()
}

func (w DemoController) Editable(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/editable").WriteTpl()
}

func (w DemoController) Event(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/event").WriteTpl()
}

func (w DemoController) Export(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/export").WriteTpl()
}

func (w DemoController) FixedColumns(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/fixedColumns").WriteTpl()
}

func (w DemoController) Footer(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/footer").WriteTpl()
}

func (w DemoController) GroupHeader(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/groupHeader").WriteTpl()
}

func (w DemoController) Image(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/image").WriteTpl()
}

func (w DemoController) Multi(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/multi").WriteTpl()
}

func (w DemoController) Other(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/other").WriteTpl()
}

func (w DemoController) PageGo(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/pageGo").WriteTpl()
}

func (w DemoController) Params(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/params").WriteTpl()
}

func (w DemoController) Remember(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/remember").WriteTpl()
}

func (w DemoController) Recorder(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/recorder").WriteTpl()
}

func (w DemoController) Search(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/table/lv_sql").WriteTpl()
}

type us struct {
	UserId      int64   `json:"userId"`
	UserCode    string  `json:"userCode"`
	UserName    string  `json:"userName"`
	UserSex     string  `json:"userName"`
	UserPhone   string  `json:"userPhone"`
	UserEmail   string  `json:"userEmail"`
	UserBalance float64 `json:"userBalance"`
	Status      string  `json:"status"`
	CreateTime  string  `json:"createTime"`
}

func (w DemoController) List(c *gin.Context) {
	var rows = make([]us, 0)
	for i := 1; i <= 10; i++ {
		var tmp us
		tmp.UserId = int64(i)
		tmp.UserName = "测试" + string(i)
		tmp.Status = "0"
		tmp.CreateTime = "2020-01-12 02:02:02"
		tmp.UserBalance = 100
		tmp.UserCode = "100000" + string(i)
		tmp.UserSex = "0"
		tmp.UserPhone = "15888888888"
		tmp.UserEmail = "111@qq.com"
		rows = append(rows, tmp)
	}
	c.JSON(http.StatusOK, model.TableDataInfo{
		Code:  200,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}
