/*******************************************************************************
 * Copyright 2017 Dell Inc.
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

type ThingModelDataBaseRequest struct {
	First bool    `json:"first"`
	Last  bool    `json:"last"`
	Range []int64 `json:"range"`
}

type ThingModelPropertyDataRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	ThingModelDataBaseRequest
	DeviceId string ` json:"deviceId"`
	Code     string `json:"code"`
}

type ThingModelEventDataRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	ThingModelDataBaseRequest
	DeviceId  string ` json:"deviceId"`
	EventCode string `json:"eventCode"`
	EventType string `json:"eventType"`
}

type ThingModelEventDataResponse struct {
	EventCode  string                 `json:"event_code"`
	EventType  string                 `json:"event_type"`
	OutputData map[string]interface{} `json:"output_data"`
	ReportTime int64                  `json:"report_time"`
	Name       string                 `json:"name"`
}

type ThingModelEventDataResponseArray []ThingModelEventDataResponse

func (array ThingModelEventDataResponseArray) Len() int {
	return len(array)
}

func (array ThingModelEventDataResponseArray) Less(i, j int) bool {
	return array[i].ReportTime > array[j].ReportTime //从小到大， 若为大于号，则从大到小
}

func (array ThingModelEventDataResponseArray) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

type ThingModelServiceDataResponse struct {
	ReportTime  int64                  `json:"report_time"`
	Code        string                 `json:"code"`
	ServiceName string                 `json:"service_name"`
	InputData   map[string]interface{} `json:"input_data"`
	OutputData  map[string]interface{} `json:"output_data"`
}

type ThingModelServiceDataResponseArray []ThingModelServiceDataResponse

func (array ThingModelServiceDataResponseArray) Len() int {
	return len(array)
}

func (array ThingModelServiceDataResponseArray) Less(i, j int) bool {
	return array[i].ReportTime > array[j].ReportTime //从小到大， 若为大于号，则从大到小
}

func (array ThingModelServiceDataResponseArray) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

type ThingModelDataResponse struct {
	ReportData
	Code       string `json:"code"`
	DataType   string `json:"data_type"`
	Unit       string `json:"unit"`
	Name       string `json:"name"`
	AccessMode string `json:"access_mode"`
}

type ThingModelPropertyDataResponse struct {
	ReportData interface{} `json:"report_data"`
	Code       string      `json:"code"`
	DataType   string      `json:"data_type"`
	Unit       string      `json:"unit"`
	Name       string      `json:"name"`
	AccessMode string      `json:"access_mode"`
}

type ThingModelServiceDataRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	ThingModelDataBaseRequest
	DeviceId string ` json:"deviceId"`
	Code     string `json:"code"`
}
