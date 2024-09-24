package model

//
//type SysLoginInfo struct {
//	InfoId        int64     `json:"info_id" xorm:"not null pk autoincr comment('访问ID') BIGINT(20)"`
//	LoginName     string    `json:"login_name" xorm:"default '' comment('登录账号') VARCHAR(50)"`
//	Ipaddr        string    `json:"ipaddr" xorm:"default '' comment('登录IP地址') VARCHAR(50)"`
//	LoginLocation string    `json:"login_location" xorm:"default '' comment('登录地点') VARCHAR(255)"`
//	Browser       string    `json:"browser" xorm:"default '' comment('浏览器类型') VARCHAR(50)"`
//	Os            string    `json:"os" xorm:"default '' comment('操作系统') VARCHAR(50)"`
//	Status        string    `json:"status" xorm:"default '0' comment('登录状态（0成功 1失败）') CHAR(1)"`
//	Msg           string    `json:"msg" xorm:"default '' comment('提示消息') VARCHAR(255)"`
//	LoginTime     time.Time `json:"login_time" xorm:"comment('访问时间') DATETIME"`
//}

// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-08-12 21:03:00 +0800 CST
// 生成人：lv
// ==========================================================================

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"time"
)

// SysLogininfor 系统访问记录
type SysLoginInfo struct {
	InfoId        int64     `gorm:"type:bigint(20);primary_key;auto_increment;访问ID;" json:"infoId"`
	LoginName     string    `gorm:"type:varchar(50);comment:登录账号;" json:"loginName"`
	Ipaddr        string    `gorm:"type:varchar(50);comment:登录IP地址;" json:"ipaddr"`
	LoginLocation string    `gorm:"type:varchar(255);comment:登录地点;" json:"loginLocation"`
	Browser       string    `gorm:"type:varchar(50);comment:浏览器类型;" json:"browser"`
	Os            string    `gorm:"type:varchar(50);comment:操作系统;" json:"os"`
	Status        string    `gorm:"type:char(1);comment:登录状态（0成功 1失败）;" json:"status"`
	Msg           string    `gorm:"type:varchar(255);comment:提示消息;" json:"msg"`
	LoginTime     time.Time `gorm:"type:datetime;comment:访问时间;" json:"loginTime" time_format:"2006-01-02 15:04:05"`
}

// 映射数据表
func (SysLoginInfo) TableName() string {
	return "sys_logininfor"
}

// Insert 插入数据
func (r *SysLoginInfo) Insert() error {
	return lv_db.GetMasterGorm().Save(r).Error
}

// Update 更新数据
func (r *SysLoginInfo) Update() error {
	return lv_db.GetMasterGorm().Updates(r).Error
}

// Delete 删除
func (r *SysLoginInfo) Delete() error {
	return lv_db.GetMasterGorm().Delete(r).Error
}

// 查
func (e *SysLoginInfo) FindById() error {
	err := lv_db.GetMasterGorm().Take(e, e.InfoId).Error
	return err
}
