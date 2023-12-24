package user

import (
	"lostvip.com/db"
	"time"
)

type SysUser struct {
	UserId      int64     `json:"user_id" xorm:"not null pk autoincr comment('用户ID') BIGINT(20)"`
	DeptId      int64     `json:"dept_id" xorm:"comment('部门ID') BIGINT(20)"`
	LoginName   string    `json:"login_name" xorm:"not null comment('登录账号') VARCHAR(30)"`
	UserName    string    `json:"user_name" xorm:"not null comment('用户昵称') VARCHAR(30)"`
	UserType    string    `json:"user_type" xorm:"default '00' comment('用户类型（00系统用户）') VARCHAR(2)"`
	Email       string    `json:"email" xorm:"default '' comment('用户邮箱') VARCHAR(50)"`
	Phonenumber string    `json:"phonenumber" xorm:"default '' comment('手机号码') VARCHAR(11)"`
	Sex         string    `json:"sex" xorm:"default '0' comment('用户性别（0男 1女 2未知）') CHAR(1)"`
	Avatar      string    `json:"avatar" xorm:"default '' comment('头像路径') VARCHAR(100)"`
	Password    string    `json:"password" xorm:"default '' comment('密码') VARCHAR(50)"`
	Salt        string    `json:"salt" xorm:"default '' comment('盐加密') VARCHAR(20)"`
	Status      string    `json:"status" xorm:"default '0' comment('帐号状态（0正常 1停用）') CHAR(1)"`
	DelFlag     string    `json:"del_flag" xorm:"default '0' comment('删除标志（0代表存在 2代表删除）') CHAR(1)"`
	LoginIp     string    `json:"login_ip" xorm:"default '' comment('最后登录IP') VARCHAR(50)"`
	LoginDate   time.Time `json:"login_date" xorm:"comment('最后登录时间') DATETIME"`
	CreateBy    string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime  time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy    string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime  time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark      string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
	//
	TenantId int64 `json:"tenant_id" xorm:"default 0 comment('租户id') BIGINT(20)"`
}

// 映射数据表
func TableName() string {
	return "sys_user"
}

// 插入数据
func (r *SysUser) Insert() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *SysUser) Update() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).ID(r.UserId).Update(r)
}

// 删除
func (r *SysUser) Delete() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).ID(r.UserId).Delete(r)
}

// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).In("user_id", ids).Delete(new(SysUser))
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *SysUser) FindOne() (bool, error) {
	return db.GetInstance().Engine().Table(TableName()).Get(r)
}

// 根据条件查询
func Find(where, order string) ([]SysUser, error) {
	var list []SysUser
	err := db.GetInstance().Engine().Table(TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

// 指定字段集合查询
func FindIn(column string, args ...interface{}) ([]SysUser, error) {
	var list []SysUser
	err := db.GetInstance().Engine().Table(TableName()).In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func FindNotIn(column string, args ...interface{}) ([]SysUser, error) {
	var list []SysUser
	err := db.GetInstance().Engine().Table(TableName()).NotIn(column, args).Find(&list)
	return list, err
}
