package lv_web

import (
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/web/dto"
	"net/http"
)

// 通用api响应
type TableResp struct {
	t *dto.TableDataInfo
	c *gin.Context
}

// 返回一个成功的消息体
func BuildTable(c *gin.Context, total any, rows interface{}) *TableResp {
	msg := dto.TableDataInfo{
		Code:  dto.SUCCESS,
		Msg:   "操作成功",
		Total: total,
		Rows:  rows,
	}
	a := TableResp{
		t: &msg,
		c: c,
	}
	return &a
}

// 输出json到客户端
func (resp *TableResp) WriteJsonExit() {
	resp.c.JSON(http.StatusOK, resp.t)
	resp.c.Abort()
}
