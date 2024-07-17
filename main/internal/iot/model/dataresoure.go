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

import "main/internal/iot/pkg/constants"

type DataResource struct {
	Timestamps `gorm:"embedded"`
	Id         string                     `json:"id" gorm:"id;primaryKey;not null;type:string;size:255;comment:主键"`
	Name       string                     `json:"name" gorm:"type:string;size:255;comment:名字"`
	Type       constants.DataResourceType `json:"type"  gorm:"type:string;size:50;comment:类型"`
	Health     bool                       `json:"health" gorm:"comment:验证"`
	Option     MapStringInterface         `json:"option" gorm:"type:text;comment:资源内容"`
}

func (d *DataResource) TableName() string {
	return "data_resource"
}

func (d *DataResource) Get() interface{} {
	return *d
}
