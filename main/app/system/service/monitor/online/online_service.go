// ==========================================================================
// LV自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-02-17 14:03:51
// 生成人：yunjie
// ==========================================================================
package online

import (
	"github.com/lv_framework/utils/lv_web"
	"main/app/system/model/monitor/online"
)

// 批量删除数据
func DeleteRecordNotInIds(ids []string) int64 {
	result, _ := online.DeleteNotIn(ids...)
	return result
}

// 根据条件分页查询数据
func SelectListByPage(params *online.SelectPageReq) ([]online.UserOnline, *lv_web.Paging, error) {
	return online.SelectListByPage(params)
}
