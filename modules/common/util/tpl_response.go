package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TplResp 通用tpl响应
type TplResp struct {
	c   *gin.Context
	tpl string
}

// BuildTpl 返回一个tpl响应
func BuildTpl(c *gin.Context, tpl string) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: tpl,
	}
	return &t
}

// ErrorTpl 返回一个错误的tpl响应
//
//	Deprecated
func ErrorTpl(c *gin.Context) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: "error/error.html",
	}
	return &t
}

// ForbiddenTpl 返回一个无操作权限tpl响应
// Deprecated
func ForbiddenTpl(c *gin.Context) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: "error/unauth.html",
	}
	return &t
}

// WriteTpl 输出页面模板
func (resp *TplResp) WriteTpl(params ...gin.H) {
	//session := sessions.Default(resp.c)
	//uid := session.Get(model_cmn.USER_ID)
	uid := 1
	//s := time.Now().UnixMilli()
	if params == nil || len(params) == 0 {
		resp.c.HTML(http.StatusOK, resp.tpl, gin.H{"uid": uid})
	} else {
		params[0]["uid"] = uid
		resp.c.HTML(http.StatusOK, resp.tpl, params[0])
	}
	resp.c.Abort()
}
