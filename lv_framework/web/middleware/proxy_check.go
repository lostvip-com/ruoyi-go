package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/utils/lv_net"
	"lostvip.com/lv_global"
	"strings"
)

func IfProxyForward() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里的具体实现方式要依据你的实际业务情况决定
		uri := c.Request.RequestURI
		isPorxyEnable := lv_global.Config().IsProxyEnabled()
		if !isPorxyEnable { //不支持代理
			c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		}
		//支持代理
		prefix, hostAndPort := getProxyPrefixAndHostPort(uri)
		if hostAndPort != "" {
			uriNew := strings.TrimLeft(uri, prefix)
			lv_net.ProxyWithUrlDiff(c, hostAndPort, uriNew)
			return
		}

	}
}

func getProxyPrefixAndHostPort(uri string) (string, string) {
	instance := lv_global.Config()
	mp := instance.GetProxyMap()
	for k, v := range *mp {
		fmt.Println(k, "==============", v)
		if strings.HasPrefix(uri, k) {
			return k, v
		}
	}
	return "", ""
}
