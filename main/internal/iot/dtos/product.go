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
	"github.com/winc-link/hummingbird/internal/models"
	"github.com/winc-link/hummingbird/internal/pkg/constants"
)

type ProductSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	Platform                 string `schema:"platform,omitempty"`
	Name                     string `schema:"name,omitempty"`
	ProductId                string `schema:"product_id,omitempty"`
	CloudInstanceId          string `schema:"cloud_instance_id,omitempty"`
	//DeviceLibraryId          string `schema:"deviceLibraryId,omitempty"`
}

type ProductSearchQueryResponse struct {
	Id           string `json:"id"`
	ProductId    string `json:"product_id"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	NodeType     string `json:"node_type"`
	Platform     string `json:"platform"`
	Status       string `json:"status"`
	CreatedAt    int64  `json:"created_at"`
	CategoryName string `json:"category_name"`
}

func ProductResponseFromModel(p models.Product) ProductSearchQueryResponse {
	return ProductSearchQueryResponse{
		Id:        p.Id,
		ProductId: p.CloudProductId,
		Name:      p.Name,
		Key:       p.Key,
		NodeType:  string(p.NodeType),
		Platform:  string(p.Platform),
		Status:    string(p.Status),
		CreatedAt: p.Created,
	}
}

type ProductSearchByIdResponse struct {
	Id              string      `json:"id"`
	Name            string      `json:"name"`
	Key             string      `json:"key"`
	CloudProductId  string      `json:"cloud_product_id"`
	CloudInstanceId string      `json:"cloud_instance_id"`
	Platform        string      `json:"platform"`
	Protocol        string      `json:"protocol"`
	NodeType        string      `json:"node_type"`
	NetType         string      `json:"net_type"`
	DataFormat      string      `json:"data_format"`
	Factory         string      `json:"factory"`
	Description     string      `json:"description"`
	Status          string      `json:"status"`
	CreatedAt       int64       `json:"created_at"`
	LastSyncTime    int64       `json:"last_sync_time"`
	Properties      interface{} `json:"properties"`
	Events          interface{} `json:"events"`
	Actions         interface{} `json:"actions"`
}

func ProductSearchByIdFromModel(p models.Product) ProductSearchByIdResponse {
	return ProductSearchByIdResponse{
		Id:              p.Id,
		Name:            p.Name,
		CloudProductId:  p.CloudProductId,
		CloudInstanceId: p.CloudInstanceId,
		Platform:        string(p.Platform),
		Protocol:        p.Protocol,
		Key:             p.Key,
		NodeType:        string(p.NodeType),
		NetType:         string(p.NetType),
		DataFormat:      p.DataFormat,
		Factory:         p.Factory,
		Description:     p.Description,
		CreatedAt:       p.Created,
		LastSyncTime:    p.LastSyncTime,
		Status:          string(p.Status),
		Properties:      p.Properties,
		Events:          p.Events,
		Actions:         p.Actions,
	}
}

type ProductSearchByIdOpenApiResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	Platform    string `json:"platform"`
	Protocol    string `json:"protocol"`
	NodeType    string `json:"node_type"`
	NetType     string `json:"net_type"`
	DataFormat  string `json:"data_format"`
	Factory     string `json:"factory"`
	Status      string `json:"status"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	//Properties  []OpenApiProperties `json:"properties"`
	//Events      []OpenApiEvents     `json:"events"`
	//Actions     []OpenApiActions    `json:"services"`
}

type OpenApiProperties struct {
	Id          string          `json:"id"`
	ProductId   string          `json:"product_id"`  // 产品ID
	Name        string          `json:"name"`        // 属性名称
	Code        string          `json:"code"`        // 标识符
	AccessMode  string          `json:"access_mode"` // 数据传输类型
	Required    bool            `json:"required"`
	TypeSpec    OpenApiTypeSpec `json:"type_spec"` // 数据属性
	Description string          `json:"description"`
	CreatedAt   int64           `json:"created_at"`
}

type OpenApiEvents struct {
	Id           string                `json:"id"`
	ProductId    string                `json:"product_id"`
	Name         string                `json:"name"` // 功能名称
	Code         string                `json:"code"` // 标识符
	EventType    string                `json:"event_type"`
	Required     bool                  `json:"required"`
	OutputParams []OpenApiOutPutParams `json:"output_params"`
	Description  string                `json:"description"`
	CreatedAt    int64                 `json:"created_at"`
}

type OpenApiOutPutParams struct {
	Code     string          `json:"code"`
	Name     string          `json:"name"`
	TypeSpec OpenApiTypeSpec `json:"type_spec"`
}

type OpenApiInPutParams struct {
	Code     string          `json:"code"`
	Name     string          `json:"name"`
	TypeSpec OpenApiTypeSpec `json:"type_spec"`
}

type OpenApiActions struct {
	Id           string                `json:"id"`
	ProductId    string                `json:"product_id"`
	Name         string                `json:"name"` // 功能名称
	Code         string                `json:"code"` // 标识符
	Required     bool                  `json:"required"`
	CallType     constants.CallType    `json:"call_type"`
	InputParams  []OpenApiInPutParams  `json:"input_params"`  // 输入参数
	OutputParams []OpenApiOutPutParams `json:"output_params"` // 输出参数
	CreatedAt    int64                 `json:"created_at"`
	Description  string                `json:"description"`
}

type OpenApiTypeSpec struct {
	Type  constants.SpecsType `json:"type,omitempty"`
	Specs string              `json:"specs,omitempty"`
}

func ProductSearchByIdOpenApiFromModel(p models.Product) ProductSearchByIdOpenApiResponse {
	return ProductSearchByIdOpenApiResponse{
		Id:          p.Id,
		Name:        p.Name,
		Key:         p.Key,
		Platform:    string(p.Platform),
		Protocol:    p.Protocol,
		NodeType:    string(p.NodeType),
		NetType:     string(p.NetType),
		DataFormat:  p.DataFormat,
		Factory:     p.Factory,
		Status:      string(p.Status),
		Description: p.Description,
		CreatedAt:   p.Created,
	}
}

type ProductSearchOpenApiResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	Platform    string `json:"platform"`
	Protocol    string `json:"protocol"`
	NodeType    string `json:"node_type"`
	NetType     string `json:"net_type"`
	DataFormat  string `json:"data_format"`
	Factory     string `json:"factory"`
	Status      string `json:"status"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

func ProductSearchOpenApiFromModel(p models.Product) ProductSearchOpenApiResponse {
	return ProductSearchOpenApiResponse{
		Id:          p.Id,
		Name:        p.Name,
		Key:         p.Key,
		Platform:    string(p.Platform),
		Protocol:    p.Protocol,
		NodeType:    string(p.NodeType),
		NetType:     string(p.NetType),
		DataFormat:  p.DataFormat,
		Factory:     p.Factory,
		Status:      string(p.Status),
		Description: p.Description,
		CreatedAt:   p.Created,
	}
}

type ProductSyncRequest struct {
	CloudInstanceId string `json:"cloud_instance_id"`
}

type ProductSyncByIdRequest struct {
	ProductId string `json:"product_id"`
}

type ProductAddRequest struct {
	Name string `json:"name"` //产品名字
	//Platform           string `json:"platform"`
	Key                string `json:"key"`
	CategoryTemplateId string `json:"category_template_id"` //如果是自定义 id固定传递"1"
	Protocol           string `json:"protocol"`             //协议
	NodeType           string `json:"node_type"`            //节点类型
	NetType            string `json:"net_type"`             //联网模式
	DataFormat         string `json:"data_format"`          //数据类型
	Factory            string `json:"factory"`              //厂家
	Description        string `json:"description"`          //描述
}

type OpenApiAddProductRequest struct {
	Name        string `json:"name"`        //产品名字
	Protocol    string `json:"protocol"`    //协议
	NodeType    string `json:"node_type"`   //节点类型
	NetType     string `json:"net_type"`    //联网模式
	DataFormat  string `json:"data_format"` //数据类型
	Factory     string `json:"factory"`     //厂家
	Description string `json:"description"` //描述
}

type OpenApiUpdateProductRequest struct {
	Id          string  `json:"id"`
	Name        *string `json:"name"`        //产品名字
	Protocol    *string `json:"protocol"`    //协议
	NodeType    *string `json:"node_type"`   //节点类型
	NetType     *string `json:"net_type"`    //联网模式
	DataFormat  *string `json:"data_format"` //数据类型
	Factory     *string `json:"factory"`     //厂家
	Description *string `json:"description"` //描述
}
