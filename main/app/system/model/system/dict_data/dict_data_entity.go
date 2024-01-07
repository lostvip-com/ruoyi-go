package dict_data

import (
	"lostvip.com/db"
	"time"
)

type Entity struct {
	DictCode   int64     `json:"dictCode" xorm:"not null pk autoincr comment('字典编码') BIGINT(20)"`
	DictSort   int       `json:"dictSort" xorm:"default 0 comment('字典排序') INT(4)"`
	DictLabel  string    `json:"dictLabel" xorm:"default '' comment('字典标签') VARCHAR(100)"`
	DictValue  string    `json:"dictValue" xorm:"default '' comment('字典键值') VARCHAR(100)"`
	DictType   string    `json:"dictType" xorm:"default '' comment('字典类型') VARCHAR(100)"`
	CssClass   string    `json:"cssClass" xorm:"comment('样式属性（其他样式扩展）') VARCHAR(100)"`
	ListClass  string    `json:"listClass" xorm:"comment('表格回显样式') VARCHAR(100)"`
	IsDefault  string    `json:"isDefault" xorm:"default 'N' comment('是否默认（Y是 N否）') CHAR(1)"`
	Status     string    `json:"status" xorm:"default '0' comment('状态（0正常 1停用）') CHAR(1)"`
	CreateBy   string    `json:"createBy" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"createTime" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"updateBy" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"updateTime" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
}

// 映射数据表
func TableName() string {
	return "sys_dict_data"
}

// 插入数据
func (r *Entity) Insert() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *Entity) Update() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).ID(r.DictCode).Update(r)
}

// 删除
func (r *Entity) Delete() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).ID(r.DictCode).Delete(r)
}

// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).In("dictCode", ids).Delete(new(Entity))
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
