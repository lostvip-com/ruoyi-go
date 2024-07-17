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
package dtos

import (
	"github.com/winc-link/hummingbird/internal/models"
	"github.com/winc-link/hummingbird/internal/pkg/constants"
)

type ThingModelAddOrUpdateReq struct {
	Id             string `json:"id"`
	ProductId      string `json:"product_id"`
	ThingModelType string `json:"thing_model_type"`
	//ModelName      string                `json:"model_name"`
	Name        string                `json:"name"`
	Code        string                `json:"code"`
	Description string                `json:"description"`
	Tag         string                `json:"tag"`
	Property    *ThingModelProperties `json:"property"`
	Event       *ThingModelEvents     `json:"event"`
	Action      *ThingModelActions    `json:"action"`
}

type ThingModelProperties struct {
	AccessModel string              `json:"access_model"`
	Require     bool                `json:"require"`
	DataType    constants.SpecsType `json:"type"`
	TypeSpec    interface{}         `json:"specs"`
}

type ThingModelEventAction struct {
	Code     string              `json:"code"`
	Name     string              `json:"name"`
	DataType constants.SpecsType `json:"type"`
	TypeSpec interface{}         `json:"specs"`
}

type ThingModelEvents struct {
	EventType   string                  `json:"event_type"`
	OutPutParam []ThingModelEventAction `json:"output_param"`
}

type ThingModelActions struct {
	CallType    constants.CallType      `json:"call_type"`
	InPutParam  []ThingModelEventAction `json:"input_param"`
	OutPutParam []ThingModelEventAction `json:"output_param"`
}

type ThingModelDeleteReq struct {
	ThingModelId   string `json:"thing_model_id"`
	ThingModelType string `json:"thing_model_type"`
}

type SystemThingModelSearchReq struct {
	ModelName      string `schema:"modelName"`
	ThingModelType string `schema:"thingModelType"`
}

type OpenApiThingModelProperties struct {
	Id          string          `json:"id"`
	Name        string          `json:"name"`        // 属性名称
	Code        string          `json:"code"`        // 标识符
	AccessMode  string          `json:"access_mode"` // 数据传输类型
	Require     bool            `json:"require"`
	TypeSpec    models.TypeSpec `json:"type_spec"` // 数据属性
	Description string          `json:"description"`
}

type OpenApiThingModelEvents struct {
	Id           string              `json:"id"`
	EventType    string              `json:"event_type"`
	Name         string              `json:"name"` // 功能名称
	Code         string              `json:"code"` // 标识符
	Description  string              `json:"description"`
	Require      bool                `json:"require"`
	OutputParams models.OutPutParams `json:"output_params"`
}

type OpenApiThingModelServices struct {
	Id           string              `json:"id"`
	Name         string              `json:"name"` // 功能名称
	Code         string              `json:"code"` // 标识符
	Description  string              `json:"description"`
	Require      bool                `json:"require"`
	CallType     constants.CallType  `json:"call_type"`
	InputParams  models.InPutParams  `json:"input_params"`  // 输入参数
	OutputParams models.OutPutParams `json:"output_params"` // 输出参数
}

type OpenApiThingModelAddOrUpdateReq struct {
	ProductId  string                        `json:"product_id"`
	Properties []OpenApiThingModelProperties `json:"properties"`
	Events     []OpenApiThingModelEvents     `json:"events"`
	Services   []OpenApiThingModelServices   `json:"services"`
}

type OpenApiQueryThingModelReq struct {
	ProductId string `schema:"product_id,omitempty"`
}

type OpenApiQueryThingModel struct {
	Properties []OpenApiThingModelProperties `json:"properties"`
	Events     []OpenApiThingModelEvents     `json:"events"`
	Services   []OpenApiThingModelServices   `json:"services"`
}

type OpenApiThingModelDeleteReq struct {
	ProductId   string   `json:"product_id"`
	PropertyIds []string `json:"property_ids"`
	EventIds    []string `json:"event_ids"`
	ServiceIds  []string `json:"service_ids"`
}

type OpenApiSetDeviceThingModel struct {
	DeviceId string                 `json:"deviceId"`
	Item     map[string]interface{} `json:"item"`
}

type OpenApiQueryDevicePropertyData struct {
}
