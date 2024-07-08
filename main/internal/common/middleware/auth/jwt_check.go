package auth

//
//import (
//	"context"
//	"github.com/gin-gonic/gin"
//	"github.com/lv_framework/myredis"
//	"github.com/lv_framework/utils/token"
//	"net/http"
//	"main/internal/global"
//	"strings"
//	"time"
//)
//
//// ////////////////////////////////////////////////////////////////
//// 移动端使用
//// TokenCheck 基于JWT的认证中间件
//// /////////////////////////////////////////////////////////////////
//func JWTCheckMiddleware() func(c *gin.Context) {
//	return func(c *gin.Context) {
//		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
//		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
//		// 这里的具体实现方式要依据你的实际业务情况决定
//		authHeader := c.Request.Header.Get("Authorization")
//		// 按空格分割
//		tokenStr := ""
//		if authHeader != "" { //1先从header取
//			parts := strings.SplitN(authHeader, " ", 2)
//			if len(parts) > 1 {
//				tokenStr = parts[1]
//			}
//		} else {
//			tokenStr = c.Request.Header.Get("token")
//		}
//
//		if tokenStr == "" { //2从url取
//			tokenStr = c.Query("token")
//		}
//		cookie, err := c.Request.Cookie("token")
//		if err == nil { //3 从cookie取
//			tokenStr = cookie.Value
//		}
//
//		if tokenStr == "" {
//			//c.JSON(http.StatusOK, gin.H{
//			//	"code": 500,
//			//	"msg":  "获取token失败",
//			//})
//			//c.Abort()
//			c.Redirect(http.StatusFound, global.GetConfigInstance().GetContextPath()+"/login")
//			return
//		}
//		//续期一小时
//		myredis.GetInstance().Expire(context.Background(), tokenStr, time.Hour)
//		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
//		t, err := token.VerifyAuthToken(tokenStr)
//		if err != nil {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 2005,
//				"msg":  "无效的Token",
//			})
//			c.Abort()
//			return
//		}
//		// 将当前请求的uid信息保存到请求的上下文c上
//		c.Set("tenantId", t.Claim.TenantId)
//		c.Set("userId", t.Claim.UserId)
//		c.Set("loginName", t.Claim.Subject)
//		//判断是否有新的token生成
//		if t.NewToken != "" {
//			c.Writer.Header().Set("nt", t.NewToken)
//		}
//		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
//	}
//}
