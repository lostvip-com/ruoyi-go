package model

import (
	"common/model"
	"database/sql/driver"
	"main/internal/iot_dev/pkg/constants"
)

type AlertRule struct {
	Id              int64 `gorm:"type:bigint;primaryKey"`
	Name            string
	DeviceId        int64                     `gorm:"type:bigint;index"`
	AlertType       constants.AlertType       //告警类型
	AlertLevel      constants.AlertLevel      //告警级别
	Status          constants.RuleStatus      //状态 启动或者禁用
	Condition       constants.WorkerCondition //执行条件
	SubRule         SubRule
	Notify          Notify
	SilenceTime     int64 //静默时间
	Description     string
	model.BaseModel `gorm:"embedded"`
}

func (a *AlertRule) EkuiperRule() bool {
	if len(a.SubRule) > 0 {
		if a.SubRule[0].Trigger == constants.DeviceDataTrigger {
			return true
		}
	}
	return false
}

type SubRule []Rule

type Rule struct {
	Trigger   constants.Trigger `json:"trigger"` //触发方式
	ProductId string            `json:"product_id"`
	DeviceId  string            `json:"device_id"`
	Option    MapStringString   `json:"option"`
}

func (c SubRule) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *SubRule) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}

type Notify []SubNotify

type SubNotify struct {
	Name            constants.AlertWay `json:"name"` //告警方式
	Option          MapStringString    `json:"option"`
	StartEffectTime string             `json:"start_effect_time"` //生效开始时间
	EndEffectTime   string             `json:"end_effect_time"`   //生效结束时间
}

func (c Notify) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *Notify) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}

func (d *AlertRule) TableName() string {
	return "iot_alert_rule"
}

func (d *AlertRule) Get() interface{} {
	return *d
}
