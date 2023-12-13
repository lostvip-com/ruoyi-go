package lv_web

import (
	"math"
)

type Paging struct {
	PageNum   int //当前页
	Pagesize  int //每页条数
	Total     int //总条数
	PageCount int //总页数
	StartNum  int //起始行
}

// 创建分页
func CreatePaging(pageNum, pagesize, total int) *Paging {
	if pageNum < 1 {
		pageNum = 1
	}
	if pagesize < 1 {
		pagesize = 10
	}

	paging := new(Paging)
	paging.PageNum = pageNum
	paging.Pagesize = pagesize
	paging.Total = total
	paging.PageCount = int(math.Ceil(float64(total) / float64(pagesize)))
	paging.StartNum = pagesize * (pageNum - 1)
	return paging
}

// 创建分页
func CreatePage(pageNum, pagesize int) *Paging {
	if pageNum < 1 {
		pageNum = 1
	}
	if pagesize < 1 {
		pagesize = 10
	}
	paging := new(Paging)
	paging.PageNum = pageNum
	paging.Pagesize = pagesize
	paging.StartNum = pagesize * (pageNum - 1)
	return paging
}
