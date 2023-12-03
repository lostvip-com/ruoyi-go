package job

import (
	"lostvip.com/db"
	"time"
)

type Entity struct {
	JobId          int64     `json:"job_id" xorm:"not null pk autoincr comment('任务ID') BIGINT(20)"`
	JobName        string    `json:"job_name" xorm:"not null default '' comment('任务名称') VARCHAR(64)"`
	JobParams      string    `json:"job_params" xorm:"default '' comment('参数') VARCHAR(200)"`
	JobGroup       string    `json:"job_group" xorm:"not null default 'DEFAULT' comment('任务组名') VARCHAR(64)"`
	InvokeTarget   string    `json:"invoke_target" xorm:"not null comment('调用目标字符串') VARCHAR(500)"`
	CronExpression string    `json:"cron_expression" xorm:"default '' comment('cron执行表达式') VARCHAR(255)"`
	MisfirePolicy  string    `json:"misfire_policy" xorm:"default '1' comment('计划执行策略（1多次执行 2执行一次）') VARCHAR(20)"`
	Concurrent     string    `json:"concurrent" xorm:"default '1' comment('是否并发执行（0允许 1禁止）') CHAR(1)"`
	Status         string    `json:"status" xorm:"default '0' comment('状态（0正常 1暂停）') CHAR(1)"`
	CreateBy       string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy       string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime     time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark         string    `json:"remark" xorm:"default '' comment('备注信息') VARCHAR(500)"`
}

// 映射数据表
func TableName() string {
	return "sys_job"
}

// 插入数据
func (r *Entity) Insert() (int64, error) {
	return db.Instance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *Entity) Update() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.JobId).Update(r)
}

// 删除
func (r *Entity) Delete() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.JobId).Delete(r)
}

// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(TableName()).In("job_id", ids).Delete(new(Entity))
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
