package tool

import (
	"errors"
	"lostvip.com/db"
)

// 数据库字符串类型
var COLUMNTYPE_STR = []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext"}

// 数据库时间类型
var COLUMNTYPE_TIME = []string{"datetime", "time", "date", "timestamp"}

// 数据库数字类型
var COLUMNTYPE_NUMBER = []string{"tinyint", "smallint", "mediumint", "int", "number", "integer", "bigint", "float", "float", "double", "decimal"}

// 页面不需要编辑字段
var COLUMNNAME_NOT_EDIT = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time"}

// 页面不需要显示的列表字段
var COLUMNNAME_NOT_LIST = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time"}

// 页面不需要查询字段
var COLUMNNAME_NOT_QUERY = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time", "remark"}

// 查询业务字段列表
func SelectGenTableColumnListByTableId(tableId int64) ([]Entity, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	var result []Entity

	model := db.Table(TableName()).Alias("t").Where("table_id=?", tableId).OrderBy("sort")
	model.Find(&result)
	return result, nil
}

// 根据表名称查询列信息
func SelectDbTableColumnsByName(tableName string) ([]Entity, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	var result []Entity

	model := db.Table("information_schema.columns")
	model.Where("table_schema = (select database())")
	model.Where("table_name=?", tableName).OrderBy("ordinal_position")
	model.Select("column_name, (case when (is_nullable = 'no' && column_key != 'PRI') then '1' else null end) as is_required, (case when column_key = 'PRI' then '1' else '0' end) as is_pk, ordinal_position as sort, column_comment, (case when extra = 'auto_increment' then '1' else '0' end) as is_increment, column_type")
	model.Find(&result)
	return result, nil
}

// 判断string 是否存在在数组中
func IsExistInArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// 判断是否是数据库字符串类型
func IsStringObject(value string) bool {
	return IsExistInArray(value, COLUMNTYPE_STR)
}

// 判断是否是数据库时间类型
func IsTimeObject(value string) bool {
	return IsExistInArray(value, COLUMNTYPE_TIME)
}

// 判断是否是数据库数字类型
func IsNumberObject(value string) bool {
	return IsExistInArray(value, COLUMNTYPE_NUMBER)
}

// 页面不需要编辑字段
func IsNotEdit(value string) bool {
	return !IsExistInArray(value, COLUMNNAME_NOT_EDIT)
}

// 页面不需要显示的列表字段
func IsNotList(value string) bool {
	return !IsExistInArray(value, COLUMNNAME_NOT_LIST)
}

// 页面不需要查询字段
func IsNotQuery(value string) bool {
	return IsExistInArray(value, COLUMNNAME_NOT_QUERY)
}
