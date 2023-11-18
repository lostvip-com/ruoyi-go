package job_log

import (
	"lostvip.com/db"
	"time"
)

type Entity struct {
	JobLogId      int64     `json:"job_log_id" xorm:"not null pk autoincr comment('任务日志ID') BIGINT(20)"`
	JobName       string    `json:"job_name" xorm:"not null comment('任务名称') VARCHAR(64)"`
	JobGroup      string    `json:"job_group" xorm:"not null comment('任务组名') VARCHAR(64)"`
	InvokeTarget  string    `json:"invoke_target" xorm:"not null comment('调用目标字符串') VARCHAR(500)"`
	JobMessage    string    `json:"job_message" xorm:"comment('日志信息') VARCHAR(500)"`
	Status        string    `json:"status" xorm:"default '0' comment('执行状态（0正常 1失败）') CHAR(1)"`
	ExceptionInfo string    `json:"exception_info" xorm:"default '' comment('异常信息') VARCHAR(2000)"`
	CreateTime    time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
}

// 映射数据表
func TableName() string {
	return "sys_job_log"
}

// 插入数据
func (r *Entity) Insert() (int64, error) {
	return db.Instance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *Entity) Update() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.JobLogId).Update(r)
}

// 删除
func (r *Entity) Delete() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.JobLogId).Delete(r)
}

// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(TableName()).In("job_log_id", ids).Delete(new(Entity))
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Entity) FindOne() (bool, error) {
	return db.Instance().Engine().Table(TableName()).Get(r)
}

// 根据条件查询
func Find(where, order string) ([]Entity, error) {
	var list []Entity
	err := db.Instance().Engine().Table(TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

// 指定字段集合查询
func FindIn(column string, args ...interface{}) ([]Entity, error) {
	var list []Entity
	err := db.Instance().Engine().Table(TableName()).In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func FindNotIn(column string, args ...interface{}) ([]Entity, error) {
	var list []Entity
	err := db.Instance().Engine().Table(TableName()).NotIn(column, args).Find(&list)
	return list, err
}
