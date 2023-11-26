// ==========================================================================
// LV自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-11-26 16:27:18 +0800 CST
// 生成人：lv
// ==========================================================================
package service

import (
		"errors"
    	"github.com/gin-gonic/gin"
    	"lostvip.com/utils/convert"
    	"lostvip.com/utils/page"
    	"robvi/app/modules/biz/model"
    	"robvi/app/modules/sys/service"
    	"time"
)
type DpcTaskService struct{}
//根据主键查询数据
func (svc DpcTaskService) SelectRecordById(id int64) (*model.DpcTask, error) {
	entity := &model.DpcTask{ Id: id}
	_, err := entity.FindOne()
	return entity, err
}

//根据主键删除数据
func (svc DpcTaskService) DeleteRecordById(id int64) bool {
	rs, err := (&model.DpcTask{ Id: id}).Delete()
	if err == nil {
		if rs > 0 {
			return true
		}
	}
	return false
}

//批量删除数据记录
func (svc DpcTaskService) DeleteRecordByIds(ids string) int64 {
	ida := convert.ToInt64Array(ids, ",")
    var entity *model.DpcTask
	result, err := entity.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

//添加数据
func (svc DpcTaskService) AddSave(req *model.AddDpcTaskReq, c *gin.Context) (int64, error) {
	var entity *model.DpcTask
	
	   
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
	entity.CreateBy = ""

	var userService service.UserService
	user := userService.GetProfile(c)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	_, err := entity.Insert()
	return entity.Id, err
}

//修改数据
func (svc DpcTaskService) EditSave(req *model.EditDpcTaskReq, c *gin.Context) (int64, error) {
	entity := &model.DpcTask{ Id: req.Id }
	ok, err := entity.FindOne()

	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}

	   
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
	entity.UpdateBy = ""

	var userService service.UserService
	user := userService.GetProfile(c)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return entity.Update()
}

//根据条件查询数据
func (svc DpcTaskService) SelectListAll(params *model.PageDpcTaskReq) ([]model.DpcTask, error) {
	entity := model.DpcTask{}
	return entity.SelectListAll(params)
}

//根据条件分页查询数据
func (svc DpcTaskService) SelectListByPage(params *model.PageDpcTaskReq) ([]model.DpcTask, *page.Paging, error) {
	entity := model.DpcTask{}
	return entity.SelectListByPage(params)
}

// 导出excel
func (svc DpcTaskService) Export(param *model.PageDpcTaskReq) (string, error) {
	head := []string{  "ID" ,"工号" ,"密码" ,"项  目  号" ,"任务内容" ,"开始日期" ,"结束日期" ,"本月工时" ,"自动提交" ,"任务状态，ready running end" ,"排序，大的优先" ,"更新时间" ,"创建时间"}
	col := []string{  "id" ,"username" ,"password" ,"prj_code" ,"task_content" ,"start_date" ,"end_date" ,"work_days" ,"auto_submit" ,"status" ,"sort" ,"update_time" ,"create_time"}
	entity := model.DpcTask{}
	return entity.SelectListExport(param, head, col)
}