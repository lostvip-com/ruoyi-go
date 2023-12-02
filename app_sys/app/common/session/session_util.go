package session

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"lostvip.com/cache/myredis"
	"lostvip.com/utils/lv_net"
	"robvi/app/global"
	userModel "robvi/app/modules/sys/model/system/user"
	"sync"
)

// 用户session列表
var SessionList sync.Map

// 判断是否是系统管理员
func IsAdmin(userId int64) bool {
	if userId == 1 {
		return true
	} else {
		return false
	}
}

func IsAdminUser(user *userModel.SysUser) bool {
	if user.UserId == 1 {
		return true
	} else {
		return false
	}
}

// 判断用户是否已经登录
func IsSignedIn(c *gin.Context) bool {
	_, exist := c.Get(global.USER_ID)
	if exist {
		return true
	}
	return false
}

// 获得用户信息详情
func GetProfile(c *gin.Context) *userModel.SysUser {
	tokenStr := lv_net.GetParam(c, "token")

	userId := myredis.GetInstance().HMGet(context.Background(), "login:"+tokenStr, "userId").String()
	if userId == "" {
		return nil
	}
	user := userModel.SysUser{}
	user.UserId = cast.ToInt64(userId)
	_, err := user.FindOne()
	if err != nil {
		return nil
	}
	//err := json.Unmarshal([]byte(s), &u)
	if err != nil {
		return nil
	}
	return &user
}

// 获得用户信息详情
func GetTenantId(c *gin.Context) int64 {
	u := GetProfile(c)
	return u.TenantId
}
