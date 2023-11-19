package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	db "lostvip.com/db"
	"lostvip.com/myredis"
	"lostvip.com/server"
	"robvi/app/global"
	_ "robvi/app/modules"
	"robvi/app/modules/sys/model/module/tenant"
	"robvi/app/modules/sys/model/system/dept"
	"robvi/app/modules/sys/model/system/post"
	"robvi/app/modules/sys/model/system/user"
	"time"
)

// @title LV 自动生成API文档
// @version 1.0
// @description 生成文档请在调试模式下进行<a href="/tool/swagger?a=r">重新生成文档</a>

// @host localhost
// @BasePath /api
func main() {
	//str, _ := gmd5.Encrypt("123456")
	//println(">>>>>>>>>>>>>: " + str)
	cfg := global.GetConfigInstance()
	if cfg.IsDebug() {
		gin.SetMode("debug")
		db.Instance().Engine().ShowSQL(true)
	}
	db.Instance().Engine().Sync2(dept.SysDept{}, tenant.SysTenant{}, user.SysUser{}, post.SysPost{})
	//reid
	ctx := context.Background()
	redis := myredis.GetInstance()
	data := map[string]any{"test": "123"}
	redis.HMSet(ctx, "mapKey1", data)

	fieldMap := make(map[string]interface{})
	fieldMap["field1"] = "val1"
	fieldMap["field2"] = "val2"
	redis.HMSet(ctx, "key", fieldMap)
	redis.Expire(ctx, "key", 100*time.Second)
	fmt.Println("------------myredis----------------------123")
	data1 := redis.HGet(ctx, "mapKey1", "test")
	fmt.Println(data1)
	//后台服务状态
	httpSvr := server.New("0.0.0.0:" + cast.ToString(cfg.GetServerPort()))
	httpSvr.Start()
	fmt.Println("------------exit----------------------")

}
