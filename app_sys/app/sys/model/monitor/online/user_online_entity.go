package online

import (
	"lostvip.com/db"
	"time"
)

type UserOnline struct {
	Sessionid      string    `json:"sessionId"  xorm:"not null default '' comment('用户会话id') VARCHAR(250)"`
	Token          string    `json:"token"      xorm:"default '' comment('token token') VARCHAR(255)"`
	LoginName      string    `json:"login_name" xorm:"default '' comment('登录账号') VARCHAR(50)"`
	DeptName       string    `json:"dept_name"  xorm:"default '' comment('部门名称') VARCHAR(50)"`
	Ipaddr         string    `json:"ipaddr" xorm:"default '' comment('登录IP地址') VARCHAR(50)"`
	LoginLocation  string    `json:"login_location" xorm:"default '' comment('登录地点') VARCHAR(255)"`
	Browser        string    `json:"browser" xorm:"default '' comment('浏览器类型') VARCHAR(50)"`
	Os             string    `json:"os" xorm:"default '' comment('操作系统') VARCHAR(50)"`
	Status         string    `json:"status" xorm:"default '' comment('在线状态on_line在线off_line离线') VARCHAR(10)"`
	StartTimestamp time.Time `json:"start_timestamp" xorm:"comment('session创建时间') DATETIME"`
	LastAccessTime time.Time `json:"last_access_time" xorm:"comment('session最后访问时间') DATETIME"`
	ExpireTime     int       `json:"expire_time" xorm:"default 0 comment('超时时间，单位为分钟') INT(5)"`
}

// 映射数据表
func TableName() string {
	return "sys_user_online"
}

// 插入数据
func (r *UserOnline) Insert() (int64, error) {
	return db.Instance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *UserOnline) Update() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.Sessionid).Update(r)
}

// 删除
func (r *UserOnline) Delete() (int64, error) {
	rs, err := db.Instance().Engine().Exec("delete from sys_user_online where sessionId = ?", r.Sessionid)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

// 批量删除
func DeleteBatch(ids ...string) (int64, error) {
	return db.Instance().Engine().Table(TableName()).In("sessionId", ids).Delete(new(UserOnline))
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *UserOnline) FindOne() (bool, error) {
	return db.Instance().Engine().Table(TableName()).Get(r)
}

// 根据条件查询
func Find(where, order string) ([]UserOnline, error) {
	var list []UserOnline
	err := db.Instance().Engine().Table(TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

// 指定字段集合查询
func FindIn(column string, args ...interface{}) ([]UserOnline, error) {
	var list []UserOnline
	err := db.Instance().Engine().Table(TableName()).In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func FindNotIn(column string, args ...interface{}) ([]UserOnline, error) {
	var list []UserOnline
	err := db.Instance().Engine().Table(TableName()).NotIn(column, args).Find(&list)
	return list, err
}
