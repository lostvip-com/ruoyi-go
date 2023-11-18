package logininfor

import (
	"lostvip.com/db"
	"time"
)

type Entity struct {
	InfoId        int64     `json:"info_id" xorm:"not null pk autoincr comment('访问ID') BIGINT(20)"`
	LoginName     string    `json:"login_name" xorm:"default '' comment('登录账号') VARCHAR(50)"`
	Ipaddr        string    `json:"ipaddr" xorm:"default '' comment('登录IP地址') VARCHAR(50)"`
	LoginLocation string    `json:"login_location" xorm:"default '' comment('登录地点') VARCHAR(255)"`
	Browser       string    `json:"browser" xorm:"default '' comment('浏览器类型') VARCHAR(50)"`
	Os            string    `json:"os" xorm:"default '' comment('操作系统') VARCHAR(50)"`
	Status        string    `json:"status" xorm:"default '0' comment('登录状态（0成功 1失败）') CHAR(1)"`
	Msg           string    `json:"msg" xorm:"default '' comment('提示消息') VARCHAR(255)"`
	LoginTime     time.Time `json:"login_time" xorm:"comment('访问时间') DATETIME"`
}

// 映射数据表
func TableName() string {
	return "sys_logininfor"
}

// 插入数据
func (r *Entity) Insert() (int64, error) {
	return db.Instance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *Entity) Update() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.InfoId).Update(r)
}

// 删除
func (r *Entity) Delete() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.InfoId).Delete(r)
}

// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(TableName()).In("info_id", ids).Delete(new(Entity))
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
