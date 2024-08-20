// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-08-04 10:51:58 +0800 CST
// 生成人：lv
// ==========================================================================
package model

import (
	cm_model "common/model"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/namedsql"
)

// SysDictType 字典类型
type SysDictType struct {
	DictId   int64  `gorm:"type:bigint(20);primary_key;auto_increment;字典主键;" json:"dictId"`
	DictName string `gorm:"type:varchar(100);comment:字典名称;" json:"dictName"`
	DictType string `gorm:"type:varchar(100);comment:字典类型;" json:"dictType"`
	Status   string `gorm:"type:char(1);comment:状态（0正常 1停用）;" json:"status"`
	Remark   string `gorm:"type:varchar(500);comment:备注;" json:"remark"`
	cm_model.BaseModel
}

func (e *SysDictType) TableName() string {
	return "sys_dict_type"
}

// 增
func (e *SysDictType) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysDictType) FindById(id int64) (*SysDictType, error) {
	err := lv_db.GetMasterGorm().Take(e, id).Error
	return e, err
}

// 查第一条
func (e *SysDictType) FindOne() (*SysDictType, error) {
	tb := lv_db.GetMasterGorm().Table(e.TableName())

	if e.DictName != "" {
		tb = tb.Where("dict_name=?", e.DictName)
	}
	if e.DictType != "" {
		tb = tb.Where("dict_type=?", e.DictType)
	}
	if e.Status != "" {
		tb = tb.Where("status=?", e.Status)
	}
	err := tb.First(e).Error
	return e, err
}

// 改
func (e *SysDictType) Updates() error {
	return lv_db.GetMasterGorm().Updates(e).Error
}

// 删
func (e *SysDictType) Delete() error {
	return lv_db.GetMasterGorm().Delete(e).Error
}

func (e *SysDictType) Count() (int64, error) {
	sql := " select count(*) from sys_dict_type where 1=1 "
	if e.DictType != "" {
		sql += " and dict_type = @DictType "
	}
	return namedsql.Count(lv_db.GetMasterGorm(), sql, e)
}
