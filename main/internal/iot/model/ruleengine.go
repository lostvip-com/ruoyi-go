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

type RuleEngine struct {
	Timestamps     `gorm:"embedded"`
	Id             string                     `gorm:"id;primaryKey;not null;type:string;size:255;comment:主键"`
	Name           string                     `gorm:"type:string;size:255;comment:名字"`
	Description    string                     `gorm:"type:text;comment:描述"`
	Status         constants.RuleEngineStatus `gorm:"type:string;size:50;comment:状态"`
	Filter         Filter
	DataResourceId string       `gorm:"type:string;size:255;comment:资源ID"`
	DataResource   DataResource `gorm:"foreignKey:DataResourceId"`
}

func (d *RuleEngine) TableName() string {
	return "rule_engine"
}

func (d *RuleEngine) Get() interface{} {
	return *d
}

type Filter struct {
	MessageSource string `json:"message_source" gorm:"type:string;size:255;comment:消息源"`
	SelectName    string `json:"select_name" gorm:"type:string;size:255;comment:选择字段"`
	Condition     string `json:"condition" gorm:"type:string;size:255;comment:条件"`
	Sql           string `json:"sql" gorm:"type:string;size:255;comment:sql"`
}

func (c Filter) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *Filter) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}
