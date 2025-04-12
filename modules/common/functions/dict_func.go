package functions

import (
	"common/common_vo"
	dictDataModel "common/models"
	"encoding/json"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"html/template"
	"strings"
	"system/dao"
)

type DictService struct {
}

// 根据字典类型和字典键值查询字典数据信息
func DictLabel(dictType string, dictValue interface{}) template.HTML {
	result := ""
	dictData := &dictDataModel.SysDictData{DictType: dictType, DictValue: lv_conv.String(dictValue)}
	dictData, err := dictData.FindOne()
	if err == nil {
		result = dictData.DictLabel

	}
	return template.HTML(result)
}

// 通用的字典单选框控件  dictType 字典类别  value 默认值
func DictRadio(dictType, name string, value interface{}) template.HTML {
	result, err := SelectDictDataByType(dictType)
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

// GetDictTypeSelect 通用的字典下拉框控件  字典类别   html控件id  html控件name html控件class  html控件value  html控件空值标签 是否可以多选
func GetDictTypeSelect(dictType, id, name, className, value, emptyLabel, multiple string) template.HTML {

	result, err := SelectDictDataByType(dictType)
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

// DictSelect Select 为GetDictTypeSelect 的简化版，只支持单选下拉，不支持指定样式
func DictSelect(dictType, htmlField, selectedVal string) template.HTML {

	result, err := SelectDictDataByType(dictType)
	if err != nil {
		return ""
	}

	if result == nil || len(result) <= 0 {
		return ""
	}
	htmlstr := `<select id="` + htmlField + `" name="` + htmlField + `" class="form-control m-b">`

	htmlstr += `<option value="">` + `请选择` + `</option>`

	for _, item := range result {
		if strings.Compare(item.DictValue, selectedVal) == 0 {
			htmlstr += `<option selected value="` + item.DictValue + `">` + item.DictLabel + `</option>`
		} else {
			htmlstr += `<option value="` + item.DictValue + `">` + item.DictLabel + `</option>`
		}
	}

	htmlstr += `</select>`

	return template.HTML(htmlstr)
}

// 通用的字典下拉框控件
func DictType(dictType string) template.JS {
	result := make([]dictDataModel.SysDictData, 0)
	rs, err := SelectDictDataByType(dictType)
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
func SelectDictDataByType(dictType string) ([]dictDataModel.SysDictData, error) {
	var dictDataModel dao.DictDataDao
	var req = common_vo.SelectDictDataPageReq{DictType: dictType, Status: "0"}
	return dictDataModel.SelectListAll(&req)
}
