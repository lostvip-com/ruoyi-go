package lv_conv

import (
	"encoding/json"
	"github.com/spf13/cast"
)

func ToJsonStr(e interface{}) (string, error) {
	//格式化
	//b, err := json.MarshalIndent(user, "", "  ")
	b, err := json.Marshal(e)
	if err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

func ToStructPtr(jsonStr string, ptr any) error {
	err := json.Unmarshal([]byte(jsonStr), ptr)
	return err
}

func ToMap(jsonStr string) map[string]any {
	var result = map[string]any{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return result
	}
	return result
}

func StructToMap(entity any) map[string]any {
	jsonStr, err := ToJsonStr(entity)
	if err != nil {
		return nil
	}
	return ToMap(jsonStr)
}

func StructToMapStr(entity any) (map[string]string, error) {
	jsonStr, err := ToJsonStr(entity)
	if err != nil {
		return nil, err
	}
	mp := ToMap(jsonStr)
	//值全转转为字符串格式
	result := make(map[string]string)
	for k, v := range mp {
		result[k] = cast.ToString(v)
	}
	return result, err
}
