package dtos

import "main/internal/iot_dev/pkg/constants"

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
