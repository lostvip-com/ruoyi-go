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
package model

import "common/model"

type ThingModelTemplate struct {
	Id              int64  `gorm:"id;primaryKey;not null;size:32;comment:主键"`
	CategoryKey     string `gorm:"size:32;comment:品类key"`
	CategoryName    string `gorm:"size:32;comment:品类名字"`
	ThingModelJSON  string `gorm:"type:text;comment:物模型信息"`
	model.BaseModel `gorm:"embedded"`
}

func (d *ThingModelTemplate) TableName() string {
	return "iot_thing_model_tpl"
}

func (d *ThingModelTemplate) Get() interface{} {
	return *d
}
