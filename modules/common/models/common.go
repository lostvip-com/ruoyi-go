package models

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

type SliceString []string

func (c SliceString) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *SliceString) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}

type MapStringString map[string]string

func (c MapStringString) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *MapStringString) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}

type MapStringInterface map[string]interface{}

func (c MapStringInterface) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *MapStringInterface) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}

type MapStringInterfaceUseNumber map[string]interface{}

func (c MapStringInterfaceUseNumber) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *MapStringInterfaceUseNumber) Scan(value interface{}) error {
	return GormScanWrapUseNumber(value, c)
}

type MapStringSliceString map[string][]string

func (c MapStringSliceString) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *MapStringSliceString) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}

// 全局处理入库数据
func GormValueWrap(c interface{}) (driver.Value, error) {
	str, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	if str == nil || string(str) == "null" {
		return "", nil
	}
	return string(str), nil
}

func GormScanWrap(value interface{}, ojb interface{}) error {

	switch v := value.(type) {
	case string:
		if v == "" {
			return nil
		}
		err := json.Unmarshal([]byte(v), &ojb)
		if err != nil {
			return fmt.Errorf("failed to unmarshal json str: %v, err: %v", v, err)
		}
	case []byte:
		if string(v) == "" {
			return nil
		}
		err := json.Unmarshal(v, &ojb)
		if err != nil {
			return fmt.Errorf("failed to unmarshal json str: %v, err: %v", v, err)
		}
	default:
		fmt.Errorf("field is not string value: %v, type: %v", value, reflect.TypeOf(value))

	}

	return nil
}

func GormScanWrapUseNumber(value interface{}, ojb interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("field is not string value: %v, type: %v", value, reflect.TypeOf(value))
	}
	if str == "" {
		return nil
	}

	read := bytes.NewBufferString(str)
	decoder := json.NewDecoder(read)
	decoder.UseNumber()
	err := decoder.Decode(&ojb)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json str: %v, err: %v", str, err)
	}
	return nil
}
