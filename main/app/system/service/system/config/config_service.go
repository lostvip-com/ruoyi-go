package config

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"lostvip.com/cache/myredis"
	"lostvip.com/conf"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	"robvi/app/system/model/system/config"
	"robvi/app/system/service"
	"strings"

	"time"
)

func GetCopyright() string {
	ctx := conf.Config().GetConf("copyright")
	return ctx
}

func GetCtx() string {
	ctxPath := conf.Config().GetContextPath()
	return ctxPath
}

func GetCtxPath(url string) string {
	ctxPath := conf.Config().GetContextPath()
	if !strings.HasPrefix(url, "http") {
		if strings.HasPrefix(url, "/") {
			ctxPath = ctxPath + url
		} else {
			ctxPath = ctxPath + "/" + url
		}
	}
	return ctxPath
}
func GetOssUrl() string {
	ossUrl := GetValueByKey("sys.resource.url")
	if ossUrl == "" {
		ossUrl = "/static"
	} else if !strings.HasPrefix(ossUrl, "http") {
		ossUrl = conf.Config().GetContextPath() + ossUrl
	}
	return ossUrl
}

// 根据用户id和权限字符串判断是否有此权限
func AddInt(a, b int) int {
	return a + b
}

// 根据键获取值
func GetValueByKey(key string) string {
	//从缓存读取
	result := myredis.GetInstance().Get(context.Background(), key).Val()
	if result == "" {
		entity := &config.Entity{ConfigKey: key}
		ok, _ := entity.FindOne()
		if !ok {
			return ""
		}
		result = entity.ConfigValue
		myredis.GetInstance().SetEx(context.Background(), key, result, 10*time.Minute)
	}

	return result
}

// 根据主键查询数据
func SelectRecordById(id int64) (*config.Entity, error) {
	entity := &config.Entity{ConfigId: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除数据
func DeleteRecordById(id int64) bool {
	entity := &config.Entity{ConfigId: id}
	ok, _ := entity.FindOne()
	if ok {
		result, err := entity.Delete()
		if err == nil {
			if result > 0 {
				//从缓存删除
				myredis.GetInstance().Del(context.Background(), entity.ConfigKey)
				return true
			}
		}
	}
	return false
}

// 批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	idarr := lv_conv.ToInt64Array(ids, ",")
	list, _ := config.FindIn("config_id", idarr)
	rs, err := config.DeleteBatch(idarr...)
	if err != nil {
		return 0
	}

	if len(list) > 0 {
		for _, item := range list {
			//从缓存删除
			myredis.GetInstance().Del(context.Background(), item.ConfigKey)
		}
	}

	return rs
}

// 添加数据
func AddSave(req *config.AddReq, c *gin.Context) (int64, error) {
	var entity config.Entity
	entity.ConfigName = req.ConfigName
	entity.ConfigKey = req.ConfigKey
	entity.ConfigType = req.ConfigType
	entity.ConfigValue = req.ConfigValue
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()
	entity.CreateBy = ""
	var userService service.UserService
	user := userService.GetProfile(c)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	_, err := entity.Insert()
	return entity.ConfigId, err
}

// 修改数据
func EditSave(req *config.EditReq, c *gin.Context) (int64, error) {
	entity := &config.Entity{ConfigId: req.ConfigId}
	ok, err := entity.FindOne()

	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}

	entity.ConfigName = req.ConfigName
	entity.ConfigKey = req.ConfigKey
	entity.ConfigValue = req.ConfigValue
	entity.Remark = req.Remark
	entity.ConfigType = req.ConfigType
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""
	var userService service.UserService
	user := userService.GetProfile(c)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	rs, err := entity.Update()

	if err != nil {
		return 0, err
	}

	//保存到缓存
	myredis.GetInstance().SetEx(context.Background(), entity.ConfigKey, entity.ConfigValue, 10*time.Minute)

	return rs, nil
}

// 根据条件分页查询角色数据
func SelectListAll(params *config.SelectPageReq) ([]config.Entity, error) {
	return config.SelectListAll(params)
}

// 根据条件分页查询角色数据
func SelectListByPage(params *config.SelectPageReq) ([]config.Entity, *lv_web.Paging, error) {
	return config.SelectListByPage(params)
}

// 导出excel
func Export(param *config.SelectPageReq) (string, error) {
	head := []string{"参数主键", "参数名称", "参数键名", "参数键值", "系统内置（Y是 N否）", "状态"}
	col := []string{"config_id", "config_name", "config_key", "config_value", "config_type"}
	return config.SelectListExport(param, head, col)
}

// 检查角色名是否唯一
func CheckConfigKeyUniqueAll(configKey string) string {
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
func CheckConfigKeyUnique(configKey string, configId int64) string {
	entity, err := config.CheckPostCodeUniqueAll(configKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.ConfigId > 0 && entity.ConfigId != configId {
		return "1"
	}
	return "0"
}
