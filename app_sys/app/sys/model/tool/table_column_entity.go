package tool

import (
	"lostvip.com/db"
	"time"
)

type Entity struct {
	ColumnId      int64     `json:"column_id" xorm:"not null pk autoincr comment('编号') BIGINT(20)"`
	TableId       int64     `json:"table_id" xorm:"comment('归属表编号') BIGINT(20)"`
	ColumnName    string    `json:"column_name" xorm:"comment('列名称') VARCHAR(200)"`
	ColumnComment string    `json:"column_comment" xorm:"comment('列描述') VARCHAR(500)"`
	ColumnType    string    `json:"column_type" xorm:"comment('列类型') VARCHAR(100)"`
	GoType        string    `json:"go_type" xorm:"comment('Go类型') VARCHAR(500)"`
	GoField       string    `json:"go_field" xorm:"comment('Go字段名') VARCHAR(200)"`
	HtmlField     string    `json:"html_field" xorm:"comment('html字段名') VARCHAR(200)"`
	IsPk          string    `json:"is_pk" xorm:"comment('是否主键（1是）') CHAR(1)"`
	IsIncrement   string    `json:"is_increment" xorm:"comment('是否自增（1是）') CHAR(1)"`
	IsRequired    string    `json:"is_required" xorm:"comment('是否必填（1是）') CHAR(1)"`
	IsInsert      string    `json:"is_insert" xorm:"comment('是否为插入字段（1是）') CHAR(1)"`
	IsEdit        string    `json:"is_edit" xorm:"comment('是否编辑字段（1是）') CHAR(1)"`
	IsList        string    `json:"is_list" xorm:"comment('是否列表字段（1是）') CHAR(1)"`
	IsQuery       string    `json:"is_query" xorm:"comment('是否查询字段（1是）') CHAR(1)"`
	QueryType     string    `json:"query_type" xorm:"default 'EQ' comment('查询方式（等于、不等于、大于、小于、范围）') VARCHAR(200)"`
	HtmlType      string    `json:"html_type" xorm:"comment('显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）') VARCHAR(200)"`
	DictType      string    `json:"dict_type" xorm:"default '' comment('字典类型') VARCHAR(200)"`
	Sort          int       `json:"sort" xorm:"comment('排序') INT(11)"`
	CreateBy      string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime    time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy      string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime    time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
}

// 映射数据表
func TableName() string {
	return "gen_table_column"
}

// 插入数据
func (r *Entity) Insert() (int64, error) {
	return db.Instance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *Entity) Update() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.ColumnId).Update(r)
}

// 删除
func (r *Entity) Delete() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.ColumnId).Delete(r)
}

// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(TableName()).In("column_id", ids).Delete(new(Entity))
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
