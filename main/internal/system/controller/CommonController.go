package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"io"
	"os"
)

type CommonController struct{}

// 下载 public/upload 文件头像之类
func (w *CommonController) Download(c *gin.Context) {
	fileName := c.Query("fileName")
	//delete := c.Query("delete")
	if fileName == "" {
		util.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	curDir, err := os.Getwd()
	filepath := curDir + "/static/upload/" + fileName
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		util.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	b, _ := io.ReadAll(file)
	c.Writer.Header().Add("Content-Disposition", "attachment")
	c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Writer.Write(b)
	c.Abort()
}
