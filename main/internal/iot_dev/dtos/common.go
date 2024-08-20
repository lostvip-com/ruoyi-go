package dtos

import "strings"

type PageRequest struct {
	NameLike string `json:"nameLike" form:"nameLike"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

type BaseSearchConditionQuery struct {
	Page     int    `schema:"page,omitempty" form:"page"`
	PageSize int    `schema:"pageSize,omitempty" form:"pageSize" json:"pageSize"`
	Id       string `schema:"id,omitempty" form:"id"`
	Ids      string `schema:"ids,omitempty" form:"ids"`
	LikeId   string `schema:"likeId,omitempty" form:"likeId"`
	Name     string `schema:"name,omitempty" form:"name"`
	NameLike string `schema:"nameLike,omitempty" form:"nameLike"`
	IsAll    bool   `schema:"isAll,omitempty" form:"isAll"`
	OrderBy  string `schema:"orderBy,omitempty" form:"orderBy"`
}

func (req BaseSearchConditionQuery) GetPage() (int, int) {
	var (
		offset = (req.Page - 1) * req.PageSize
		limit  = req.PageSize
	)
	if req.Page == 0 && req.PageSize == 0 {
		offset = 0
		limit = -1
	}
	if req.IsAll {
		offset = 0
		limit = -1
	}
	return offset, limit
}

func ApiParamsStringToArray(str string) []string {
	return strings.Split(str, ",")
}

type ApiOrderBy struct {
	Key    string
	IsDesc bool
}

func ApiParamsStringToOrderBy(str string) []ApiOrderBy {
	orderBys := make([]ApiOrderBy, 0)
	arr := strings.Split(str, ",")
	if len(arr) <= 0 {
		return nil
	}
	for _, v := range arr {
		vArr := strings.Split(v, ":")
		if len(vArr) <= 1 {
			continue
		}
		switch vArr[1] {
		case "desc":
			orderBys = append(orderBys, ApiOrderBy{
				Key:    vArr[0],
				IsDesc: true,
			})
		case "asc":
			orderBys = append(orderBys, ApiOrderBy{
				Key:    vArr[0],
				IsDesc: false,
			})
		default:
			continue
		}
	}
	return orderBys
}

func ApiParamsArrayToString(arr []string) string {
	return strings.Join(arr, ",")
}
