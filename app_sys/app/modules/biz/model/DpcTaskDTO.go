// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-11-26 16:27:18 +0800 CST
// 生成人：lv
// ==========================================================================
package model

import (
	"errors"
	"lostvip.com/db"
	"lostvip.com/utils/excel"
	"lostvip.com/utils/page"
	"time"
	"xorm.io/builder"
)

// 新增页面请求参数
type AddDpcTaskReq struct {
	Username    string    `form:"username" binding:"required"`
	Password    string    `form:"password" `
	PrjCode     string    `form:"prjCode" `
	TaskContent string    `form:"taskContent" `
	StartDate   time.Time `form:"startDate" `
	EndDate     time.Time `form:"endDate" `
	WorkDays    int64     `form:"workDays" `
	AutoSubmit  string    `form:"autoSubmit" `
	Status      string    `form:"status" binding:"required"`
	Sort        int64     `form:"sort" `
}

// 修改页面请求参数
type EditDpcTaskReq struct {
	Id          int64     `form:"id" binding:"required"`
	Username    string    `form:"username" binding:"required"`
	Password    string    `form:"password" `
	PrjCode     string    `form:"prjCode" `
	TaskContent string    `form:"taskContent" `
	StartDate   time.Time `form:"startDate" `
	EndDate     time.Time `form:"endDate" `
	WorkDays    int64     `form:"workDays" `
	AutoSubmit  string    `form:"autoSubmit" `
	Status      string    `form:"status" binding:"required"`
	Sort        int64     `form:"sort" `
}

// 分页请求参数
type PageDpcTaskReq struct {
	Username    string `form:"username"`    //工号
	Password    string `form:"password"`    //密码
	PrjCode     string `form:"prjCode"`     //项  目  号
	TaskContent string `form:"taskContent"` //任务内容
	BeginTime   string `form:"beginTime"`   //开始时间
	EndTime     string `form:"endTime"`     //结束时间
	PageNum     int    `form:"pageNum"`     //当前页码
	PageSize    int    `form:"pageSize"`    //每页数
}

// 根据条件分页查询数据
func (e *DpcTask) SelectListByPage(param *PageDpcTaskReq) ([]DpcTask, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)

	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table("gen_table").Alias("t")

	if param != nil {

		if param.Username != "" {
			model.Where("t.username like ?", "%"+param.Username+"%")
		}

		if param.Password != "" {
			model.Where("t.password like ?", "%"+param.Password+"%")
		}

		if param.PrjCode != "" {
			model.Where("t.prj_code like ?", "%"+param.PrjCode+"%")
		}

		if param.TaskContent != "" {
			model.Where("t.task_content like ?", "%"+param.TaskContent+"%")
		}
		if param.BeginTime != "" {
			model.Where("t.create_time >= ?", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("t.create_time<=?", param.EndTime)
		}
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))
	model.Limit(p.Pagesize, p.StartNum)

	var result []DpcTask
	err = model.Find(&result)
	return result, p, err
}

// 导出excel
func (e *DpcTask) SelectListExport(param *PageDpcTaskReq, head, col []string) (string, error) {
	db := db.Instance().Engine()
	build := builder.Select(col...).From("gen_table", "t")

	if param != nil {

		if param.Username != "" {
			build.Where(builder.Like{"t.username", param.Username})
		}

		if param.Password != "" {
			build.Where(builder.Like{"t.password", param.Password})
		}

		if param.PrjCode != "" {
			build.Where(builder.Like{"t.prj_code", param.PrjCode})
		}

		if param.TaskContent != "" {
			build.Where(builder.Like{"t.task_content", param.TaskContent})
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
func (e *DpcTask) SelectListAll(param *PageDpcTaskReq) ([]DpcTask, error) {
	db := db.Instance().Engine()
	model := db.Table("gen_table").Alias("t")

	if param != nil {

		if param.Username != "" {
			model.Where("t.username like ?", "%"+param.Username+"%")
		}

		if param.Password != "" {
			model.Where("t.password like ?", "%"+param.Password+"%")
		}

		if param.PrjCode != "" {
			model.Where("t.prj_code like ?", "%"+param.PrjCode+"%")
		}

		if param.TaskContent != "" {
			model.Where("t.task_content like ?", "%"+param.TaskContent+"%")
		}
		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []DpcTask
	err := model.Find(&result)
	return result, err
}
