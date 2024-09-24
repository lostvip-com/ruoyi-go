// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-08-04 10:51:57 +0800 CST
// 生成人：lv
// ==========================================================================
package model

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/namedsql"
)

// SysDictData 字典数据
type SysDictData struct {
	DictCode  int64  `gorm:"type:bigint(20);primary_key;auto_increment;字典编码;" json:"dictCode"`
	DictSort  int    `gorm:"type:int(11);comment:字典排序;" json:"dictSort"`
	DictLabel string `gorm:"type:varchar(100);comment:字典标签;" json:"dictLabel"`
	DictValue string `gorm:"type:varchar(100);comment:字典键值;" json:"dictValue"`
	DictType  string `gorm:"type:varchar(100);comment:字典类型;" json:"dictType"`
	CssClass  string `gorm:"type:varchar(100);comment:样式属性（其他样式扩展）;" json:"cssClass"`
	ListClass string `gorm:"type:varchar(100);comment:表格回显样式;" json:"listClass"`
	IsDefault string `gorm:"type:char(1);comment:是否默认（Y是 N否）;" json:"isDefault"`
	Status    string `gorm:"type:char(1);comment:状态（0正常 1停用）;" json:"status"`
	Remark    string `gorm:"type:varchar(500);comment:备注;" json:"remark"`
	BaseModel
}

func (e *SysDictData) TableName() string {
	return "sys_dict_data"
}

// 增
func (e *SysDictData) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysDictData) FindById(id int64) (*SysDictData, error) {
	err := lv_db.GetMasterGorm().Take(e, id).Error
	return e, err
}

// 查第一条
func (e *SysDictData) FindOne() (*SysDictData, error) {
	tb := lv_db.GetMasterGorm().Table(e.TableName())

	if e.DictLabel != "" {
		tb = tb.Where("dict_label=?", e.DictLabel)
	}
	if e.DictValue != "" {
		tb = tb.Where("dict_value=?", e.DictValue)
	}
	if e.DictType != "" {
		tb = tb.Where("dict_type=?", e.DictType)
	}
	err := tb.First(e).Error
	return e, err
}

// 改
func (e *SysDictData) Updates() error {
	return lv_db.GetMasterGorm().Updates(e).Error
}

// 删
func (e *SysDictData) Delete() error {
	return lv_db.GetMasterGorm().Delete(e).Error
}

func (e *SysDictData) Count() (int64, error) {
	sql := " select count(*) from sys_dict_data where del_flag = 0 "

	if e.DictSort != 0 {
		sql += " and dict_sort = @DictSort "
	}
	if e.DictLabel != "" {
		sql += " and dict_label = @DictLabel "
	}
	if e.DictValue != "" {
		sql += " and dict_value = @DictValue "
	}
	if e.DictType != "" {
		sql += " and dict_type = @DictType "
	}

	return namedsql.Count(lv_db.GetMasterGorm(), sql, e)
}
