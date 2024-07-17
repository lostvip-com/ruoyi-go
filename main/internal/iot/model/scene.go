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
	"main/internal/iot/pkg/timer/jobs"
)

type Scene struct {
	Timestamps  `gorm:"embedded"`
	Id          string                `json:"id" gorm:"id;primaryKey;not null;type:string;size:255;comment:主键"`
	Name        string                `json:"name" gorm:"type:string;size:255;comment:名字"`
	Description string                `json:"description" gorm:"type:text;comment:描述"`
	Status      constants.SceneStatus `json:"status" gorm:"type:string;size:50;comment:状态"`
	Conditions  Conditions            `json:"conditions" gorm:"type:text;comment:条件"`
	Actions     Actions2              `json:"actions" gorm:"type:text;comment:动作"`
}

func (d *Scene) TableName() string {
	return "scene"
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
