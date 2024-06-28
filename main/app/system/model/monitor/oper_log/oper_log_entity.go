package oper_log

import (
	"github.com/lv_framework/db"
	"time"
)

type Entity struct {
	OperId        int64     `json:"oper_id" xorm:"not null pk autoincr comment('日志主键') BIGINT(20)"`
	Title         string    `json:"title" xorm:"default '' comment('模块标题') VARCHAR(50)"`
	BusinessType  int       `json:"business_type" xorm:"default 0 comment('业务类型（0其它 1新增 2修改 3删除）') INT(2)"`
	Method        string    `json:"method" xorm:"default '' comment('方法名称') VARCHAR(100)"`
	RequestMethod string    `json:"request_method" xorm:"default '' comment('请求方式') VARCHAR(10)"`
	OperatorType  int       `json:"operator_type" xorm:"default 0 comment('操作类别（0其它 1后台用户 2手机端用户）') INT(1)"`
	OperName      string    `json:"oper_name" xorm:"default '' comment('操作人员') VARCHAR(50)"`
	DeptName      string    `json:"dept_name" xorm:"default '' comment('部门名称') VARCHAR(50)"`
	OperUrl       string    `json:"oper_url" xorm:"default '' comment('请求URL') VARCHAR(255)"`
	OperIp        string    `json:"oper_ip" xorm:"default '' comment('主机地址') VARCHAR(50)"`
	OperLocation  string    `json:"oper_location" xorm:"default '' comment('操作地点') VARCHAR(255)"`
	OperParam     string    `json:"oper_param" xorm:"default '' comment('请求参数') VARCHAR(2000)"`
	JsonResult    string    `json:"json_result" xorm:"default '' comment('返回参数') VARCHAR(2000)"`
	Status        int       `json:"status" xorm:"default 0 comment('操作状态（0正常 1异常）') INT(1)"`
	ErrorMsg      string    `json:"error_msg" xorm:"default '' comment('错误消息') VARCHAR(2000)"`
	OperTime      time.Time `json:"oper_time" xorm:"comment('操作时间') DATETIME"`
}

// 映射数据表
func TableName() string {
	return "sys_oper_log"
}

// 插入数据
func (r *Entity) Insert() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *Entity) Update() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).ID(r.OperId).Update(r)
}

// 删除
func (r *Entity) Delete() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).ID(r.OperId).Delete(r)
}

// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).In("oper_id", ids).Delete(new(Entity))
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Entity) FindOne() (bool, error) {
	return db.GetInstance().Engine().Table(TableName()).Get(r)
}

// 根据条件查询
func Find(where, order string) ([]Entity, error) {
	var list []Entity
	err := db.GetInstance().Engine().Table(TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

// 指定字段集合查询
func FindIn(column string, args ...interface{}) ([]Entity, error) {
	var list []Entity
	err := db.GetInstance().Engine().Table(TableName()).In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func FindNotIn(column string, args ...interface{}) ([]Entity, error) {
	var list []Entity
	err := db.GetInstance().Engine().Table(TableName()).NotIn(column, args).Find(&list)
	return list, err
}
