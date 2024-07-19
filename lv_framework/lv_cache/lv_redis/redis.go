package lv_redis

import (
	"context"
	"fmt"
	"github.com/lostvip-com/lv_framework/conf"
	"github.com/lostvip-com/lv_framework/logme"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	redisClient *RedisClient
)

type RedisClient struct {
	c *redis.Client
}

//
//// 获取缓存单例
//func GetInstance() lv_cache.ICache {
//	if rdb == nil {
//		rdb = NewInstance(0)
//	}
//	return rdb
//}

// 获取缓存单例
func NewRedisClient(indexDb int) *RedisClient {
	conf := conf.ConfigDefault{}
	addr := conf.GetValueStr("go.redis.host")
	port := conf.GetValueStr("go.redis.port")
	password := conf.GetValueStr("go.redis.password")
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: password, // 没有密码，默认值
		DB:       indexDb,  // 默认DB 0
	})
	redisClient = new(RedisClient)
	redisClient.c = rdb
	if redisClient.c.Ping(context.Background()).Val() == "" {
		msg := ` 
			  ------------>连接 reids 错误：
			  无法链接到redis!!!! 检查相关配置:
			  host: %v
			  port: %v
			  password: %v
             `
		host := conf.GetValueStr("go.redis.host")
		logme.Error(fmt.Sprintf(msg, host, conf.GetValueStr("go.redis.port"), conf.GetValueStr("go.redis.password")))
		panic("redis 错误:" + host + " port:" + port)
	}
	return redisClient
}

func (rcc *RedisClient) HMSet(key string, mp map[string]any, expiration time.Duration) error {
	err := rcc.c.HSet(context.Background(), key, mp).Err()
	if err != nil {
		return err
	}
	err = rcc.c.Expire(context.Background(), key, expiration).Err()
	return err
}

func (rcc *RedisClient) Expire(key string, duration time.Duration) error {
	return rcc.c.Expire(context.Background(), key, duration).Err()
}

func (rcc *RedisClient) Exists(key string) (int64, error) {
	return rcc.c.Exists(context.Background(), key).Result()
}

func (rcc *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	rcc.c.Set(context.Background(), key, value, expiration)
	return nil
}

func (rcc *RedisClient) Get(key string) (data string, err error) {
	return rcc.c.Get(context.Background(), key).Result()
}

func (rcc *RedisClient) Del(keys ...string) error {
	var err error = nil
	for _, key := range keys {
		err0 := rcc.c.Del(context.Background(), key).Err()
		if err != nil {
			err = err0
		}
	}
	return err
}

func (rcc *RedisClient) HSet(key string, values ...interface{}) error {
	err := rcc.c.HSet(context.Background(), key, values...).Err()
	return err
}

func (rcc *RedisClient) HGet(key, field string) (string, error) {
	return rcc.c.HGet(context.Background(), key, field).Result()
}

func (rcc *RedisClient) HDel(key string, fields ...string) error {
	return rcc.c.HDel(context.Background(), key, fields...).Err()
}

func (rcc *RedisClient) HGetAll(key string) (map[string]string, error) {
	return rcc.c.HGetAll(context.Background(), key).Result()
}

func (rcc *RedisClient) Close() {
	rcc.c.Close()
}
