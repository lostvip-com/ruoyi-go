package service

import (
	"common/model"
	"github.com/lostvip-com/lv_framework/lv_cache"
	"github.com/spf13/cast"
)

var cacheService *CacheService

func init() {
	GetCacheService()
}

type CacheService struct {
}

func GetCacheService() *CacheService {
	if cacheService == nil {
		cacheService = &CacheService{}
	}
	return cacheService
}

func (svc *CacheService) GetSysDeptName(deptId int64) string {
	key := "sys:dept:" + cast.ToString(deptId)
	deptName, _ := lv_cache.GetCacheClient().HGet(key, "deptName")
	if deptName == "" {
		dept := new(model.SysDept)
		dept, err := dept.FindById(deptId)
		if err == nil {
			lv_cache.GetCacheClient().HSet(key, dept)
		}
		deptName = dept.DeptName
	}
	return deptName
}

func (svc *CacheService) GetDictLabel(dictType string, dictValue string) string {
	key := "sys:dict:" + dictType
	label, _ := lv_cache.GetCacheClient().HGet(key, dictValue)
	if label == "" {
		dictData := &model.SysDictData{DictType: dictType, DictValue: dictValue}
		dictData, err := dictData.FindOne()
		if err == nil {
			label = dictData.DictLabel
			lv_cache.GetCacheClient().HSet(key, dictValue, label)
		}
	}
	return label
}
