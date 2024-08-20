package dtos

import (
	"main/internal/iot_dev/pkg/constants"
)

type ThingModelTemplate struct {
	Properties []ThingModelTemplateProperties `json:"properties"`
	Events     []ThingModelTemplateEvents     `json:"events"`
	Services   []ThingModelTemplateServices   `json:"services"`
}

type ThingModelTemplateArray struct {
	ChildDataType string `json:"childDataType"`
	Size          int    `json:"size"`
}

type ThingModelTemplateIntOrFloat struct {
	Max      string `json:"max"`
	Min      string `json:"min"`
	Step     string `json:"step"`
	Unit     string `json:"unit"`
	UnitName string `json:"unitName"`
}

type ThingModelTemplateBool struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ThingModelTemplateText struct {
	Length int `json:"length"`
}

type ThingModelTemplateDate struct {
	Length string `json:"length"`
}

type ThingModelTemplateEnum struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ThingModelTemplateStruct struct {
	ChildDataType string      `json:"childDataType"`
	ChildName     string      `json:"childName"`
	Identifier    string      `json:"identifier"`
	ChildSpecsDTO interface{} `json:"childSpecsDTO"`
}

//------------------------------------------------------------
// INT DOUBLE FLOAT TEXT ARRAY=> dataSpecs
//BOOL ENUM STRUCT=> dataSpecsList

type ThingModelTemplateProperties struct {
	Name          string      `json:"name"`
	Identifier    string      `json:"identifier"`
	DataType      string      `json:"dataType"`
	Description   string      `json:"description"`
	Required      bool        `json:"required"`
	RwFlag        string      `json:"rwFlag"`
	DataSpecs     interface{} `json:"dataSpecs"`
	DataSpecsList interface{} `json:"dataSpecsList"`
}

// ASYNC SYNC
type ThingModelTemplateServices struct {
	ServiceName string                                 `json:"serviceName"`
	Identifier  string                                 `json:"identifier"`
	Description string                                 `json:"description"`
	Required    bool                                   `json:"required"`
	CallType    constants.CallType                     `json:"callType"`
	InputParams []ThingModelTemplateServicesInputParam `json:"inputParams"`
	OutParams   []ThingModelTemplateServicesOutParam   `json:"outParams"`
}

type ThingModelTemplateServicesInputParam struct {
	Name          string      `json:"name"`
	Identifier    string      `json:"identifier"`
	DataType      string      `json:"dataType"`
	DataSpecs     interface{} `json:"dataSpecs"`
	DataSpecsList interface{} `json:"dataSpecsList"`
}

type ThingModelTemplateServicesOutParam struct {
	Name          string      `json:"name"`
	Identifier    string      `json:"identifier"`
	DataType      string      `json:"dataType"`
	DataSpecs     interface{} `json:"dataSpecs"`
	DataSpecsList interface{} `json:"dataSpecsList"`
}

type ThingModelTemplateEvents struct {
	EventName   string                               `json:"eventName"`
	EventType   string                               `json:"eventType"`
	Identifier  string                               `json:"identifier"`
	Description string                               `json:"description"`
	Required    bool                                 `json:"required"`
	OutputData  []ThingModelTemplateEventsOutputData `json:"outputData"`
}

type ThingModelTemplateEventsOutputData struct {
	Name          string      `json:"name"`
	Identifier    string      `json:"identifier"`
	DataType      string      `json:"dataType"`
	Required      bool        `json:"required"`
	DataSpecs     interface{} `json:"dataSpecs"`
	DataSpecsList interface{} `json:"dataSpecsList"`
}
