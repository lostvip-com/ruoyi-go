package dict_type

import (
	"lostvip.com/db"
	"time"
)

type Entity struct {
	DictId     int64     `json:"dict_id" xorm:"not null pk autoincr comment('字典主键') BIGINT(20)"`
	DictName   string    `json:"dict_name" xorm:"default '' comment('字典名称') VARCHAR(100)"`
	DictType   string    `json:"dict_type" xorm:"default '' comment('字典类型') unique VARCHAR(100)"`
	Status     string    `json:"status" xorm:"default '0' comment('状态（0正常 1停用）') CHAR(1)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
}

// 映射数据表
func TableName() string {
	return "sys_dict_type"
}

// 插入数据
func (r *Entity) Insert() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *Entity) Update() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).ID(r.DictId).Update(r)
}

// 删除
func (r *Entity) Delete() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).ID(r.DictId).Delete(r)
}

// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).In("dict_id", ids).Delete(new(Entity))
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
