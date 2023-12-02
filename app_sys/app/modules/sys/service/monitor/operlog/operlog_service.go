package operlog

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_net"
	"lostvip.com/utils/lv_web"
	"robvi/app/modules/sys/model"
	"robvi/app/modules/sys/model/monitor/oper_log"
	"robvi/app/modules/sys/service"
	"time"
)

// 新增记录
func Add(c *gin.Context, title, inContent string, outContent *model.CommonRes) error {
	var userService service.UserService
	user := userService.GetProfile(c)
	if user == nil {
		return errors.New("用户未登陆")
	}

	var operLog oper_log.Entity

	outJson, _ := json.Marshal(outContent)
	outJsonStr := string(outJson)

	operLog.Title = title
	operLog.OperParam = inContent
	operLog.JsonResult = outJsonStr
	operLog.BusinessType = int(outContent.Btype)
	//操作类别（0其它 1后台用户 2手机端用户）
	operLog.OperatorType = 1
	//操作状态（0正常 1异常）
	if outContent.Code == 0 {
		operLog.Status = 0
	} else {
		operLog.Status = 1
	}

	operLog.OperName = user.LoginName
	operLog.RequestMethod = c.Request.Method

	//获取用户部门
	var deptServic service.DeptService
	dept := deptServic.SelectDeptById(user.DeptId)

	if dept != nil {
		operLog.DeptName = dept.DeptName
	} else {
		operLog.DeptName = ""
	}

	operLog.OperUrl = c.Request.URL.Path
	operLog.Method = c.Request.Method
	operLog.OperIp = c.ClientIP()

	operLog.OperLocation = lv_net.GetCityByIp(operLog.OperIp)
	operLog.OperTime = time.Now()

	_, err := operLog.Insert()
	return err
}

// 根据条件分页查询用户列表
func SelectPageList(param *oper_log.SelectPageReq) (*[]oper_log.Entity, *lv_web.Paging, error) {
	return oper_log.SelectPageList(param)
}

// 根据主键查询用户信息
func SelectRecordById(id int64) (*oper_log.Entity, error) {
	entity := &oper_log.Entity{OperId: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除用户信息
func DeleteRecordById(id int64) bool {
	entity := &oper_log.Entity{OperId: id}
	result, err := entity.Delete()
	if err == nil && result > 0 {
		return true
	}

	return false
}

// 批量删除记录
func DeleteRecordByIds(ids string) int64 {
	idarr := lv_conv.ToInt64Array(ids, ",")
	result, _ := oper_log.DeleteBatch(idarr...)
	return result
}

// 清空记录
func DeleteRecordAll() (int64, error) {
	return oper_log.DeleteAll()
}

// 导出excel
func Export(param *oper_log.SelectPageReq) (string, error) {
	head := []string{"日志主键", "模块标题", "业务类型", "方法名称", "请求方式", "操作类别", "操作人员", "部门名称", "请求URL", "主机地址", "操作地点", "请求参数", "返回参数", "操作状态", "操作时间"}
	col := []string{"oper_id", "title", "business_type", "method", "request_method", "operator_type", "oper_name", "dept_name", "oper_url", "oper_ip", "oper_location", "oper_param", "json_result", "status", "error_msg", "oper_time"}
	return oper_log.SelectExportList(param, head, col)
}
