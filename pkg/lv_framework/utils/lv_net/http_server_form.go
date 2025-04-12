package lv_net

import (
	"github.com/gin-gonic/gin"
)

func GetParam(c *gin.Context, key string) string {
	v := c.Query(key)
	if v == "" {
		v = c.PostForm(key)
		if v == "" {
			v = c.GetHeader(key)
			if v == "" {
				cookie, err := c.Request.Cookie(key)
				if err == nil { //3 从cookie取
					v = cookie.Value
				}
			}
		}
	}
	return v
}
