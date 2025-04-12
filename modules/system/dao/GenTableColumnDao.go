package dao

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"system/model"
)

type GenTableColumnDao struct {
}

// 查询业务字段列表
func (dao GenTableColumnDao) SelectGenTableColumnListByTableId(tableId int64) ([]model.GenTableColumn, error) {
	db := lv_db.GetMasterGorm()
	var result []model.GenTableColumn
	model := db.Table("gen_table_column t").Where("table_id=?", tableId).Order(" sort desc ")
	err := model.Find(&result).Error
	return result, err
}

// SelectDbTableColumnsByName 根据表名称查询列信息
func (dao GenTableColumnDao) SelectDbTableColumnsByName(tableName string) ([]model.GenTableColumn, error) {
	db := lv_db.GetMasterGorm()
	var result []model.GenTableColumn
	model := db.Table("information_schema.columns")
	model.Select("column_name as 'column_name', (case when (is_nullable = 'no' && column_key != 'PRI') then '1' else null end) as is_required, (case when column_key = 'PRI' then '1' else '0' end) as is_pk, ordinal_position as sort, column_comment 'column_comment', (case when extra = 'auto_increment' then '1' else '0' end) as is_increment, column_type as 'column_type'")
	model.Where("table_schema = (select database())")
	model.Where("table_name=?", tableName).Order("ordinal_position desc ")
	model.Find(&result)
	return result, nil
}
