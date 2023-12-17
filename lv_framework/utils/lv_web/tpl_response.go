package lv_web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 通用tpl响应
type TplResp struct {
	c   *gin.Context
	tpl string
}

// 返回一个tpl响应
func BuildTpl(c *gin.Context, tpl string) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: tpl,
	}
	return &t
}

// 返回一个错误的tpl响应
func ErrorTpl(c *gin.Context) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: "error/error.html",
	}
	return &t
}

// 返回一个无操作权限tpl响应
func ForbiddenTpl(c *gin.Context) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: "error/unauth.html",
	}
	return &t
}

// 输出页面模板
func (resp *TplResp) WriteTpl(params ...gin.H) {
	//session := sessions.Default(resp.c)
	//uid := session.Get(model_cmn.USER_ID)
	uid := 1
	if params == nil || len(params) == 0 {
		resp.c.HTML(http.StatusOK, resp.tpl, gin.H{"uid": uid})
	} else {
		params[0]["uid"] = uid
		resp.c.HTML(http.StatusOK, resp.tpl, params[0])
	}
	//resp.c.Abort()
}
