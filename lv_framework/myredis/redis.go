package myredis

import (
	"github.com/redis/go-redis/v9"
	"lostvip.com/conf"
	"sync"
)

var (
	rdb  *redis.Client
	once sync.Once
)

// 获取缓存单例
func GetInstance() *redis.Client {
	once.Do(func() {
		vipperCfg := conf.Config().GetVipperCfg()
		addr := vipperCfg.GetString("go.redis.host")
		port := vipperCfg.GetString("go.redis.port")
		password := vipperCfg.GetString("go.redis.password")
		rdb = redis.NewClient(&redis.Options{
			Addr:     addr + ":" + port,
			Password: password, // 没有密码，默认值
			DB:       0,        // 默认DB 0
		})
	})
	return rdb
}
