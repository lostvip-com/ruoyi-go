package dict_data

import (
	"errors"
	"lostvip.com/db"
	"lostvip.com/utils/excel"
	"lostvip.com/utils/page"
	"xorm.io/builder"
)

// Fill with you ideas below.
// 新增页面请求参数
type AddReq struct {
	DictLabel string `form:"dictLabel"  binding:"required"`
	DictValue string `form:"dictValue"  binding:"required"`
	DictType  string `form:"dictType"  binding:"required"`
	DictSort  int    `form:"dictSort"  binding:"required"`
	CssClass  string `form:"cssClass"`
	ListClass string `form:"listClass" binding:"required"`
	IsDefault string `form:"isDefault" binding:"required"`
	Status    string `form:"status"    binding:"required"`
	Remark    string `form:"remark"`
}

// 修改页面请求参数
type EditReq struct {
	DictCode  int64  `form:"dictCode" binding:"required"`
	DictLabel string `form:"dictLabel"  binding:"required"`
	DictValue string `form:"dictValue"  binding:"required"`
	DictType  string `form:"dictType"`
	DictSort  int    `form:"dictSort"  binding:"required"`
	CssClass  string `form:"cssClass"`
	ListClass string `form:"listClass" binding:"required"`
	IsDefault string `form:"isDefault" binding:"required"`
	Status    string `form:"status"    binding:"required"`
	Remark    string `form:"remark"`
}

// 分页请求参数
type SelectPageReq struct {
	DictType  string `form:"dictType"`  //字典名称
	DictLabel string `form:"dictLabel"` //字典标签
	Status    string `form:"status"`    //状态
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	PageNum   int    `form:"pageNum"`   //当前页码
	PageSize  int    `form:"pageSize"`  //每页数
}

// 根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) (*[]Entity, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("t")

	if param != nil {
		if param.DictLabel != "" {
			model.Where("t.dict_label like ?", "%"+param.DictLabel+"%")
		}

		if param.Status != "" {
			model.Where("t.status = ", param.Status)
		}

		if param.DictType != "" {
			model.Where("t.dict_type like ?", "%"+param.DictType+"%")
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

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

	model.Limit(p.Pagesize, p.StartNum)

	var result []Entity
	model.Find(&result)
	return &result, p, nil
}

// 导出excel
func SelectListExport(param *SelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(TableName(), "t")

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

	path, err := excel.DownlaodExcel(head, arr)

	return path, err
}

// 获取所有数据
func SelectListAll(param *SelectPageReq) ([]Entity, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("t")

	if param != nil {
		if param.DictLabel != "" {
			model.Where("t.dict_label like ?", "%"+param.DictLabel+"%")
		}

		if param.Status != "" {
			model.Where("t.status = ", param.Status)
		}

		if param.DictType != "" {
			model.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []Entity
	model.Find(&result)
	return result, nil
}
