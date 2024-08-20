package model

import (
	"common/model"
	"github.com/lostvip-com/lv_framework/lv_db"
)

type IotPrdEvent struct {
	Id              int64  `json:"id" gorm:"type:bigint(20);primaryKey;not null;size:22;comment:主键"`
	ProductId       int64  `json:"product_id" gorm:"size:32;comment:产品ID"`
	EventType       string `json:"event_type" gorm:"size:32;comment:事件类型"`
	Code            string `json:"code" gorm:"size:32;comment:标识符"`
	Tag             string `json:"tag" gorm:"size:50;comment:标签"`
	Name            string `json:"name" gorm:"size:32;comment:名字"`
	Remark          string `json:"remark" gorm:"comment:描述"`
	OutputParams    string `json:"output_params" gorm:"type:text;comment:输入参数"`
	model.BaseModel `gorm:"embedded"`
}

func (table *IotPrdEvent) TableName() string {
	return "iot_prd_event"
}

// 增
func (e *IotPrdEvent) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// 查
func (e *IotPrdEvent) FindById() error {
	err := lv_db.GetMasterGorm().Take(e, e.Id).Error
	return err
}

// 查第一条
func (e *IotPrdEvent) FindOne() error {
	tb := lv_db.GetMasterGorm().Table(e.TableName())

	if e.ProductId != 0 {
		tb = tb.Where("product_id=?", e.ProductId)
	}
	if e.EventType != "" {
		tb = tb.Where("event_type=?", e.EventType)
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

// 改
func (e *IotPrdEvent) Updates() error {
	return lv_db.GetMasterGorm().Updates(e).Error
}

// 删
func (e *IotPrdEvent) Delete() error {
	return lv_db.GetMasterGorm().Delete(e).Error
}
