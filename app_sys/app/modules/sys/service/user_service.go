package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"lostvip.com/db"
	"lostvip.com/myredis"
	"lostvip.com/utils/convert"
	"lostvip.com/utils/lib_net"
	"lostvip.com/utils/lib_secret"
	"lostvip.com/utils/page"
	"lostvip.com/utils/random"
	"robvi/app/common/session"
	"robvi/app/global"
	"robvi/app/modules/sys/model/monitor/online"
	user2 "robvi/app/modules/sys/model/system/user"
	"robvi/app/modules/sys/model/system/user_post"
	"robvi/app/modules/sys/model/system/user_role"
	"strings"
	"time"
)

type UserService struct{}

// 根据主键查询用户信息
func (svc UserService) SelectRecordById(id int64) (*user2.SysUser, error) {
	entity := &user2.SysUser{UserId: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据条件分页查询用户列表
func (svc UserService) SelectRecordList(param *user2.SelectPageReq) ([]user2.UserListEntity, *page.Paging, error) {
	var deptService DeptService
	var d = deptService.SelectDeptById(param.DeptId)
	if d != nil { //数据权限
		param.Ancestors = d.Ancestors
	}
	return user2.SelectPageList(param)
}

// 导出excel
func (svc UserService) Export(param *user2.SelectPageReq) (string, error) {
	head := []string{"用户名", "呢称", "Email", "电话号码", "性别", "部门", "领导", "状态", "删除标记", "创建人", "创建时间", "备注"}
	col := []string{"u.login_name", "u.user_name", "u.email", "u.phonenumber", "u.sex", "d.dept_name", "d.leader", "u.status", "u.del_flag", "u.create_by", "u.create_time", "u.remark"}
	return user2.SelectExportList(param, head, col)
}

// 新增用户
func (svc UserService) AddSave(req *user2.AddReq, c *gin.Context) (int64, error) {
	var u user2.SysUser
	u.LoginName = req.LoginName
	u.UserName = req.UserName
	u.Email = req.Email
	u.Phonenumber = req.Phonenumber
	u.Status = req.Status
	u.Sex = req.Sex
	u.DeptId = req.DeptId
	u.Remark = req.Remark

	//生成密码
	newSalt := random.GenerateSubId(6)
	newToken := req.LoginName + req.Password + newSalt
	newToken = lib_secret.MustEncryptString(newToken)

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

	_, err = session.Table(user2.TableName()).Insert(&u)

	if err != nil || u.UserId <= 0 {
		session.Rollback()
		return 0, err
	}

	//增加岗位数据
	if req.PostIds != "" {
		postIds := convert.ToInt64Array(req.PostIds, ",")
		userPosts := make([]user_post.Entity, 0)
		for i := range postIds {
			if postIds[i] > 0 {
				var userPost user_post.Entity
				userPost.UserId = u.UserId
				userPost.PostId = postIds[i]
				userPosts = append(userPosts, userPost)
			}
		}
		if len(userPosts) > 0 {
			_, err := session.Table(user_post.TableName()).Insert(userPosts)
			if err != nil {
				session.Rollback()
				return 0, err
			}
		}

	}

	//增加角色数据
	if req.RoleIds != "" {
		roleIds := convert.ToInt64Array(req.RoleIds, ",")
		userRoles := make([]user_role.Entity, 0)
		for i := range roleIds {
			if roleIds[i] > 0 {
				var userRole user_role.Entity
				userRole.UserId = u.UserId
				userRole.RoleId = roleIds[i]
				userRoles = append(userRoles, userRole)
			}
		}
		if len(userRoles) > 0 {
			_, err := session.Table(user_role.TableName()).Insert(userRoles)
			if err != nil {
				session.Rollback()
				return 0, err
			}
		}
	}

	return u.UserId, session.Commit()
}

// 新增用户
func (svc UserService) EditSave(req *user2.EditReq, c *gin.Context) (int64, error) {
	u := &user2.SysUser{UserId: req.UserId}
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

	_, tanErr = session.Table(user2.TableName()).ID(u.UserId).Update(u)

	if tanErr != nil {
		session.Rollback()
		return 0, tanErr
	}

	//增加岗位数据
	if req.PostIds != "" {
		postIds := convert.ToInt64Array(req.PostIds, ",")
		userPosts := make([]user_post.Entity, 0)
		for i := range postIds {
			if postIds[i] > 0 {
				var userPost user_post.Entity
				userPost.UserId = u.UserId
				userPost.PostId = postIds[i]
				userPosts = append(userPosts, userPost)
			}
		}
		if len(userPosts) > 0 {
			session.Exec("delete from sys_user_post where user_id=?", u.UserId)
			_, tanErr = session.Table(user_post.TableName()).Insert(userPosts)
			if tanErr != nil {
				session.Rollback()
				return 0, err
			}
		}

	}

	//增加角色数据
	if req.RoleIds != "" {
		roleIds := convert.ToInt64Array(req.RoleIds, ",")
		userRoles := make([]user_role.Entity, 0)
		for i := range roleIds {
			if roleIds[i] > 0 {
				var userRole user_role.Entity
				userRole.UserId = u.UserId
				userRole.RoleId = roleIds[i]
				userRoles = append(userRoles, userRole)
			}
		}
		if len(userRoles) > 0 {
			session.Exec("delete from sys_user_role where user_id=?", u.UserId)
			_, err := session.Table(user_role.TableName()).Insert(userRoles)
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
	entity := &user2.SysUser{UserId: id}
	result, _ := entity.Delete()
	if result > 0 {
		return true
	}
	return false
}

// 批量删除用户记录
func (svc UserService) DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, _ := user2.DeleteBatch(idarr...)
	user_role.DeleteBatch(idarr...)
	user_post.DeleteBatch(idarr...)
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
func (svc UserService) SignIn(loginnName, password string) (*user2.SysUser, error) {
	//查询用户信息
	user := user2.SysUser{LoginName: loginnName}
	ok, err := user.FindOne()

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("用户名或者密码错误")
	}

	//校验密码
	pwdNew := user.LoginName + password + user.Salt

	pwdNew = lib_secret.MustEncryptString(pwdNew)

	if strings.Compare(user.Password, pwdNew) == -1 {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}

// 清空用户菜单缓存

// 用户注销
func (svc UserService) SignOut(c *gin.Context) error {
	//user := svc.GetProfile(c)
	tokenStr := lib_net.GetParam(c, "token")
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
	entity := user2.SysUser{LoginName: loginName}
	if ok, err := entity.FindOne(); err != nil || !ok {
		return false
	} else {
		return true
	}
}

// 检查登陆名是否存在,存在返回true,否则false
func (svc UserService) CheckNickName(userName string) bool {
	entity := user2.SysUser{UserName: userName}
	if ok, err := entity.FindOne(); err != nil || !ok {
		return false
	} else {
		return true
	}
}

// 检查登陆名是否存在,存在返回true,否则false
func (svc UserService) CheckLoginName(loginName string) bool {
	entity := user2.SysUser{LoginName: loginName}
	if ok, err := entity.FindOne(); err != nil || !ok {
		return false
	} else {
		return true
	}
}

// 获得用户信息详情
func (svc UserService) GetProfile(c *gin.Context) *user2.SysUser {

	var u = session.GetProfile(c)
	return u
}

// 更新用户信息详情
func (svc UserService) UpdateProfile(profile *user2.ProfileReq, c *gin.Context) error {
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
func (svc UserService) UpdatePassword(profile *user2.PasswordReq, c *gin.Context) error {
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
	token = lib_secret.MustEncryptString(token)

	if token != user.Password {
		return errors.New("原密码不正确")
	}

	//新校验密码
	newSalt := random.GenerateSubId(6)
	newToken := user.LoginName + profile.NewPassword + newSalt
	newToken = lib_secret.MustEncryptString(newToken)

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
func (svc UserService) ResetPassword(params *user2.ResetPwdReq) (bool, error) {
	user := user2.SysUser{UserId: params.UserId}
	if ok, err := user.FindOne(); err != nil || !ok {
		return false, errors.New("用户不存在")
	}
	//新校验密码
	newSalt := random.GenerateSubId(6)
	newToken := user.LoginName + params.Password + newSalt
	newToken = lib_secret.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken
	if _, err := user.Update(); err != nil {
		return false, errors.New("保存数据失败")
	}
	return true, nil
}

// 校验密码是否正确
func (svc UserService) CheckPassword(user *user2.SysUser, password string) bool {
	if user == nil || user.UserId <= 0 {
		return false
	}

	//校验密码
	token := user.LoginName + password + user.Salt
	token = lib_secret.MustEncryptString(token)

	if strings.Compare(token, user.Password) == 0 {
		return true
	} else {
		return false
	}
}

// 检查邮箱是否已使用
func (svc UserService) CheckEmailUnique(userId int64, email string) bool {
	return user2.CheckEmailUnique(userId, email)
}

// 检查邮箱是否存在,存在返回true,否则false
func (svc UserService) CheckEmailUniqueAll(email string) bool {
	return user2.CheckEmailUniqueAll(email)
}

// 检查手机号是否已使用,存在返回true,否则false
func (svc UserService) CheckPhoneUnique(userId int64, phone string) bool {
	return user2.CheckPhoneUnique(userId, phone)
}

// 检查手机号是否已使用 ,存在返回true,否则false
func (svc UserService) CheckPhoneUniqueAll(phone string) bool {
	return user2.CheckPhoneUniqueAll(phone)
}

// 根据登陆名查询用户信息
func (svc UserService) SelectUserByLoginName(loginName string) (*user2.SysUser, error) {
	return user2.SelectUserByLoginName(loginName)
}

// 根据手机号查询用户信息
func (svc UserService) SelectUserByPhoneNumber(phonenumber string) (*user2.SysUser, error) {
	return user2.SelectUserByPhoneNumber(phonenumber)
}

// 查询已分配用户角色列表
func (svc UserService) SelectAllocatedList(roleId int64, loginName, phonenumber string) ([]user2.SysUser, error) {
	return user2.SelectAllocatedList(roleId, loginName, phonenumber)
}

// 查询未分配用户角色列表
func (svc UserService) SelectUnallocatedList(roleId int64, loginName, phonenumber string) ([]user2.SysUser, error) {
	return user2.SelectUnallocatedList(roleId, loginName, phonenumber)
}
