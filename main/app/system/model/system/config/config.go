package config

import (
	"errors"
	"lostvip.com/db"
	"lostvip.com/utils/lv_office"
	"lostvip.com/utils/lv_web"
	"xorm.io/builder"
)

// Fill with you ideas below.
// 新增页面请求参数
type AddReq struct {
	ConfigName  string `form:"configName"  binding:"required"`
	ConfigKey   string `form:"configKey"  binding:"required"`
	ConfigValue string `form:"configValue"  binding:"required"`
	ConfigType  string `form:"configType"    binding:"required"`
	Remark      string `form:"remark"`
}

// 修改页面请求参数
type EditReq struct {
	ConfigId    int64  `form:"configId" binding:"required"`
	ConfigName  string `form:"configName"  binding:"required"`
	ConfigKey   string `form:"configKey"  binding:"required"`
	ConfigValue string `form:"configValue"  binding:"required"`
	ConfigType  string `form:"configType"    binding:"required"`
	Remark      string `form:"remark"`
}

// 分页请求参数
type SelectPageReq struct {
	ConfigName string `form:"configName"` //参数名称
	ConfigKey  string `form:"configKey"`  //参数键名
	ConfigType string `form:"configType"` //状态
	BeginTime  string `form:"beginTime"`  //开始时间
	EndTime    string `form:"endTime"`    //结束时间
	PageNum    int    `form:"pageNum"`    //当前页码
	PageSize   int    `form:"pageSize"`   //每页数
}

// 检查参数键名请求参数
type CheckConfigKeyReq struct {
	ConfigId  int64  `form:"configId"  binding:"required"`
	ConfigKey string `form:"configKey"  binding:"required"`
}

// 检查参数键名请求参数
type CheckPostCodeALLReq struct {
	ConfigKey string `form:"configKey"  binding:"required"`
}

// 根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) ([]Entity, *lv_web.Paging, error) {
	db := db.GetInstance().Engine()
	p := new(lv_web.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("t")

	if param != nil {
		if param.ConfigName != "" {
			model.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}

		if param.ConfigType != "" {
			model.Where("t.config_type = ?", param.ConfigType)
		}

		if param.ConfigKey != "" {
			model.Where("t.config_key like ?", "%"+param.ConfigKey+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	totalModel := model.Clone()

	total, err := totalModel.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = lv_web.CreatePaging(param.PageNum, param.PageSize, int(total))

	model.Limit(p.Pagesize, p.StartNum)

	var result []Entity
	err = model.Find(&result)
	return result, p, err
}

// 导出excel
func SelectListExport(param *SelectPageReq, head, col []string) (string, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(TableName(), "t")

	if param != nil {
		if param.ConfigName != "" {
			build.Where(builder.Like{"t.config_name", param.ConfigName})
		}

		if param.ConfigType != "" {
			build.Where(builder.Eq{"t.status": param.ConfigType})
		}

		if param.ConfigKey != "" {
			build.Where(builder.Like{"t.config_key", param.ConfigKey})
		}

		if param.BeginTime != "" {
			build.Where(builder.Gte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.BeginTime + "','%y%m%d')"})
		}

		if param.EndTime != "" {
			build.Where(builder.Lte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.EndTime + "','%y%m%d')"})
		}
	}

	sqlStr, _, _ := build.ToSQL()
	arr, err := db.SQL(sqlStr).QuerySliceString()

	path, err := lv_office.DownlaodExcel(head, arr)

	return path, err
}

// 获取所有数据
func SelectListAll(param *SelectPageReq) ([]Entity, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("t")

	if param != nil {
		if param.ConfigName != "" {
			model.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}

		if param.ConfigType != "" {
			model.Where("t.status = ", param.ConfigType)
		}

		if param.ConfigKey != "" {
			model.Where("t.config_key like ?", "%"+param.ConfigKey+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []Entity
	err := model.Find(&result)
	return result, err
}

// 校验参数键名是否唯一
func CheckPostCodeUniqueAll(configKey string) (*Entity, error) {
	var entity Entity
	entity.ConfigKey = configKey
	ok, err := entity.FindOne()
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}
