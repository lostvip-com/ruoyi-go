package service

import (
	"github.com/lostvip-com/lv_framework/lv_cache"
	"github.com/spf13/cast"
	"main/internal/iot_dev/model"
)

var cacheService *DeviceCacheService

type DeviceCacheService struct {
}

func GetDeviceCacheService() *DeviceCacheService {
	if cacheService == nil {
		cacheService = &DeviceCacheService{}
	}
	return cacheService
}

func (svc *DeviceCacheService) GetSysDevName(deviceId int64) string {
	key := "sys:dev:" + cast.ToString(deviceId)
	deptName, _ := lv_cache.GetCacheClient().HGet(key, "deptName")
	if deptName == "" {
		dev := new(model.IotDevice)
		dev, err := dev.FindById(deviceId)
		if err == nil {
			lv_cache.GetCacheClient().HSet(key, dev)
		}
		deptName = dev.Name
	}
	return deptName
}
