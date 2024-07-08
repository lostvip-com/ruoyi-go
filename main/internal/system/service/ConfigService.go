package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/utils/lv_conv"
	"github.com/lv_framework/utils/lv_err"
	"github.com/lv_framework/utils/lv_office"
	"lostvip.com/cache/lv_redis"
	dao2 "main/internal/system/dao"
	"main/internal/system/model"
	"main/internal/system/vo"
	"time"
)

type ConfigService struct {
}

// 根据键获取值
func (svc *ConfigService) GetValueByKey(key string) string {
	//从缓存读取
	result := lv_redis.GetInstance().Get(context.Background(), key).Val()
	if result == "" {
		entity := &model.SysConfig{ConfigKey: key}
		err := entity.FindOne()
		lv_err.HasErrAndPanic(err)
		result = entity.ConfigValue
		lv_redis.GetInstance().SetEx(context.Background(), key, result, 10*time.Minute)
	}

	return result
}

// 根据主键查询数据
func (svc *ConfigService) SelectRecordById(id int64) (*model.SysConfig, error) {
	entity := &model.SysConfig{ConfigId: id}
	err := entity.FindOne()
	return entity, err
}

// 根据主键删除数据
func (svc *ConfigService) DeleteRecordById(id int64) bool {
	entity := &model.SysConfig{ConfigId: id}
	err := entity.FindOne()
	lv_err.HasErrAndPanic(err)
	err = entity.Delete()
	if err == nil {
		//从缓存删除
		lv_redis.GetInstance().Del(context.Background(), entity.ConfigKey)
		return true
	}
	return false
}

// 批量删除数据记录
func (svc *ConfigService) DeleteRecordByIds(ids string) {
	idarr := lv_conv.ToInt64Array(ids, ",")
	var dao dao2.ConfigDao
	list, err := dao.FindIn("config_id", idarr)
	lv_err.HasErrAndPanic(err)
	err = dao.DeleteBatch(idarr...)
	lv_err.HasErrAndPanic(err)
	for _, item := range list {
		//从缓存删除
		lv_redis.GetInstance().Del(context.Background(), item.ConfigKey)
	}
}

// 添加数据
func (svc *ConfigService) AddSave(req *vo.AddConfigReq, c *gin.Context) (int64, error) {
	var entity model.SysConfig
	entity.ConfigName = req.ConfigName
	entity.ConfigKey = req.ConfigKey
	entity.ConfigType = req.ConfigType
	entity.ConfigValue = req.ConfigValue
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()
	entity.CreateBy = ""
	var userService UserService
	user := userService.GetProfile(c)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	err := entity.Save()
	return entity.ConfigId, err
}

// 修改数据
func (svc *ConfigService) EditSave(req *vo.EditConfigReq, c *gin.Context) {
	entity := &model.SysConfig{ConfigId: req.ConfigId}
	err := entity.FindOne()
	lv_err.HasErrAndPanic(err)
	entity.ConfigName = req.ConfigName
	entity.ConfigKey = req.ConfigKey
	entity.ConfigValue = req.ConfigValue
	entity.Remark = req.Remark
	entity.ConfigType = req.ConfigType
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""
	var userService UserService
	user := userService.GetProfile(c)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	err = entity.Update()
	lv_err.HasErrAndPanic(err)
	//保存到缓存
	lv_redis.GetInstance().SetEx(context.Background(), entity.ConfigKey, entity.ConfigValue, 10*time.Minute)
}

// 根据条件分页查询角色数据
func (svc *ConfigService) SelectListAll(params *vo.SelectConfigPageReq) ([]model.SysConfig, error) {
	var config dao2.ConfigDao
	return config.SelectListAll(params)
}

// 根据条件分页查询角色数据
func (svc *ConfigService) SelectListByPage(params *vo.SelectConfigPageReq) (*[]map[string]string, int64, error) {
	var config dao2.ConfigDao
	return config.SelectPageList(params)
}

// 导出excel
func (svc *ConfigService) Export(param *vo.SelectConfigPageReq) (string, error) {
	head := []string{"参数主键", "参数名称", "参数键名", "参数键值", "系统内置（Y是 N否）", "状态"}
	col := []string{"config_id", "config_name", "config_key", "config_value", "config_type"}
	var d dao2.ConfigDao
	listMap, err := d.SelectExportList(param)
	lv_err.HasErrAndPanic(err)
	return lv_office.DownlaodExcelByListMap(&head, &col, listMap)
}

// 检查角色名是否唯一
func (svc *ConfigService) CheckConfigKeyUniqueAll(configKey string) string {
	var config dao2.ConfigDao
	entity, err := config.CheckPostCodeUniqueAll(configKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.ConfigId > 0 {
		return "1"
	}
	return "0"
}

// 检查岗位名称是否唯一
func (svc *ConfigService) CheckConfigKeyUnique(configKey string, configId int64) string {
	var config dao2.ConfigDao
	entity, err := config.CheckPostCodeUniqueAll(configKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.ConfigId > 0 && entity.ConfigId != configId {
		return "1"
	}
	return "0"
}
