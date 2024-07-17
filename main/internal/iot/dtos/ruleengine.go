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

package dtos

import (
	"github.com/winc-link/hummingbird/internal/models"
)

type RuleEngineRequest struct {
	Name           string `json:"name"`        //名字
	Description    string `json:"description"` //描述
	Filter         Filter `json:"filter"`
	DataResourceId string `json:"data_resource_id"`
}

func (r RuleEngineRequest) BuildEkuiperSql() string {
	return r.Filter.Sql
}

type Filter struct {
	MessageSource string `json:"message_source"`
	SelectName    string `json:"select_name"`
	Condition     string `json:"condition"`
	Sql           string `json:"sql"`
}

type RuleEngineUpdateRequest struct {
	Id             string  `json:"id"`
	Name           *string `json:"name"`        //名字
	Description    *string `json:"description"` //描述
	Filter         *Filter `json:"filter"`
	DataResourceId *string `json:"data_resource_id"`
}

func ReplaceRuleEngineModelFields(ds *models.RuleEngine, patch RuleEngineUpdateRequest) {
	if patch.Name != nil {
		ds.Name = *patch.Name
	}
	if patch.Description != nil {
		ds.Description = *patch.Description
	}
	if patch.Filter != nil {
		ds.Filter = models.Filter(*patch.Filter)
	}
	if patch.DataResourceId != nil {
		ds.DataResourceId = *patch.DataResourceId
	}

}

func (r RuleEngineUpdateRequest) BuildEkuiperSql() string {
	return r.Filter.Sql
}

type RuleEngineFieldUpdateRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	//AlertLevel  constants.AlertLevel `json:"alert_level"`
	Description string `json:"description"`
}

type RuleEngineResponse struct {
	Id             string           `json:"id"`
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	Filter         Filter           `json:"filter"`
	Created        int64            `json:"created"`
	DataResourceId string           `json:"data_resource_id"`
	DataResource   DataResourceInfo `json:"dataResource"`
	Modified       int64            `json:"modified"`
}

type RuleEngineSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	Name                     string `schema:"name,omitempty"`
	Status                   string `schema:"status,omitempty"`
}

type RuleEngineSearchQueryResponse struct {
	Id           string           `json:"id"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Created      int64            `json:"created"`
	Status       string           `json:"status"`
	ResourceType string           `json:"resource_type"`
	DataResource DataResourceInfo `json:"dataResource"`
}

func RuleEngineSearchQueryResponseFromModel(p models.RuleEngine) RuleEngineSearchQueryResponse {
	var dataResource DataResourceInfo
	dataResource.Name = p.DataResource.Name
	dataResource.Type = string(p.DataResource.Type)
	dataResource.Option = p.DataResource.Option
	return RuleEngineSearchQueryResponse{
		Id:           p.Id,
		Name:         p.Name,
		Description:  p.Description,
		Created:      p.Created,
		Status:       string(p.Status),
		DataResource: dataResource,
	}
}
