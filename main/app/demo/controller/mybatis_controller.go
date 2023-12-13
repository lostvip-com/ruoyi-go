package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"lostvip.com/cache/myredis"
	"lostvip.com/db"
	"lostvip.com/utils/lv_web"
	"robvi/app/system/model/system/post"
	"time"
)

// 相对于mapper目录的路径
var SQL_FILE_POST = "sys_post/sys_post_mapper.tpl"

/**
 * 基于ibatis 的分页查询演示
 */
func (w DemoController) Mybatis3(c *gin.Context) {
	req := post.SelectPageReq{PostName: "%test%", PageSize: 200, PageNum: 1}
	if err := c.ShouldBind(&req); err != nil { //获取参数
		lv_web.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	resp := db.GetPageByNamedSql[post.SysPost](SQL_FILE_POST, "test", req)
	lv_web.PageOK(c, resp)
}

func (w DemoController) Mybatis1(c *gin.Context) {
	req := post.SelectPageReq{PostName: "%test%", PageSize: 200, PageNum: 1}
	if err := c.ShouldBind(&req); err != nil { //获取参数
		lv_web.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	ibatis := db.NewIBatis(SQL_FILE_POST)
	sql := ibatis.GetLimitSql("test", req)
	list := db.ListByNamedSql[post.SysPost](sql, req)
	count := db.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_web.PageOK2(c, list, count)
}

func (w DemoController) Mybatis2(c *gin.Context) {
	req := post.SelectPageReq{PostName: "%test%", PageSize: 200, PageNum: 1}
	if err := c.ShouldBind(&req); err != nil { //获取参数
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("patient管理", req).WriteJsonExit()
		return
	}

	ibatis := db.NewIBatis(SQL_FILE_POST)
	sql := ibatis.GetLimitSql("test", req)
	listMap, err := db.GetInstance().ListByNamedSql(sql, req, true)
	count, err := db.GetInstance().CountByNamedSql(ibatis.GetCountSql(), req)
	if err != nil {
		panic(err)
	}
	lv_web.PageOK2(c, listMap, count)
}

func (w DemoController) TestRedis(c *gin.Context) {
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
	lv_web.Sucess(c, data1)
}
