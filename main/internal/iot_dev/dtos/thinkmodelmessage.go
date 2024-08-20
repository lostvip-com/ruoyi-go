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

import (
	"encoding/json"
)

type MessageBus struct {
	DeviceId    string      `json:"deviceId"`
	MessageType string      `json:"messageType"`
	Data        interface{} `json:"data"`
}

type DevicePropertyReport struct {
	MsgId   string `json:"msgId"`
	Version string `json:"version"`
	//Time    int64  `json:"time"`
	Sys struct {
		Ack int `json:"ack"`
	} `json:"sys"`
	Data map[string]ReportData `json:"data"`
}

type DeviceGetPropertyResponse struct {
	MsgId string                  `json:"msgId"`
	Data  []EffectivePropertyData `json:"data"`
}

type DevicePropertySetResponse struct {
	MsgId string                `json:"msgId"`
	Data  DevicePropertySetData `json:"data"`
}

type DevicePropertySetData struct {
	ErrorMessage string `json:"errorMessage"`
	Code         uint32 `json:"code"`
	Success      bool   `json:"success"`
}

type DeviceBatchReport struct {
	MsgId   string `json:"msgId"`
	Version string `json:"version"`
	Time    int64  `json:"time"`
	Sys     struct {
		Ack int `json:"ack"`
	} `json:"sys"`
	Data BatchData `json:"data"`
}

type BatchData struct {
	Properties map[string]BatchProperty `json:"properties"`
	Events     map[string]BatchEvent    `json:"events"`
}

type BatchProperty struct {
	Value interface{} `json:"value"`
	//Time  int64       `json:"time"`
}
type BatchEvent struct {
	//EventTime    int64                  `json:"eventTime"`
	OutputParams map[string]interface{} `json:"outputParams"`
}

type ReportData struct {
	Value interface{} `json:"value"`
	Time  int64       `json:"time"`
}

func (r *ReportData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type EdgeXDeviceEventReport struct {
	MsgId   string `json:"msgId"`
	Version string `json:"version"`
	Sys     struct {
		Ack int `json:"ack"`
	} `json:"sys"`
	Data EventData `json:"data"`
}

type EventData struct {
	EventCode    string                 `json:"eventCode"`
	EventTime    int64                  `json:"eventTime"`
	OutputParams map[string]interface{} `json:"outputParams"`
}

func (r *EventData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// operationType:SERVICE_EXECUTE data:"topic:\"/sys/hq85KDcqGI6/test_008/thing/service/DeleteAlgorithmModel\"
// message:\"{\\\"id\\\":\\\"2130175210\\\",\\\"version\\\":\\\"1.0.0\\\",\\\"code\\\":\\\"DeleteAlgorithmModel\\\",\\\"params\\\":{\\\"ForceDelete\\\":3}}\""
type InvokeDeviceService struct {
	MsgId   string      `json:"msgId"`
	Version string      `json:"version"`
	Time    int64       `json:"time"`
	Data    ServiceData `json:"data"`
}

type ServiceData struct {
	Code        string                 `json:"code"`
	InputParams map[string]interface{} `json:"inputParams"`
}

func (r *ServiceData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *InvokeDeviceService) ToString() string {
	s, _ := json.Marshal(r)
	return string(s)
}

type DeviceGetPropertyData struct {
	MsgId   string   `json:"msgId"`
	Version string   `json:"version"`
	Time    int64    `json:"time"`
	Data    []string `json:"data"`
}

func (r *DeviceGetPropertyData) ToString() string {
	s, _ := json.Marshal(r)
	return string(s)
}

type SaveServiceIssueData struct {
	MsgId        string                 `json:"msgId"`
	Code         string                 `json:"code"`
	Time         int64                  `json:"time"`
	InputParams  map[string]interface{} `json:"inputParams"`
	OutputParams map[string]interface{} `json:"outputParams"`
}

func (r *SaveServiceIssueData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ServiceExecResponse struct {
	MsgId        string                 `json:"msgId"`
	Code         string                 `json:"code"`
	Time         int64                  `json:"time"`
	OutputParams map[string]interface{} `json:"data"`
}

func (r *ServiceExecResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PropertySet struct {
	Version string                 `json:"version"`
	MsgId   string                 `json:"msgId"`
	Time    int64                  `json:"time"`
	Params  map[string]interface{} `json:"data"`
}

func (r *PropertySet) ToString() string {
	s, _ := json.Marshal(r)
	return string(s)
}
