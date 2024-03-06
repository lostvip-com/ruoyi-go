package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"lostvip.com/cache/myredis"
	"lostvip.com/db/lvbatis"
	"lostvip.com/db/lvdao"
	"lostvip.com/utils/lv_err"
	"lostvip.com/utils/lv_web"
	"main/app/system/model"
	"main/app/system/vo"
	"time"
)

// 相对于mapper目录的路径
var SQL_FILE_POST = "sys_post/sys_post_mapper.sql"

func (w DemoController) MybatisMap(c *gin.Context) {
	req := vo.SelectPostPageReq{}
	if err := c.ShouldBind(&req); err != nil { //获取参数
		lv_web.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}

	ibatis := lvbatis.NewInstance(SQL_FILE_POST)
	sql, err := ibatis.GetLimitSql("listSql", &req)
	lv_err.HasErrAndPanic(err)
	listMap, err := lvdao.ListMapByNamedSql(sql, &req, true)
	lv_err.HasErrAndPanic(err)
	count, err := lvdao.CountByNamedSql(ibatis.GetCountSql(), &req)
	lv_err.HasErrAndPanic(err)
	lv_web.PageOK2(c, listMap, count)
}

func (w DemoController) MybatisStruct(c *gin.Context) {
	req := vo.SelectPostPageReq{}
	if err := c.ShouldBind(&req); err != nil { //获取参数
		lv_web.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	ibatis := lvbatis.NewInstance(SQL_FILE_POST)
	sql, err := ibatis.GetLimitSql("listSql", &req)
	lv_err.HasErrAndPanic(err)
	list, err := lvdao.ListByNamedSql[model.SysPost](sql, &req)
	lv_err.HasErrAndPanic(err)
	count, err := lvdao.CountByNamedSql(ibatis.GetCountSql(), &req)
	lv_err.HasErrAndPanic(err)
	lv_web.PageOK2(c, list, count)
}

/**
 * 基于ibatis 的分页查询演示
 */
func (w DemoController) MybatisStructPage(c *gin.Context) {
	req := vo.SelectPostPageReq{}
	if err := c.ShouldBind(&req); err != nil { //获取参数
		lv_web.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	resp := lvdao.GetPageByNamedSql[model.SysPost](SQL_FILE_POST, "listSql", &req)
	lv_web.PageOK(c, resp)
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
	lv_web.SucessData(c, data1)
}
