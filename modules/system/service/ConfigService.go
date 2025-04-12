package service

import (
	"common/common_vo"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_cache"
	"github.com/lostvip-com/lv_framework/lv_cache/lv_ram"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_office"
	"github.com/spf13/cast"
	dao2 "system/dao"
	"system/model"
	"time"
)

type ConfigService struct {
}

// 根据键获取值
func (svc *ConfigService) GetValueFromCache(key string) string {
	//从缓存读取
	result, err := lv_cache.GetCacheClient().Get(key)
	if err != nil {
		entity := &model.SysConfig{ConfigKey: key}
		err := entity.FindOne()
		lv_err.HasErrAndPanic(err)
		result = entity.ConfigValue
		lv_cache.GetCacheClient().Set(key, result, 10*time.Minute)
	}

	return result
}

// 根据键获取值
func (svc *ConfigService) GetValueFromRam(key string) string {
	// 这里为了提高速度使用内在缓存
	result, err := lv_ram.GetRamCacheClient().Get(key)
	if err != nil {
		entity := &model.SysConfig{ConfigKey: key}
		err := entity.FindOne()
		if err != nil {
			panic(errors.New("获取配置失败,检查配置表：sys_config 中是否存在配置项：" + key))
		}
		result = entity.ConfigValue
		//内存续期
		lv_ram.GetRamCacheClient().Set(key, result, 10*time.Minute)
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
		lv_cache.GetCacheClient().Del(entity.ConfigKey)
		return true
	}
	return false
}

// 批量删除数据记录
func (svc *ConfigService) DeleteRecordByIds(ids string) {
	idarr := lv_conv.ToInt64Array(ids, ",")
	cfg := new(model.SysConfig)
	for _, id := range idarr {
		cfg, err := cfg.FindById(cast.ToInt64(id))
		lv_err.HasErrAndPanic(err)
		cfg.Delete()
		//从缓存删除
		lv_cache.GetCacheClient().Del(cfg.ConfigKey)
	}
}

// 添加数据
func (svc *ConfigService) AddSave(req *common_vo.AddConfigReq, c *gin.Context) (int64, error) {
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
func (svc *ConfigService) EditSave(req *common_vo.EditConfigReq, c *gin.Context) {
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
	lv_cache.GetCacheClient().Set(entity.ConfigKey, entity.ConfigValue, 10*time.Minute)
}

// 根据条件分页查询角色数据
func (svc *ConfigService) SelectListAll(params *common_vo.SelectConfigPageReq) ([]model.SysConfig, error) {
	var config dao2.ConfigDao
	return config.SelectListAll(params)
}

// 根据条件分页查询角色数据
func (svc *ConfigService) SelectListByPage(params *common_vo.SelectConfigPageReq) (*[]map[string]string, int64, error) {
	var config dao2.ConfigDao
	return config.SelectPageList(params)
}

// 导出excel
func (svc *ConfigService) Export(param *common_vo.SelectConfigPageReq) (string, error) {
	head := []string{"参数主键", "参数名称", "参数键名", "参数键值", "系统内置（Y是 N否）", "状态"}
	col := []string{"config_id", "config_name", "config_key", "config_value", "config_type"}
	var d dao2.ConfigDao
	listMap, err := d.SelectExportList(param)
	lv_err.HasErrAndPanic(err)
	return lv_office.DownlaodExcelByListMapStr(&head, &col, listMap)
}

func (svc *ConfigService) CountKey(key string) (int64, error) {
	var config dao2.ConfigDao
	return config.Count(key)

}
