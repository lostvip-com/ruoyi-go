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

type UnitRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	UnitName                 string `schema:"unitName" json:"unitName"`
}

type UnitResponse struct {
	Id       string `json:"id"`
	Symbol   string `json:"symbol"`
	UnitName string `json:"unit_name"`
}

type CosUnitTemplateResponse struct {
	UnitName string `json:"Name"`
	Symbol   string `json:"Symbol"`
}

func UnitTemplateResponseFromModel(unitModel models.Unit) UnitResponse {
	return UnitResponse{
		Id:       unitModel.Id,
		Symbol:   unitModel.Symbol,
		UnitName: unitModel.UnitName,
	}
}

type UnitTemplateSyncRequest struct {
	VersionName string `json:"version_name"`
}
