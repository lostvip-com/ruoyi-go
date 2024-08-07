package logininfor

import (
	"github.com/lostvip-com/lv_framework/lv_cache"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	"github.com/spf13/cast"
	"main/internal/common/global"
	logininfor2 "main/internal/system/model/monitor/logininfor"
	"time"
)

const USER_NOPASS_TIME string = "user_nopass_"
const USER_LOCK string = "user_lock_"

// 根据条件分页查询用户列表
func SelectPageList(param *logininfor2.SelectPageReq) (*[]logininfor2.Entity, *lv_web.Paging, error) {
	return logininfor2.SelectPageList(param)
}

// 根据主键查询用户信息
func SelectRecordById(id int64) (*logininfor2.Entity, error) {
	entity := &logininfor2.Entity{InfoId: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除用户信息
func DeleteRecordById(id int64) bool {
	entity := &logininfor2.Entity{InfoId: id}
	result, err := entity.Delete()
	if err == nil && result > 0 {
		return true
	}

	return false
}

// 批量删除记录
func DeleteRecordByIds(ids string) int64 {
	idarr := lv_conv.ToInt64Array(ids, ",")
	result, _ := logininfor2.DeleteBatch(idarr...)
	return result
}

// 清空记录
func DeleteRecordAll() (int64, error) {
	return logininfor2.DeleteAll()
}

// 导出excel
func Export(param *logininfor2.SelectPageReq) (string, error) {
	head := []string{"访问编号", "登录名称", "登录地址", "登录地点", "浏览器", "操作系统", "登录状态", "操作信息", "登录时间"}
	col := []string{"info_id", "login_name", "ipaddr", "login_location", "browser", "os", "status", "msg", "login_time"}
	return logininfor2.SelectExportList(param, head, col)
}

// 记录密码尝试次数
func SetPasswordCounts(loginName string) int {
	curTimes := 0
	curTimeObj, err := lv_cache.GetCacheClient().Get(USER_NOPASS_TIME + loginName)
	if err == nil {
		curTimes = cast.ToInt(curTimeObj)
	}
	curTimes = curTimes + 1
	lv_cache.GetCacheClient().Set(USER_NOPASS_TIME+loginName, curTimes, 1*time.Minute)

	if curTimes >= global.ErrTimes2Lock {
		Lock(loginName)
	}
	return curTimes
}

// 记录密码尝试次数
func GetPasswordCounts(loginName string) int {
	curTimes := 0
	curTimeObj, err := lv_cache.GetCacheClient().Get(USER_NOPASS_TIME + loginName)
	if err != nil {
		curTimes = cast.ToInt(curTimeObj)
	}
	return curTimes
}

// 移除密码错误次数
func RemovePasswordCounts(loginName string) {
	lv_cache.GetCacheClient().Del(USER_NOPASS_TIME + loginName)
}

// 锁定账号
func Lock(loginName string) {
	lv_cache.GetCacheClient().Set(USER_LOCK+loginName, true, 30*time.Minute)
}

// 解除锁定
func Unlock(loginName string) {
	lv_cache.GetCacheClient().Del(USER_LOCK + loginName)
}

// 检查账号是否锁定
func CheckLock(loginName string) bool {
	result := false
	rs, _ := lv_cache.GetCacheClient().Get(USER_LOCK + loginName)
	if rs != "" {
		result = true
	}
	return result
}
