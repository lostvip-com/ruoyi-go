// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-07-19 17:16:13 +0800 CST
// 生成人：lv
// ==========================================================================
package model

import (
	"common/model"
	"github.com/lostvip-com/lv_framework/lv_db"
)

type IotProduct struct {
	Id           int64  `gorm:"type:bigint(20);primary_key;auto_increment;主键;" json:"id"`
	Code         string `gorm:"type:varchar(32);comment:产品编码,对应可监控类型ID;" json:"code"`
	Name         string `gorm:"type:varchar(32);comment:名字;" json:"name"`
	Protocol     string `gorm:"type:varchar(32);comment:协议;" json:"protocol"`
	NodeType     string `gorm:"type:varchar(32);comment:节点类型;" json:"nodeType"`
	Manufacturer string `gorm:"type:varchar(32);comment:生产厂商;" json:"manufacturer"`
	Active       string `gorm:"type:char(1);default:1;comment:是否启用;column:active;" json:"active"`
	Description  string `gorm:"type:text;comment:描述;" json:"description"`
	//
	NodeTypeName string `gorm:"-" json:"nodeTypeName"`
	ProtocolName string `gorm:"-" json:"protocolName"`
	model.BaseModel
}

func (e *IotProduct) TableName() string {
	return "iot_product"
}

// 增
func (e *IotProduct) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// 查
func (e *IotProduct) FindById() error {
	err := lv_db.GetMasterGorm().Take(e, e.Id).Error
	return err
}

// 查第一条
func (e *IotProduct) FindOne() error {
	tb := lv_db.GetMasterGorm().Table(e.TableName())

	if e.Name != "" {
		tb = tb.Where("name like ?", "%"+e.Name+"%")
	}

	err := tb.First(e).Error
	return err
}

// 改
func (e *IotProduct) Updates() error {
	return lv_db.GetMasterGorm().Updates(e).Error
}

// 删
func (e *IotProduct) Delete() error {
	return lv_db.GetMasterGorm().Delete(e).Error
}
