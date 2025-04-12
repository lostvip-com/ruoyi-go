package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"io"
	"os"
)

type CommonController struct{}

// DownloadTmp 从临时目录下载，如excell等动态生成的数据（默认下载）
func (w *CommonController) DownloadTmp(c *gin.Context) {
	fileName := c.Query("fileName")
	filepath := lv_global.Config().GetTmpPath() + "/" + fileName
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		lv_log.Error(err)
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

// DownloadUpload 从upload目录下载,下载 public/upload 文件头像之类
func (w *CommonController) DownloadUpload(c *gin.Context) {
	fileName := c.Query("fileName")
	filepath := lv_global.Config().GetUploadPath() + "/" + fileName
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		lv_log.Error(err)
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
