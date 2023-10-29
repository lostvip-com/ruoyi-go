package logininfor

import (
	"lostvip.com/utils/convert"
	"lostvip.com/utils/gconv"
	"lostvip.com/utils/page"
	"robvi/app/common/cache"
	"robvi/app/global"
	"robvi/app/modules/sys/model/monitor/logininfor"
	"time"
)

const USER_NOPASS_TIME string = "user_nopass_"
const USER_LOCK string = "user_lock_"

// 根据条件分页查询用户列表
func SelectPageList(param *logininfor.SelectPageReq) (*[]logininfor.Entity, *page.Paging, error) {
	return logininfor.SelectPageList(param)
}

// 根据主键查询用户信息
func SelectRecordById(id int64) (*logininfor.Entity, error) {
	entity := &logininfor.Entity{InfoId: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除用户信息
func DeleteRecordById(id int64) bool {
	entity := &logininfor.Entity{InfoId: id}
	result, err := entity.Delete()
	if err == nil && result > 0 {
		return true
	}

	return false
}

// 批量删除记录
func DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, _ := logininfor.DeleteBatch(idarr...)
	return result
}

// 清空记录
func DeleteRecordAll() (int64, error) {
	return logininfor.DeleteAll()
}

// 导出excel
func Export(param *logininfor.SelectPageReq) (string, error) {
	head := []string{"访问编号", "登录名称", "登录地址", "登录地点", "浏览器", "操作系统", "登录状态", "操作信息", "登录时间"}
	col := []string{"info_id", "login_name", "ipaddr", "login_location", "browser", "os", "status", "msg", "login_time"}
	return logininfor.SelectExportList(param, head, col)
}

// 记录密码尝试次数
func SetPasswordCounts(loginName string) int {
	curTimes := 0
	curTimeObj, _ := cache.Instance().Get(USER_NOPASS_TIME + loginName)
	if curTimeObj != nil {
		curTimes = gconv.Int(curTimeObj)
	}
	curTimes = curTimes + 1
	cache.Instance().Set(USER_NOPASS_TIME+loginName, curTimes, 1*time.Minute)

	if curTimes >= global.ErrTimes2Lock {
		Lock(loginName)
	}
	return curTimes
}

// 记录密码尝试次数
func GetPasswordCounts(loginName string) int {
	curTimes := 0
	curTimeObj, _ := cache.Instance().Get(USER_NOPASS_TIME + loginName)
	if curTimeObj != nil {
		curTimes = gconv.Int(curTimeObj)
	}
	return curTimes
}

// 移除密码错误次数
func RemovePasswordCounts(loginName string) {
	cache.Instance().Delete(USER_NOPASS_TIME + loginName)
}

// 锁定账号
func Lock(loginName string) {
	cache.Instance().Set(USER_LOCK+loginName, true, 30*time.Minute)
}

// 解除锁定
func Unlock(loginName string) {
	cache.Instance().Delete(USER_LOCK + loginName)
}

// 检查账号是否锁定
func CheckLock(loginName string) bool {
	result := false
	rs, _ := cache.Instance().Get(USER_LOCK + loginName)
	if rs != nil {
		result = true
	}
	return result
}
