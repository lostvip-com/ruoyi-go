package global

import (
	"sync"
)

//SESSION前缀
const (
	USER_ID           = "userId"
	USER_SESSION_MARK = "user_info"
)

// 登陆用户的菜单列表缓存前缀
const MENU_CACHE = "menu_cache"
//用户session列表
var SessionList sync.Map

// 判断用户是否已经登录
func GetSessions() *sync.Map {
	return &SessionList
}
