package user

import (
	"errors"
	"lostvip.com/db"
	"lostvip.com/utils/excel"
	"lostvip.com/utils/page"
	"time"
	"xorm.io/builder"
)

// Fill with you ideas below.

// 修改用户资料请求参数
type ProfileReq struct {
	UserName    string `form:"userName"  binding:"required,min=5,max=30"`
	Phonenumber string `form:"phonenumber"  binding:"required,len=10"`
	Email       string `form:"email"  binding:"required,email"`
	Sex         string `form:"sex"  binding:"required"`
}

// 修改密码请求参数
type PasswordReq struct {
	OldPassword string `form:"oldPassword" binding:"required,min=5,max=30"`
	NewPassword string `form:"newPassword" binding:"required,min=5,max=30"`
	Confirm     string `form:"confirm" binding:"required,min=5,max=30"`
}

// 重置密码请求参数
type ResetPwdReq struct {
	UserId   int64  `form:"userId"  binding:"required"`
	Password string `form:"password" binding:"required,min=5,max=30"`
}

// 检查email请求参数
type CheckEmailReq struct {
	UserId int64  `form:"userId"  binding:"required,min=5,max=30"`
	Email  string `form:"email"  binding:"required,email"`
}

// 检查email请求参数
type CheckEmailAllReq struct {
	Email string `form:"email"  binding:"required,email"`
}

// 检查phone请求参数
type CheckLoginNameReq struct {
	LoginName string `form:"loginName"  binding:"required"`
}

// 检查phone请求参数
type CheckPhoneReq struct {
	UserId      int64  `form:"userId"  binding:"required`
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
}

// 检查phone请求参数
type CheckPhoneAllReq struct {
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
}

// 检查密码请求参数
type CheckPasswordReq struct {
	Password string `form:"password"  binding:"required"`
}

// 查询用户列表请求参数
type SelectPageReq struct {
	LoginName   string `form:"loginName"`     //登陆名
	Status      string `form:"status"`        //状态
	Phonenumber string `form:"phonenumber"`   //手机号码
	BeginTime   string `form:"beginTime"`     //数据范围
	EndTime     string `form:"endTime"`       //开始时间
	DeptId      int64  `form:"deptId"`        //结束时间
	PageNum     int    `form:"pageNum"`       //当前页码
	PageSize    int    `form:"pageSize"`      //每页数
	SortName    string `form:"orderByColumn"` //排序字段
	SortOrder   string `form:"isAsc"`         //排序方式
	Ancestors   string `form:"ancestors"`     //排序方式
	//
	TenantId int64 `form:"tenantId"`
}

// 用户列表数据结构
type UserListEntity struct {
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
	LoginIp     string    `json:"login_ip" xorm:"default '' comment('最后登陆IP') VARCHAR(50)"`
	LoginDate   time.Time `json:"login_date" xorm:"comment('最后登陆时间') DATETIME"`
	CreateBy    string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime  time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy    string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime  time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark      string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
	DeptName    string    `json:"dept_name"` // 部门名称
	Leader      string    `json:"leader"`    // 负责人
}

// 新增用户资料请求参数
type AddReq struct {
	UserName    string `form:"userName"  binding:"required,min=5,max=30"`
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
	Email       string `form:"email"  binding:"required,email"`
	LoginName   string `form:"loginName"  binding:"required"`
	Password    string `form:"password"  binding:"required,min=5,max=30"`
	DeptId      int64  `form:"deptId" binding:"required`
	Sex         string `form:"sex"  binding:"required"`
	Status      string `form:"status"`
	RoleIds     string `form:"roleIds"`
	PostIds     string `form:"postIds"`
	Remark      string `form:"remark"`
}

// 新增用户资料请求参数
type EditReq struct {
	UserId      int64  `form:"userId" binding:"required`
	UserName    string `form:"userName"  binding:"required,min=5,max=30"`
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
	Email       string `form:"email"  binding:"required,email"`
	DeptId      int64  `form:"deptId" binding:"required`
	Sex         string `form:"sex"  binding:"required"`
	Status      string `form:"status"`
	RoleIds     string `form:"roleIds"`
	PostIds     string `form:"postIds"`
	Remark      string `form:"remark"`
}

// 根据条件分页查询用户列表
func SelectPageList(param *SelectPageReq) ([]UserListEntity, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("u").Join("LEFT", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id")
	model.Where(" u.del_flag = '0' ")

	if param != nil {
		if param.LoginName != "" {
			model.Where("u.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.Phonenumber != "" {
			model.Where("u.phonenumber like ?", "%"+param.Phonenumber+"%")
		}

		if param.Status != "" {
			model.Where("u.status = ?", param.Status)
		}
		if param.TenantId != 0 {
			model.Where("u.tenant_id = ?", param.TenantId)
		}
		if param.BeginTime != "" {
			model.Where("u.create_time >= ? ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("u.create_time <= ? ", param.EndTime)
		}

		if param.Ancestors != "" {
			model.Where(" u.dept_id IN ( SELECT t.dept_id FROM sys_dept t WHERE t.ancestors like ? )", param.Ancestors+"%")
		}
	}

	totalModel := model.Clone()

	total, err := totalModel.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

	model.Select("u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber, u.password,u.sex, u.salt, u.status, u.del_flag, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark,d.dept_name, d.leader")

	model.OrderBy("u." + param.SortName + " " + param.SortOrder + " ")

	model.Limit(p.Pagesize, p.StartNum)

	var result []UserListEntity
	err = model.Find(&result)
	return result, p, err
}

// 导出excel
func SelectExportList(param *SelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()
	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(TableName(), "u").LeftJoin("sys_dept d", "u.dept_id = d.dept_id").Where(builder.Expr("u.del_flag = '0'"))

	if param != nil {
		if param.LoginName != "" {
			build.Where(builder.Like{"u.login_name", param.LoginName})
		}

		if param.Phonenumber != "" {
			build.Where(builder.Like{"u.phonenumber", param.Phonenumber})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"u.status": param.Status})
		}

		if param.BeginTime != "" {
			build.Where(builder.Gte{"u.create_time": param.BeginTime})
		}

		if param.EndTime != "" {
			build.Where(builder.Lte{"u.create_time": param.EndTime})
		}

		if param.Ancestors != "" {
			deptIdSql := " SELECT t.dept_id FROM sys_dept t WHERE t.ancestors like ?  "
			build.Where(builder.In("u.dept_id", builder.Expr(deptIdSql, param.Ancestors+"%")))
		}
	}

	sqlStr, _, _ := build.ToSQL()
	arr, err := db.SQL(sqlStr).QuerySliceString()

	path, err := excel.DownlaodExcel(head, arr)

	return path, err
}

// 根据条件分页查询已分配用户角色列表
func SelectAllocatedList(roleId int64, loginName, phonenumber string) ([]SysUser, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("u")
	model.Join("Left", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id")
	model.Join("Left", []string{"sys_user_role", "ur"}, " u.user_id = ur.user_id")
	model.Join("Left", []string{"sys_role", "r"}, "r.role_id = ur.role_id")

	model.Where("u.del_flag =?", 0)
	model.Where("r.role_id = ?", roleId)

	if loginName != "" {
		model.Where("u.login_name like ?", "%"+loginName+"%")
	}

	if phonenumber != "" {
		model.Where("u.phonenumber like ?", "%"+phonenumber+"%")
	}

	var result []SysUser

	model.Select("distinct u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber,u.status, u.create_time")

	err := model.Find(&result)
	return result, err
}

// 根据条件分页查询未分配用户角色列表
func SelectUnallocatedList(roleId int64, loginName, phonenumber string) ([]SysUser, error) {
	db := db.Instance().Engine()
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("u")
	model.Join("Left", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id")
	model.Join("Left", []string{"sys_user_role", "ur"}, "u.user_id = ur.user_id")
	model.Join("Left", []string{"sys_role", "r"}, "r.role_id = ur.role_id")

	model.Where("u.user_id not in (select u.user_id from sys_user u inner join sys_user_role ur on u.user_id = ur.user_id and ur.role_id = ?)", roleId)

	if loginName != "" {
		model.Where("u.login_name like ?", "%"+loginName+"%")
	}

	if phonenumber != "" {
		model.Where("u.phonenumber like ?", "%"+phonenumber+"%")
	}

	model.Select("distinct u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber, u.status, u.create_time")

	var result []SysUser
	err := model.Find(&result)
	return result, err
}

// 检查邮箱是否已使用
func CheckEmailUnique(userId int64, email string) bool {
	db := db.Instance().Engine()
	if db == nil {
		return false
	}
	rs, err := db.Table(TableName()).Where("email=? AND user_id<>?", email, userId).Count()
	if err != nil {
		return false
	}

	if rs > 0 {
		return true
	} else {
		return false
	}
}

// 检查邮箱是否存在,存在返回true,否则false
func CheckEmailUniqueAll(email string) bool {
	db := db.Instance().Engine()
	if db == nil {
		return false
	}
	rs, err := db.Table(TableName()).Where("email=?", email).Count()
	if err != nil {
		return false
	}

	if rs > 0 {
		return true
	} else {
		return false
	}
}

// 检查手机号是否已使用,存在返回true,否则false
func CheckPhoneUnique(userId int64, phone string) bool {
	db := db.Instance().Engine()
	if db == nil {
		return false
	}
	rs, err := db.Table(TableName()).Where("phonenumber = ? AND user_id<>?", phone, userId).Count()
	if err != nil {
		return false
	}

	if rs > 0 {
		return true
	} else {
		return false
	}
}

// 检查手机号是否已使用 ,存在返回true,否则false
func CheckPhoneUniqueAll(phone string) bool {
	db := db.Instance().Engine()
	if db == nil {
		return false
	}
	rs, err := db.Table(TableName()).Where("phonenumber = ?", phone).Count()
	if err != nil {
		return false
	}

	if rs > 0 {
		return true
	} else {
		return false
	}
}

// 根据登陆名查询用户信息
func SelectUserByLoginName(loginName string) (*SysUser, error) {
	var entity SysUser
	entity.LoginName = loginName
	ok, err := entity.FindOne()
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}

// 根据手机号查询用户信息
func SelectUserByPhoneNumber(phonenumber string) (*SysUser, error) {
	var entity SysUser
	entity.Phonenumber = phonenumber
	ok, err := entity.FindOne()
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}
