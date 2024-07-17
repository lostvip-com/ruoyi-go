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

type ThingModelTemplate struct {
	Timestamps     `gorm:"embedded"`
	Id             string `gorm:"id;primaryKey;not null;type:string;size:255;comment:主键"`
	CategoryKey    string `gorm:"type:string;size:255;comment:品类key"`
	CategoryName   string `gorm:"type:string;size:255;comment:品类名字"`
	ThingModelJSON string `gorm:"type:text;comment:物模型信息"`
}

func (d *ThingModelTemplate) TableName() string {
	return "thing_model_template"
}

func (d *ThingModelTemplate) Get() interface{} {
	return *d
}
