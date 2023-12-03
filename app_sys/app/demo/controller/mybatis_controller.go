package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"lostvip.com/cache/myredis"
	"lostvip.com/db"
	"lostvip.com/db/ibatis"
	"lostvip.com/logme"
	"lostvip.com/utils/lv_web"
	model2 "robvi/app/biz/model"
	"time"
)

// 相对于mapper目录的路径
var SQL_FILE_PATIENT_MAPPER = "patient/his_patient_mapper.tpl"

func (w DemoController) TestMybatis1(c *gin.Context) {
	req := model2.PageHisPatientReq{Name: "test", PageNum: 1, PageSize: 100}
	if err := c.ShouldBind(req); err != nil { //获取参数
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("patient管理", req).WriteJsonExit()
		return
	}

	ibatis, err := ibatis.NewIBatis(SQL_FILE_PATIENT_MAPPER)
	list := []model2.HisPatient{}
	err = db.GetInstanceMaster().SQL(ibatis.GetLimitSql("test", req)).Find(&list)
	total, err := db.GetInstanceMaster().SQL(ibatis.GetCountSql()).Count()
	if err != nil {
		panic(err)
	} else {
		logme.Log.Info(list)
	}
	lv_web.PageOK(c, list, total, "")
}

func (w DemoController) TestMybatisStr2(c *gin.Context) {
	p := map[string]any{"name": "test", "pageSize": 200, "pageNum": 1}
	ibatis, err := ibatis.NewIBatis(SQL_FILE_PATIENT_MAPPER)
	list, err := db.GetInstanceMaster().QueryString(ibatis.GetLimitSql("test", p))
	if err != nil {
		panic(err)
	} else {
		logme.Log.Info(list)
	}
	total, err := db.GetInstanceMaster().SQL(ibatis.GetCountSql()).Count()
	if err != nil {
		panic(err)
	} else {
		logme.Log.Info(list)
	}
	lv_web.PageOK(c, list, total, "")
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
