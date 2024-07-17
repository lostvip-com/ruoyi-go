/*******************************************************************************
 * Copyright 2017 Dell Inc.
 * Copyright (c) 2019 Intel Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
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
