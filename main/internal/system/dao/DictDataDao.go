package dao

import (
	"errors"
	"github.com/lv_framework/db"
	"github.com/lv_framework/utils/lv_office"
	"github.com/lv_framework/utils/lv_web"
	"main/internal/system/model"
	"main/internal/system/vo"
	"xorm.io/builder"
)

type DictDataDao struct {
}

// 根据条件分页查询数据
func (dao *DictDataDao) SelectListByPage(param *vo.SelectDictDataPageReq) (*[]model.SysDictData, *lv_web.Paging, error) {
	db := db.GetInstance().Engine()
	p := new(lv_web.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	tb := db.Table("sys_dict_data").Alias("t")

	if param != nil {
		if param.DictLabel != "" {
			tb.Where("t.dict_label like ?", "%"+param.DictLabel+"%")
		}

		if param.Status != "" {
			tb.Where("t.status = ", param.Status)
		}

		if param.DictType != "" {
			tb.Where("t.dict_type like ?", "%"+param.DictType+"%")
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

	tb.Limit(p.PageSize, p.StartNum)

	var result []model.SysDictData
	tb.Find(&result)
	return &result, p, nil
}

// 导出excel
func (dao *DictDataDao) SelectListExport(param *vo.SelectDictDataPageReq, head, col []string) (string, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From("sys_dict_data", "t")

	if param != nil {
		if param.DictLabel != "" {
			build.Where(builder.Like{"t.dict_label", param.DictLabel})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"t.status": param.Status})
		}

		if param.DictType != "" {
			build.Where(builder.Like{"t.dict_type", param.DictType})
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
func (dao *DictDataDao) SelectListAll(param *vo.SelectDictDataPageReq) ([]model.SysDictData, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	tb := db.Table("sys_dict_data").Alias("t")

	if param != nil {
		if param.DictLabel != "" {
			tb.Where("t.dict_label like ?", "%"+param.DictLabel+"%")
		}

		if param.Status != "" {
			tb.Where("t.status = ", param.Status)
		}

		if param.DictType != "" {
			tb.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.BeginTime != "" {
			tb.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			tb.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []model.SysDictData
	tb.Find(&result)
	return result, nil
}
