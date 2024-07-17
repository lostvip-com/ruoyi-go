/*******************************************************************************
 * Copyright 2017 Dell Inc.
 * Copyright (c) 2019 Intel Corporation
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
	"gorm.io/gorm"
	"main/internal/iot/pkg/constants"
)

type MqttAuth struct {
	Timestamps   `gorm:"embedded"`
	Id           string                 `gorm:"id;primaryKey;not null;type:string;size:255;comment:主键"`
	ResourceId   string                 `gorm:"type:string;size:255;comment:资源ID"`
	ResourceType constants.ResourceType `gorm:"type:string;size:255;comment:资源类型"`
	ClientId     string                 `gorm:"uniqueIndex;size:255;comment:客户端ID"`
	UserName     string                 `gorm:"type:string;size:255;comment:用户名"`
	Password     string                 `gorm:"type:string;size:255;comment:密码"`
}

func (d *MqttAuth) TableName() string {
	return "mqtt_auth"
}

func (d *MqttAuth) Get() interface{} {
	return *d
}

func (d *MqttAuth) BeforeCreate(tx *gorm.DB) (err error) {

	return nil
}
