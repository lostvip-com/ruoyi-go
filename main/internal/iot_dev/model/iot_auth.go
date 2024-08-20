package model

import (
	"common/model"
	"gorm.io/gorm"
	"main/internal/iot_dev/pkg/constants"
)

type MqttAuth struct {
	Id              string                 `gorm:"id;primaryKey;not null;type:string;size:32;comment:主键"`
	ResourceId      string                 `gorm:"type:string;size:32;comment:资源ID"`
	ResourceType    constants.ResourceType `gorm:"type:string;size:32;comment:资源类型"`
	ClientId        string                 `gorm:"uniqueIndex;size:32;comment:客户端ID"`
	UserName        string                 `gorm:"type:string;size:32;comment:用户名"`
	Password        string                 `gorm:"type:string;size:32;comment:密码"`
	model.BaseModel `gorm:"embedded"`
}

func (d *MqttAuth) TableName() string {
	return "iot_auth"
}

func (d *MqttAuth) Get() interface{} {
	return *d
}

func (d *MqttAuth) BeforeCreate(tx *gorm.DB) (err error) {

	return nil
}
