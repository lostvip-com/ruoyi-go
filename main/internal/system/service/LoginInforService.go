package service

import (
	"github.com/lostvip-com/lv_framework/lv_cache"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/namedsql"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_office"
	"github.com/spf13/cast"
	"main/internal/common/global"
	"main/internal/system/model"
	"main/internal/system/vo"
	"time"
	"xorm.io/builder"
)

const USER_NOPASS_TIME string = "user_nopass_"
const USER_LOCK string = "user_lock_"

type LoginInforService struct {
}

// SelectPageList 根据条件分页查询用户列表
func (svc LoginInforService) SelectPageList(param *vo.LoginInfoPageReq) (*[]model.SysLoginInfo, int64, error) {
	db := lv_db.GetMasterGorm()
	tb := db.Table("sys_logininfor")
	if param != nil {
		if param.LoginName != "" {
			tb.Where("login_name like ?", "%"+param.LoginName+"%")
		}
		if param.Ipaddr != "" {
			tb.Where("ipaddr like ?", "%"+param.Ipaddr+"%")
		}
		if param.Status != "" {
			tb.Where("status = ?", param.Status)
		}

		if param.BeginTime != "" {
			tb.Where("login_time >= ?", param.BeginTime)
		}
		if param.EndTime != "" {
			tb.Where("login_time <= ?", param.EndTime)
		}
	}
	var total int64
	tb = tb.Count(&total).Offset(param.GetStartNum()).Limit(param.GetPageSize())
	if param.SortName != "" {
		tb.Order("info_id " + param.SortOrder)
	}
	var result []model.SysLoginInfo
	err := tb.Find(&result).Error
	return &result, total, err
}

// SelectRecordById 根据主键查询用户信息
func (svc LoginInforService) SelectRecordById(id int64) (*model.SysLoginInfo, error) {
	entity := &model.SysLoginInfo{InfoId: id}
	err := entity.FindById()
	return entity, err
}

// DeleteRecordById 根据主键删除用户信息
func (svc LoginInforService) DeleteRecordById(id int64) bool {
	entity := &model.SysLoginInfo{InfoId: id}
	err := entity.Delete()
	if err == nil {
		return true
	}

	return false
}

// DeleteRecordByIds 批量删除记录
func (svc LoginInforService) DeleteRecordByIds(ids string) error {
	idarr := lv_conv.ToInt64Array(ids, ",")
	err := lv_db.GetMasterGorm().Exec("delete from sys_logininfor where info_id in ? ", idarr).Error
	return err
}

// 清空记录
func (svc LoginInforService) DeleteRecordAll() error {
	db := lv_db.GetMasterGorm()
	err := db.Exec("delete from sys_logininfor").Error
	return err
}

// 导出excel
func (svc LoginInforService) Export(param *vo.LoginInfoPageReq) (string, error) {
	head := []string{"访问编号", "登录名称", "登录地址", "登录地点", "浏览器", "操作系统", "登录状态", "操作信息", "登录时间"}
	col := []string{"info_id", "login_name", "ipaddr", "login_location", "browser", "os", "status", "msg", "login_time"}
	db := lv_db.GetMasterGorm()
	build := builder.Select(col...).From(" sys_logininfor ", "t")
	if param != nil {
		if param.LoginName != "" {
			build.Where(builder.Like{"t.login_name", param.LoginName})
		}

		if param.Ipaddr != "" {
			build.Where(builder.Like{"t.ipaddr", param.Ipaddr})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"t.status": param.Status})
		}

		if param.BeginTime != "" {
			build.Where(builder.Gte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.BeginTime + "','%y%m%d')"})
		}

		if param.EndTime != "" {
			build.Where(builder.Lte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.EndTime + "','%y%m%d')"})
		}
	}
	sqlStr, err := build.ToBoundSQL()
	arr, err := namedsql.ListArrStr(db, sqlStr, nil)
	path, err := lv_office.DownlaodExcel(head, *arr)
	return path, err
}

// 记录密码尝试次数
func (svc LoginInforService) SetPasswordCounts(loginName string) int {
	curTimes := 0
	curTimeObj, err := lv_cache.GetCacheClient().Get(USER_NOPASS_TIME + loginName)
	if err == nil {
		curTimes = cast.ToInt(curTimeObj)
	}
	curTimes = curTimes + 1
	lv_cache.GetCacheClient().Set(USER_NOPASS_TIME+loginName, curTimes, 1*time.Minute)

	if curTimes >= global.ErrTimes2Lock {
		svc.Lock(loginName)
	}
	return curTimes
}

// 记录密码尝试次数
func (svc LoginInforService) GetPasswordCounts(loginName string) int {
	curTimes := 0
	curTimeObj, err := lv_cache.GetCacheClient().Get(USER_NOPASS_TIME + loginName)
	if err != nil {
		curTimes = cast.ToInt(curTimeObj)
	}
	return curTimes
}

// 移除密码错误次数
func (svc LoginInforService) RemovePasswordCounts(loginName string) {
	lv_cache.GetCacheClient().Del(USER_NOPASS_TIME + loginName)
}

// 锁定账号
func (svc LoginInforService) Lock(loginName string) {
	lv_cache.GetCacheClient().Set(USER_LOCK+loginName, true, 30*time.Minute)
}

// 解除锁定
func (svc LoginInforService) Unlock(loginName string) {
	lv_cache.GetCacheClient().Del(USER_LOCK + loginName)
}

// 检查账号是否锁定
func (svc LoginInforService) CheckLock(loginName string) bool {
	result := false
	rs, _ := lv_cache.GetCacheClient().Get(USER_LOCK + loginName)
	if rs != "" {
		result = true
	}
	return result
}
