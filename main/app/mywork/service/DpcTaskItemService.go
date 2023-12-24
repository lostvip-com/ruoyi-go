// ==========================================================================
// LV自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-12-24 16:29:05 +0800 CST
// 生成人：lv
// ==========================================================================
package service

import (
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_logic"
	"lostvip.com/utils/lv_office"
	"robvi/app/mywork/dao"
	"robvi/app/mywork/model"
	"robvi/app/mywork/vo"
	"robvi/app/system/model/system/user"
	"time"
)

type DpcTaskItemService struct{}

// FindById 根据主键查询数据
func (svc DpcTaskItemService) FindById(id int64) (*model.DpcTaskItem, error) {
	entity := &model.DpcTaskItem{Id: id}
	err := entity.FindById()
	return entity, err
}

// DeleteById 根据主键删除数据
func (svc DpcTaskItemService) DeleteById(id int64) error {
	err := (&model.DpcTaskItem{Id: id}).Delete()
	return err
}

// DeleteByIds 批量删除数据记录
func (svc DpcTaskItemService) DeleteByIds(ids string) int64 {
	ida := lv_conv.ToInt64Array(ids, ",")
	var d dao.DpcTaskItemDao
	rowsAffected := d.DeleteByIds(ida)
	return rowsAffected
}

// AddSave 添加数据
func (svc DpcTaskItemService) AddSave(req *vo.AddDpcTaskItemReq, user *user.SysUser) (int64, error) {
	var entity = new(model.DpcTaskItem)

	entity.TaskId = req.TaskId
	entity.Action = req.Action
	entity.Content = req.Content
	entity.IdxXpath = req.IdxXpath
	entity.IdxImg1 = req.IdxImg1
	entity.IdxImg2 = req.IdxImg2
	entity.IdxImg3 = req.IdxImg3
	entity.Status = req.Status
	entity.Sort = req.Sort

	entity.CreateTime = time.Now()
	entity.UpdateTime = entity.CreateTime
	entity.CreateBy = user.LoginName

	err := entity.Save()
	lv_logic.HasErrAndPanic(err)
	return entity.Id, err
}

// EditSave 修改数据
func (svc DpcTaskItemService) EditSave(req *vo.EditDpcTaskItemReq, user *user.SysUser) error {
	entity := &model.DpcTaskItem{Id: req.Id}
	err := entity.FindById()
	lv_logic.HasErrAndPanic(err)

	entity.TaskId = req.TaskId
	entity.Action = req.Action
	entity.Content = req.Content
	entity.IdxXpath = req.IdxXpath
	entity.IdxImg1 = req.IdxImg1
	entity.IdxImg2 = req.IdxImg2
	entity.IdxImg3 = req.IdxImg3
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.UpdateTime = time.Now()
	entity.UpdateBy = user.LoginName
	return entity.Updates()
}

// ListByPage 根据条件分页查询数据
func (svc DpcTaskItemService) ListByPage(params *vo.PageDpcTaskItemReq) (*[]model.DpcTaskItem, int64, error) {
	var d dao.DpcTaskItemDao
	return d.ListByPage(params)
}

// ExportAll 导出excel
func (svc DpcTaskItemService) ExportAll(param *vo.PageDpcTaskItemReq) (string, error) {
	var d dao.DpcTaskItemDao
	listMap, err := d.ListAll(param)
	lv_logic.HasErrAndPanic(err)
	heads := []string{"ID", "任务ID", "click dbclick write write_enter", "内容", "html xpath", "图片1", "图片2", "图片3", "操作状态，ready running end", "排序，大的优先", "删除标识1删除0未删除", "创建人", "创建时间", "更新者", "更新时间"}
	cols := []string{"id", "task_id", "action", "content", "idx_xpath", "idx_img1", "idx_img2", "idx_img3", "status", "sort", "del_flag", "create_by", "create_time", "update_by", "update_time"}
	url, err := lv_office.DownlaodExcelByListMap(&heads, &cols, listMap)
	return url, err
}
