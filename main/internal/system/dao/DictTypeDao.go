package dao

import (
	"common/common_vo"
	"errors"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/namedsql"
	"github.com/lostvip-com/lv_framework/utils/lv_office"
	"main/internal/system/model"
	"xorm.io/builder"
)

type DictTypeDao struct {
}

// 根据条件分页查询数据
func (dao *DictTypeDao) SelectListByPage(param *common_vo.DictTypePageReq) ([]model.SysDictType, int64, error) {
	db := lv_db.GetMasterGorm()
	if db == nil {
		return nil, 0, errors.New("获取数据库连接失败")
	}
	tb := db.Table("sys_dict_type t")
	if param != nil {
		if param.DictName != "" {
			tb.Where("t.dict_name like ?", "%"+param.DictName+"%")
		}
		if param.DictType != "" {
			tb.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}
		if param.Status != "" {
			tb.Where("t.status = ", param.Status)
		}
		if param.BeginTime != "" {
			tb.Where("t.create_time >= ?", param.BeginTime)
		}
		if param.EndTime != "" {
			tb.Where("t.create_time <= ?", param.EndTime)
		}
	}
	var total int64
	tb = tb.Count(&total)
	var result []model.SysDictType
	tb = tb.Order(param.SortName + " " + param.SortOrder).Offset(param.GetStartNum()).Limit(param.GetPageSize())
	tb = tb.Find(&result)
	return result, total, tb.Error
}

// 导出excel
func (dao *DictTypeDao) SelectListExport(param *common_vo.DictTypePageReq, head, col []string) (string, error) {
	gdb := lv_db.GetMasterGorm()
	if gdb == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From("sys_dict_type", "t")
	if param != nil {
		if param.DictName != "" {
			build.Where(builder.Like{"t.dict_name", param.DictName})
		}

		if param.DictType != "" {
			build.Where(builder.Like{"t.dict_type", param.DictType})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"t.status": param.Status})
		}

		if param.BeginTime != "" {
			build.Where(builder.Gte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.BeginTime + "','%y%m%d')"})
		}

		if param.EndTime != "" {
			build.Where(builder.Lte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.EndTime + "','%y%m%d')"})
		}
	}

	sqlStr, _, _ := build.ToSQL()
	arr, err := namedsql.ListArrStr(gdb, sqlStr, map[string]any{})

	path, err := lv_office.DownlaodExcel(head, *arr)

	return path, err
}

// 获取所有数据
func (dao *DictTypeDao) SelectListAll(param *common_vo.DictTypePageReq) ([]model.SysDictType, error) {
	gdb := lv_db.GetMasterGorm()
	if gdb == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	tb := gdb.Table("sys_dict_type t")

	if param != nil {
		if param.DictName != "" {
			tb.Where("t.dict_name like ?", "%"+param.DictName+"%")
		}

		if param.DictType != "" {
			tb.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.Status != "" {
			tb.Where("t.status = ", param.Status)
		}

		if param.BeginTime != "" {
			tb.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			tb.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	var result []model.SysDictType
	tb = tb.Find(&result)
	return result, tb.Error
}
