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
type DpcTaskItem struct {
	Id              int64  `json:"id" xorm:"not null pk autoincr comment('ID') bigint(20)"`
	TaskId          int64  `json:"task_id" xorm:"not null comment('任务ID') bigint(20)"`
	Action          string `json:"action" xorm:"comment('click dbclick write write_enter') varchar(16)"`
	Content         string `json:"content" xorm:"comment('内容') varchar(16)"`
	Img1            string `json:"img1" xorm:"comment('图片1') varchar(128)"`
	Img2            string `json:"img2" xorm:"comment('图片2') varchar(128)"`
	Img3            string `json:"img3" xorm:"comment('图片3') varchar(128)"`
	Status          string `json:"status" xorm:"comment('操作状态，ready running end') varchar(16)"`
	Sort            int64  `json:"sort" xorm:"comment('排序，大的优先') int(11)"`
	model.BaseModel `xorm:"extends"`
}

// 映射数据表
func (e *DpcTaskItem) TableName() string {
	return "dpc_task_item"
}

// 插入数据
func (e *DpcTaskItem) Insert() (int64, error) {
	return db.Instance().Engine().Table(e.TableName()).Insert(e)
}

// 更新数据
func (e *DpcTaskItem) Update() (int64, error) {
	return db.Instance().Engine().Table(e.TableName()).ID(e.Id).Update(e)
}

// 删除
func (e *DpcTaskItem) Delete() (int64, error) {
	return db.Instance().Engine().Table(e.TableName()).ID(e.Id).Delete(e)
}

// 批量删除
func (e *DpcTaskItem) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(e.TableName()).In("id", ids).Delete(new(DpcTaskItem))
}

// 根据结构体中已有的非空数据来获得单条数据
func (e *DpcTaskItem) FindOne() (bool, error) {
	return db.Instance().Engine().Table(e.TableName()).Get(e)
}

// 根据条件查询
func (e *DpcTaskItem) Find(where, order string) ([]DpcTaskItem, error) {
	var list []DpcTaskItem
	err := db.Instance().Engine().Table(e.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

// 指定字段集合查询
func (e *DpcTaskItem) FindIn(column string, args ...interface{}) ([]DpcTaskItem, error) {
	var list []DpcTaskItem
	err := db.Instance().Engine().Table(e.TableName()).In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func (e *DpcTaskItem) FindNotIn(column string, args ...interface{}) ([]DpcTaskItem, error) {
	var list []DpcTaskItem
	err := db.Instance().Engine().Table(e.TableName()).NotIn(column, args).Find(&list)
	return list, err
}
