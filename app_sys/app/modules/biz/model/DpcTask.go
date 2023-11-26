// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2023-11-26 16:27:18 +0800 CST
// 生成人：lv
// ==========================================================================

package model

import (
    "robvi/app/common/model"
    "lostvip.com/db"
	"time"
)

// 数据表映射结构体
type DpcTask struct { 
	 Id       int64         `json:"id" xorm:"not null pk autoincr comment('ID') int(11)"`    
	 Username    string         `json:"username" xorm:"comment('工号') varchar(16)"`    
	 Password    string         `json:"password" xorm:"comment('密码') varchar(32)"`    
	 PrjCode    string         `json:"prj_code" xorm:"comment('项  目  号') varchar(16)"`    
	 TaskContent    string         `json:"task_content" xorm:"comment('任务内容') varchar(64)"`    
	 StartDate    time.Time         `json:"start_date" xorm:"comment('开始日期') date"`    
	 EndDate    time.Time         `json:"end_date" xorm:"comment('结束日期') date"`    
	 WorkDays    int64         `json:"work_days" xorm:"comment('本月工时') int(11)"`    
	 AutoSubmit    string         `json:"auto_submit" xorm:"comment('自动提交') char(1)"`    
	 Status    string         `json:"status" xorm:"comment('任务状态，ready running end') varchar(16)"`    
	 Sort    int64         `json:"sort" xorm:"comment('排序，大的优先') int(11)"`    
	 UpdateTime    time.Time         `json:"update_time" xorm:"comment('更新时间') datetime"`    
	 CreateTime    time.Time         `json:"create_time" xorm:"comment('创建时间') datetime"`    
    model.BaseModel
}

//映射数据表
func (e *DpcTask) TableName() string {
	return "gen_table"
}

// 插入数据
func (e *DpcTask) Insert() (int64, error) {
	return db.Instance().Engine().Table(e.TableName()).Insert(e)
}

// 更新数据
func (e *DpcTask) Update() (int64, error) {
	return db.Instance().Engine().Table(e.TableName()).ID(e.Id).Update(e)
}

// 删除
func (e *DpcTask) Delete() (int64, error) {
	return db.Instance().Engine().Table(e.TableName()).ID(e.Id).Delete(e)
}

//批量删除
func (e *DpcTask) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(e.TableName()).In("id", ids).Delete(new(DpcTask))
}

// 根据结构体中已有的非空数据来获得单条数据
func (e *DpcTask) FindOne() (bool, error) {
	return db.Instance().Engine().Table(e.TableName()).Get(e)
}

// 根据条件查询
func (e *DpcTask) Find(where, order string) ([]DpcTask, error) {
	var list []DpcTask
	err := db.Instance().Engine().Table(e.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (e *DpcTask) FindIn(column string, args ...interface{}) ([]DpcTask, error) {
	var list []DpcTask
	err := db.Instance().Engine().Table(e.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (e *DpcTask)  FindNotIn(column string, args ...interface{}) ([]DpcTask, error) {
	var list []DpcTask
	err := db.Instance().Engine().Table(e.TableName()).NotIn(column, args).Find(&list)
	return list, err
}