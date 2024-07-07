package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/cache/myredis"
	"github.com/lv_framework/db"
	"github.com/lv_framework/utils/lv_conv"
	"github.com/lv_framework/utils/lv_err"
	"github.com/lv_framework/utils/lv_gen"
	"github.com/lv_framework/utils/lv_net"
	"github.com/lv_framework/utils/lv_office"
	"github.com/lv_framework/utils/lv_secret"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"main/app/system/dao"
	"main/app/system/model"
	"main/app/system/vo"
	"strings"
	"time"
)

type UserService struct{}

// 根据主键查询用户信息
func (svc *UserService) SelectRecordById(id int64) (*model.SysUser, error) {
	entity := &model.SysUser{UserId: id}
	err := entity.FindOne()
	return entity, err
}

// 根据条件分页查询用户列表
func (svc UserService) SelectRecordList(param *vo.SelectUserPageReq) (*[]map[string]string, int64, error) {
	var deptService DeptService
	var dept = deptService.SelectDeptById(param.DeptId)
	if dept != nil { //数据权限
		param.Ancestors = dept.Ancestors
	}
	var d dao.SysUserDao
	return d.SelectPageList(param)
}

// 导出excel
func (svc UserService) Export(param *vo.SelectUserPageReq) (string, error) {
	head := []string{"用户名", "呢称", "Email", "电话号码", "性别", "部门", "领导", "状态", "删除标记", "创建人", "创建时间", "备注"}
	col := []string{"loginName", "userName", "u.email", "phonenumber", "sex", "deptName", "leader", "status", "delFlag", "createBy", "createTime", "Remark"}
	var d dao.SysUserDao
	listMap, err := d.SelectExportList(param)
	lv_err.HasErrAndPanic(err)
	return lv_office.DownlaodExcelByListMap(&head, &col, listMap)
}

// 新增用户
func (svc UserService) AddSave(req *vo.AddUserReq, c *gin.Context) (int64, error) {
	var u model.SysUser
	u.LoginName = req.LoginName
	u.UserName = req.UserName
	u.Email = req.Email
	u.Phonenumber = req.Phonenumber
	u.Status = req.Status
	u.Sex = req.Sex
	u.DeptId = req.DeptId
	u.Remark = req.Remark

	//生成密码
	newSalt := lv_gen.GenerateSubId(6)
	newToken := req.LoginName + req.Password + newSalt
	newToken = lv_secret.MustEncryptString(newToken)
	u.Salt = newSalt
	u.Password = newToken
	u.CreateTime = time.Now()
	createUser := svc.GetProfile(c)

	if createUser != nil {
		u.CreateBy = createUser.LoginName
	}
	u.DelFlag = 0

	err := db.GetMasterGorm().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&u).Error; err != nil {
			return err
		}
		//增加岗位数据
		if req.PostIds != "" {
			postIds := lv_conv.ToInt64Array(req.PostIds, ",")
			userPosts := make([]model.SysUserPost, 0)
			for i := range postIds {
				if postIds[i] > 0 {
					var userPost model.SysUserPost
					userPost.UserId = u.UserId
					userPost.PostId = postIds[i]
					userPosts = append(userPosts, userPost)
				}
			} //end for
			if len(userPosts) > 0 {
				if err := tx.CreateInBatches(userPosts, 1).Error; err != nil {
					return err
				}
			}
		}
		//增加角色数据
		if req.RoleIds != "" {
			roleIds := lv_conv.ToInt64Array(req.RoleIds, ",")
			userRoles := make([]model.SysUserRole, 0)
			for i := range roleIds {
				if roleIds[i] > 0 {
					var userRole model.SysUserRole
					userRole.UserId = u.UserId
					userRole.RoleId = roleIds[i]
					userRoles = append(userRoles, userRole)
				}
			}
			if len(userRoles) > 0 {
				if err := tx.CreateInBatches(userRoles, 1).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})

	return u.UserId, err
}

// 新增用户
func (svc UserService) EditSave(req *vo.EditUserReq, c *gin.Context) error {
	userPtr := &model.SysUser{UserId: req.UserId}
	err := userPtr.FindOne()
	if err != nil {
		return err
	}
	userPtr.UserName = req.UserName
	userPtr.Email = req.Email
	userPtr.Phonenumber = req.Phonenumber
	userPtr.Status = req.Status
	userPtr.Sex = req.Sex
	userPtr.DeptId = req.DeptId
	userPtr.Remark = req.Remark
	userPtr.UpdateTime = time.Now()
	updateUser := svc.GetProfile(c)

	if updateUser != nil {
		userPtr.UpdateBy = updateUser.LoginName
	}
	err = db.GetMasterGorm().Transaction(func(tx *gorm.DB) error {
		if err := tx.Updates(userPtr).Error; err != nil {
			return err
		}
		//增加岗位数据
		if req.PostIds != "" {
			postIds := lv_conv.ToInt64Array(req.PostIds, ",")
			userPosts := make([]model.SysUserPost, 0)
			for i := range postIds {
				if postIds[i] > 0 {
					var userPost model.SysUserPost
					userPost.UserId = userPtr.UserId
					userPost.PostId = postIds[i]
					userPosts = append(userPosts, userPost)
				}
			} //end for
			if len(userPosts) > 0 {
				tx.Exec("delete from sys_user_post where user_id=?", userPtr.UserId)
				if err := tx.Save(userPosts).Error; err != nil {
					return err
				}
			}
		}
		//增加角色数据
		if req.RoleIds != "" {
			roleIds := lv_conv.ToInt64Array(req.RoleIds, ",")
			userRoles := make([]model.SysUserRole, 0)
			for i := range roleIds {
				if roleIds[i] > 0 {
					var userRole model.SysUserRole
					userRole.UserId = userPtr.UserId
					userRole.RoleId = roleIds[i]
					userRoles = append(userRoles, userRole)
				}
			} //end for
			if len(userRoles) > 0 {
				tx.Exec("delete from sys_user_role where user_id=?", userPtr.UserId)
				if err := tx.Save(userRoles).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})

	return err
}

// 根据主键删除用户信息
func (svc UserService) DeleteRecordById(id int64) error {
	entity := &model.SysUser{UserId: id}
	err := entity.Delete()
	return err
}

// 批量删除用户记录
func (svc UserService) DeleteRecordByIds(ids string) error {
	idarr := lv_conv.ToInt64Array(ids, ",")
	idarr = lv_conv.RemoveOne(idarr, 1) //去掉admin的id
	if len(idarr) == 0 {
		return errors.New("ids can not be empty ")
	}
	err := db.GetMasterGorm().Transaction(func(tx *gorm.DB) error {
		err := tx.Table("sys_user").Where("user_id in ? and user_id!=1 ", idarr).Update("del_flag", 1).Error
		if err != nil {
			return err
		}
		err = tx.Table("sys_user_post").Where("user_id in ? and user_id!=1 ", idarr).Update("del_flag", 1).Error
		if err != nil {
			return err
		}
		err = tx.Table("sys_user_role").Where("user_id in ? and user_id!=1 ", idarr).Update("del_flag", 1).Error
		if err != nil {
			return err
		}
		return err
	})
	return err
}

// 判断是否是系统管理员
func (svc UserService) IsAdmin(userId int64) bool {
	if userId == 1 {
		return true
	} else {
		return false
	}
}

func (svc UserService) IsSignedIn(tokenStr string) bool {
	key := "login:" + tokenStr
	yes := myredis.GetInstance().Exists(context.Background(), key).Val()
	fmt.Println("===============" + string(yes))
	return yes > 0
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func (svc UserService) SignIn(loginnName, password string) (*model.SysUser, error) {
	//查询用户信息
	user := model.SysUser{LoginName: loginnName}
	err := user.FindOne()

	if err != nil {
		return nil, err
	}
	//校验密码
	pwdNew := user.LoginName + password + user.Salt

	pwdNew = lv_secret.MustEncryptString(pwdNew)

	if strings.Compare(user.Password, pwdNew) == -1 {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}

// 清空用户菜单缓存

// 用户注销
func (svc UserService) SignOut(c *gin.Context) error {
	//user := svc.GetProfile(c)
	tokenStr := lv_net.GetParam(c, "token")
	myredis.GetInstance().Del(context.Background(), "login:"+tokenStr)
	//userId := c.MustGet("userId").(int64)
	return nil
}

// 强退用户
func (svc UserService) ForceLogout(c *gin.Context) error {
	svc.SignOut(c)
	return nil
}

// 检查账号是否符合规范,存在返回false,否则true
func (svc UserService) CheckPassport(loginName string) bool {
	entity := model.SysUser{LoginName: loginName}
	if err := entity.FindOne(); err != nil {
		return false
	} else {
		return true
	}
}

// 检查登录名是否存在,存在返回true,否则false
func (svc UserService) CheckNickName(userName string) bool {
	entity := model.SysUser{UserName: userName}
	if err := entity.FindOne(); err != nil {
		return false
	} else {
		return true
	}
}

// 检查登录名是否存在,存在返回true,否则false
func (svc UserService) CheckLoginName(loginName string) bool {
	entity := model.SysUser{LoginName: loginName}
	if err := entity.FindOne(); err != nil {
		return false
	} else {
		return true
	}
}

// 获得用户信息详情
func (svc UserService) GetProfile(c *gin.Context) *model.SysUser {
	token := lv_net.GetParam(c, "token")
	key := "login:" + token
	userId := myredis.GetInstance().HGet(context.Background(), key, "userId").Val()
	roleKeys := myredis.GetInstance().HGet(context.Background(), key, "roleKeys").Val()
	u := new(model.SysUser)
	u.UserId = cast.ToInt64(userId)
	err := u.FindOne()
	if err != nil {
		panic(err)
	} else {
		u.RoleKeys = roleKeys
	}
	return u
}

// 更新用户信息详情
func (svc UserService) UpdateProfile(profile *vo.ProfileReq, c *gin.Context) error {
	user := svc.GetProfile(c)

	if profile.UserName != "" {
		user.UserName = profile.UserName
	}

	if profile.Email != "" {
		user.Email = profile.Email
	}

	if profile.Phonenumber != "" {
		user.Phonenumber = profile.Phonenumber
	}

	if profile.Sex != "" {
		user.Sex = profile.Sex
	}

	err := user.Updates()
	if err != nil {
		return errors.New("保存数据失败")
	}

	//SaveUserToSession(user, c)
	return nil
}

// 更新用户头像
func (svc UserService) UpdateAvatar(avatar string, c *gin.Context) error {
	user := svc.GetProfile(c)

	if avatar != "" {
		user.Avatar = avatar
	}

	err := user.Updates()
	if err != nil {
		return errors.New("保存数据失败")
	}

	//SaveUserToSession(user, c)
	return nil
}

// 修改用户密码
func (svc UserService) UpdatePassword(profile *vo.PasswordReq, c *gin.Context) error {
	user := svc.GetProfile(c)

	if profile.OldPassword == "" {
		return errors.New("旧密码不能为空")
	}

	if profile.NewPassword == "" {
		return errors.New("新密码不能为空")
	}

	if profile.Confirm == "" {
		return errors.New("确认密码不能为空")
	}

	if profile.NewPassword == profile.OldPassword {
		return errors.New("新旧密码不能相同")
	}

	if profile.Confirm != profile.NewPassword {
		return errors.New("确认密码不一致")
	}

	//校验密码
	token := user.LoginName + profile.OldPassword + user.Salt
	token = lv_secret.MustEncryptString(token)

	if token != user.Password {
		return errors.New("原密码不正确")
	}

	//新校验密码
	newSalt := lv_gen.GenerateSubId(6)
	newToken := user.LoginName + profile.NewPassword + newSalt
	newToken = lv_secret.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken

	err := user.Updates()
	if err != nil {
		return errors.New("保存数据失败")
	}

	//SaveUserToSession(user, c)
	return nil
}

// 重置用户密码
func (svc UserService) ResetPassword(params *vo.ResetPwdReq) (bool, error) {
	user := model.SysUser{UserId: params.UserId}
	if err := user.FindOne(); err != nil {
		return false, errors.New("用户不存在")
	}
	//新校验密码
	newSalt := lv_gen.GenerateSubId(6)
	newToken := user.LoginName + params.Password + newSalt
	newToken = lv_secret.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken
	if err := user.Updates(); err != nil {
		return false, errors.New("保存数据失败")
	}
	return true, nil
}

// 校验密码是否正确
func (svc UserService) CheckPassword(user *model.SysUser, password string) bool {
	if user == nil || user.UserId <= 0 {
		return false
	}

	//校验密码
	token := user.LoginName + password + user.Salt
	token = lv_secret.MustEncryptString(token)

	if strings.Compare(token, user.Password) == 0 {
		return true
	} else {
		return false
	}
}

// 检查邮箱是否已使用
func (svc UserService) CheckEmailUnique(userId int64, email string) bool {
	var vo dao.SysUserDao
	return vo.CheckEmailUnique(userId, email)
}

// 检查邮箱是否存在,存在返回true,否则false
func (svc UserService) CheckEmailUniqueAll(email string) bool {
	var vo dao.SysUserDao
	return vo.CheckEmailUniqueAll(email)
}

// 检查手机号是否已使用,存在返回true,否则false
func (svc UserService) CheckPhoneUnique(userId int64, phone string) bool {
	var vo dao.SysUserDao
	return vo.CheckPhoneUnique(userId, phone)
}

// 检查手机号是否已使用 ,存在返回true,否则false
func (svc UserService) CheckPhoneUniqueAll(phone string) bool {
	var vo dao.SysUserDao
	return vo.CheckPhoneUniqueAll(phone)
}

// 根据登录名查询用户信息
func (svc UserService) SelectUserByLoginName(loginName string) (*model.SysUser, error) {
	var vo dao.SysUserDao
	return vo.SelectUserByLoginName(loginName)
}

// 根据手机号查询用户信息
func (svc UserService) SelectUserByPhoneNumber(phonenumber string) (*model.SysUser, error) {
	var vo dao.SysUserDao
	return vo.SelectUserByPhoneNumber(phonenumber)
}

// 查询已分配用户角色列表
func (svc UserService) SelectAllocatedList(roleId int64, loginName, phonenumber string) (*[]map[string]string, error) {
	var vo dao.SysUserDao
	return vo.SelectAllocatedList(roleId, loginName, phonenumber)
}

// 查询未分配用户角色列表
func (svc UserService) SelectUnallocatedList(roleId int64, loginName, phonenumber string) (*[]map[string]string, error) {
	var vo dao.SysUserDao
	return vo.SelectUnallocatedList(roleId, loginName, phonenumber)
}

// 查询未分配用户角色列表
func (svc UserService) GetRoleKeys(userId int64) (string, error) {
	if userId == 1 {
		return "admin", nil
	}
	var sql = " SELECT GROUP_CONCAT(r.role_key) roles from sys_user_role ur,sys_role r where ur.user_id=? and ur.role_id = r.role_id "
	var roles string
	err := db.GetMasterGorm().Raw(sql, userId).Scan(&roles).Error
	return roles, err
}
