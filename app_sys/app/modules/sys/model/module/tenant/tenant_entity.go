// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2021-06-29 22:21:21 +0800 CST
// 生成路径: app/model/modules/tenant/SysTenant.go
// 生成人：robnote
// ==========================================================================

package tenant

import (
	"lostvip.com/db"
	"time"
)

// 数据表映射结构体
type SysTenant struct {
	Id         int64     `json:"id" xorm:"not null pk autoincr comment('id') bigint"`
	DelFlag    string    `json:"del_flag" xorm:"comment('删除标志（0代表存在 2代表删除）') char(1)"`
	CreateBy   string    `json:"create_by" xorm:"comment('创建者') varchar(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') datetime"`
	UpdateBy   string    `json:"update_by" xorm:"comment('更新者') varchar(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') datetime"`
	Name       string    `json:"name" xorm:"comment('商户名称') varchar(32)"`
	Address    string    `json:"address" xorm:"comment('联系地址') varchar(64)"`
	Manager    string    `json:"manager" xorm:"comment('负责人') varchar(32)"`
	Phone      string    `json:"phone" xorm:"comment('联系电话') varchar(18)"`
	Remark     string    `json:"remark" xorm:"comment('备注信息') varchar(255)"`
	StartTime  time.Time `json:"start_time" xorm:"comment('起租时间') datetime"`
	EndTime    time.Time `json:"end_time" xorm:"comment('结束时间') datetime"`
	Email      string    `json:"email" xorm:"comment('安全邮箱') varchar(255)"`
}

// 映射数据表
func TableName() string {
	return "sys_tenant"
}

// 插入数据
func (e *SysTenant) Insert() (int64, error) {
	return db.Instance().Engine().Table(TableName()).Insert(e)
}

// 更新数据
func (e *SysTenant) Update() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(e.Id).Update(e)
}

// 删除
func (e *SysTenant) Delete() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(e.Id).Delete(e)
}

// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(TableName()).In("id", ids).Delete(new(SysTenant))
}

// 根据结构体中已有的非空数据来获得单条数据
func (e *SysTenant) FindOne() (bool, error) {
	return db.Instance().Engine().Table(TableName()).Get(e)
}

// 根据条件查询
func Find(where, order string) ([]SysTenant, error) {
	var list []SysTenant
	err := db.Instance().Engine().Table(TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

// 指定字段集合查询
func FindIn(column string, args ...interface{}) ([]SysTenant, error) {
	var list []SysTenant
	err := db.Instance().Engine().Table(TableName()).In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func FindNotIn(column string, args ...interface{}) ([]SysTenant, error) {
	var list []SysTenant
	err := db.Instance().Engine().Table(TableName()).NotIn(column, args).Find(&list)
	return list, err
}
