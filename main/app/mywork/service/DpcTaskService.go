// ==========================================================================
// LV自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-12-17 19:23:47 +0800 CST
// 生成人：lv
// ==========================================================================
package service

import (
	"errors"
    "github.com/gin-gonic/gin"
    "lostvip.com/utils/lv_conv"
    "lostvip.com/utils/lv_web"
    "robvi/app/mywork/model"
    "robvi/app/mywork/dao"
    "robvi/app/mywork/dto"
    "robvi/app/system/service"
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
	ida := lv_conv.ToInt64Array(ids, ",")
    var dao dao.DpcTaskDao
	result, err := dao.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

//添加数据
func (svc DpcTaskService) AddSave(req *dto.AddDpcTaskReq, c *gin.Context) (int64, error) {
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
	  entity.DelFlag = req.DelFlag  
	 
	 
	 
	 
	entity.CreateTime = time.Now()
	entity.CreateBy = ""

	var userService service.UserService
	user := userService.GetProfile(c)

	if user != nil {
	    entity.CreateBy = user.LoginName
	}

	_, err := entity.Insert()
	if err != nil {
        panic(err)
    }
	return entity.Id, err
}

//修改数据
func (svc DpcTaskService) EditSave(req *dto.EditDpcTaskReq, c *gin.Context) (int64, error) {
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
func (svc DpcTaskService) SelectListAll(params *dto.PageDpcTaskReq) ([]model.DpcTask, error) {
	var dao dao.DpcTaskDao
	return dao.SelectListAll(params)
}

//根据条件分页查询数据
func (svc DpcTaskService) SelectListByPage(params *dto.PageDpcTaskReq) ([]model.DpcTask, *lv_web.Paging, error) {
	var dao dao.DpcTaskDao
	return dao.SelectListByPage(params)
}

// 导出excel
func (svc DpcTaskService) Export(param *dto.PageDpcTaskReq) (string, error) {
	head := []string{  "ID" ,"工号" ,"密码" ,"项  目  号" ,"任务内容" ,"开始日期" ,"结束日期" ,"本月工时" ,"自动提交" ,"任务状态" ,"排序，大的优先" ,"删除标识1删除0未删除" ,"创建人" ,"创建时间" ,"更新者" ,"更新时间"}
	col := []string{  "id" ,"username" ,"password" ,"prj_code" ,"task_content" ,"start_date" ,"end_date" ,"work_days" ,"auto_submit" ,"status" ,"sort" ,"del_flag" ,"create_by" ,"create_time" ,"update_by" ,"update_time"}
	var dao dao.DpcTaskDao
	return dao.SelectListExport(param, head, col)
}