package auth

import (
	"common/myconf"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_net"
	"net/http"
	"system/service"
)

// TokenCheckMiddleware 基于token的认证中间件
func TokenCheck() func(c *gin.Context) {
	session := service.SessionService{}
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里的具体实现方式要依据你的实际业务情况决定
		lv_log.Info("URL------------>" + c.Request.RequestURI)
		tokenStr := lv_net.GetParam(c, "token")
		isSignIn := session.IsSignedIn(tokenStr)
		if !isSignIn {
			c.Redirect(http.StatusFound, myconf.GetConfigInstance().GetContextPath()+"/login")
			return
		}
		//续期一小时
		session.Refresh(tokenStr)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
