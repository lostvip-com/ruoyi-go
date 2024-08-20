// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-07-30 21:59:38 +0800 CST
// 生成人：lv
// ==========================================================================
package model

import (
	"common/model"
	"github.com/lostvip-com/lv_framework/lv_db"
)

// IotPrdProperty 属性
type IotPrdProperty struct {
	Id              int64  `gorm:"type:type:bigint(20);primary_key;auto_increment;主键;" json:"id"`
	ProductId       int64  `gorm:"type:bigint(20);comment:产品ID;" json:"productId"`
	Name            string `gorm:"type:varchar(32);comment:名字;" json:"name"`
	Code            string `gorm:"type:varchar(32);comment:标识符;" json:"code"`
	Tag             string `gorm:"type:varchar(50);comment:标签;" json:"tag"`
	AccessMode      string `gorm:"type:varchar(50);comment:读写模型;" json:"accessMode"`
	Type            string `gorm:"type:varchar(50);comment:数据类型;" json:"type"`
	Precision       int16  `gorm:"size:10;default:1;comment:倍率" json:"precision"`
	Unit            string `gorm:"type:varchar(50);comment:单位;" json:"unit"`
	Remark          string `gorm:"type:longtext;comment:描述;" json:"remark"`
	model.BaseModel `gorm:"embedded"`
}

func (e *IotPrdProperty) TableName() string {
	return "iot_prd_property"
}

// Save 增
func (e *IotPrdProperty) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// FindById 查
func (e *IotPrdProperty) FindById() error {
	err := lv_db.GetMasterGorm().Take(e, e.Id).Error
	return err
}

// FindOne 查第一条
func (e *IotPrdProperty) FindOne() error {
	tb := lv_db.GetMasterGorm().Table(e.TableName())

	if e.ProductId != 0 {
		tb = tb.Where("product_id=?", e.ProductId)
	}
	if e.Name != "" {
		tb = tb.Where("name=?", e.Name)
	}
	if e.Code != "" {
		tb = tb.Where("code=?", e.Code)
	}
	if e.Tag != "" {
		tb = tb.Where("tag=?", e.Tag)
	}
	err := tb.First(e).Error
	return err
}

// Updates 改
func (e *IotPrdProperty) Updates() error {
	return lv_db.GetMasterGorm().Updates(e).Error
}

// Delete 删
func (e *IotPrdProperty) Delete() error {
	return lv_db.GetMasterGorm().Delete(e).Error
}
