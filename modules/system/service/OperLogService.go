package service

import (
	"common/util"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"system/model"
	"system/vo"
	"time"
)

type OperLogService struct {
}

// Add  新增记录
func (svc OperLogService) Add(c *gin.Context, title, inContent string, outContent *lv_dto.CommonRes) error {
	var userService UserService
	user := userService.GetProfile(c)
	if user == nil {
		return errors.New("用户未登录")
	}

	var operLog model.SysOperLog

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
	var deptServic DeptService
	dept := deptServic.SelectDeptById(user.DeptId)

	if dept != nil {
		operLog.DeptName = dept.DeptName
	} else {
		operLog.DeptName = ""
	}
	operLog.OperUrl = c.Request.URL.Path
	operLog.Method = c.Request.Method
	operLog.OperIp = c.ClientIP()
	operLog.OperLocation = util.GetCityByIp(operLog.OperIp)
	operLog.OperTime = time.Now()
	err := operLog.Save()
	return err
}

// 根据条件分页查询用户列表
func (svc OperLogService) SelectPageList(param *vo.OperLogPageReq) (*[]model.SysOperLog, int64, error) {
	db := lv_db.GetMasterGorm()
	tb := db.Table("sys_oper_log")
	if param != nil {
		if param.Title != "" {
			tb.Where("title like ?", "%"+param.Title+"%")
		}

		if param.OperName != "" {
			tb.Where("oper_name like ?", "%"+param.OperName+"%")
		}

		if param.Status != "" {
			tb.Where("status = ?", param.Status)
		}

		if param.BusinessTypes >= 0 {
			tb.Where("business_type = ?", param.BusinessTypes)
		}

		if param.BeginTime != "" {
			tb.Where("date_format(oper_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			tb.Where("date_format(oper_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}
	}
	var result []model.SysOperLog
	var total int64
	err := tb.Count(&total).Offset(param.GetStartNum()).Limit(param.GetPageSize()).Order("oper_id desc").Find(&result).Error
	return &result, total, err
}

// 根据主键查询用户信息
func (svc OperLogService) SelectRecordById(id int64) (*model.SysOperLog, error) {
	entity := &model.SysOperLog{OperId: id}
	_, err := entity.FindOne()
	return entity, err
}

// DeleteRecordById 根据主键删除用户信息
func (svc OperLogService) DeleteRecordById(id int64) bool {
	entity := &model.SysOperLog{OperId: id}
	err := entity.Delete()
	if err == nil {
		return true
	}
	return false
}

// DeleteRecordByIds 批量删除记录
func (svc OperLogService) DeleteRecordByIds(ids string) error {
	idarr := lv_conv.ToInt64Array(ids, ",")
	err := lv_db.GetMasterGorm().Delete(&model.SysOperLog{}, "oper_id in (?)", idarr).Error
	return err
}

// 清空记录
func (svc OperLogService) DeleteRecordAll() error {
	err := lv_db.GetMasterGorm().Exec("delete from sys_oper_log").Error
	return err
}
