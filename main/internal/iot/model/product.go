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

import "main/internal/iot/pkg/constants"

type Product struct {
	Timestamps      `gorm:"embedded"`
	Id              string                    `gorm:"id;primaryKey;not null;type:string;size:255;comment:主键"`
	Name            string                    `gorm:"type:string;size:255;comment:名字"`
	Key             string                    `gorm:"type:string;size:255;comment:产品标识"`
	CloudProductId  string                    `gorm:"type:string;size:255;comment:云产品ID"`
	CloudInstanceId string                    `gorm:"index;type:string;size:255;comment:云实例ID"`
	Platform        constants.IotPlatform     `gorm:"type:string;size:255;comment:平台"`
	Protocol        string                    `gorm:"type:string;size:255;comment:协议"`
	NodeType        constants.ProductNodeType `gorm:"type:string;size:255;comment:节点类型"`
	NetType         constants.ProductNetType  `gorm:"type:string;size:255;comment:网络类型"`
	DataFormat      string                    `gorm:"type:string;size:255;comment:数据类型"`
	LastSyncTime    int64                     `gorm:"comment:最后一次同步时间"`
	Factory         string                    `gorm:"type:string;size:255;comment:工厂名称"`
	Description     string                    `gorm:"type:text;comment:描述"`
	Status          constants.ProductStatus   `gorm:"type:string;size:255;comment:产品状态"`
	Extra           MapStringString           `gorm:"type:string;size:255;comment:扩展字段"`
	Properties      []Properties              `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // 物模型的属性列表
	Events          []Events                  `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // 物模型的事件列表
	Actions         []Actions                 `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // 物模型的动作列表
}

func (d *Product) TableName() string {
	return "product"
}

func (d *Product) Get() interface{} {
	return *d
}
