// ==========================================================================
// LV自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-02-17 14:03:51
// 生成路径: app/service/online/online_service.go
// 生成人：yunjie
// ==========================================================================
package online

import (
	"lostvip.com/utils/lv_web"
	"robvi/app/sys/model/monitor/online"
	"strings"
)

// 根据主键查询数据
func SelectRecordById(id string) (*online.UserOnline, error) {
	entity := &online.UserOnline{Sessionid: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除数据
func DeleteRecordById(id string) bool {
	entity := &online.UserOnline{Sessionid: id}
	result, err := entity.Delete()
	if err == nil && result > 0 {
		return true
	}

	return false
}

// 批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	idarr := strings.Split(ids, ",")
	result, _ := online.DeleteBatch(idarr...)
	return result
}

// 批量删除数据
func DeleteRecordNotInIds(ids []string) int64 {
	result, _ := online.DeleteNotIn(ids...)
	return result
}

// 添加数据
func AddSave(entity online.UserOnline) (int64, error) {
	return entity.Insert()
}

// 根据条件查询数据
func SelectListAll(params *online.SelectPageReq) ([]online.UserOnline, error) {
	return online.SelectListAll(params)
}

// 根据条件分页查询数据
func SelectListByPage(params *online.SelectPageReq) ([]online.UserOnline, *lv_web.Paging, error) {
	return online.SelectListByPage(params)
}
