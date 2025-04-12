package pages

import (
	util2 "common/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_cache"
	"github.com/lostvip-com/lv_framework/lv_db/lv_batis"
	"github.com/lostvip-com/lv_framework/lv_db/lv_dao"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"system/model"
	"system/vo"
	"time"
)

// 相对于mapper目录的路径
var SQL_FILE_POST = "sys_post/sys_post_mapper.sql"

func (w DemoController) MybatisMap(c *gin.Context) {
	req := vo.SelectPostPageReq{}
	if err := c.ShouldBind(&req); err != nil { //获取参数
		util2.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}

	ibatis := lv_batis.NewInstance(SQL_FILE_POST)
	sql, err := ibatis.GetLimitSql("listSql", &req)
	lv_err.HasErrAndPanic(err)
	listMap, err := lv_dao.ListMapStrByNamedSql(sql, &req, true)
	lv_err.HasErrAndPanic(err)
	count, err := lv_dao.CountByNamedSql(ibatis.GetCountSql(), &req)
	lv_err.HasErrAndPanic(err)
	util2.PageOK2(c, listMap, count)
}

func (w DemoController) MybatisStruct(c *gin.Context) {
	req := vo.SelectPostPageReq{}
	if err := c.ShouldBind(&req); err != nil { //获取参数
		util2.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	ibatis := lv_batis.NewInstance(SQL_FILE_POST)
	sql, err := ibatis.GetLimitSql("listSql", &req)
	lv_err.HasErrAndPanic(err)
	list, err := lv_dao.ListByNamedSql[model.SysPost](sql, &req)
	lv_err.HasErrAndPanic(err)
	count, err := lv_dao.CountByNamedSql(ibatis.GetCountSql(), &req)
	lv_err.HasErrAndPanic(err)
	util2.PageOK2(c, list, count)
}

/**
 * 基于ibatis 的分页查询演示
 */
func (w DemoController) MybatisStructPage(c *gin.Context) {
	req := vo.SelectPostPageReq{}
	if err := c.ShouldBind(&req); err != nil { //获取参数
		util2.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	resp := lv_dao.GetPageByNamedSql[model.SysPost](SQL_FILE_POST, "listSql", &req)
	util2.PageOK(c, resp)
}

func (w DemoController) TestRedis(c *gin.Context) {
	//reid
	redis := lv_cache.GetCacheClient()
	data := map[string]any{"test": "123"}
	redis.HSet("mapKey1", data)

	fieldMap := make(map[string]interface{})
	fieldMap["field1"] = "val1"
	fieldMap["field2"] = "val2"
	redis.HSet("key", fieldMap)
	redis.Expire("key", 100*time.Second)
	fmt.Println("------------myredis----------------------123")
	data1, _ := redis.HGet("mapKey1", "test")
	fmt.Println(data1)
	util2.SucessData(c, data1)
}
