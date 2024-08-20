/*******************************************************************************
 * IOT数据来源
 *
 *******************************************************************************/

package model

import (
	"common/model"
	"main/internal/iot_dev/pkg/constants"
)

// DataResource 数据源
type DataResource struct {
	Id     int64                      `json:"id" gorm:"id;primaryKey;not null;size:32;comment:主键"`
	Name   string                     `json:"name" gorm:"size:32;comment:名字"`
	Type   constants.DataResourceType `json:"type"  gorm:"size:50;comment:类型"`
	Health bool                       `json:"health" gorm:"comment:验证"`
	// {"bodyType":"json","headers":{"test":"test","test1":"test1"},"method":"get","sendSingle":true,"timeout":100000,"url":"http://www.baidu.com"}
	Option          MapStringInterface `json:"option" gorm:"comment:资源内容"`
	model.BaseModel `gorm:"embedded"`
}

func (d *DataResource) TableName() string {
	return "iot_data_resource"
}

func (d *DataResource) Get() interface{} {
	return *d
}
