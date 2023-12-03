package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"lostvip.com/cache/myredis"
	"lostvip.com/db"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_gen"
	"lostvip.com/utils/lv_net"
	"lostvip.com/utils/lv_secret"
	"lostvip.com/utils/lv_web"
	"robvi/app/common/global"
	"robvi/app/common/session"
	"robvi/app/system/model/monitor/online"
	"robvi/app/system/model/system/user"
	user_post2 "robvi/app/system/model/system/user_post"
	user_role2 "robvi/app/system/model/system/user_role"
	"strings"
	"time"
)

type UserService struct{}

// 根据主键查询用户信息
func (svc UserService) SelectRecordById(id int64) (*user.SysUser, error) {
	entity := &user.SysUser{UserId: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据条件分页查询用户列表
func (svc UserService) SelectRecordList(param *user.SelectPageReq) ([]user.UserListEntity, *lv_web.Paging, error) {
	var deptService DeptService
	var d = deptService.SelectDeptById(param.DeptId)
	if d != nil { //数据权限
		param.Ancestors = d.Ancestors
	}
	return user.SelectPageList(param)
}

// 导出excel
func (svc UserService) Export(param *user.SelectPageReq) (string, error) {
	head := []string{"用户名", "呢称", "Email", "电话号码", "性别", "部门", "领导", "状态", "删除标记", "创建人", "创建时间", "备注"}
	col := []string{"u.login_name", "u.user_name", "u.email", "u.phonenumber", "u.sex", "d.dept_name", "d.leader", "u.status", "u.del_flag", "u.create_by", "u.create_time", "u.remark"}
	return user.SelectExportList(param, head, col)
}

// 新增用户
func (svc UserService) AddSave(req *user.AddReq, c *gin.Context) (int64, error) {
	var u user.SysUser
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

	u.DelFlag = "0"

	session := db.Instance().Engine().NewSession()
	err := session.Begin()

	_, err = session.Table(user.TableName()).Insert(&u)

	if err != nil || u.UserId <= 0 {
		session.Rollback()
		return 0, err
	}

	//增加岗位数据
	if req.PostIds != "" {
		postIds := lv_conv.ToInt64Array(req.PostIds, ",")
		userPosts := make([]user_post2.Entity, 0)
		for i := range postIds {
			if postIds[i] > 0 {
				var userPost user_post2.Entity
				userPost.UserId = u.UserId
				userPost.PostId = postIds[i]
				userPosts = append(userPosts, userPost)
			}
		}
		if len(userPosts) > 0 {
			_, err := session.Table(user_post2.TableName()).Insert(userPosts)
			if err != nil {
				session.Rollback()
				return 0, err
			}
		}

	}

	//增加角色数据
	if req.RoleIds != "" {
		roleIds := lv_conv.ToInt64Array(req.RoleIds, ",")
		userRoles := make([]user_role2.Entity, 0)
		for i := range roleIds {
			if roleIds[i] > 0 {
				var userRole user_role2.Entity
				userRole.UserId = u.UserId
				userRole.RoleId = roleIds[i]
				userRoles = append(userRoles, userRole)
			}
		}
		if len(userRoles) > 0 {
			_, err := session.Table(user_role2.TableName()).Insert(userRoles)
			if err != nil {
				session.Rollback()
				return 0, err
			}
		}
	}

	return u.UserId, session.Commit()
}

// 新增用户
func (svc UserService) EditSave(req *user.EditReq, c *gin.Context) (int64, error) {
	u := &user.SysUser{UserId: req.UserId}
	ok, err := u.FindOne()
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errors.New("数据不存在")
	}

	u.UserName = req.UserName
	u.Email = req.Email
	u.Phonenumber = req.Phonenumber
	u.Status = req.Status
	u.Sex = req.Sex
	u.DeptId = req.DeptId
	u.Remark = req.Remark

	u.UpdateTime = time.Now()

	updateUser := svc.GetProfile(c)

	if updateUser != nil {
		u.UpdateBy = updateUser.LoginName
	}

	session := db.Instance().Engine().NewSession()
	tanErr := session.Begin()

	_, tanErr = session.Table(user.TableName()).ID(u.UserId).Update(u)

	if tanErr != nil {
		session.Rollback()
		return 0, tanErr
	}

	//增加岗位数据
	if req.PostIds != "" {
		postIds := lv_conv.ToInt64Array(req.PostIds, ",")
		userPosts := make([]user_post2.Entity, 0)
		for i := range postIds {
			if postIds[i] > 0 {
				var userPost user_post2.Entity
				userPost.UserId = u.UserId
				userPost.PostId = postIds[i]
				userPosts = append(userPosts, userPost)
			}
		}
		if len(userPosts) > 0 {
			session.Exec("delete from sys_user_post where user_id=?", u.UserId)
			_, tanErr = session.Table(user_post2.TableName()).Insert(userPosts)
			if tanErr != nil {
				session.Rollback()
				return 0, err
			}
		}

	}

	//增加角色数据
	if req.RoleIds != "" {
		roleIds := lv_conv.ToInt64Array(req.RoleIds, ",")
		userRoles := make([]user_role2.Entity, 0)
		for i := range roleIds {
			if roleIds[i] > 0 {
				var userRole user_role2.Entity
				userRole.UserId = u.UserId
				userRole.RoleId = roleIds[i]
				userRoles = append(userRoles, userRole)
			}
		}
		if len(userRoles) > 0 {
			session.Exec("delete from sys_user_role where user_id=?", u.UserId)
			_, err := session.Table(user_role2.TableName()).Insert(userRoles)
			if tanErr != nil {
				session.Rollback()
				return 0, err
			}
		}
	}

	return 1, session.Commit()
}

// 根据主键删除用户信息
func (svc UserService) DeleteRecordById(id int64) bool {
	entity := &user.SysUser{UserId: id}
	result, _ := entity.Delete()
	if result > 0 {
		return true
	}
	return false
}

// 批量删除用户记录
func (svc UserService) DeleteRecordByIds(ids string) int64 {
	idarr := lv_conv.ToInt64Array(ids, ",")
	result, _ := user.DeleteBatch(idarr...)
	user_role2.DeleteBatch(idarr...)
	user_post2.DeleteBatch(idarr...)
	return result
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
func (svc UserService) SignIn(loginnName, password string) (*user.SysUser, error) {
	//查询用户信息
	user := user.SysUser{LoginName: loginnName}
	ok, err := user.FindOne()

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("用户名或者密码错误")
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
func (svc UserService) ForceLogout(sessionId string) error {
	var tmp interface{}
	if v, ok := global.SessionList.Load(sessionId); ok {
		tmp = v
	}

	if tmp != nil {
		c := tmp.(*gin.Context)
		if c != nil {
			svc.SignOut(c)
			global.SessionList.Delete(sessionId)
			entity := online.UserOnline{Sessionid: sessionId}
			entity.Delete()
		}
	}

	return nil
}

// 检查账号是否符合规范,存在返回false,否则true
func (svc UserService) CheckPassport(loginName string) bool {
	entity := user.SysUser{LoginName: loginName}
	if ok, err := entity.FindOne(); err != nil || !ok {
		return false
	} else {
		return true
	}
}

// 检查登陆名是否存在,存在返回true,否则false
func (svc UserService) CheckNickName(userName string) bool {
	entity := user.SysUser{UserName: userName}
	if ok, err := entity.FindOne(); err != nil || !ok {
		return false
	} else {
		return true
	}
}

// 检查登陆名是否存在,存在返回true,否则false
func (svc UserService) CheckLoginName(loginName string) bool {
	entity := user.SysUser{LoginName: loginName}
	if ok, err := entity.FindOne(); err != nil || !ok {
		return false
	} else {
		return true
	}
}

// 获得用户信息详情
func (svc UserService) GetProfile(c *gin.Context) *user.SysUser {

	var u = session.GetProfile(c)
	return u
}

// 更新用户信息详情
func (svc UserService) UpdateProfile(profile *user.ProfileReq, c *gin.Context) error {
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

	_, err := user.Update()
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

	_, err := user.Update()
	if err != nil {
		return errors.New("保存数据失败")
	}

	//SaveUserToSession(user, c)
	return nil
}

// 修改用户密码
func (svc UserService) UpdatePassword(profile *user.PasswordReq, c *gin.Context) error {
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

	_, err := user.Update()
	if err != nil {
		return errors.New("保存数据失败")
	}

	//SaveUserToSession(user, c)
	return nil
}

// 重置用户密码
func (svc UserService) ResetPassword(params *user.ResetPwdReq) (bool, error) {
	user := user.SysUser{UserId: params.UserId}
	if ok, err := user.FindOne(); err != nil || !ok {
		return false, errors.New("用户不存在")
	}
	//新校验密码
	newSalt := lv_gen.GenerateSubId(6)
	newToken := user.LoginName + params.Password + newSalt
	newToken = lv_secret.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken
	if _, err := user.Update(); err != nil {
		return false, errors.New("保存数据失败")
	}
	return true, nil
}

// 校验密码是否正确
func (svc UserService) CheckPassword(user *user.SysUser, password string) bool {
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
	return user.CheckEmailUnique(userId, email)
}

// 检查邮箱是否存在,存在返回true,否则false
func (svc UserService) CheckEmailUniqueAll(email string) bool {
	return user.CheckEmailUniqueAll(email)
}

// 检查手机号是否已使用,存在返回true,否则false
func (svc UserService) CheckPhoneUnique(userId int64, phone string) bool {
	return user.CheckPhoneUnique(userId, phone)
}

// 检查手机号是否已使用 ,存在返回true,否则false
func (svc UserService) CheckPhoneUniqueAll(phone string) bool {
	return user.CheckPhoneUniqueAll(phone)
}

// 根据登陆名查询用户信息
func (svc UserService) SelectUserByLoginName(loginName string) (*user.SysUser, error) {
	return user.SelectUserByLoginName(loginName)
}

// 根据手机号查询用户信息
func (svc UserService) SelectUserByPhoneNumber(phonenumber string) (*user.SysUser, error) {
	return user.SelectUserByPhoneNumber(phonenumber)
}

// 查询已分配用户角色列表
func (svc UserService) SelectAllocatedList(roleId int64, loginName, phonenumber string) ([]user.SysUser, error) {
	return user.SelectAllocatedList(roleId, loginName, phonenumber)
}

// 查询未分配用户角色列表
func (svc UserService) SelectUnallocatedList(roleId int64, loginName, phonenumber string) ([]user.SysUser, error) {
	return user.SelectUnallocatedList(roleId, loginName, phonenumber)
}
