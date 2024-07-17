// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-01-03 21:50:54 +0800 CST
// 生成人：lv
// ==========================================================================

package model

import (
	"github.com/lostvip-com/lv_framework/db"
	"time"
)

// DpcTask DpcTask
type DpcTask struct {
	Id          int64     `gorm:"type:bigint(20);primary_key;auto_increment;;"     json:"id"  form:"id"`
	Username    string    `gorm:"type:varchar(32);comment:工号;" json:"username" form:"username"`
	Password    string    `gorm:"type:varchar(32);comment:密码;" json:"password" form:"password"`
	PrjCode     string    `gorm:"type:varchar(32);comment:项  目  号;" json:"prjCode" form:"prjCode"`
	TaskContent string    `gorm:"type:varchar(128);comment:任务内容;" json:"taskContent" form:"taskContent"`
	StartDate   time.Time `gorm:"type:date;comment:开始日期;" json:"startDate" form:"startDate" time_format:"2006-01-02"`
	EndDate     time.Time `gorm:"type:date;comment:结束日期;" json:"endDate" form:"endDate" time_format:"2006-01-02"`
	WorkDays    int64     `gorm:"type:bigint(20);comment:本月工时;" json:"workDays" form:"workDays"`
	AutoSubmit  string    `gorm:"type:char(1);comment:自动提交;" json:"autoSubmit" form:"autoSubmit"`
	Status      string    `gorm:"type:varchar(8);comment:任务状态;" json:"status" form:"status"`
	Sort        int64     `gorm:"type:bigint(20);comment:排序，大的优先;" json:"sort" form:"sort"`
	UpdateBy    string    `gorm:"type:varchar(32);comment:更新者;" json:"updateBy" form:"updateBy"`
	UpdateTime  time.Time `gorm:"type:datetime;comment:更新时间;" json:"updateTime" form:"updateTime" time_format:"2006-01-02 15:04:05"`
	CreateTime  time.Time `gorm:"type:datetime;comment:创建时间;"  column:create_time; time_format:"2006-01-02 15:04:05" json:"createTime"  `
	CreateBy    string    `gorm:"type:varchar(32);comment:创建人;" column:create_by; json:"createBy" form:"createBy"`
	DelFlag     int       `gorm:"type:tinyint(1);default:0;comment:删除标记;" column:del_flag; json:"delFlag"`
}

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
