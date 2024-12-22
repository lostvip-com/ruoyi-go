package lv_cache

import (
	"github.com/lostvip-com/lv_framework/lv_cache/lv_ram"
	"github.com/lostvip-com/lv_framework/lv_cache/lv_redis"
	"github.com/lostvip-com/lv_framework/lv_global"
	"time"
)

// 支持set、hashSet操作
type ICache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (value string, err error)
	Del(key ...string) error

	HSet(key string, values ...interface{}) error
	HMSet(key string, mp map[string]any, duration time.Duration) error
	HGet(key, field string) (string, error)
	HDel(key string, fields ...string) error
	HGetAll(key string) (map[string]string, error)
	Exists(key string) (int64, error)
	Close()
	Expire(key string, duration time.Duration) error
}

var cacheClient ICache //主数据库

func GetCacheClient() ICache {
	if cacheClient == nil {
		var config = lv_global.Config()
		cacheType := config.GetVipperCfg().GetString("go.cache")
		if cacheType == "redis" {
			cacheClient = lv_redis.GetInstance(0)
		} else {
			cacheClient = lv_ram.GetRamCacheClient()
		}
	}
	return cacheClient
}
