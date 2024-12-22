package lv_conv

import (
	"github.com/goccy/go-json"
)

func ToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

func ToStructPtr(jsonStr string, ptr any) error {
	err := json.Unmarshal([]byte(jsonStr), ptr)
	return err
}

func ToMap(jsonStr string) (map[string]any, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
