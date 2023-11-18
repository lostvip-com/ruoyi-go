// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-02-18 15:44:13
// 生成路径: app/model/modules/log/log.go
// 生成人：yunjie
// ==========================================================================
package job_log

import (
	"errors"
	"lostvip.com/db"
	"lostvip.com/utils/excel"
	"lostvip.com/utils/page"
	"xorm.io/builder"
)

// 分页请求参数
type SelectPageReq struct {
	JobLogId      int64  `form:"jobLogId"`      //任务日志ID
	JobName       string `form:"jobName"`       //任务名称
	JobGroup      string `form:"jobGroup"`      //任务组名
	InvokeTarget  string `form:"invokeTarget"`  //调用目标字符串
	JobMessage    string `form:"jobMessage"`    //日志信息
	Status        string `form:"status"`        //执行状态（0正常 1失败）
	ExceptionInfo string `form:"exceptionInfo"` //异常信息
	BeginTime     string `form:"beginTime"`     //开始时间
	EndTime       string `form:"endTime"`       //结束时间
	PageNum       int    `form:"pageNum"`       //当前页码
	PageSize      int    `form:"pageSize"`      //每页数
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

		if param.JobName != "" {
			model.Where("t.job_name like ?", "%"+param.JobName+"%")
		}

		if param.JobGroup != "" {
			model.Where("t.job_group = ?", param.JobGroup)
		}

		if param.InvokeTarget != "" {
			model.Where("t.invoke_target = ?", param.InvokeTarget)
		}

		if param.JobMessage != "" {
			model.Where("t.job_message = ?", param.JobMessage)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
		}

		if param.ExceptionInfo != "" {
			model.Where("t.exception_info = ?", param.ExceptionInfo)
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

	err = model.Find(&result)

	if err != nil {
		return nil, p, errors.New("读取数据失败")
	}
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

		if param.JobName != "" {
			build.Where(builder.Like{"t.job_name", param.JobName})
		}

		if param.JobGroup != "" {
			build.Where(builder.Eq{"t.job_group": param.JobGroup})
		}

		if param.InvokeTarget != "" {
			build.Where(builder.Eq{"t.invoke_target": param.InvokeTarget})
		}

		if param.JobMessage != "" {
			build.Where(builder.Eq{"t.job_message": param.JobMessage})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"t.status": param.Status})
		}

		if param.ExceptionInfo != "" {
			build.Where(builder.Eq{"t.exception_info": param.ExceptionInfo})
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

		if param.JobName != "" {
			model.Where("t.job_name like ?", "%"+param.JobName+"%")
		}

		if param.JobGroup != "" {
			model.Where("t.job_group = ?", param.JobGroup)
		}

		if param.InvokeTarget != "" {
			model.Where("t.invoke_target = ?", param.InvokeTarget)
		}

		if param.JobMessage != "" {
			model.Where("t.job_message = ?", param.JobMessage)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
		}

		if param.ExceptionInfo != "" {
			model.Where("t.exception_info = ?", param.ExceptionInfo)
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

	if err != nil {
		return nil, errors.New("读取数据失败")
	}
	return result, nil
}
