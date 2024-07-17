// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-02-28 14:21:50 +0800 CST
// 生成人：lv
// ==========================================================================
package model

import (
	"github.com/lostvip-com/lv_framework/db"
	"time"
)

// AppGenParams 参量字典
type AppGenParams struct {
	Id            int64     `gorm:"type:bigint;primary_key;auto_increment;ID;" json:"id"`
	UseFlag       string    `gorm:"type:char(1);comment:是否使用;" json:"useFlag"`
	ParamNo       int       `gorm:"type:int;comment:参量号;" json:"paramNo"`
	ParamName     string    `gorm:"type:varchar(255);comment:参量名称;" json:"paramName"`
	ParamType     string    `gorm:"type:varchar(255);comment:参量类型;" json:"paramType"`
	Unit          string    `gorm:"type:varchar(255);comment:单位;" json:"unit"`
	Remark        string    `gorm:"type:varchar(255);comment:备注信息;" json:"remark"`
	MonitorTypeId int       `gorm:"type:int;comment:监控类型;" json:"monitorTypeId"`
	CreateTime    time.Time `gorm:"type:datetime;comment:创建时间;column:create_time;" json:"createTime" time_format:"2006-01-02 15:04:05"`
	UpdateTime    time.Time `gorm:"type:datetime;comment:创建时间;column:update_time;" json:"updateTime" time_format:"2006-01-02 15:04:05"`
	CreateBy      string    `gorm:"type:varchar(32);comment:创建人;column:create_by;"  json:"createBy"`
	UpdateBy      string    `gorm:"type:varchar(32);comment:创建人;column:update_by;"  json:"updateBy"`
	DelFlag       int       `gorm:"type:tinyint(1);default:0;comment:删除标记;column:del_flag;" json:"delFlag"`
}

func (e *AppGenParams) TableName() string {
	return "app_gen_params"
}

// 增
func (e *AppGenParams) Save() error {
	return db.GetMasterGorm().Save(e).Error
}

// 查
func (e *AppGenParams) FindById() error {
	err := db.GetMasterGorm().Take(e, e.Id).Error
	return err
}

// 查第一条
func (e *AppGenParams) FindOne() error {
	tb := db.GetMasterGorm().Table(e.TableName())

	if e.ParamNo != 0 {
		tb = tb.Where("param_no=?", e.ParamNo)
	}
	if e.ParamName != "" {
		tb = tb.Where("param_name=?", e.ParamName)
	}
	if e.ParamType != "" {
		tb = tb.Where("param_type=?", e.ParamType)
	}
	if e.MonitorTypeId != 0 {
		tb = tb.Where("monitor_type_id=?", e.MonitorTypeId)
	}
	err := tb.First(e).Error
	return err
}

// 改
func (e *AppGenParams) Updates() error {
	return db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 删
func (e *AppGenParams) Delete() error {
	return db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}
