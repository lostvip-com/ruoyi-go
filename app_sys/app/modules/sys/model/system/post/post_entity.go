package post

import (
	"lostvip.com/db"
	"time"
)

type SysPost struct {
	PostId     int64     `json:"post_id" xorm:"not null pk autoincr comment('岗位ID') BIGINT(20)"`
	PostCode   string    `json:"post_code" xorm:"not null comment('岗位编码') VARCHAR(64)"`
	PostName   string    `json:"post_name" xorm:"not null comment('岗位名称') VARCHAR(50)"`
	PostSort   int       `json:"post_sort" xorm:"not null comment('显示顺序') INT(4)"`
	Status     string    `json:"status" xorm:"not null comment('状态（0正常 1停用）') CHAR(1)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`

	TenantId int64 `json:"tenant_id" xorm:"default 0 comment('租户id') BIGINT(20)"`
}

// 映射数据表
func TableName() string {
	return "sys_post"
}

// 插入数据
func (r *SysPost) Insert() (int64, error) {
	return db.Instance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *SysPost) Update() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.PostId).Update(r)
}

// 删除
func (r *SysPost) Delete() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.PostId).Delete(r)
}

// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(TableName()).In("post_id", ids).Delete(new(SysPost))
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *SysPost) FindOne() (bool, error) {
	return db.Instance().Engine().Table(TableName()).Get(r)
}

// 根据条件查询
func Find(where, order string) ([]SysPost, error) {
	var list []SysPost
	err := db.Instance().Engine().Table(TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

// 指定字段集合查询
func FindIn(column string, args ...interface{}) ([]SysPost, error) {
	var list []SysPost
	err := db.Instance().Engine().Table(TableName()).In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func FindNotIn(column string, args ...interface{}) ([]SysPost, error) {
	var list []SysPost
	err := db.Instance().Engine().Table(TableName()).NotIn(column, args).Find(&list)
	return list, err
}
