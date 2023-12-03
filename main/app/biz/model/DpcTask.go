// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2023-11-26 16:27:18 +0800 CST
// 生成人：lv
// ==========================================================================

package model

import (
	"lostvip.com/db"
	"robvi/app/common/model"
)

// 数据表映射结构体
type DpcTask struct {
	Id              int64  `json:"id"          xorm:"not null pk autoincr comment('ID') int(11)"`
	Title           string `json:"title"       xorm:"comment('任务名称') varchar(16)"`
	AutoSubmit      string `json:"auto_submit" xorm:"comment('自动提交') char(1)"`
	Status          string `json:"status"      xorm:"comment('任务状态，ready running end') varchar(16)"`
	Sort            int64  `json:"sort"        xorm:"comment('排序，大的优先') int(11)"`
	RunTime         string `json:"run_time"    xorm:"comment('执行时间') varchar(32)"`
	Repeat          string `json:"repeat"      xorm:"comment('重复方式') varchar(32)"`
	model.BaseModel `xorm:"extends"`
	Comment         string `json:"comment"     xorm:"comment('任务表')"` // 添加表注释字段
}

// 映射数据表
func (e *DpcTask) TableName() string {
	return "dpc_task"
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

// 批量删除
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

// 指定字段集合查询
func (e *DpcTask) FindIn(column string, args ...interface{}) ([]DpcTask, error) {
	var list []DpcTask
	err := db.Instance().Engine().Table(e.TableName()).In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func (e *DpcTask) FindNotIn(column string, args ...interface{}) ([]DpcTask, error) {
	var list []DpcTask
	err := db.Instance().Engine().Table(e.TableName()).NotIn(column, args).Find(&list)
	return list, err
}
