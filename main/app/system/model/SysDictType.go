package model

import (
	"lostvip.com/db"
	"time"
)

type DictType struct {
	DictId     int64     `json:"dictId" xorm:"not null pk autoincr comment('字典主键') BIGINT(20)"`
	DictName   string    `json:"dictName" xorm:"default '' comment('字典名称') VARCHAR(100)"`
	DictType   string    `json:"dictType" xorm:"default '' comment('字典类型') unique VARCHAR(100)"`
	Status     string    `json:"status" xorm:"default '0' comment('状态（0正常 1停用）') CHAR(1)"`
	CreateBy   string    `json:"createBy" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"createTime" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"updateBy" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"updateTime" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
}

// 映射数据表
func (e *DictType) TableName() string {
	return "sys_dict_type"
}

// 插入数据
func (r *DictType) Insert() (int64, error) {
	return db.GetInstance().Engine().Table("sys_dict_type").Insert(r)
}

// 更新数据
func (r *DictType) Update() (int64, error) {
	return db.GetInstance().Engine().Table("sys_dict_type").ID(r.DictId).Update(r)
}

// 删除
func (r *DictType) Delete() (int64, error) {
	return db.GetInstance().Engine().Table("sys_dict_type").ID(r.DictId).Delete(r)
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *DictType) FindOne() (bool, error) {
	return db.GetInstance().Engine().Table("sys_dict_type").Get(r)
}

// 根据条件查询
func (r *DictType) Find(where, order string) ([]DictType, error) {
	var list []DictType
	err := db.GetInstance().Engine().Table("sys_dict_type").Where(where).OrderBy(order).Find(&list)
	return list, err
}
