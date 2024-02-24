package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_web"
	"lostvip.com/web/dto"
	"net/http"
)

type OperateController struct{}

func (w OperateController) Add(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/operate/add").WriteTpl()
}

func (w OperateController) Detail(c *gin.Context) {
	var tmp us
	tmp.UserId = 1
	tmp.UserName = "测试1"
	tmp.Status = "0"
	tmp.CreateTime = "2020-01-12 02:02:02"
	tmp.UserBalance = 100
	tmp.UserCode = "1000001"
	tmp.UserSex = "0"
	tmp.UserPhone = "15888888888"
	tmp.UserEmail = "111@qq.com"
	lv_web.BuildTpl(c, "demo/operate/detail").WriteTpl(gin.H{"user": tmp})
}

type User struct {
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

func (w OperateController) EditSave(c *gin.Context) {
	var tmp us
	tmp.UserId = 1
	tmp.UserName = "测试1"
	tmp.Status = "0"
	tmp.CreateTime = "2020-01-12 02:02:02"
	tmp.UserBalance = 100
	tmp.UserCode = "1000001"
	tmp.UserSex = "0"
	tmp.UserPhone = "15888888888"
	tmp.UserEmail = "111@qq.com"
	lv_web.SucessResp(c).SetData(tmp).Log("demo演示", gin.H{"UserId": 1}).WriteJsonExit()
}

func (w OperateController) Edit(c *gin.Context) {
	var tmp us
	tmp.UserId = 1
	tmp.UserName = "测试1"
	tmp.Status = "0"
	tmp.CreateTime = "2020-01-12 02:02:02"
	tmp.UserBalance = 100
	tmp.UserCode = "1000001"
	tmp.UserSex = "0"
	tmp.UserPhone = "15888888888"
	tmp.UserEmail = "111@qq.com"
	lv_web.BuildTpl(c, "demo/operate/edit").WriteTpl(gin.H{"user": tmp})
}

func (w OperateController) Other(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/operate/other").WriteTpl()
}

func (w OperateController) Table(c *gin.Context) {
	lv_web.BuildTpl(c, "demo/operate/table").WriteTpl()
}

func (w OperateController) List(c *gin.Context) {
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
	c.JSON(http.StatusOK, dto.TableDataInfo{
		Code:  200,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}
