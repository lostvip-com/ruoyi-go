// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-12-24 16:29:05 +0800 CST
// 生成人：lv
// ==========================================================================
package dao

import (
	"lostvip.com/db"
	"lostvip.com/db/lvbatis"
	"lostvip.com/lvdao"
	"lostvip.com/utils/lv_logic"
	"lostvip.com/utils/lv_reflect"
	"robvi/app/mywork/model"
	"robvi/app/mywork/vo"
)

// 新增页面请求参数
type DpcTaskItemDao struct{}

// 根据条件分页查询数据
func (e DpcTaskItemDao) ListByPage(req *vo.PageDpcTaskItemReq) (*[]model.DpcTaskItem, int64, error) {
	ibatis := lvbatis.NewInstance("mywork/dpc_task_item_mapper.tpl")
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	tagName := lv_reflect.GetMethodName()
	limitSQL, err := ibatis.GetLimitSql(tagName, req)
	//查询数据
	rows, err := lvdao.ListByNamedSql[model.DpcTaskItem](limitSQL, req)
	lv_logic.HasErrAndPanic(err)
	count, err := lvdao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_logic.HasErrAndPanic(err)
	return rows, count, nil
}

// ListAll 导出excel使用
func (e DpcTaskItemDao) ListAll(req *vo.PageDpcTaskItemReq) (*[]map[string]string, error) {
	ibatis := lvbatis.NewInstance("mywork/dpc_task_item_mapper.tpl")
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	sql, err := ibatis.GetLimitSql(lv_reflect.GetMethodName(), req)
	lv_logic.HasErrAndPanic(err)

	arr, err := lvdao.ListMapByNamedSql(sql, req, false)
	return arr, err
}

// FindByWhere 根据条件查询
func (e DpcTaskItemDao) Find(where, order string) (*[]model.DpcTaskItem, error) {
	var list []model.DpcTaskItem
	err := db.GetMasterGorm().Table("dpc_task_item").Where(where).Order(order).Find(&list).Error
	return &list, err
}

func (e DpcTaskItemDao) DeleteByIds(ida []int64) int64 {
	var entity model.DpcTaskItem
	db := db.GetMasterGorm().Find(&entity, ida).Update("del_flag", 1)
	if db.Error != nil {
		panic(db.Error)
	}
	return db.RowsAffected
}
