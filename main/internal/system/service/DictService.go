package service

import (
	"encoding/json"
	"github.com/lv_framework/utils/lv_conv"
	"html/template"
	dictDataModel "main/internal/system/model"
	"strings"
)

type DictService struct {
}

// 根据字典类型和字典键值查询字典数据信息
func (svc *DictService) GetDictLabel(dictType string, dictValue interface{}) template.HTML {
	result := ""
	dictData := &dictDataModel.SysDictData{DictType: dictType, DictValue: lv_conv.String(dictValue)}
	ok, _ := dictData.FindOne()
	if ok {
		result = dictData.DictLabel

	}
	return template.HTML(result)
}

// 通用的字典单选框控件  dictType 字典类别  value 默认值
func (svc *DictService) GetDictTypeRadio(dictType, name string, value interface{}) template.HTML {
	result, err := svc.SelectDictDataByType(dictType)
	if err != nil {
		return ""
	}

	if result == nil || len(result) <= 0 {
		return ""
	}

	htmlstr := ``

	for _, item := range result {
		if strings.Compare(item.DictValue, lv_conv.String(value)) == 0 {
			htmlstr += `<div class="radio-box"><option value="` + item.DictValue + `">` + item.DictLabel + `</option>`
			htmlstr += `<input type="radio" id="` + lv_conv.String(item.DictCode) + `" name="` + name + `" value="` + item.DictValue + `"
                           checked="checked">
                    <label for="` + lv_conv.String(item.DictCode) + `" text="` + item.DictLabel + `"></label></div>`
		} else {
			htmlstr += `<div class="radio-box"><option value="` + item.DictValue + `">` + item.DictLabel + `</option>`
			htmlstr += `<input type="radio" id="` + lv_conv.String(item.DictCode) + `" name="` + name + `" value="` + item.DictValue + `">
                    <label for="` + lv_conv.String(item.DictCode) + `" text="` + item.DictLabel + `"></label></div>`
		}
	}

	htmlstr += ``
	return template.HTML(htmlstr)
}

// 通用的字典下拉框控件  字典类别   html控件id  html控件name html控件class  html控件value  html控件空值标签 是否可以多选
func (svc *DictService) GetDictTypeSelect(dictType, id, name, className, value, emptyLabel, multiple string) template.HTML {

	result, err := svc.SelectDictDataByType(dictType)
	if err != nil {
		return ""
	}

	if result == nil || len(result) <= 0 {
		return ""
	}

	htmlstr := `<select id="` + id + `" name="` + name + `" class="` + className + `" ` + multiple + `>`

	if emptyLabel != "" {
		htmlstr += `<option value="">` + emptyLabel + `</option>`
	}

	for _, item := range result {
		if strings.Compare(item.DictValue, value) == 0 {
			htmlstr += `<option selected value="` + item.DictValue + `">` + item.DictLabel + `</option>`
		} else {
			htmlstr += `<option value="` + item.DictValue + `">` + item.DictLabel + `</option>`
		}
	}

	htmlstr += `</select>`

	return template.HTML(htmlstr)
}

// 通用的字典下拉框控件
func (svc *DictService) GetDictTypeData(dictType string) template.JS {
	result := make([]dictDataModel.SysDictData, 0)
	rs, err := svc.SelectDictDataByType(dictType)
	if err == nil || len(rs) > 0 {
		result = rs
	}

	jsonstr := ""

	jsonbyte, err := json.Marshal(result)

	if err == nil {
		jsonstr = string(jsonbyte)
	}

	return template.JS(jsonstr)
}

// 根据字典类型查询字典数据
func (svc *DictService) SelectDictDataByType(dictType string) ([]dictDataModel.SysDictData, error) {
	return dictDataModel.Find("status = '0' and dict_type = '"+dictType+"'", "dict_sort asc")
}
