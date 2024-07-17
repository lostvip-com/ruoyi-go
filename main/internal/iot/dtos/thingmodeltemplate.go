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
	"github.com/winc-link/hummingbird/internal/models"
	"github.com/winc-link/hummingbird/internal/pkg/constants"
	"github.com/winc-link/hummingbird/internal/pkg/utils"
	"strconv"
	"strings"
	"time"
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

func (t *ThingModelTemplateProperties) TransformModelTypeSpec() (typeSpec models.TypeSpec) {
	return getThingModelTemplateTypeSpec(t.DataType, t.DataSpecs, t.DataSpecsList)
}

func getModelSpecByDataType(specType constants.SpecsType, dto interface{}) string {
	switch specType {
	case constants.SpecsTypeInt, constants.SpecsTypeFloat:
		thingModelTemplateIntOrFloat := new(ThingModelTemplateIntOrFloat)
		b, _ := json.Marshal(dto)
		err := json.Unmarshal(b, thingModelTemplateIntOrFloat)
		if err != nil {
			return ""
		}
		var modelTypeSpecIntOrFloat models.TypeSpecIntOrFloat
		modelTypeSpecIntOrFloat.Max = thingModelTemplateIntOrFloat.Max
		modelTypeSpecIntOrFloat.Min = thingModelTemplateIntOrFloat.Min
		modelTypeSpecIntOrFloat.Step = thingModelTemplateIntOrFloat.Step
		modelTypeSpecIntOrFloat.Unit = thingModelTemplateIntOrFloat.Unit
		modelTypeSpecIntOrFloat.UnitName = thingModelTemplateIntOrFloat.UnitName
		return modelTypeSpecIntOrFloat.TransformTostring()
	case constants.SpecsTypeText:
		thingModelTemplateText := new(ThingModelTemplateText)
		b, _ := json.Marshal(dto)
		err := json.Unmarshal(b, thingModelTemplateText)
		if err != nil {
			return ""
		}
		var modelTypeSpecText models.TypeSpecText
		modelTypeSpecText.Length = strconv.Itoa(thingModelTemplateText.Length)
		return modelTypeSpecText.TransformTostring()
	case constants.SpecsTypeDate:
		var modelTypeSpecText models.TypeSpecDate
		return modelTypeSpecText.TransformTostring()
	case constants.SpecsTypeBool:
		thingModelTemplateBool := new([]ThingModelTemplateBool)
		b, _ := json.Marshal(dto)
		err := json.Unmarshal(b, thingModelTemplateBool)
		if err != nil {
			return ""
		}
		var modelTypeSpecBool models.TypeSpecBool
		modelTypeSpecBool = make(map[string]string)
		if thingModelTemplateBool != nil {
			for _, templateBool := range *thingModelTemplateBool {
				modelTypeSpecBool[strconv.Itoa(templateBool.Value)] = templateBool.Name
			}
		}
		return modelTypeSpecBool.TransformTostring()
	case constants.SpecsTypeEnum:
		thingModelTemplateEnum := new([]ThingModelTemplateEnum)
		b, _ := json.Marshal(dto)
		err := json.Unmarshal(b, thingModelTemplateEnum)
		if err != nil {
			return ""
		}
		var modelTypeSpecEunm models.TypeSpecEnum
		modelTypeSpecEunm = make(map[string]string)
		if thingModelTemplateEnum != nil {
			for _, templateEunm := range *thingModelTemplateEnum {
				modelTypeSpecEunm[strconv.Itoa(templateEunm.Value)] = templateEunm.Name
			}
		}
		return modelTypeSpecEunm.TransformTostring()
	case constants.SpecsTypeArray:
		thingModelTemplateArray := new(ThingModelTemplateArray)
		b, _ := json.Marshal(dto)
		err := json.Unmarshal(b, thingModelTemplateArray)
		if err != nil {
			return ""
		}
		var modelTypeSpecArray models.TypeSpecArray
		modelTypeSpecArray.Size = strconv.Itoa(thingModelTemplateArray.Size)
		modelTypeSpecArray.Item = models.Item{
			Type: thingModelTemplateArray.ChildDataType,
		}
		return modelTypeSpecArray.TransformTostring()
	}
	return ""
}

func transformSpecType(specType string) constants.SpecsType {
	var specs constants.SpecsType
	switch specType {
	case "INT":
		specs = constants.SpecsTypeInt
	case "DOUBLE", "FLOAT":
		specs = constants.SpecsTypeFloat
	case "TEXT":
		specs = constants.SpecsTypeText
	case "ARRAY":
		specs = constants.SpecsTypeArray
	case "BOOL":
		specs = constants.SpecsTypeBool
	case "ENUM":
		specs = constants.SpecsTypeEnum
	case "STRUCT":
		specs = constants.SpecsTypeStruct
	case "DATE":
		specs = constants.SpecsTypeDate
	}
	return specs
}

//ASYNC SYNC
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

func (t *ThingModelTemplateServices) TransformModelInPutParams() (inPutParams models.InPutParams) {
	for _, datum := range t.InputParams {
		var inputOutput models.InputOutput
		inputOutput.Code = datum.Identifier
		inputOutput.Name = datum.Name
		inputOutput.TypeSpec = getThingModelTemplateTypeSpec(datum.DataType, datum.DataSpecs, datum.DataSpecsList)
		inPutParams = append(inPutParams, inputOutput)
	}
	return
}

func (t *ThingModelTemplateServices) TransformModelOutPutParams() (outPutParams models.OutPutParams) {
	for _, datum := range t.OutParams {
		var inputOutput models.InputOutput
		inputOutput.Code = datum.Identifier
		inputOutput.Name = datum.Name
		inputOutput.TypeSpec = getThingModelTemplateTypeSpec(datum.DataType, datum.DataSpecs, datum.DataSpecsList)
		outPutParams = append(outPutParams, inputOutput)
	}
	return
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

func (t *ThingModelTemplateEvents) TransformModelOutputParams() (outPutParams models.OutPutParams) {
	for _, datum := range t.OutputData {
		var inputOutput models.InputOutput
		inputOutput.Code = datum.Identifier
		inputOutput.Name = datum.Name
		inputOutput.TypeSpec = getThingModelTemplateTypeSpec(datum.DataType, datum.DataSpecs, datum.DataSpecsList)
		outPutParams = append(outPutParams, inputOutput)
	}
	return outPutParams
}

func getThingModelTemplateTypeSpec(dataType string, dataSpecs, dataSpecsList interface{}) (typeSpec models.TypeSpec) {
	typeSpec.Type = transformSpecType(dataType)
	switch typeSpec.Type {
	case constants.SpecsTypeInt, constants.SpecsTypeFloat:
		thingModelTemplateIntOrFloat := new(ThingModelTemplateIntOrFloat)
		b, _ := json.Marshal(dataSpecs)
		err := json.Unmarshal(b, thingModelTemplateIntOrFloat)
		if err != nil {
			return
		}
		var modelTypeSpecIntOrFloat models.TypeSpecIntOrFloat
		modelTypeSpecIntOrFloat.Max = thingModelTemplateIntOrFloat.Max
		modelTypeSpecIntOrFloat.Min = thingModelTemplateIntOrFloat.Min
		modelTypeSpecIntOrFloat.Step = thingModelTemplateIntOrFloat.Step
		modelTypeSpecIntOrFloat.Unit = thingModelTemplateIntOrFloat.Unit
		modelTypeSpecIntOrFloat.UnitName = thingModelTemplateIntOrFloat.UnitName
		typeSpec.Specs = modelTypeSpecIntOrFloat.TransformTostring()

	case constants.SpecsTypeText:
		thingModelTemplateText := new(ThingModelTemplateText)
		b, _ := json.Marshal(dataSpecs)
		err := json.Unmarshal(b, thingModelTemplateText)
		if err != nil {
			return
		}
		var modelTypeSpecText models.TypeSpecText
		modelTypeSpecText.Length = strconv.Itoa(thingModelTemplateText.Length)
		typeSpec.Specs = modelTypeSpecText.TransformTostring()
	case constants.SpecsTypeDate:
		var modelTypeSpecText models.TypeSpecDate
		typeSpec.Specs = modelTypeSpecText.TransformTostring()
	case constants.SpecsTypeBool:
		thingModelTemplateBool := new([]ThingModelTemplateBool)
		b, _ := json.Marshal(dataSpecsList)
		err := json.Unmarshal(b, thingModelTemplateBool)
		if err != nil {
			return
		}
		var modelTypeSpecBool models.TypeSpecBool
		modelTypeSpecBool = make(map[string]string)
		if thingModelTemplateBool != nil {
			for _, templateBool := range *thingModelTemplateBool {
				modelTypeSpecBool[strconv.Itoa(templateBool.Value)] = templateBool.Name
			}
		}
		typeSpec.Specs = modelTypeSpecBool.TransformTostring()
	case constants.SpecsTypeEnum:
		thingModelTemplateEnum := new([]ThingModelTemplateEnum)
		b, _ := json.Marshal(dataSpecsList)
		err := json.Unmarshal(b, thingModelTemplateEnum)
		if err != nil {
			return
		}
		var modelTypeSpecEunm models.TypeSpecEnum
		modelTypeSpecEunm = make(map[string]string)
		if thingModelTemplateEnum != nil {
			for _, templateEunm := range *thingModelTemplateEnum {
				modelTypeSpecEunm[strconv.Itoa(templateEunm.Value)] = templateEunm.Name
			}
		}
		typeSpec.Specs = modelTypeSpecEunm.TransformTostring()
	case constants.SpecsTypeArray:
		thingModelTemplateArray := new(ThingModelTemplateArray)
		b, _ := json.Marshal(dataSpecs)
		err := json.Unmarshal(b, thingModelTemplateArray)
		if err != nil {
			return
		}
		var modelTypeSpecArray models.TypeSpecArray
		modelTypeSpecArray.Size = strconv.Itoa(thingModelTemplateArray.Size)
		modelTypeSpecArray.Item = models.Item{
			Type: strings.ToLower(thingModelTemplateArray.ChildDataType),
		}
		typeSpec.Specs = modelTypeSpecArray.TransformTostring()
	case constants.SpecsTypeStruct:
		thingModelTemplateStruct := new([]ThingModelTemplateStruct)
		b, _ := json.Marshal(dataSpecsList)
		err := json.Unmarshal(b, thingModelTemplateStruct)
		if err != nil {
			return
		}
		var modelTypeSpecStruct []models.TypeSpecStruct
		if thingModelTemplateStruct != nil {
			for _, templateStruct := range *thingModelTemplateStruct {
				modelTypeSpecStruct = append(modelTypeSpecStruct, models.TypeSpecStruct{
					Code: templateStruct.Identifier,
					Name: templateStruct.ChildName,
					DataType: models.TypeSpec{
						Type:  transformSpecType(templateStruct.ChildDataType),
						Specs: getModelSpecByDataType(transformSpecType(templateStruct.ChildDataType), templateStruct.ChildSpecsDTO),
					},
				})
			}
		}

		bm, _ := json.Marshal(modelTypeSpecStruct)
		typeSpec.Specs = string(bm)
	}

	return
}

func GetModelPropertyEventActionByThingModelTemplate(thingModelJSON string) (properties []models.Properties, events []models.Events, actions []models.Actions) {
	thingModelTemplate := new(ThingModelTemplate)
	err := json.Unmarshal([]byte(thingModelJSON), thingModelTemplate)
	if err == nil {
		for _, property := range thingModelTemplate.Properties {
			var accessMode string
			if property.RwFlag == "READ_ONLY" {
				accessMode = "R"
			} else if property.RwFlag == "READ_WRITE" {
				accessMode = "RW"
			} else if property.RwFlag == "WRITE_ONLY" {
				accessMode = "W"
			}
			properties = append(properties, models.Properties{
				Id:          utils.RandomNum(),
				Name:        property.Name,
				Code:        property.Identifier,
				AccessMode:  accessMode,
				Require:     property.Required,
				Description: property.Description,
				TypeSpec:    property.TransformModelTypeSpec(),
				Tag:         string(constants.TagNameSystem),
				Timestamps: models.Timestamps{
					Created: time.Now().UnixMilli(),
				},
			})
		}
		for _, event := range thingModelTemplate.Events {
			var eventType constants.EventType
			if event.EventType == "ALERT_EVENT_TYPE" {
				eventType = constants.EventTypeAlert
			} else if event.EventType == "INFO_EVENT_TYPE" {
				eventType = constants.EventTypeInfo
			} else if event.EventType == "ERROR_EVENT_TYPE" {
				eventType = constants.EventTypeError
			}
			events = append(events, models.Events{
				Id:           utils.RandomNum(),
				Name:         event.EventName,
				EventType:    string(eventType),
				Code:         event.Identifier,
				Require:      event.Required,
				Description:  event.Description,
				OutputParams: event.TransformModelOutputParams(),
				Tag:          string(constants.TagNameSystem),
				Timestamps: models.Timestamps{
					Created: time.Now().UnixMilli(),
				},
			})
		}
		for _, service := range thingModelTemplate.Services {
			actions = append(actions, models.Actions{
				Id:           utils.RandomNum(),
				Name:         service.ServiceName,
				Code:         service.Identifier,
				CallType:     service.CallType,
				Require:      service.Required,
				Description:  service.Description,
				InputParams:  service.TransformModelInPutParams(),
				OutputParams: service.TransformModelOutPutParams(),
				Tag:          string(constants.TagNameSystem),
				Timestamps: models.Timestamps{
					Created: time.Now().UnixMilli(),
				},
			})
		}
	}
	return
}
