package logininfor

import (
	"errors"
	"lostvip.com/db"
	"lostvip.com/utils/excel"
	"lostvip.com/utils/page"
	"xorm.io/builder"
)

// Fill with you ideas below.
// 查询列表请求参数
type SelectPageReq struct {
	LoginName     string `form:"loginName"`     //登陆名
	Status        string `form:"status"`        //状态
	Ipaddr        string `form:"ipaddr"`        //登录地址
	BeginTime     string `form:"beginTime"`     //数据范围
	EndTime       string `form:"endTime"`       //开始时间
	PageNum       int    `form:"pageNum"`       //当前页码
	PageSize      int    `form:"pageSize"`      //每页数
	OrderByColumn string `form:"orderByColumn"` //排序字段
	IsAsc         string `form:"isAsc"`         //排序方式
}

// 根据条件分页查询用户列表
func SelectPageList(param *SelectPageReq) (*[]Entity, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName())

	if param != nil {
		if param.LoginName != "" {
			model.Where("login_name like ?", "%"+param.LoginName+"%")
		}

		if param.Ipaddr != "" {
			model.Where("ipaddr like ?", "%"+param.Ipaddr+"%")
		}

		if param.Status != "" {
			model.Where("status = ?", param.Status)
		}

		if param.BeginTime != "" {
			model.Where("date_format(login_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(login_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}
	}

	totalModel := model.Clone()

	total, err := totalModel.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

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
	db := db.Instance().Engine()
	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(TableName(), "t")

	if param != nil {
		if param.LoginName != "" {
			build.Where(builder.Like{"t.login_name", param.LoginName})
		}

		if param.Ipaddr != "" {
			build.Where(builder.Like{"t.ipaddr", param.Ipaddr})
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

	path, err := excel.DownlaodExcel(head, arr)

	return path, err
}

// 清空记录
func DeleteAll() (int64, error) {
	db := db.Instance().Engine()
	if db == nil {
		return 0, errors.New("获取数据库连接失败")
	}

	rs, _ := db.Exec("delete from sys_logininfor")

	return rs.RowsAffected()
}
