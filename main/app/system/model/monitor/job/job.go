// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-02-18 15:44:13
// 生成路径: app/model/job/job.go
// 生成人：yunjie
// ==========================================================================
package job

import (
	"errors"
	"lostvip.com/db"
	"lostvip.com/utils/lv_office"
	"lostvip.com/utils/lv_web"
	"xorm.io/builder"
)

// 新增页面请求参数
type AddReq struct {
	JobName        string `form:"jobName" `
	JobParams      string `form:"jobParams"` // 任务参数
	JobGroup       string `form:"jobGroup" `
	InvokeTarget   string `form:"invokeTarget" `
	CronExpression string `form:"cronExpression" `
	MisfirePolicy  string `form:"misfirePolicy" `
	Concurrent     string `form:"concurrent" `
	Status         string `form:"status" binding:"required"`
	Remark         string `form:"remark" `
}

// 修改页面请求参数
type EditReq struct {
	JobName        string `form:"jobName" `
	JobParams      string `form:"jobParams"` // 任务参数
	JobGroup       string `form:"jobGroup" `
	JobId          int64  `form:"jobId" binding:"required"`
	InvokeTarget   string `form:"invokeTarget" `
	CronExpression string `form:"cronExpression" `
	MisfirePolicy  string `form:"misfirePolicy" `
	Concurrent     string `form:"concurrent" `
	Status         string `form:"status" binding:"required"`
	Remark         string `form:"remark" `
}

// 分页请求参数
type SelectPageReq struct {
	JobId          int64  `form:"jobId"`          //任务ID
	JobName        string `form:"jobName"`        //任务名称
	JobGroup       string `form:"jobGroup"`       //任务组名
	InvokeTarget   string `form:"invokeTarget"`   //调用目标字符串
	CronExpression string `form:"cronExpression"` //cron执行表达式
	MisfirePolicy  string `form:"misfirePolicy"`  //计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Concurrent     string `form:"concurrent"`     //是否并发执行（0允许 1禁止）
	Status         string `form:"status"`         //状态（0正常 1暂停）
	BeginTime      string `form:"beginTime"`      //开始时间
	EndTime        string `form:"endTime"`        //结束时间
	PageNum        int    `form:"pageNum"`        //当前页码
	PageSize       int    `form:"pageSize"`       //每页数
	OrderByColumn  string `form:"orderByColumn"`  //排序字段
	IsAsc          string `form:"isAsc"`          //排序方式
}

// 根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) (*[]Entity, *lv_web.Paging, error) {
	db := db.Instance().Engine()
	p := new(lv_web.Paging)
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

		if param.CronExpression != "" {
			model.Where("t.cron_expression = ?", param.CronExpression)
		}

		if param.MisfirePolicy != "" {
			model.Where("t.misfire_policy = ?", param.MisfirePolicy)
		}

		if param.Concurrent != "" {
			model.Where("t.concurrent = ?", param.Concurrent)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
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

		if param.CronExpression != "" {
			build.Where(builder.Eq{"t.cron_expression": param.CronExpression})
		}

		if param.MisfirePolicy != "" {
			build.Where(builder.Eq{"t.misfire_policy": param.MisfirePolicy})
		}

		if param.Concurrent != "" {
			build.Where(builder.Eq{"t.concurrent": param.Concurrent})
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

		if param.CronExpression != "" {
			model.Where("t.cron_expression = ?", param.CronExpression)
		}

		if param.MisfirePolicy != "" {
			model.Where("t.misfire_policy = ?", param.MisfirePolicy)
		}

		if param.Concurrent != "" {
			model.Where("t.concurrent = ?", param.Concurrent)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
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

// 批量修改状态
func UpdateState(ids, status string) (int64, error) {
	db := db.Instance().Engine()

	if db == nil {
		return 0, errors.New("获取数据库连接失败")
	}

	rs, err := db.Exec("update sys_job set status=? where job_id in (?)", status, ids)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}
