package tool

import (
	"lostvip.com/utils/lv_conv"
	tool2 "robvi/app/modules/sys/model/tool"
)

type TableColumnService struct {
}

// 新增业务字段
func (svc TableColumnService) Insert(entity *tool2.Entity) (int64, error) {
	_, err := entity.Insert()
	if err != nil {
		return 0, err
	}
	return entity.ColumnId, err
}

// 修改业务字段
func (svc TableColumnService) Update(entity *tool2.Entity) (int64, error) {
	return entity.Update()
}

// 根据主键查询数据
func (svc TableColumnService) SelectRecordById(id int64) (*tool2.Entity, error) {
	entity := &tool2.Entity{ColumnId: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除数据
func (svc TableColumnService) DeleteRecordById(id int64) bool {
	rs, err := (&tool2.Entity{ColumnId: id}).Delete()
	if err == nil && rs > 0 {
		return true
	}
	return false
}

// 批量删除数据记录
func (svc TableColumnService) DeleteRecordByIds(ids string) int64 {
	idarr := lv_conv.ToInt64Array(ids, ",")
	result, err := tool2.DeleteBatch(idarr...)
	if err != nil {
		return 0
	}
	return result
}

// 查询业务字段列表
func (svc TableColumnService) SelectGenTableColumnListByTableId(tableId int64) ([]tool2.Entity, error) {
	return tool2.SelectGenTableColumnListByTableId(tableId)
}

// 根据表名称查询列信息
func (svc TableColumnService) SelectDbTableColumnsByName(tableName string) ([]tool2.Entity, error) {
	return tool2.SelectDbTableColumnsByName(tableName)
}
