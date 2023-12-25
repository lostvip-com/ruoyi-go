package dict_type

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
	DictName string `form:"dictName"  binding:"required"`
	DictType string `form:"dictType"  binding:"required"`
	Status   string `form:"status"  binding:"required"`
	Remark   string `form:"remark"`
}

// 修改页面请求参数
type EditReq struct {
	DictId   int64  `form:"dictId" binding:"required"`
	DictName string `form:"dictName"  binding:"required"`
	DictType string `form:"dictType"  binding:"required"`
	Status   string `form:"status"  binding:"required"`
	Remark   string `form:"remark"`
}

// 分页请求参数
type SelectPageReq struct {
	DictName      string `form:"dictName"`      //字典名称
	DictType      string `form:"dictType"`      //字典类型
	Status        string `form:"status"`        //字典状态
	BeginTime     string `form:"beginTime"`     //开始时间
	EndTime       string `form:"endTime"`       //结束时间
	OrderByColumn string `form:"orderByColumn"` //排序字段
	IsAsc         string `form:"isAsc"`         //排序方式
	PageNum       int    `form:"pageNum"`       //当前页码
	PageSize      int    `form:"pageSize"`      //每页数
}

// 检查字典类型请求参数
type CheckDictTypeReq struct {
	DictId   int64  `form:"dictId"  binding:"required"`
	DictType string `form:"dictType"  binding:"required"`
}

// 检查字典类型请求参数
type CheckDictTypeALLReq struct {
	DictType string `form:"dictType"  binding:"required"`
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
		if param.DictName != "" {
			model.Where("t.dict_name like ?", "%"+param.DictName+"%")
		}

		if param.DictType != "" {
			model.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.Status != "" {
			model.Where("t.status = ", param.Status)
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

	if param.OrderByColumn != "" {
		model.OrderBy(param.OrderByColumn + " " + param.IsAsc + " ")
	}

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
func SelectListAll(param *SelectPageReq) ([]Entity, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("t")

	if param != nil {
		if param.DictName != "" {
			model.Where("t.dict_name like ?", "%"+param.DictName+"%")
		}

		if param.DictType != "" {
			model.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.Status != "" {
			model.Where("t.status = ", param.Status)
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

// 校验字典类型是否唯一
func CheckDictTypeUniqueAll(dictType string) (*Entity, error) {
	var entity Entity
	entity.DictType = dictType
	ok, err := entity.FindOne()
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}
