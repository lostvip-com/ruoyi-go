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

type SceneAddRequest struct {
	Name        string `json:"name"`        //名字
	Description string `json:"description"` //描述
}

type SceneUpdateRequest struct {
	Id         string      `json:"id"`
	Conditions []Condition `json:"conditions"`
	Actions    []Action    `json:"actions"`
}

func ReplaceSceneModelFields(scene *models.Scene, req SceneUpdateRequest) {
	//scene.Conditions = req.Conditions

	var modelConditions models.Conditions
	for _, condition := range req.Conditions {
		modelConditions = append(modelConditions, models.Condition{
			ConditionType: condition.ConditionType,
			Option:        condition.Option,
		})
	}
	scene.Conditions = modelConditions

	var modelAction models.Actions2
	for _, action := range req.Actions {
		modelAction = append(modelAction, models.Action{
			ProductName: action.ProductName,
			ProductID:   action.ProductID,
			DeviceName:  action.DeviceName,
			DeviceID:    action.DeviceID,
			Code:        action.Code,
			DataType:    action.DataType,
			Value:       action.Value,
		})
	}
	scene.Actions = modelAction
}

type Condition struct {
	ConditionType string            `json:"condition_type"`
	Option        map[string]string `json:"option"`
}

type Action struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	DeviceID    string `json:"device_id"`
	DeviceName  string `json:"device_name"`
	Code        string `json:"code"`
	DataType    string `json:"data_type"`
	Value       string `json:"value"`
}

type SceneSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	Name                     string `json:"name"`
	Status                   string `json:"status"`
}

type SceneLogSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	StartAt                  int64  `schema:"start_time"`
	EndAt                    int64  `schema:"end_time"`
	SceneId                  string `json:"scene_id"`
}
