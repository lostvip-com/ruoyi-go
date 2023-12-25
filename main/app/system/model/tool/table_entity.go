package tool

import (
	"lostvip.com/db"
	"time"
)

type GenTable struct {
	TableId        int64     `json:"table_id" xorm:"not null pk autoincr comment('编号') BIGINT(20)"`
	TbName         string    `json:"table_name"  xorm:"default '' comment('表名称') VARCHAR(200) 'table_name' "`
	TableComment   string    `json:"table_comment" xorm:"default '' comment('表描述') VARCHAR(500)"`
	ClassName      string    `json:"class_name" xorm:"default '' comment('实体类名称') VARCHAR(100)"`
	TplCategory    string    `json:"tpl_category" xorm:"default 'crud' comment('使用的模板（crud单表操作 tree树表操作）') VARCHAR(200)"`
	PackageName    string    `json:"package_name" xorm:"comment('生成包路径') VARCHAR(100)"`
	ModuleName     string    `json:"module_name" xorm:"comment('生成模块名') VARCHAR(30)"`
	BusinessName   string    `json:"business_name" xorm:"comment('生成业务名') VARCHAR(30)"`
	FunctionName   string    `json:"function_name" xorm:"comment('生成功能名') VARCHAR(50)"`
	FunctionAuthor string    `json:"function_author" xorm:"comment('生成功能作者') VARCHAR(50)"`
	Options        string    `json:"options" xorm:"comment('其它生成选项') VARCHAR(1000)"`
	CreateBy       string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy       string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime     time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark         string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
	HasEditTime    string    `xorm:"-"` //1需要导入time.Time 0 不需要
}

// 映射数据表
func (r *GenTable) TableName() string {
	return "gen_table"
}

// 插入数据
func (r *GenTable) Insert() (int64, error) {
	return db.GetInstance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *GenTable) Update() (int64, error) {
	return db.GetInstance().Engine().Table(r.TableName()).ID(r.TableId).Update(r)
}

// 删除
func (r *GenTable) Delete() (int64, error) {
	return db.GetInstance().Engine().Table(r.TableName()).ID(r.TableId).Delete(r)
}

// 批量删除
func (r *GenTable) DeleteBatch(ids ...int64) (int64, error) {
	return db.GetInstance().Engine().Table(r.TableName()).In("table_id", ids).Delete(new(GenTable))
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *GenTable) FindOne() (bool, error) {
	return db.GetInstance().Engine().Table(r.TableName()).Get(r)
}

// 根据条件查询
func (r *GenTable) Find(where, order string) ([]GenTable, error) {
	var list []GenTable
	err := db.GetInstance().Engine().Table(r.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

// 指定字段集合查询
func (r *GenTable) FindIn(column string, args ...interface{}) ([]GenTable, error) {
	var list []GenTable
	err := db.GetInstance().Engine().Table(r.TableName()).In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func (r *GenTable) FindNotIn(column string, args ...interface{}) ([]GenTable, error) {
	var list []GenTable
	err := db.GetInstance().Engine().Table(r.TableName()).NotIn(column, args).Find(&list)
	return list, err
}
