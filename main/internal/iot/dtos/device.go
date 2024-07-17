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
	"encoding/json"
	models "main/internal/iot/model"
	"main/internal/iot/pkg/constants"
	"main/internal/iot/pkg/utils"
)

type DeviceSyncRequest struct {
	CloudInstanceId string `json:"cloud_instance_id"`
	DriveInstanceId string `json:"driver_instance_id"`
}

type DeviceSyncByIdRequest struct {
	DeviceId string `json:"device_id"`
}

type DeviceStatusRequest struct {
	DeviceId string `json:"device_id"`
}

type DeviceSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	Platform                 string `schema:"platform,omitempty"`
	Name                     string `schema:"name,omitempty"`
	ProductId                string `schema:"product_id,omitempty"`
	CloudProductId           string `schema:"cloud_product_id,omitempty"`
	CloudInstanceId          string `schema:"cloud_instance_id,omitempty"`
	DriveInstanceId          string `schema:"drive_instance_id,omitempty"`
	Status                   string `schema:"status,omitempty"`
}

type DeviceSearchQueryResponse struct {
	Id                string                 `json:"id"`
	Name              string                 `json:"name"`
	ProductId         string                 `json:"product_id"`
	Status            constants.DeviceStatus `json:"status"`
	Platform          constants.IotPlatform  `json:"platform"`
	CloudInstanceId   string                 `json:"cloud_instance_id"`
	CloudProductId    string                 `json:"cloud_product_id"`
	DriverServiceName string                 `json:"driver_service_name"`
	ProductName       string                 `json:"product_name"`
	LastSyncTime      int64                  `json:"last_sync_time"`
	LastOnlineTime    int64                  `json:"last_online_time"`
	DriveInstanceId   string                 `json:"drive_instance_id"`
	Created           int64                  `json:"created"`
	Description       string                 `json:"description"`
}

func DeviceResponseFromModel(p models.Device, deviceServiceName string) DeviceSearchQueryResponse {
	return DeviceSearchQueryResponse{
		Id:                p.Id,
		ProductId:         p.ProductId,
		Name:              p.Name,
		Platform:          p.Platform,
		Status:            p.Status,
		DriverServiceName: deviceServiceName,
		CloudInstanceId:   p.CloudInstanceId,
		CloudProductId:    p.CloudProductId,
		ProductName:       p.Product.Name,
		LastSyncTime:      p.LastSyncTime,
		LastOnlineTime:    p.LastOnlineTime,
		DriveInstanceId:   p.DriveInstanceId,
		Created:           p.Created,
		Description:       p.Description,
	}
}

type OpenApiDeviceStatus struct {
	Status constants.DeviceStatus `json:"status"`
}

type OpenApiDeviceInfoResponse struct {
	Id          string                 `json:"id"`
	Name        string                 `json:"name"`
	Platform    constants.IotPlatform  `json:"platform"`
	Status      constants.DeviceStatus `json:"status"`
	Description string                 `json:"description"`
	ProductId   string                 `json:"product_id"`
	ProductName string                 `json:"product_name"`
	//Secret         string                 `json:"secret"`
	LastOnlineTime int64 `json:"last_online_time"`
	Created        int64 `json:"created_at"`
}

func OpenApiDeviceInfoResponseFromModel(p models.Device) OpenApiDeviceInfoResponse {
	return OpenApiDeviceInfoResponse{
		Id:          p.Id,
		Name:        p.Name,
		Platform:    p.Platform,
		Status:      p.Status,
		Description: p.Description,
		ProductId:   p.ProductId,
		ProductName: p.Product.Name,
		//Secret:         p.Secret,
		LastOnlineTime: p.LastOnlineTime,
		Created:        p.Created,
	}
}

type DeviceInfoResponse struct {
	Id                string                 `json:"id"`
	CloudDeviceId     string                 `json:"cloud_device_id"`
	CloudProductId    string                 `json:"cloud_product_id"`
	CloudInstanceId   string                 `json:"cloud_instance_id"`
	Name              string                 `json:"name"`
	Status            constants.DeviceStatus `json:"status"`
	Description       string                 `json:"description"`
	ProductId         string                 `json:"product_id"`
	ProductName       string                 `json:"product_name"`
	Secret            string                 `json:"secret"`
	Platform          constants.IotPlatform  `json:"platform"`
	DeviceServiceName string                 `json:"device_service_name"`
	LastSyncTime      int64                  `json:"last_sync_time"`
	LastOnlineTime    int64                  `json:"last_online_time"`
	Created           int64                  `json:"create_at"`
}

func DeviceInfoResponseFromModel(p models.Device, deviceServiceName string) DeviceInfoResponse {
	return DeviceInfoResponse{
		Id:                p.Id,
		CloudDeviceId:     p.CloudDeviceId,
		CloudProductId:    p.CloudProductId,
		Name:              p.Name,
		Status:            p.Status,
		Description:       p.Description,
		ProductId:         p.ProductId,
		ProductName:       p.Product.Name,
		Secret:            p.Secret,
		Platform:          p.Platform,
		DeviceServiceName: deviceServiceName,
		LastSyncTime:      p.LastSyncTime,
		LastOnlineTime:    p.LastOnlineTime,
		Created:           p.Created,
		CloudInstanceId:   p.CloudInstanceId,
	}
}

type DeviceReportPropertiesValueSearchRequest struct {
	DeviceId string `json:"device_id"`
}

type PropertyInfo struct {
	Code     string `json:"code,omitempty"`
	Value    string `json:"value,omitempty"`
	DataType string `json:"dataType,omitempty"`
	Time     string `json:"time,omitempty"`
	Unit     string `json:"unit,omitempty"`
	Name     string `json:"name,omitempty"`
}

type DeviceReportPropertiesValueSearchResponse struct {
	PropertyInfoList []PropertyInfo `json:"property_info_list"`
}

type DeviceAddRequest struct {
	DeviceId         string                `json:"device_id"`
	Name             string                `json:"name"`
	ProductId        string                `json:"product_id"`
	Description      string                `json:"description"`
	Platform         constants.IotPlatform `json:"platform"`
	DriverInstanceId string                `json:"driver_instance_id"`
	//CloudDeviceId   string                 `json:"cloud_device_id"`
	//CloudProductId  string                 `json:"cloud_product_id"`
	//CloudInstanceId string                 `gorm:"index"`
	//Status          constants.DeviceStatus `json:"status"`
	//Secret          string                 `json:"secret"`
}

type DeviceAuthInfoResponse struct {
	ClientId string `json:"clientId"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"mqttHostUrl"`
	Port     int    `json:"port"`
}

func DeviceAuthInfoResponseFromModel(p models.MqttAuth) DeviceAuthInfoResponse {
	ip, _ := utils.GetOutBoundIP()
	return DeviceAuthInfoResponse{
		ClientId: p.ClientId,
		UserName: p.UserName,
		Password: p.Password,
		Host:     ip,
		Port:     58090,
	}
}

type DeviceUpdateRequest struct {
	Id              string  `json:"id"`
	Description     *string `json:"description"`
	Name            *string `json:"name"`
	InstallLocation *string `json:"install_location"`
	DriveInstanceId *string `json:"drive_instance_id"`
}

func ReplaceDeviceModelFields(ds *models.Device, patch DeviceUpdateRequest) {
	if patch.Description != nil {
		ds.Description = *patch.Description
	}
	if patch.Name != nil {
		ds.Name = *patch.Name
	}
	if patch.DriveInstanceId != nil {
		ds.DriveInstanceId = *patch.DriveInstanceId
	}

	if patch.InstallLocation != nil {
		ds.InstallLocation = *patch.InstallLocation
	}
}

type DeviceUpdateOrCreateCallBack struct {
	Id              string                 `json:"id"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	ProductId       string                 `json:"product_id"`
	Status          constants.DeviceStatus `json:"status"`
	Platform        constants.IotPlatform  `json:"platform"`
	DriveInstanceId string                 `json:"drive_instance_id"`
}

type DeviceDeleteCallBack struct {
	Id              string `json:"id"`
	DriveInstanceId string `json:"drive_instance_id"`
}

type DeviceImportTemplateRequest struct {
}

type DevicesImport struct {
	ProductId        string `schema:"product_id,omitempty"`
	DriverInstanceId string `schema:"driver_instance_id,omitempty"`
}

type DeviceBatchDelete struct {
	DeviceIds []string `json:"device_ids"`
}

type DevicesBindDriver struct {
	DeviceIds        []string `json:"device_ids"`
	DriverInstanceId string   `json:"driver_instance_id,omitempty"`
}

type DevicesBindProductId struct {
	ProductId        string `json:"product_id"`
	DriverInstanceId string `json:"driver_instance_id,omitempty"`
}

type DevicesUnBindDriver struct {
	DeviceIds []string `json:"device_ids"`
}

type AddMqttAuthInfoRequest struct {
	Id           string `json:"id"`
	ClientId     string `json:"client_id"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	ResourceId   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}

type DeviceExecRes struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

func (d *DeviceExecRes) ToString() string {
	b, _ := json.Marshal(d)
	return string(b)
}

type JobAction struct {
	ActionType  string      `json:"actionType"`
	ProductId   string      `json:"productId"`
	ProductName string      `json:"product_name"`
	DeviceId    string      `json:"deviceId"`
	DeviceName  string      `json:"deviceName"`
	Code        string      `json:"code"`
	DateType    string      `json:"dateType"`
	Value       interface{} `json:"value"`
}

type InvokeDeviceServiceReq struct {
	DeviceId string                 `json:"deviceId"`
	Code     string                 `json:"code"`
	Items    map[string]interface{} `json:"inputParams"`
}

type DeviceEffectivePropertyDataReq struct {
	DeviceId string   `json:"deviceId"`
	Codes    []string `json:"codes"`
}

type DeviceEffectivePropertyDataResponse struct {
	Data []EffectivePropertyData `json:"propertyInfo"`
}

type EffectivePropertyData struct {
	Code  string      `json:"code"`
	Value interface{} `json:"value"`
	Time  int64       `json:"time"`
}
