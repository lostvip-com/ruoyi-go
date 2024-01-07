package oper_log

import (
	"errors"
	"lostvip.com/db"
	"lostvip.com/utils/lv_office"
	"lostvip.com/utils/lv_web"
	"xorm.io/builder"
)

// 查询列表请求参数
type SelectPageReq struct {
	Title         string `form:"title"`         //系统模块
	OperName      string `form:"operName"`      //操作人员
	BusinessTypes int    `form:"businessTypes"` //操作类型
	Status        string `form:"status"`        //操作类型
	BeginTime     string `form:"beginTime"`     //数据范围
	EndTime       string `form:"endTime"`       //开始时间
	PageNum       int    `form:"pageNum"`       //当前页码
	PageSize      int    `form:"pageSize"`      //每页数
	OrderByColumn string `form:"orderByColumn"` //排序字段
	IsAsc         string `form:"isAsc"`         //排序方式
}

// 根据条件分页查询用户列表
func SelectPageList(param *SelectPageReq) (*[]Entity, *lv_web.Paging, error) {
	db := db.GetInstance().Engine()
	p := new(lv_web.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName())

	if param != nil {
		if param.Title != "" {
			model.Where("title like ?", "%"+param.Title+"%")
		}

		if param.OperName != "" {
			model.Where("oper_name like ?", "%"+param.OperName+"%")
		}

		if param.Status != "" {
			model.Where("status = ?", param.Status)
		}

		if param.BusinessTypes >= 0 {
			model.Where("status = ?", param.BusinessTypes)
		}

		if param.BeginTime != "" {
			model.Where("date_format(oper_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(oper_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}
	}

	totalModel := model.Clone()

	total, err := totalModel.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = lv_web.CreatePaging(param.PageNum, param.PageSize, int(total))

	if param.OrderByColumn != "" {
		model.OrderBy(param.OrderByColumn + " " + param.IsAsc + " ")
	}

	model.Limit(p.Pagesize, p.StartNum)

	var result []Entity

	err = model.Find(&result)
	return &result, p, nil
}

// 导出excel
func SelectExportList(param *SelectPageReq, head, col []string) (string, error) {
	db := db.GetInstance().Engine()
	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(TableName())

	if param != nil {
		if param.Title != "" {
			build.Where(builder.Like{"title", param.Title})
		}

		if param.OperName != "" {
			build.Where(builder.Like{"oper_name", param.OperName})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"status": param.Status})
		}

		if param.BusinessTypes >= 0 {
			build.Where(builder.Eq{"business_type": param.BusinessTypes})
		}

		if param.BeginTime != "" {
			build.Where(builder.Gte{"date_format(create_time,'%y%m%d')": "date_format('" + param.BeginTime + "','%y%m%d')"})
		}

		if param.EndTime != "" {
			build.Where(builder.Lte{"date_format(create_time,'%y%m%d')": "date_format('" + param.EndTime + "','%y%m%d')"})
		}
	}

	sqlStr, _, _ := build.ToSQL()
	arr, err := db.SQL(sqlStr).QuerySliceString()

	path, err := lv_office.DownlaodExcel(head, arr)

	return path, err
}

// 清空记录
func DeleteAll() (int64, error) {
	db := db.GetInstance().Engine()
	if db == nil {
		return 0, errors.New("获取数据库连接失败")
	}

	rs, _ := db.Exec("delete from sys_oper_log")

	return rs.RowsAffected()
}
