package service

import (
	"errors"
	"github.com/lostvip-com/lv_framework/lv_cache"
	"github.com/lostvip-com/lv_framework/utils/lv_net"
	"github.com/lostvip-com/lv_framework/utils/lv_secret"
	"github.com/mssola/user_agent"
	"main/internal/system/model"
	"strings"
	"time"
)

type SessionService struct{}

func (svc *SessionService) IsSignedIn(tokenStr string) bool {
	key := "login:" + tokenStr
	num, err := lv_cache.GetCacheClient().Exists(key)
	return err == nil && num > 0
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func (svc *SessionService) SignIn(loginnName, password string) (*model.SysUser, error) {
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
func (svc *SessionService) SignOut(tokenStr string) error {
	lv_cache.GetCacheClient().Del("login:" + tokenStr)
	entity := model.SysUserOnline{SessionId: "login:" + tokenStr}
	entity.Delete()
	return nil
}

// 强退用户
func (svc *SessionService) ForceLogout(token string) error {
	svc.SignOut(token)
	entity := model.SysUserOnline{SessionId: token}
	entity.Delete()
	return nil
}

func (svc *SessionService) SaveUserToSession(token string, user *model.SysUser, roleKeys string) error {
	//记录到redis
	fieldMap := make(map[string]interface{})
	fieldMap["userName"] = user.UserName
	fieldMap["userId"] = user.UserId
	fieldMap["loginName"] = user.LoginName
	fieldMap["avatar"] = user.Avatar
	fieldMap["roleKeys"] = roleKeys
	fieldMap["deptId"] = user.DeptId
	fieldMap["tenantId"] = user.TenantId //租户ID
	//其它
	key := "login:" + token
	err := lv_cache.GetCacheClient().HSet(key, fieldMap)
	if err != nil {
		return errors.New("redis 故障！" + err.Error())
	}
	err = lv_cache.GetCacheClient().Expire(key, time.Hour)
	if err != nil {
		return errors.New("redis 故障！" + err.Error())
	}
	return err
}

func (svc *SessionService) SaveLoginLog2DB(token string, user *model.SysUser, userAgent string, ip string) error {
	//save to lv_db
	ua := user_agent.New(userAgent)
	os := ua.OS()
	browser, _ := ua.Browser()
	//移除登录次数记录
	var logininforService LoginInforService
	logininforService.RemovePasswordCounts(user.UserName)
	//
	var userOnline model.SysUserOnline
	// 保存用户信息到session
	loginLocation := lv_net.GetCityByIp(ip)
	userOnline.LoginName = user.UserName
	userOnline.Browser = browser
	userOnline.Os = os
	userOnline.DeptName = ""
	userOnline.Ipaddr = ip
	userOnline.ExpireTime = 1440
	userOnline.StartTimestamp = time.Now()
	userOnline.LastAccessTime = time.Now()
	userOnline.CreateTime = userOnline.StartTimestamp
	userOnline.Status = "on_line"
	userOnline.LoginLocation = loginLocation
	userOnline.SessionId = token
	err := userOnline.Delete()
	if err != nil {
		return err
	}
	err = userOnline.Save()
	if err != nil {
		return err
	}
	return err
}

func (svc *SessionService) Refresh(token string) {
	token = "login:" + token
	lv_cache.GetCacheClient().Expire(token, 8*time.Hour)
}
