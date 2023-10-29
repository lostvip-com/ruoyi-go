package redis

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"lostvip.com/conf"
)

var MyCache *Cache

type Cache struct {
	RedisPool *redis.Pool //创建redis连接池
}

func InitRedis() {
	vipperCfg := conf.Config().GetVipperCfg()
	MyCache = new(Cache)
	MyCache.RedisPool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 5, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			addr := vipperCfg.GetString("go.redis.host")
			port := vipperCfg.GetString("go.redis.port")
			password := vipperCfg.GetString("go.redis.password")
			conn, err := redis.Dial("tcp", addr+":"+port)
			if password != "" {
				err = conn.Send("auth", password)
			}
			if err != nil {
				panic(err)
			}
			return conn, err
		},
	}
}
func showErr(tag string, msg interface{}) {
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXX tag:", tag, msg)
}
func (e *Cache) SetInt(k, v string, seconds int) {
	defer func() { //错误拦截要在产?错误前设置
		err := recover()
		if err != nil {
			showErr(" redis ", err)
		}
	}()
	if e == nil {
		panic("未初始化redis")
	}
	c := e.RedisPool.Get() //从连接池，取一个链接
	defer c.Close()
	_, err := redis.Int(c.Do("Set", v, seconds))
	if err != nil {
		panic(err)
	}
}

func (e *Cache) GetInt(key string) int {
	defer func() { //错误拦截要在产?错误前设置
		err := recover()
		if err != nil {
			showErr(" redis ", err)
		}
	}()
	if e == nil {
		panic("未初始化redis")
	}
	c := e.RedisPool.Get() //从连接池，取一个链接
	defer c.Close()
	r, err := redis.Int(c.Do("Get", key))
	if err != nil {
		showErr(" GetInt ", err)
	}
	return r
}

func (e *Cache) Incr(key string) int {
	defer func() { //错误拦截要在产?错误前设置
		err := recover()
		if err != nil {
			showErr(" redis ", err)
		}
	}()
	if e == nil {
		panic("未初始化redis")
	}
	c := e.RedisPool.Get() //从连接池，取一个链接
	r, err := redis.Int(c.Do("Incr", key))
	if err != nil {
		showErr(" Incr ", err)
	}
	c.Close()
	return r
}

// --------------------------------------------------------
func (e *Cache) HSetString(key, field, v string) {
	defer func() { //错误拦截要在产?错误前设置
		err := recover()
		if err != nil {
			showErr(" redis ", err)
		}
	}()
	if e == nil {
		panic("未初始化redis")
	}
	c := e.RedisPool.Get() //从连接池，取一个链接
	defer c.Close()
	_, err := redis.String(c.Do("HSet", key, field, v))
	if err != nil {
		showErr(" HSetString ", err)
	}
}
func (e *Cache) HGetString(key, field string) string {
	defer func() { //错误拦截要在产?错误前设置
		err := recover()
		if err != nil {
			fmt.Println("XXXXXXXX", err)
		}
	}()
	if e == nil {
		panic("未初始化redis")
	}
	c := e.RedisPool.Get() //从连接池，取一个链接
	defer c.Close()
	r, err := redis.String(c.Do("HGet", key, field))
	if err != nil {
		showErr(" HGetString ", err)
	}
	return r
}
