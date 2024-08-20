package model

import "common/model"

type AlertType int

type AlertLevel int

type AlertContent struct {
	ServiceName     string     `json:"name"`                        // 服务名
	Type            AlertType  `json:"type" binding:"oneof=1 2"`    // 告警类型
	Level           AlertLevel `json:"level" binding:"oneof=1 2 3"` // 告警级别
	Time            int64      `json:"time"`                        // 告警时间
	Content         string     `json:"content"`                     // 告警内容
	model.BaseModel `gorm:"embedded"`
}

func (table *AlertContent) TableName() string {
	return "iot_alert_content"
}

func (table *AlertContent) Get() interface{} {
	return *table
}
