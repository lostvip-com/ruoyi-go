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

type LanguageSDK struct {
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Sort        int    `json:"sort"`
	Addr        string `json:"addr"`
	Description string `json:"description"`
}

type LanguageSDKSyncRequest struct {
	VersionName string `json:"version_name"`
}

type LanguageSDKSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
}

type LanguageSDKSearchResponse struct {
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Addr        string `json:"addr"`
	Description string `json:"description"`
}
