package dao

import (
	"github.com/spf13/cast"
	"lostvip.com/db"
	"lostvip.com/namedsql"
	"robvi/app/system/model"
	"robvi/app/system/vo"
)

// Fill with you ideas below.

// 修改用户资料请求参数
type SysUserDao struct {
}

func (e SysUserDao) DeleteByIds(ida []int64) int64 {
	db := db.GetMasterGorm().Table("sys_user").Where("user_id in ? and user_id!=1 ", ida).Update("del_flag", 1)
	if db.Error != nil {
		panic(db.Error)
	}
	return db.RowsAffected
}

// 根据条件分页查询用户列表
func (d SysUserDao) SelectPageList(param *vo.SelectUserPageReq) (*[]map[string]string, int64, error) {
	db := db.GetMasterGorm()
	sqlParams, sql := d.GetSql(param)
	limitSql := sql + " order by u.user_id desc "
	limitSql += "  limit " + cast.ToString(param.GetStartNum()) + "," + cast.ToString(param.GetPageSize())
	result, err := namedsql.ListMap(db, limitSql, sqlParams, true)
	countSql := "select count(*) from (" + sql + ") t "
	total, err := namedsql.Count(db, countSql, sqlParams)
	return result, total, err
}

func (d SysUserDao) GetSql(param *vo.SelectUserPageReq) (map[string]interface{}, string) {
	sqlParams := make(map[string]interface{})
	sql := `
            select u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber, u.password,u.sex, u.salt, u.status, u.del_flag, 
            u.login_ip, u.login_date, u.create_by, u.create_time, u.remark,d.dept_name, d.leader
            from sys_user u left join sys_dept d on  u.dept_id = d.dept_id where u.del_flag =0 
           `
	if param != nil {
		if param.LoginName != "" {
			sql += " and  u.login_name like @loginName "
			sqlParams["loginName"] = "%" + param.LoginName + "%"
		}

		if param.Phonenumber != "" {
			sql += " and  u.phonenumber like @phonenumber "
			sqlParams["phonenumber"] = "%" + param.Phonenumber + "%"
		}

		if param.Status != "" {
			sql += " and  u.status = @status "
			sqlParams["status"] = param.Status
		}
		if param.TenantId != 0 {
			if param.TenantId != 0 {
				sql += " and  u.tenant_id = @TenantId "
				sqlParams["TenantId"] = param.TenantId
			}
		}
		if param.BeginTime != "" {
			sql += " and  u.create_time >= @BeginTime "
			sqlParams["BeginTime"] = param.BeginTime
		}
		if param.EndTime != "" {
			sql += " and  u.create_time <= @EndTime "
			sqlParams["EndTime"] = param.EndTime
		}
		if param.Ancestors != "" {
			sql += " and u.dept_id IN ( SELECT t.dept_id FROM sys_dept t WHERE t.ancestors like @Ancestors )"
			sqlParams["Ancestors"] = param.Ancestors + "%"
		}
	}
	return sqlParams, sql
}

// 导出excel
func (d SysUserDao) SelectExportList(param *vo.SelectUserPageReq) (*[]map[string]string, error) {
	db := db.GetMasterGorm()
	sqlParams, sql := d.GetSql(param)
	limitSql := sql + " order by u.user_id desc "
	result, err := namedsql.ListMap(db, limitSql, &sqlParams, true)
	return result, err
}

// 根据条件分页查询已分配用户角色列表
func (d SysUserDao) SelectAllocatedList(roleId int64, loginName, phonenumber string) (*[]map[string]string, error) {
	db := db.GetMasterGorm()
	sqlParams := make(map[string]interface{})
	sql := `
            select distinct u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber,u.status, u.create_time
            from sys_user u 
            left join sys_dept d on  u.dept_id = d.dept_id 
            left join sys_user_role ur on  u.user_id = ur.user_id
             left join sys_role r on  r.role_id = ur.role_id
            where u.del_flag =0 and  r.role_id = ` + cast.ToString(roleId)

	if loginName != "" {
		sql += " and u.login_name like @loginName "
		sqlParams["loginName"] = "%" + loginName + "%"
	}

	if phonenumber != "" {
		sql += " and  u.phonenumber like @phonenumber "
		sqlParams["phonenumber"] = "%" + phonenumber + "%"
	}

	result, err := namedsql.ListMap(db, sql, &sqlParams, true)
	return result, err
}

// 根据条件分页查询未分配用户角色列表
func (d SysUserDao) SelectUnallocatedList(roleId int64, loginName, phonenumber string) (*[]map[string]string, error) {
	db := db.GetMasterGorm()
	sqlParams := make(map[string]interface{})
	sql := `
            select distinct u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber,u.status, u.create_time
            from sys_user u 
            where u.del_flag =0  `
	sql += " and u.user_id not in (select u.user_id from sys_user u inner join sys_user_role ur on u.user_id = ur.user_id and ur.role_id = " + cast.ToString(roleId) + ") "
	if loginName != "" {
		sql += " and u.login_name like @loginName "
		sqlParams["loginName"] = "%" + loginName + "%"
	}

	if phonenumber != "" {
		sql += " and  u.phonenumber like @phonenumber "
		sqlParams["phonenumber"] = "%" + phonenumber + "%"
	}

	result, err := namedsql.ListMap(db, sql, &sqlParams, true)
	return result, err
}

//func (d SysUserDao) SelectUnallocatedList(roleId int64, loginName, phonenumber string) ([]model2.SysUser, error) {
//	db := db.GetInstance().Engine()
//	if db == nil {
//		return nil, errors.New("获取数据库连接失败")
//	}
//
//	model := db.Table("sys_user").Alias("u")
//	model.Join("Left", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id")
//	model.Join("Left", []string{"sys_user_role", "ur"}, "u.user_id = ur.user_id")
//	model.Join("Left", []string{"sys_role", "r"}, "r.role_id = ur.role_id")
//
//	model.Where("u.user_id not in (select u.user_id from sys_user u inner join sys_user_role ur on u.user_id = ur.user_id and ur.role_id = ?)", roleId)
//
//	if loginName != "" {
//		model.Where("u.login_name like ?", "%"+loginName+"%")
//	}
//
//	if phonenumber != "" {
//		model.Where("u.phonenumber like ?", "%"+phonenumber+"%")
//	}
//
//	model.Select("distinct u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber, u.status, u.create_time")
//
//	var result []model.SysUser
//	err := model.Find(&result)
//	return result, err
//}

// 检查邮箱是否已使用
func (d SysUserDao) CheckEmailUnique(userId int64, email string) bool {
	db := db.GetInstance().Engine()
	if db == nil {
		return false
	}
	rs, err := db.Table("sys_user").Where("email=? AND user_id<>?", email, userId).Count()
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
func (d SysUserDao) CheckEmailUniqueAll(email string) bool {
	db := db.GetInstance().Engine()
	if db == nil {
		return false
	}
	rs, err := db.Table("sys_user").Where("email=?", email).Count()
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
func (d SysUserDao) CheckPhoneUnique(userId int64, phone string) bool {
	db := db.GetInstance().Engine()
	if db == nil {
		return false
	}
	rs, err := db.Table("sys_user").Where("phonenumber = ? AND user_id<>?", phone, userId).Count()
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
func (d SysUserDao) CheckPhoneUniqueAll(phone string) bool {
	db := db.GetInstance().Engine()
	if db == nil {
		return false
	}
	rs, err := db.Table("sys_user").Where("phonenumber = ?", phone).Count()
	if err != nil {
		return false
	}

	if rs > 0 {
		return true
	} else {
		return false
	}
}

// 根据登录名查询用户信息
func (d SysUserDao) SelectUserByLoginName(loginName string) (*model.SysUser, error) {
	var entity = new(model.SysUser)
	entity.LoginName = loginName
	err := entity.FindOne()
	return entity, err
}

// 根据手机号查询用户信息
func (d SysUserDao) SelectUserByPhoneNumber(phonenumber string) (*model.SysUser, error) {
	var entity = new(model.SysUser)
	entity.Phonenumber = phonenumber
	err := entity.FindOne()
	return entity, err
}
