package dao

import (
	"errors"
	"github.com/lostvip-com/lv_framework/db"
	"github.com/lostvip-com/lv_framework/utils/lv_office"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	"main/internal/system/model"
	"main/internal/system/vo"
	"xorm.io/builder"
)

type DictTypeDao struct {
}

// 根据条件分页查询数据
func (dao *DictTypeDao) SelectListByPage(param *vo.SelectDictTypePageReq) ([]model.DictType, *lv_web.Paging, error) {
	db := db.GetInstance().Engine()
	p := new(lv_web.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	tb := db.Table("sys_dict_type").Alias("t")

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

	totalModel := tb.Clone()

	total, err := totalModel.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = lv_web.CreatePaging(param.PageNum, param.PageSize, int64(total))

	if param.SortName != "" {
		tb.OrderBy(param.SortName + " " + param.SortOrder + " ")
	}

	tb.Limit(p.PageSize, p.StartNum)

	var result []model.DictType
	err = tb.Find(&result)
	return result, p, err
}

// 导出excel
func (dao *DictTypeDao) SelectListExport(param *vo.SelectDictTypePageReq, head, col []string) (string, error) {
	db := db.GetInstance().Engine()

	if db == nil {
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
	arr, err := db.SQL(sqlStr).QuerySliceString()

	path, err := lv_office.DownlaodExcel(head, arr)

	return path, err
}

// 获取所有数据
func (dao *DictTypeDao) SelectListAll(param *vo.SelectDictTypePageReq) ([]model.DictType, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	tb := db.Table("sys_dict_type").Alias("t")

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

	var result []model.DictType
	err := tb.Find(&result)
	return result, err
}

// 校验字典类型是否唯一
func (dao *DictTypeDao) CheckDictTypeUniqueAll(dictType string) (*model.DictType, error) {
	var entity model.DictType
	entity.DictType = dictType
	ok, err := entity.FindOne()
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}

// 批量删除
func (dao *DictTypeDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.GetInstance().Engine().Table("sys_dict_type").In("dict_id", ids).Delete(new(model.DictType))
}

// 指定字段集合查询
func (dao *DictTypeDao) FindIn(column string, args ...interface{}) ([]model.DictType, error) {
	var list []model.DictType
	err := db.GetInstance().Engine().Table("sys_dict_type").In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func (dao *DictTypeDao) FindNotIn(column string, args ...interface{}) ([]model.DictType, error) {
	var list []model.DictType
	err := db.GetInstance().Engine().Table("sys_dict_type").NotIn(column, args).Find(&list)
	return list, err
}
