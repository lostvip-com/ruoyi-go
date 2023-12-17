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
	Id                  int64     `json:"id" xorm:"not null pk autoincr comment('ID') int(11)"`
	Username            string    `json:"username" xorm:"comment('工号') varchar(16)"`
	Password            string    `json:"password" xorm:"comment('密码') varchar(32)"`
	PrjCode             string    `json:"prj_code" xorm:"comment('项  目  号') varchar(16)"`
	TaskContent         string    `json:"task_content" xorm:"comment('任务内容') varchar(64)"`
	StartDate           time.Time `json:"start_date" xorm:"comment('开始日期') date"`
	EndDate             time.Time `json:"end_date" xorm:"comment('结束日期') date"`
	WorkDays            int64     `json:"work_days" xorm:"comment('本月工时') int(11)"`
	AutoSubmit          string    `json:"auto_submit" xorm:"comment('自动提交') char(1)"`
	Status              int64     `json:"status" xorm:"comment('任务状态') tinyint(16)"`
	Sort                int64     `json:"sort" xorm:"comment('排序，大的优先') int(11)"`
	model_cmn.BaseModel `xorm:"extends"`
}

// 映射数据表
func (e *DpcTask) TableName() string {
	return "dpc_task"
}

// 增
func (e *DpcTask) Insert() (int64, error) {
	return db.Instance().Engine().Table("dpc_task").Insert(e)
}

// 查
func (e *DpcTask) FindOne() (bool, error) {
	return db.Instance().Engine().Table("dpc_task").Get(e)
}

// 改
func (e *DpcTask) Update() (int64, error) {
	return db.Instance().Engine().Table("dpc_task").ID(e.Id).Update(e)
}

// 删
func (e *DpcTask) Delete() (int64, error) {
	return db.Instance().Engine().Table("dpc_task").ID(e.Id).Delete(e)
}
