// ==========================================================================
// LV自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-01-03 21:50:54 +0800 CST
// 生成人：lv
// ==========================================================================
package service

import (
	"github.com/lv_framework/utils/lv_conv"
	"github.com/lv_framework/utils/lv_err"
	"github.com/lv_framework/utils/lv_office"
	"main/app/mywork/dao"
	"main/app/mywork/model"
	"main/app/mywork/vo"
	"time"
)

type DpcTaskService struct{}

// FindById 根据主键查询数据
func (svc DpcTaskService) FindById(id int64) (*model.DpcTask, error) {
	entity := &model.DpcTask{Id: id}
	err := entity.FindById()
	return entity, err
}

// DeleteById 根据主键删除数据
func (svc DpcTaskService) DeleteById(id int64) error {
	err := (&model.DpcTask{Id: id}).Delete()
	return err
}

// DeleteByIds 批量删除数据记录
func (svc DpcTaskService) DeleteByIds(ids string) int64 {
	ida := lv_conv.ToInt64Array(ids, ",")
	var d dao.DpcTaskDao
	rowsAffected := d.DeleteByIds(ida)
	return rowsAffected
}

// AddSave 添加数据
func (svc DpcTaskService) AddSave(req *vo.AddDpcTaskReq) (int64, error) {
	var entity = new(model.DpcTask)

	entity.Username = req.Username
	entity.Password = req.Password
	entity.PrjCode = req.PrjCode
	entity.TaskContent = req.TaskContent
	entity.StartDate = req.StartDate
	entity.EndDate = req.EndDate
	entity.WorkDays = req.WorkDays
	entity.AutoSubmit = req.AutoSubmit
	entity.Status = req.Status
	entity.Sort = req.Sort

	entity.CreateTime = time.Now()
	entity.UpdateTime = entity.CreateTime
	entity.CreateBy = req.CreateBy
	err := entity.Save()
	lv_err.HasErrAndPanic(err)
	return entity.Id, err
}

// EditSave 修改数据
func (svc DpcTaskService) EditSave(req *vo.EditDpcTaskReq) error {
	entity := &model.DpcTask{Id: req.Id}
	err := entity.FindById()
	lv_err.HasErrAndPanic(err)

	entity.Username = req.Username
	entity.Password = req.Password
	entity.PrjCode = req.PrjCode
	entity.TaskContent = req.TaskContent
	entity.StartDate = req.StartDate
	entity.EndDate = req.EndDate
	entity.WorkDays = req.WorkDays
	entity.AutoSubmit = req.AutoSubmit
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.UpdateTime = time.Now()
	entity.UpdateBy = req.UpdateBy
	return entity.Updates()
}

// ListByPage 根据条件分页查询数据
func (svc DpcTaskService) ListByPage(params *vo.PageDpcTaskReq) (*[]model.DpcTask, int64, error) {
	var d dao.DpcTaskDao
	return d.ListByPage(params)
}

// ExportAll 导出excel
func (svc DpcTaskService) ExportAll(param *vo.PageDpcTaskReq) (string, error) {
	var d dao.DpcTaskDao
	listMap, err := d.ListAll(param, false)
	lv_err.HasErrAndPanic(err)
	heads := []string{"", "工号", "密码", "项  目  号", "任务内容", "开始日期", "结束日期", "本月工时", "自动提交", "任务状态", "排序，大的优先", "更新者", "更新时间", "创建时间", "创建人", "删除标记"}
	cols := []string{"id", "username", "password", "prj_code", "task_content", "start_date", "end_date", "work_days", "auto_submit", "status", "sort", "update_by", "update_time", "create_time", "create_by", "del_flag"}
	url, err := lv_office.DownlaodExcelByListMap(&heads, &cols, listMap)
	return url, err
}
