package router

import (
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	DELETE  = "DELETE"
	CONNECT = "CONNECT"
	TRACE   = "TRACE"
)

// 所有的路由
var GroupList = make([]*routerGroup, 0)

// 注册所有的URL标识
var PermissionMap = make(map[string]string, 0)

// 路由信息
type router struct {
	Method       string            //方法名称
	RelativePath string            //url路径
	Permiss      string            //权限字符串
	HandlerFunc  []gin.HandlerFunc //执行函数
}

// 路由组信息
type routerGroup struct {
	RelativePath string            //url路径
	Handlers     []gin.HandlerFunc //中间件
	Router       []*router         //路由信息
}

// 下级分组，继承上级分组的 路径和中间件
func (group *routerGroup) Group(relativePath string, middleware ...gin.HandlerFunc) *routerGroup {
	subGroup := New(group.RelativePath, group.Handlers...)
	subGroup.Handlers = append(subGroup.Handlers, middleware...)
	if strings.HasSuffix(subGroup.RelativePath, "/") {
		if strings.HasPrefix(relativePath, "/") {
			subGroup.RelativePath = subGroup.RelativePath + relativePath[1:]
		} else {
			subGroup.RelativePath = subGroup.RelativePath + relativePath
		}
	} else {
		if strings.HasPrefix(relativePath, "/") {
			subGroup.RelativePath = subGroup.RelativePath + relativePath
		} else {
			subGroup.RelativePath = subGroup.RelativePath + "/" + relativePath
		}
	}

	return subGroup
}

// 根据url获取权限字符串
func FindPermission(url string) string {
	return PermissionMap[url]
}

// 创建一个路由组
func New(relativePath string, middleware ...gin.HandlerFunc) *routerGroup {
	var rg routerGroup
	rg.Router = make([]*router, 0)
	rg.RelativePath = relativePath
	rg.Handlers = middleware
	GroupList = append(GroupList, &rg)
	return &rg
}

// 添加路由信息
func (group *routerGroup) Handle(method, relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	var r router
	r.Method = method
	r.Permiss = permiss
	r.RelativePath = relativePath
	r.HandlerFunc = handlers
	group.Router = append(group.Router, &r)
	if len(permiss) > 0 {
		if strings.EqualFold(relativePath, "/") {
			PermissionMap[group.RelativePath] = permiss
		} else {
			PermissionMap[group.RelativePath+relativePath] = permiss
		}
	}
	return group
}

// 添加路由信息-ANY
func (group *routerGroup) ANY(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle("ANY", relativePath, permiss, handlers...)
	return group
}

// 添加路由信息-GET
func (group *routerGroup) GET(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(GET, relativePath, permiss, handlers...)
	return group
}

// 添加路由信息-POST
func (group *routerGroup) POST(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(POST, relativePath, permiss, handlers...)
	return group
}

// 添加路由信息-OPTIONS
func (group *routerGroup) OPTIONS(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(OPTIONS, relativePath, permiss, handlers...)
	return group
}

// 添加路由信息-PUT
func (group *routerGroup) PUT(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(PUT, relativePath, permiss, handlers...)
	return group
}

// 添加路由信息-PATCH
func (group *routerGroup) PATCH(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(PATCH, relativePath, permiss, handlers...)
	return group
}

// 添加路由信息-HEAD
func (group *routerGroup) HEAD(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(HEAD, relativePath, permiss, handlers...)
	return group
}

// 添加路由信息-DELETE
func (group *routerGroup) DELETE(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(DELETE, relativePath, permiss, handlers...)
	return group
}

// 添加路由信息-CONNECT
func (group *routerGroup) CONNECT(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(CONNECT, relativePath, permiss, handlers...)
	return group
}

// 添加路由信息-TRACE
func (group *routerGroup) TRACE(relativePath, permiss string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(TRACE, relativePath, permiss, handlers...)
	return group
}
