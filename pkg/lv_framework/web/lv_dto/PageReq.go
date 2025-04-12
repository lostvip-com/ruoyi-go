package lv_dto

type PageReq struct {
	PageIndex int `json:"pageIndex" form:"pageIndex" `
	PageSize  int `json:"pageSize"  form:"pageSize"  `
}

func (e *PageReq) GetPageIndex() int {
	return e.PageIndex
}

func (e *PageReq) GetPageSize() int {
	if e.PageSize == 0 {
		e.PageSize = 10
	}
	return e.PageSize
}
