// ==========================================================================
// LV自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-02-17 14:03:51
// 生成路径: app/service/modules/online/online_service.go
// 生成人：yunjie
// ==========================================================================
package online

import (
	"lostvip.com/utils/lv_web"
	online2 "robvi/app/modules/sys/model/monitor/online"
	"strings"
)

// 根据主键查询数据
func SelectRecordById(id string) (*online2.UserOnline, error) {
	entity := &online2.UserOnline{Sessionid: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除数据
func DeleteRecordById(id string) bool {
	entity := &online2.UserOnline{Sessionid: id}
	result, err := entity.Delete()
	if err == nil && result > 0 {
		return true
	}

	return false
}

// 批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	idarr := strings.Split(ids, ",")
	result, _ := online2.DeleteBatch(idarr...)
	return result
}

// 批量删除数据
func DeleteRecordNotInIds(ids []string) int64 {
	result, _ := online2.DeleteNotIn(ids...)
	return result
}

// 添加数据
func AddSave(entity online2.UserOnline) (int64, error) {
	return entity.Insert()
}

// 根据条件查询数据
func SelectListAll(params *online2.SelectPageReq) ([]online2.UserOnline, error) {
	return online2.SelectListAll(params)
}

// 根据条件分页查询数据
func SelectListByPage(params *online2.SelectPageReq) ([]online2.UserOnline, *lv_web.Paging, error) {
	return online2.SelectListByPage(params)
}
