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

import models "main/internal/iot/model"

type CategoryTemplateRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	CategoryName             string `schema:"category_name"`
	Scene                    string `schema:"scene"`
}

type CategoryTemplateSyncRequest struct {
	VersionName string `json:"version_name"`
}

type ThingModelTemplateSyncRequest struct {
	VersionName string `json:"version_name"`
}

func CategoryTemplateResponseFromModel(m models.CategoryTemplate) CategoryTemplateResponse {
	return CategoryTemplateResponse{
		Id:           m.Id,
		CategoryName: m.CategoryName,
		CategoryKey:  m.CategoryKey,
		Scene:        m.Scene,
	}
}

type CategoryTemplateResponse struct {
	Id           string `json:"id"`
	CategoryName string `json:"category_name"` //品类名称
	CategoryKey  string `json:"category_key"`
	Scene        string `json:"scene"` //所属场景
}

type CosCategoryTemplateResponse struct {
	CategoryName string `json:"category_name"`
	CategoryKey  string `json:"category_key"`
	Scene        string `json:"scene"`
}

type ThingModelTemplateRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	CategoryKey              string `schema:"category_key"`
	CategoryName             string `schema:"category_name"`
}

type ThingModelTemplateResponse struct {
	Id             string `json:"id"`
	CategoryName   string `json:"category_name"` //品类名称
	CategoryKey    string `json:"category_key"`
	ThingModelJSON string `json:"thing_model_json"`
	//models.Properties
	Properties interface{} `json:"properties"`
	Events     interface{} `json:"events"`
	Actions    interface{} `json:"actions"`
}

func ThingModelTemplateResponseFromModel(m models.ThingModelTemplate) ThingModelTemplateResponse {
	p, e, a := GetModelPropertyEventActionByThingModelTemplate(m.ThingModelJSON)
	return ThingModelTemplateResponse{
		CategoryKey:  m.CategoryKey,
		CategoryName: m.CategoryName,
		Properties:   p,
		Events:       e,
		Actions:      a,
	}
}

type CosThingModelTemplateResponse struct {
	CategoryName   string `json:"category_name"`
	CategoryKey    string `json:"category_key"`
	ThingModelJSON string `json:"thing_model_json"`
}
