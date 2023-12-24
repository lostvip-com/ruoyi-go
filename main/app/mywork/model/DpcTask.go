// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2023-12-17 19:23:47 +0800 CST
// 生成人：lv
// ==========================================================================

package model

import (
	"lostvip.com/db"
	"robvi/app/common/model_cmn"
	"time"
)

// 数据表映射结构体
type DpcTask struct {
	Id          int64     `gorm:"primary_key;auto_increment;ID;"     json:"id"  form:"id"`
	Username    string    `gorm:"varchar(16);comment:工号;" json:"username" form:"username"`
	Password    string    `gorm:"varchar(32);comment:密码;" json:"password" form:"password"`
	PrjCode     string    `gorm:"varchar(16);comment:项  目  号;" json:"prjCode" form:"prjCode"`
	TaskContent string    `gorm:"varchar(64);comment:任务内容;" json:"taskContent" form:"taskContent"`
	StartDate   time.Time `gorm:"date;comment:开始日期;" json:"startDate" form:"startDate"`
	EndDate     time.Time `gorm:"date;comment:结束日期;" json:"endDate" form:"endDate"`
	WorkDays    int64     `gorm:"int(11);comment:本月工时;" json:"workDays" form:"workDays"`
	AutoSubmit  string    `gorm:"char(1);comment:自动提交;" json:"autoSubmit" form:"autoSubmit"`
	Status      string    `gorm:"char(1);comment:任务状态;" json:"status" form:"status"`
	Sort        int       `gorm:"int(11);comment:排序，大的优先;" json:"sort" form:"sort"`
	DelFlag     string    `gorm:"char(1);comment:删除标识1删除0未删除;" json:"delFlag" form:"delFlag"`
	CreateBy    string    `gorm:"varchar(32);comment:创建人;" json:"createBy" form:"createBy"`
	UpdateBy    string    `gorm:"varchar(32);comment:更新者;" json:"updateBy" form:"updateBy"`
	model_cmn.Model
}

// 映射数据表
func (e *DpcTask) TableName() string {
	return "dpc_task"
}

// 增
func (e *DpcTask) Save() error {
	return db.GetMasterGorm().Save(e).Error
}

// 查
func (e *DpcTask) FindById() error {
	err := db.GetMasterGorm().Find(e, e.Id).Error
	return err
}

// 改
func (e *DpcTask) Updates() error {
	return db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 删
func (e *DpcTask) Delete() error {
	return db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}
