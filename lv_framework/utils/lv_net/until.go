package lv_net

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FailAuth(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 400, "msg": msg, "data": nil})
}

func Fail(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 500, "msg": msg, "data": nil})
}

func ReadToken(c *gin.Context) (*JWTClaims, error) {
	tokenStr := c.GetHeader("token") // header
	fmt.Println("######################## c.GetHeader：", tokenStr)
	if tokenStr == "" { //get
		tokenStr = c.DefaultQuery("token", "")
		fmt.Println("######################## c.DefaultQuery：", tokenStr)
	} else if tokenStr == "" { //post
		tokenStr = c.DefaultPostForm("token", "")
		fmt.Println("######################## c.GetString：", tokenStr)
	}
	return ReadTokenInfo(tokenStr)
}

func ReadParam(c *gin.Context, key string) string {
	value := c.GetHeader(key) // header
	fmt.Println("######################## c.GetHeader：", value)
	if value == "" { //get
		value = c.DefaultQuery("token", "")
		fmt.Println("######################## c.DefaultQuery：", value)
	} else if value == "" { //post
		value = c.DefaultPostForm("token", "")
		fmt.Println("######################## c.GetString：", value)
	}
	return value
}

func ReadTokenInfo(tokenStr string) (*JWTClaims, error) {
	fmt.Println("token==================> token:", tokenStr)
	var err error
	var jwt *JWTClaims
	if tokenStr == "" || tokenStr == "null" { //js传过来的
		error := errors.New("token 为空，非法请求！")
		return jwt, error
	} else {
		jwt, err = VerifyAction(tokenStr)
	}
	return jwt, err
}
