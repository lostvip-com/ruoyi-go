package model

import (
	"common/model"
	"database/sql/driver"
	"main/internal/iot_dev/pkg/constants"
	"main/internal/iot_dev/pkg/timer/jobs"
)

type Scene struct {
	Id          int64                 `json:"id" gorm:"id;primaryKey;not null;size:32;comment:主键"`
	Name        string                `json:"name" gorm:"size:32;comment:名字"`
	Description string                `json:"description" gorm:"type:text;comment:描述"`
	Status      constants.SceneStatus `json:"status" gorm:"size:50;comment:状态"`
	Conditions  Conditions            `json:"conditions" gorm:"type:text;comment:条件"`
	Actions     Actions2              `json:"actions" gorm:"type:text;comment:动作"`

	model.BaseModel `gorm:"embedded"`
}

func (d *Scene) TableName() string {
	return "iot_scene"
}

func (d *Scene) Get() interface{} {
	return *d
}

func (d *Scene) ToRuntimeJob() (schedule *jobs.JobSchedule, err error) {
	var (
		rj = jobs.RuntimeJobStu{
			JobID:       d.Id,
			JobName:     d.Name,
			Description: d.Description,
			Status:      string(d.Status),
		}
	)

	rj.TimeData = jobs.TimeData{
		Expression: d.Conditions[0].Option["cron_expression"],
	}

	for _, action := range d.Actions {
		rj.JobData.ActionData = append(rj.JobData.ActionData, jobs.DeviceMeta{
			ProductId:   action.ProductID,
			ProductName: action.ProductName,
			DeviceId:    action.DeviceID,
			DeviceName:  action.DeviceName,
			Code:        action.Code,
			DateType:    action.DataType,
			Value:       action.Value,
		})
	}
	return jobs.NewJobSchedule(&rj)
}

type Conditions []Condition

// Condition [{"condition_type":"timer","option":{"cron_expression":"00 07 * * ","decide_condition":" temp > 100 "}}]
type Condition struct {
	ConditionType string          `json:"condition_type"`
	Option        MapStringString `json:"option"`
}

func (c Conditions) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *Conditions) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}

type Actions2 []Action

type Action struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	DeviceID    string `json:"device_id"`
	DeviceName  string `json:"device_name"`
	Code        string `json:"code"`
	DataType    string `json:"data_type"`
	Value       string `json:"value"`
}

func (c Actions2) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *Actions2) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}
