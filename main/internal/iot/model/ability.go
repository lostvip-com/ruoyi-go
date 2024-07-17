package models

import (
	"database/sql/driver"
	"encoding/json"
	"main/internal/iot/pkg/constants"
)

// 能力三件套 Properties Events Actions

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type Properties struct {
	Id          string   `json:"id" gorm:"id;primaryKey;not null;type:string;size:255;comment:主键"`
	ProductId   string   `json:"product_id" gorm:"type:string;size:255;comment:产品ID"`
	Name        string   `json:"name" gorm:"type:string;size:255;comment:名字"`
	Code        string   `json:"code" gorm:"type:string;size:255;comment:标识符"`
	AccessMode  string   `json:"access_mode" gorm:"type:string;size:50;comment:读写模型"`
	Require     bool     `json:"require" gorm:"comment:是否必须"`
	TypeSpec    TypeSpec `json:"type_spec" gorm:"type:text;comment:属性物模型详情"`
	Description string   `json:"description" gorm:"type:text;comment:描述"`
	Tag         string   `json:"tag" gorm:"type:string;size:50;comment:标签"`
	System      bool     `json:"system" gorm:"comment:系统内置"`
	Timestamps
}

func (p *Properties) TableName() string {
	return "properties"
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type Events struct {
	Id           string       `json:"id" gorm:"id;primaryKey;not null;type:string;size:255;comment:主键"`
	ProductId    string       `json:"product_id" gorm:"type:string;size:255;comment:产品ID"`
	EventType    string       `json:"event_type" gorm:"type:string;size:255;comment:事件类型"`
	Code         string       `json:"code" gorm:"type:string;size:255;comment:标识符"`
	Name         string       `json:"name" gorm:"type:string;size:255;comment:名字"`
	Description  string       `json:"description" gorm:"type:text;comment:描述"`
	Require      bool         `json:"require" gorm:"comment:是否必须"`
	OutputParams OutPutParams `json:"output_params" gorm:"type:text;comment:输入参数"`
	Tag          string       `json:"tag" gorm:"type:string;size:50;comment:标签"`
	System       bool         `json:"system" gorm:"comment:系统内置"`
	Timestamps
}

func (table *Events) TableName() string {
	return "events"
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Actions struct {
	Id           string             `json:"id" gorm:"id;primaryKey;not null;type:string;size:255;comment:主键"`
	ProductId    string             `json:"product_id" gorm:"type:string;size:255;comment:产品ID"`
	Code         string             `json:"code" gorm:"type:string;size:255;comment:标识符"`
	Name         string             `json:"name" gorm:"type:string;size:255;comment:名字"`
	Description  string             `json:"description" gorm:"type:text;comment:描述"`
	Require      bool               `json:"require" gorm:"comment:是否必须"`
	CallType     constants.CallType `json:"call_type" gorm:"type:string;size:50;comment:调用方式"`
	InputParams  InPutParams        `json:"input_params" gorm:"type:text;comment:输入参数"`  // 输入参数
	OutputParams OutPutParams       `json:"output_params" gorm:"type:text;comment:输入参数"` // 输出参数
	Tag          string             `json:"tag" gorm:"type:string;size:50;comment:标签"`
	System       bool               `json:"system" gorm:"comment:系统内置"`
	Timestamps
}

func (table *Actions) TableName() string {
	return "actions"
}

// ============================================================================================================================
// TypeSpec 物模型的属性、事件、动作的基础信息字段
type TypeSpec struct {
	Type  constants.SpecsType `json:"type,omitempty"`
	Specs string              `json:"specs,omitempty"`
}

type TypeSpecIntOrFloat struct {
	Min      string `json:"min,omitempty"`
	Max      string `json:"max,omitempty"`
	Step     string `json:"step,omitempty"`
	Unit     string `json:"unit,omitempty"`
	UnitName string `json:"unitName,omitempty"`
}

func (t *TypeSpecIntOrFloat) TransformTostring() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type TypeSpecText struct {
	Length string `json:"length,omitempty"`
}

func (t *TypeSpecText) TransformTostring() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type TypeSpecBool map[string]string

func (t *TypeSpecBool) TransformTostring() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type TypeSpecArray struct {
	Size string `json:"size,omitempty"`
	Item Item   `json:"item,omitempty"`
}

func (t *TypeSpecArray) TransformTostring() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type TypeSpecEnum map[string]string

func (t *TypeSpecEnum) TransformTostring() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type TypeSpecStruct struct {
	Code     string   `json:"code"`
	Name     string   `json:"name"`
	DataType TypeSpec `json:"data_type"`
}

func (t *TypeSpecStruct) TransformTostring() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type TypeSpecDate struct {
}

func (t *TypeSpecDate) TransformTostring() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type Item struct {
	Type string `json:"type,omitempty"`
}

func (c Item) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *Item) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}

func (c TypeSpec) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *TypeSpec) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}

type InPutParams []InputOutput

type InputOutput struct {
	Code     string   `json:"code"`
	Name     string   `json:"name"`
	TypeSpec TypeSpec `json:"type_spec"`
}

// Value Value() 方法会在将值写入数据库时被调用
func (c InPutParams) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

// Scan sql.Scanner 接口的 Scan() 方法会在从数据库读取值时被调用
func (c *InPutParams) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}

type OutPutParams []InputOutput

func (c OutPutParams) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *OutPutParams) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}
