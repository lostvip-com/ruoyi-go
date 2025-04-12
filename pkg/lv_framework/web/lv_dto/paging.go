package lv_dto

import (
	"math"
)

const PageSize int = 15

type Paging struct {
	PageNum  int   `form:"pageNum"  json:"pageNum"`  //当前页
	PageSize int   `form:"pageSize" json:"pageSize"` //每页条数
	Total    int64 `form:"total"    json:"total"`    //每页条数//总条数
	//SortOrder string `form:"sortOrder"    json:"sortOrder"` //排序
	//SortName  string `form:"sortName"    json:"sortName"`   //每页条数//总条数
	// 注意这里的注解，前端传过来的是 isAsc
	SortOrder string `form:"isAsc"          json:"sortOrder"` //排序
	SortName  string `form:"orderByColumn"  json:"sortName"`  //每页条数//总条数

	PageCount int //总页数
	StartNum  int //起始行
}

// 创建分页
func (p *Paging) GetStartNum() int {
	if p.PageNum < 1 {
		p.PageNum = 1
	}
	if p.PageSize < 1 {
		p.PageSize = PageSize
	}
	p.StartNum = p.PageSize * (p.PageNum - 1)
	return p.StartNum
}

// 创建分页
func (p *Paging) GetPageSize() int {
	if p.PageSize < 1 {
		p.PageSize = PageSize
	}
	return p.PageSize
}

// 创建分页
func CreatePaging(pageNum, pagesize int, total int64) *Paging {
	if pageNum < 1 {
		pageNum = 1
	}
	if pagesize < 1 {
		pagesize = PageSize
	}
	paging := new(Paging)
	paging.PageNum = pageNum
	paging.PageSize = pagesize
	paging.Total = total
	paging.PageCount = int(math.Ceil(float64(total) / float64(pagesize)))
	paging.StartNum = pagesize * (pageNum - 1)
	return paging
}
