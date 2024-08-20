package model

import (
	"common/model"
	"main/internal/iot_dev/pkg/constants"
)

type AlertList struct {
	model.BaseModel `gorm:"embedded"`
	Id              int64                     `gorm:"type:bigint;primaryKey;not null;type:string;comment:主键"`
	AlertRuleId     int64                     `gorm:"type:bigint;type:string;comment:告警记录ID"`
	TriggerTime     string                    `gorm:"comment:触发时间"`
	AlertResult     MapStringInterface        `json:"alert_result" gorm:"type:string;size:32;comment:告警内容"`
	AlertRule       AlertRule                 `gorm:"foreignKey:AlertRuleId"`
	Status          constants.AlertListStatus `json:"status" gorm:"type:string;size:50;comment:状态"`
	TreatedTime     string                    `gorm:"comment:处理时间"`
	Message         string                    `gorm:"type:text;comment:处理意见"`
	IsSend          bool                      `gorm:"comment:是否发送通知"`
}

func (d *AlertList) TableName() string {
	return "iot_alert_list"
}

func (d *AlertList) Get() interface{} {
	return *d
}
