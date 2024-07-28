package session

import (
	"common/cm_vo"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_cache"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_net"
	"github.com/spf13/cast"
)

// 获得用户信息详情
func GetTenantId(c *gin.Context) int64 {
	return 0
}

func GetLoginInfo(c *gin.Context) *cm_vo.LoginInfo {
	token := lv_net.GetParam(c, "token")
	key := "login:" + token
	mp, err := lv_cache.GetCacheClient().HGetAll(key)
	lv_err.HasErrAndPanic(err)
	login := new(cm_vo.LoginInfo)
	login.Username = mp["username"]
	login.Avatar = mp["avatar"]
	login.RoleKeys = mp["roleKeys"]
	deptId := mp["deptId"]
	login.DeptId = cast.ToInt64(deptId)
	login.UserId = cast.ToInt64(mp["userId"])
	login.TenantId = cast.ToInt64(mp["tenantId"])
	return login
}
