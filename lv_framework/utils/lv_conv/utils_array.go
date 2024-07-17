// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package lv_conv

import "reflect"

// IsArray checks whether given value is array/slice.
// Note that it uses reflect internally implementing this feature.
func IsArray(value interface{}) bool {
	rv := reflect.ValueOf(value)
	kind := rv.Kind()
	if kind == reflect.Ptr {
		rv = rv.Elem()
		kind = rv.Kind()
	}
	switch kind {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}

func RemoveOne(nums []int64, val int64) []int64 {
	newNums := make([]int64, 0)
	for _, num := range nums {
		if num != val {
			newNums = append(newNums, num)
		}
	}
	return newNums
}
func Remove(slice []any, element any) []any {
	var newSlice []any
	for _, v := range slice {
		if v != element {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
