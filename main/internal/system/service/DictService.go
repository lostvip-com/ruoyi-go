package service

import (
	dictDataModel "common/model"
	db2 "github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
)

type DictService struct {
}

// 根据字典类型和字典键值查询字典数据信息
func (svc *DictService) GetDictLabel(dictType string, dictValue string) string {
	dictData := &dictDataModel.SysDictData{DictType: dictType, DictValue: lv_conv.String(dictValue)}
	dictData, err := dictData.FindOne()
	if err == nil {
		dictValue = dictData.DictLabel
	}
	return dictValue
}

// 根据字典类型查询字典数据
func (svc *DictService) SelectDictDataByType(dictType string) ([]dictDataModel.SysDictData, error) {
	tb := db2.GetMasterGorm().Table("sys_dict_data")
	var list []dictDataModel.SysDictData
	err := tb.Where(&list, "status = '0' and dict_type = ? ", dictType).Order("dict_sort asc").Find(&list).Error
	return list, err
}
