package dao

import (
	"errors"
	"github.com/spf13/cast"
	"lostvip.com/db"
	"lostvip.com/db/namedsql"
	"lostvip.com/utils/lv_err"
	"main/app/system/model"
	"main/app/system/vo"
)

type ConfigDao struct {
}

// 批量删除
func (d *ConfigDao) DeleteBatch(ids ...int64) error {
	err := db.GetMasterGorm().Delete(&model.SysConfig{}, ids).Error
	return err
}

// 根据条件分页查询用户列表
func (d ConfigDao) SelectPageList(param *vo.SelectConfigPageReq) (*[]map[string]string, int64, error) {
	db := db.GetMasterGorm()
	sqlParams, sql := d.GetSql(param)
	countSql := "select count(*) from (" + sql + ") t "
	total, err := namedsql.Count(db, countSql, sqlParams)
	lv_err.HasErrAndPanic(err)
	limitSql := sql + " order by u.config_id desc "
	limitSql += "  limit " + cast.ToString(param.GetStartNum()) + "," + cast.ToString(param.GetPageSize())
	result, err := namedsql.ListMap(db, limitSql, sqlParams, true)
	lv_err.HasErrAndPanic(err)
	return result, total, err
}

func (d ConfigDao) GetSql(param *vo.SelectConfigPageReq) (map[string]interface{}, string) {
	sqlParams := make(map[string]interface{})
	sql := `
           select * from sys_config u where 1=1
          `
	if param != nil {
		if param.ConfigName != "" {
			sql += " and  u.config_name like @ConfigName "
			sqlParams["ConfigName"] = "%" + param.ConfigName + "%"
		}
		if param.ConfigKey != "" {
			sql += " and  u.config_key like @ConfigKey "
			sqlParams["ConfigKey"] = "%" + param.ConfigKey + "%"
		}
		if param.ConfigType != "" {
			sql += " and  u.config_type like @ConfigType "
			sqlParams["ConfigType"] = "%" + param.ConfigType + "%"
		}

		if param.BeginTime != "" {
			sql += " and  u.create_time >= @BeginTime "
			sqlParams["BeginTime"] = param.BeginTime
		}
		if param.EndTime != "" {
			sql += " and  u.create_time <= @EndTime "
			sqlParams["EndTime"] = param.EndTime
		}

	}
	return sqlParams, sql
}

//
//// 根据条件分页查询数据
//func (dao *ConfigDao) SelectListByPage(param *vo.SelectConfigPageReq) ([]model.SysConfig, *lv_web.Paging, error) {
//	db := db.GetInstance().Engine()
//	p := new(lv_web.Paging)
//	if db == nil {
//		return nil, p, errors.New("获取数据库连接失败")
//	}
//	tb := db.Table("sys_config").Alias("t")
//
//	if param != nil {
//		if param.ConfigName != "" {
//			tb.Where("t.config_name like ?", "%"+param.ConfigName+"%")
//		}
//
//		if param.ConfigType != "" {
//			tb.Where("t.config_type = ?", param.ConfigType)
//		}
//
//		if param.ConfigKey != "" {
//			tb.Where("t.config_key like ?", "%"+param.ConfigKey+"%")
//		}
//
//		if param.BeginTime != "" {
//			tb.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
//		}
//
//		if param.EndTime != "" {
//			tb.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
//		}
//	}
//
//	totalModel := tb.Clone()
//
//	total, err := totalModel.Count()
//
//	if err != nil {
//		return nil, p, errors.New("读取行数失败")
//	}
//
//	p = lv_web.CreatePaging(param.PageNum, param.PageSize, int64(total))
//
//	tb.Limit(p.PageSize, p.StartNum)
//
//	var result []model.SysConfig
//	err = tb.Find(&result)
//	return result, p, err
//}

// 导出excel
func (d ConfigDao) SelectExportList(param *vo.SelectConfigPageReq) (*[]map[string]string, error) {
	db := db.GetMasterGorm()
	sqlParams, sql := d.GetSql(param)
	limitSql := sql + " order by u.user_id desc "
	result, err := namedsql.ListMap(db, limitSql, &sqlParams, false)
	return result, err
}

//// 导出excel
//func (dao *ConfigDao) SelectListExport(param *vo.SelectConfigPageReq, head, col []string) (string, error) {
//	db := db.GetMasterGorm()
//	build := builder.Select(col...).From("sys_config", "t")
//	if param != nil {
//		if param.ConfigName != "" {
//			build.Where(builder.Like{"t.config_name", param.ConfigName})
//		}
//
//		if param.ConfigType != "" {
//			build.Where(builder.Eq{"t.status": param.ConfigType})
//		}
//
//		if param.ConfigKey != "" {
//			build.Where(builder.Like{"t.config_key", param.ConfigKey})
//		}
//
//		if param.BeginTime != "" {
//			build.Where(builder.Gte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.BeginTime + "','%y%m%d')"})
//		}
//
//		if param.EndTime != "" {
//			build.Where(builder.Lte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.EndTime + "','%y%m%d')"})
//		}
//	}
//
//	sqlStr, _, _ := build.ToSQL()
//	arr, err := db.Raw(sqlStr).QuerySliceString()
//
//	path, err := lv_office.DownlaodExcel(head, arr)
//
//	return path, err
//}

// 获取所有数据
func (dao *ConfigDao) SelectListAll(param *vo.SelectConfigPageReq) ([]model.SysConfig, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	tb := db.Table("sys_config").Alias("t")

	if param != nil {
		if param.ConfigName != "" {
			tb.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}

		if param.ConfigType != "" {
			tb.Where("t.status = ", param.ConfigType)
		}

		if param.ConfigKey != "" {
			tb.Where("t.config_key like ?", "%"+param.ConfigKey+"%")
		}

		if param.BeginTime != "" {
			tb.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			tb.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []model.SysConfig
	err := tb.Find(&result)
	return result, err
}

// 校验参数键名是否唯一
func (dao *ConfigDao) CheckPostCodeUniqueAll(configKey string) (*model.SysConfig, error) {
	var entity model.SysConfig
	entity.ConfigKey = configKey
	err := entity.FindOne()
	return &entity, err
}

// 指定字段集合查询
func (r *ConfigDao) FindIn(column string, args ...interface{}) ([]model.SysConfig, error) {
	var list []model.SysConfig
	err := db.GetInstance().Engine().Table("sys_config").In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func (r *ConfigDao) FindNotIn(column string, args ...interface{}) ([]model.SysConfig, error) {
	var list []model.SysConfig
	err := db.GetInstance().Engine().Table("sys_config").NotIn(column, args).Find(&list)
	return list, err
}
