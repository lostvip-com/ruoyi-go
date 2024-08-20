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
package constants

type IotPlatform string

const (
	IotPlatform_LocalIot    IotPlatform = "本地"
	IotPlatform_CustomerIot IotPlatform = "用户自定义"  //用户自定义
	IotPlatform_WinCLinkIot IotPlatform = "赢创万联"   //赢创万联
	IotPlatform_AliIot      IotPlatform = "阿里云"    //阿里
	IotPlatform_HuaweiIot   IotPlatform = "华为云"    //华为
	IotPlatform_TencentIot  IotPlatform = "腾讯云"    //腾讯
	IotPlatform_TuyaIot     IotPlatform = "涂鸦云"    //涂鸦
	IotPlatform_OneNetIot   IotPlatform = "OneNET" //中国移动
)

func (i IotPlatform) TransformToCloudInstancePlatform() string {
	switch i {
	case IotPlatform_WinCLinkIot:
		return CloudServiceWincLinkName
	case IotPlatform_AliIot:
		return CloudServiceAliyunName
	case IotPlatform_HuaweiIot:
		return CloudServiceHuaweiName
	case IotPlatform_TencentIot:
		return CloudServiceTencentName
	case IotPlatform_TuyaIot:
		return CloudServiceTuyaName
	case IotPlatform_OneNetIot:
		return CloudServiceOneNETName
	default:
		return ""
	}
}

type SpecsType string

const (
	SpecsTypeInt    SpecsType = "int"
	SpecsTypeFloat  SpecsType = "float"
	SpecsTypeText   SpecsType = "text"
	SpecsTypeDate   SpecsType = "date"
	SpecsTypeBool   SpecsType = "bool"
	SpecsTypeEnum   SpecsType = "enum"
	SpecsTypeStruct SpecsType = "struct"
	SpecsTypeArray  SpecsType = "array"
)

func (i SpecsType) AllowSendInEkuiper() bool {
	if i == SpecsTypeInt || i == SpecsTypeFloat || i == SpecsTypeText || i == SpecsTypeBool || i == SpecsTypeEnum {
		return true
	}
	return false
}

type ProductNodeType string

const (
	ProductNodeTypeUnKnow    ProductNodeType = "其他"
	ProductNodeTypeGateway   ProductNodeType = "网关"
	ProductNodeTypeDevice    ProductNodeType = "直连设备"
	ProductNodeTypeSubDevice ProductNodeType = "网关子设备"
)

type ProductStatus string

const (
	ProductRelease   ProductStatus = "已发布"
	ProductUnRelease ProductStatus = "未发布"
)

type ProductNetType string

const (
	ProductNetTypeOther    ProductNetType = "其他"
	ProductNetTypeCellular ProductNetType = "蜂窝"
	ProductNetTypeWifi     ProductNetType = "WIFI"
	ProductNetTypeEthernet ProductNetType = "以太网"
	ProductNetTypeNB       ProductNetType = "NB"
)

type ProductProtocol string

const (
	ProductProtocolMQTT  ProductProtocol = "MQTT"
	ProductProtocolCoAP  ProductProtocol = "CoAP"
	ProductProtocolLwM2M ProductProtocol = "LwM2M"
	ProductProtocolHttp  ProductProtocol = "HTTP"
	ProductProtocolOther ProductProtocol = "其他"
)

type TagName string

const (
	TagNameCustom TagName = "自定义"
	TagNameSystem TagName = "系统"
)

type EventType string

const (
	EventTypeInfo  EventType = "info"
	EventTypeAlert EventType = "alert"
	EventTypeError EventType = "error"
)

type CallType string

const (
	CallTypeSync  CallType = "SYNC"  //同步
	CallTypeAsync CallType = "ASYNC" //异步
)
