package dao

import (
	"common/common_vo"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/namedsql"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/spf13/cast"
	"main/internal/system/model"
)

type ConfigDao struct {
}

// 批量删除
func (d *ConfigDao) Count(configKey string) (total int64, err error) {
	err = lv_db.GetMasterGorm().Table("sys_config").Where("config_key=?", configKey).Count(&total).Error
	return total, err
}

// 批量删除
func (d *ConfigDao) DeleteBatch(ids ...int64) error {
	err := lv_db.GetMasterGorm().Delete(&model.SysConfig{}, ids).Error
	return err
}

// 根据条件分页查询用户列表
func (d ConfigDao) SelectPageList(param *common_vo.SelectConfigPageReq) (*[]map[string]string, int64, error) {
	db := lv_db.GetMasterGorm()
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

func (d ConfigDao) GetSql(param *common_vo.SelectConfigPageReq) (map[string]interface{}, string) {
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
//	lv_db := lv_db.GetInstance().Engine()
//	p := new(lv_web.Paging)
//	if lv_db == nil {
//		return nil, p, errors.New("获取数据库连接失败")
//	}
//	tb := lv_db.Table("sys_config").Alias("t")
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
func (d ConfigDao) SelectExportList(param *common_vo.SelectConfigPageReq) (*[]map[string]string, error) {
	db := lv_db.GetMasterGorm()
	sqlParams, sql := d.GetSql(param)
	limitSql := sql + " order by u.user_id desc "
	result, err := namedsql.ListMap(db, limitSql, &sqlParams, false)
	return result, err
}

// 获取所有数据
func (dao *ConfigDao) SelectListAll(param *common_vo.SelectConfigPageReq) ([]model.SysConfig, error) {
	db := lv_db.GetMasterGorm()

	tb := db.Table("sys_config t")

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
	err := tb.Find(&result).Error
	return result, err
}

func (dao *ConfigDao) FindOne(configKey string) (*model.SysConfig, error) {
	var entity model.SysConfig
	entity.ConfigKey = configKey
	err := entity.FindOne()
	return &entity, err
}
