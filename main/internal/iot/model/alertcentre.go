/*******************************************************************************
 * Copyright 2017.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package models

import (
	"database/sql/driver"
	"main/internal/iot/pkg/constants"
)

type AlertRule struct {
	Timestamps  `gorm:"embedded"`
	Id          string `gorm:"id;primaryKey"`
	Name        string
	DeviceId    string                    `gorm:"index"`
	AlertType   constants.AlertType       //告警类型
	AlertLevel  constants.AlertLevel      //告警级别
	Status      constants.RuleStatus      //状态 启动或者禁用
	Condition   constants.WorkerCondition //执行条件
	SubRule     SubRule
	Notify      Notify
	SilenceTime int64 //静默时间
	Description string
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
	return "alert_rule"
}

func (d *AlertRule) Get() interface{} {
	return *d
}

type AlertList struct {
	Timestamps  `gorm:"embedded"`
	Id          string                    `gorm:"id;primaryKey;not null;type:string;size:255;comment:主键"`
	AlertRuleId string                    `gorm:"type:string;size:255;comment:告警记录ID"`
	TriggerTime int64                     `gorm:"comment:触发时间"`
	AlertResult MapStringInterface        `json:"alert_result" gorm:"type:string;size:255;comment:告警内容"`
	AlertRule   AlertRule                 `gorm:"foreignKey:AlertRuleId"`
	Status      constants.AlertListStatus `json:"status" gorm:"type:string;size:50;comment:状态"`
	TreatedTime int64                     `gorm:"comment:处理时间"`
	Message     string                    `gorm:"type:text;comment:处理意见"`
	IsSend      bool                      `gorm:"comment:是否发送通知"`
}

func (d *AlertList) TableName() string {
	return "alert_list"
}

func (d *AlertList) Get() interface{} {
	return *d
}
