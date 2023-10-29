// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-06-29 23:13:41 +0800 CST
// 生成路径: app/model/modules/config/config.go
// 生成人：robnote
// ==========================================================================
package config

import (
	"errors"
	"lostvip.com/db"
	"lostvip.com/utils/excel"
	"lostvip.com/utils/page"
	"xorm.io/builder"
)

// 新增页面请求参数
type AddReq struct {
	ConfigName  string `form:"configName" binding:"required"`
	ConfigKey   string `form:"configKey" `
	ConfigValue string `form:"configValue" `
	ConfigType  string `form:"configType" `

	Remark string `form:"remark" `
}

// 修改页面请求参数
type EditReq struct {
	ConfigId    int64  `form:"configId" binding:"required"`
	ConfigName  string `form:"configName" binding:"required不能为空"`
	ConfigKey   string `form:"configKey" `
	ConfigValue string `form:"configValue" `
	ConfigType  string `form:"configType" `
	Remark      string `form:"remark" `
}

// 分页请求参数
type SelectPageReq struct {
	ConfigId    int    `form:"configId"`    //参数主键
	ConfigName  string `form:"configName"`  //参数名称
	ConfigKey   string `form:"configKey"`   //参数键名
	ConfigValue string `form:"configValue"` //参数键值
	ConfigType  string `form:"configType"`  //系统内置（Y是 N否）
	BeginTime   string `form:"beginTime"`   //开始时间
	EndTime     string `form:"endTime"`     //结束时间
	PageNum     int    `form:"pageNum"`     //当前页码
	PageSize    int    `form:"pageSize"`    //每页数
}

// 根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) ([]SysConfig, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)

	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table("sys_config").Alias("t")

	if param != nil {

		if param.ConfigId != 0 {
			model.Where("t.config_id = ?", param.ConfigId)
		}

		if param.ConfigName != "" {
			model.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}

		if param.ConfigKey != "" {
			model.Where("t.config_key = ?", param.ConfigKey)
		}

		if param.ConfigValue != "" {
			model.Where("t.config_value = ?", param.ConfigValue)
		}

		if param.ConfigType != "" {
			model.Where("t.config_type = ?", param.ConfigType)
		}
		if param.BeginTime != "" {
			model.Where("t.create_time >= ?", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("t.create_time<=?", param.EndTime)
		}
	}

	totalModel := model.Clone()
	total, err := totalModel.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

	model.Limit(p.Pagesize, p.StartNum)

	var result []SysConfig
	err = model.Find(&result)
	return result, p, err
}

// 导出excel
func SelectListExport(param *SelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From("sys_config", "t")

	if param != nil {

		if param.ConfigId != 0 {
			build.Where(builder.Eq{"t.config_id": param.ConfigId})
		}

		if param.ConfigName != "" {
			build.Where(builder.Like{"t.config_name", param.ConfigName})
		}

		if param.ConfigKey != "" {
			build.Where(builder.Eq{"t.config_key": param.ConfigKey})
		}

		if param.ConfigValue != "" {
			build.Where(builder.Eq{"t.config_value": param.ConfigValue})
		}

		if param.ConfigType != "" {
			build.Where(builder.Eq{"t.config_type": param.ConfigType})
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

	path, err := excel.DownlaodExcel(head, arr)

	return path, err
}

// 获取所有数据
func SelectListAll(param *SelectPageReq) ([]SysConfig, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table("sys_config").Alias("t")

	if param != nil {

		if param.ConfigId != 0 {
			model.Where("t.config_id = ?", param.ConfigId)
		}

		if param.ConfigName != "" {
			model.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}

		if param.ConfigKey != "" {
			model.Where("t.config_key = ?", param.ConfigKey)
		}

		if param.ConfigValue != "" {
			model.Where("t.config_value = ?", param.ConfigValue)
		}

		if param.ConfigType != "" {
			model.Where("t.config_type = ?", param.ConfigType)
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []SysConfig
	err := model.Find(&result)
	return result, err
}
