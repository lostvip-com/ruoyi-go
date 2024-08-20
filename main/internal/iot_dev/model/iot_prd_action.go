package model

import (
	"common/model"
	"github.com/lostvip-com/lv_framework/lv_db"
	"main/internal/iot_dev/pkg/constants"
)

type IotPrdAction struct {
	Id              int64              `json:"id" gorm:"type:bigint(20);primaryKey;not null;size:32;comment:主键"`
	ProductId       int64              `json:"product_id" gorm:"size:32;comment:产品ID"`
	Code            string             `json:"code" gorm:"size:32;comment:标识符"`
	Tag             string             `json:"tag" gorm:"size:50;comment:标签"`
	Name            string             `json:"name" gorm:"size:32;comment:名字"`
	CallType        constants.CallType `json:"call_type" gorm:"size:50;comment:调用方式"`
	InputParams     string             `json:"input_params" gorm:"type:text;comment:输入参数"`  // 输入参数
	OutputParams    string             `json:"output_params" gorm:"type:text;comment:输入参数"` // 输出参数
	Remark          string             `json:"remark" gorm:"size:64;comment:描述"`
	model.BaseModel `gorm:"embedded"`
}

func (table *IotPrdAction) TableName() string {
	return "iot_prd_action"
}

//
//type InPutParams []InputOutput
//
//type InputOutput struct {
//	Code     string              `json:"code"`
//	Name     string              `json:"name"`
//	Type     constants.SpecsType `json:"type,omitempty"`
//	Min      string              `json:"min,omitempty"`
//	Max      string              `json:"max,omitempty"`
//	Step     string              `json:"step,omitempty"`
//	Unit     string              `json:"unit,omitempty"`
//	UnitName string              `json:"unitName,omitempty"`
//}
//
//// Value Value() 方法会在将值写入数据库时被调用
//func (c InPutParams) Value() (driver.Value, error) {
//	return GormValueWrap(c)
//}
//
//// Scan sql.Scanner 接口的 Scan() 方法会在从数据库读取值时被调用
//func (c *InPutParams) Scan(value interface{}) error {
//	return GormScanWrap(value, c)
//}
//
//type OutPutParams []InputOutput
//
//func (c OutPutParams) Value() (driver.Value, error) {
//	return GormValueWrap(c)
//}

//func (c *OutPutParams) Scan(value interface{}) error {
//	return GormScanWrap(value, c)
//}

// 增
func (e *IotPrdAction) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// 查
func (e *IotPrdAction) FindById() error {
	err := lv_db.GetMasterGorm().Take(e, e.Id).Error
	return err
}

// 查第一条
func (e *IotPrdAction) FindOne() error {
	tb := lv_db.GetMasterGorm().Table(e.TableName())

	if e.ProductId != 0 {
		tb = tb.Where("product_id=?", e.ProductId)
	}
	if e.Code != "" {
		tb = tb.Where("code=?", e.Code)
	}
	if e.Tag != "" {
		tb = tb.Where("tag=?", e.Tag)
	}
	if e.Name != "" {
		tb = tb.Where("name=?", e.Name)
	}
	err := tb.First(e).Error
	return err
}

// 改
func (e *IotPrdAction) Updates() error {
	return lv_db.GetMasterGorm().Updates(e).Error
}

// 删
func (e *IotPrdAction) Delete() error {
	return lv_db.GetMasterGorm().Delete(e).Error
}
