package vo

import (
	"github.com/lostvip-com/lv_framework/web/lv_dto"
)

// 检查菜单名称请求参数
type CheckMenuNameReq struct {
	MenuId   int64  `form:"menuId"  binding:"required"`
	ParentId int64  `form:"parentId"  binding:"required"`
	MenuName string `form:"menuName"  binding:"required"`
}

// 检查菜单名称请求参数
type CheckMenuNameALLReq struct {
	ParentId int64  `form:"parentId"  binding:"required"`
	MenuName string `form:"menuName"  binding:"required"`
}

// 分页请求参数
type SelectMenuPageReq struct {
	MenuName  string `form:"menuName"`  //菜单名称
	Visible   string `form:"visible"`   //状态
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	PageNum   int    `form:"pageNum"`   //当前页码
	PageSize  int    `form:"pageSize"`  //每页数
	SortName  string `form:"sortName"`  //排序字段
	SortOrder string `form:"sortOrder"` //排序方式
	lv_dto.Paging
}

// 新增页面请求参数
type AddMenuReq struct {
	ParentId int64  `form:"parentId"`
	MenuType string `form:"menuType"  binding:"required"`
	MenuName string `form:"menuName"  binding:"required"`
	OrderNum int    `form:"orderNum" binding:"required"`
	Url      string `form:"url"`
	Icon     string `form:"icon"`
	Target   string `form:"target"`
	Perms    string `form:"perms""`
	Visible  string `form:"visible"`
}

// 修改页面请求参数
type EditMenuReq struct {
	MenuId   int64  `form:"menuId" binding:"required"`
	ParentId int64  `form:"parentId" `
	MenuType string `form:"menuType"  binding:"required"`
	MenuName string `form:"menuName"  binding:"required"`
	OrderNum int    `form:"orderNum" binding:"required"`
	Url      string `form:"url"`
	Icon     string `form:"icon"`
	Target   string `form:"target"`
	Perms    string `form:"perms""`
	Visible  string `form:"visible"`
}
