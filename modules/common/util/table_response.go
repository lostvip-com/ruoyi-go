package util

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"net/http"
)

// 通用api响应
type TableResp struct {
	t *lv_dto.TableDataInfo
	c *gin.Context
}

// BuildTable 返回一个成功的消息体
// Deprecated
func BuildTable(c *gin.Context, total any, rows interface{}) *TableResp {
	msg := lv_dto.TableDataInfo{
		Code:  lv_dto.SUCCESS,
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

// WriteJsonExit 输出json到客户端
// Deprecated
func (resp *TableResp) WriteJsonExit() {
	resp.c.JSON(http.StatusOK, resp.t)
	resp.c.Abort()
}
