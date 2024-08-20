package dtos

import (
	"encoding/json"
	models "main/internal/iot_dev/model"
	"main/internal/iot_dev/pkg/constants"
	"main/internal/iot_dev/pkg/utils"
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
