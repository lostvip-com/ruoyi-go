package pages

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"github.com/spf13/cast"
	"net/http"
)

func (w DemoController) Button(c *gin.Context) {
	util.BuildTpl(c, "demo/github.com/lostvip-com/lv_framework/utils/button").WriteTpl()
}

func (w DemoController) Child(c *gin.Context) {
	util.BuildTpl(c, "demo/table/child").WriteTpl()
}

func (w DemoController) Curd(c *gin.Context) {
	util.BuildTpl(c, "demo/table/curd").WriteTpl()
}

func (w DemoController) Detail(c *gin.Context) {
	util.BuildTpl(c, "demo/table/detail").WriteTpl()
}

func (w DemoController) Editable(c *gin.Context) {
	util.BuildTpl(c, "demo/table/editable").WriteTpl()
}

func (w DemoController) Event(c *gin.Context) {
	util.BuildTpl(c, "demo/table/event").WriteTpl()
}

func (w DemoController) Export(c *gin.Context) {
	util.BuildTpl(c, "demo/table/export").WriteTpl()
}

func (w DemoController) FixedColumns(c *gin.Context) {
	util.BuildTpl(c, "demo/table/fixedColumns").WriteTpl()
}

func (w DemoController) Footer(c *gin.Context) {
	util.BuildTpl(c, "demo/table/footer").WriteTpl()
}

func (w DemoController) GroupHeader(c *gin.Context) {
	util.BuildTpl(c, "demo/table/groupHeader").WriteTpl()
}

func (w DemoController) Image(c *gin.Context) {
	util.BuildTpl(c, "demo/table/image").WriteTpl()
}

func (w DemoController) Multi(c *gin.Context) {
	util.BuildTpl(c, "demo/table/multi").WriteTpl()
}

func (w DemoController) Other(c *gin.Context) {
	util.BuildTpl(c, "demo/table/other").WriteTpl()
}

func (w DemoController) PageGo(c *gin.Context) {
	util.BuildTpl(c, "demo/table/pageGo").WriteTpl()
}

func (w DemoController) Params(c *gin.Context) {
	util.BuildTpl(c, "demo/table/params").WriteTpl()
}

func (w DemoController) Remember(c *gin.Context) {
	util.BuildTpl(c, "demo/table/remember").WriteTpl()
}

func (w DemoController) Recorder(c *gin.Context) {
	util.BuildTpl(c, "demo/table/recorder").WriteTpl()
}

func (w DemoController) Search(c *gin.Context) {
	util.BuildTpl(c, "demo/table/search").WriteTpl()
}

type us struct {
	UserId      int64   `json:"userId"`
	UserCode    string  `json:"userCode"`
	UserName    string  `json:"userName"`
	UserSex     string  `json:"userSex"`
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
		tmp.UserName = "测试" + cast.ToString(i)
		tmp.Status = "0"
		tmp.CreateTime = "2020-01-12 02:02:02"
		tmp.UserBalance = 100
		tmp.UserCode = cast.ToString(i)
		tmp.UserSex = "0"
		tmp.UserPhone = "15888888888"
		tmp.UserEmail = "111@qq.com"
		rows = append(rows, tmp)
	}
	c.JSON(http.StatusOK, lv_dto.TableDataInfo{
		Code:  200,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}
