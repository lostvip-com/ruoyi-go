package myredis

import (
	"context"
	"fmt"
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
	if rdb == nil {
		rdb = NewInstance(0)
	}
	return rdb
}

// 获取缓存单例
func NewInstance(indexDb int) *redis.Client {
	vipperCfg := conf.Config().GetVipperCfg()
	addr := vipperCfg.GetString("go.redis.host")
	port := vipperCfg.GetString("go.redis.port")
	password := vipperCfg.GetString("go.redis.password")
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     addr + ":" + port,
			Password: password, // 没有密码，默认值
			DB:       indexDb,  // 默认DB 0
		})
	})
	if rdb.Ping(context.Background()).Val() == "" {
		msg := ` 
			  ------------>连接 reids 错误：
			  无法链接到redis!!!! 检查相关配置:
			  host: %v
			  port: %v
			  password: %v
             `
		panic(fmt.Sprintf(msg, vipperCfg.GetString("go.redis.host"), vipperCfg.GetString("go.redis.port"), vipperCfg.GetString("go.redis.password")))
	}
	return rdb
}
