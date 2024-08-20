/*******************************************************************************
 * 品类
 *
 *******************************************************************************/
package model

import "common/model"

type CategoryTemplate struct {
	Id              int64  `gorm:"id;primaryKey;not null;size:32;comment:主键"`
	CategoryKey     string `json:"category_key" gorm:"size:32;comment:品类标识"`
	CategoryName    string `json:"category_name" gorm:"size:32;comment:品类名字"`
	Scene           string `json:"scene" gorm:"size:32;comment:场景"`
	model.BaseModel `gorm:"embedded"`
}

func (d *CategoryTemplate) TableName() string {
	return "iot_category"
}

func (d *CategoryTemplate) Get() interface{} {
	return *d
}
