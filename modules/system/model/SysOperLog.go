// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-08-16 18:27:37 +0800 CST
// 生成人：lv
// ==========================================================================
package model

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/namedsql"
	"time"
)

// SysOperLog 操作日志记录
type SysOperLog struct {
	OperId        int64     `gorm:"type:bigint(20);primary_key;auto_increment;日志主键;" json:"operId"`
	Title         string    `gorm:"type:varchar(50);comment:模块标题;" json:"title"`
	BusinessType  int       `gorm:"type:int(11);comment:业务类型（0其它 1新增 2修改 3删除）;" json:"businessType"`
	Method        string    `gorm:"type:varchar(100);comment:方法名称;" json:"method"`
	RequestMethod string    `gorm:"type:varchar(10);comment:请求方式;" json:"requestMethod"`
	OperatorType  int       `gorm:"type:int(11);comment:操作类别（0其它 1后台用户 2手机端用户）;" json:"operatorType"`
	OperName      string    `gorm:"type:varchar(50);comment:操作人员;" json:"operName"`
	DeptName      string    `gorm:"type:varchar(50);comment:部门名称;" json:"deptName"`
	OperUrl       string    `gorm:"type:varchar(255);comment:请求URL;" json:"operUrl"`
	OperIp        string    `gorm:"type:varchar(50);comment:主机地址;" json:"operIp"`
	OperLocation  string    `gorm:"type:varchar(255);comment:操作地点;" json:"operLocation"`
	OperParam     string    `gorm:"type:varchar(2000);comment:请求参数;" json:"operParam"`
	JsonResult    string    `gorm:"type:varchar(2000);comment:返回参数;" json:"jsonResult"`
	Status        int       `gorm:"type:int(11);comment:操作状态（0正常 1异常）;" json:"status"`
	ErrorMsg      string    `gorm:"type:varchar(2000);comment:错误消息;" json:"errorMsg"`
	OperTime      time.Time `gorm:"type:datetime;comment:操作时间;" json:"operTime" time_format:"2006-01-02 15:04:05"`
	CreateTime    time.Time `gorm:"type:datetime;comment:创建时间;column:create_time;" json:"createTime" time_format:"2006-01-02 15:04:05"`
	CreateBy      string    `gorm:"type:varchar(32);comment:创建人;column:create_by;"  json:"createBy"`
	DelFlag       int       `gorm:"type:tinyint(1);default:0;comment:删除标记;column:del_flag;" json:"delFlag"`
}

func (e *SysOperLog) TableName() string {
	return "sys_oper_log"
}

// 增
func (e *SysOperLog) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysOperLog) FindById(id int64) (*SysOperLog, error) {
	err := lv_db.GetMasterGorm().Take(e, id).Error
	return e, err
}

// 查第一条
func (e *SysOperLog) FindOne() (*SysOperLog, error) {
	tb := lv_db.GetMasterGorm().Table(e.TableName())

	if e.Title != "" {
		tb = tb.Where("title=?", e.Title)
	}
	if e.BusinessType != 0 {
		tb = tb.Where("business_type=?", e.BusinessType)
	}
	if e.Method != "" {
		tb = tb.Where("method=?", e.Method)
	}
	if e.RequestMethod != "" {
		tb = tb.Where("request_method=?", e.RequestMethod)
	}
	err := tb.First(e).Error
	return e, err
}

// 改
func (e *SysOperLog) Updates() error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 删
func (e *SysOperLog) Delete() error {
	return lv_db.GetMasterGorm().Delete(e).Error
}

func (e *SysOperLog) Count() (int64, error) {
	sql := " select count(*) from sys_oper_log where del_flag = 0 "

	if e.Title != "" {
		sql += " and title = @Title "
	}
	if e.BusinessType != 0 {
		sql += " and business_type = @BusinessType "
	}
	if e.Method != "" {
		sql += " and method = @Method "
	}
	if e.RequestMethod != "" {
		sql += " and request_method = @RequestMethod "
	}

	return namedsql.Count(lv_db.GetMasterGorm(), sql, e)
}
