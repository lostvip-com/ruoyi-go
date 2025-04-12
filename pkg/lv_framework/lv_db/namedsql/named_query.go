// ///////////////////////////////////////////////////////////////////////////
// 业务逻辑处理类的基类，简单的直接在model中处理即可，不需要service
//
// //////////////////////////////////////////////////////////////////////////
package namedsql

import (
	"database/sql"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_sql"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"strings"
)

func Exec(db *gorm.DB, dmlSql string, req map[string]any) (int64, error) {
	if lv_global.IsDebug {
		db = db.Debug()
	}
	if strings.Contains(dmlSql, "@") {
		kvMap, isMap := checkAndExtractMap(req)
		if isMap {
			req = kvMap
		}
		tx := db.Exec(dmlSql, req)
		return tx.RowsAffected, tx.Error
	} else {
		tx := db.Exec(dmlSql)
		return tx.RowsAffected, tx.Error
	}
}

func GetOneMapByNamedSql(db *gorm.DB, limitSql string, req any, isCamel bool) (result *map[string]any, err error) {
	list, err := ListMapAny(db, limitSql, req, isCamel)
	var mpList []map[string]any = *list
	if err == nil {
		if mpList == nil || len(mpList) == 0 {
			err = gorm.ErrRecordNotFound
		} else {
			result = &mpList[0]
		}
	}
	return result, err
}

func toCamelMap(result *map[string]any) *map[string]any {
	mp := make(map[string]any)
	for k, v := range *result {
		mp[cast.ToString(lv_sql.ToCamel(k))] = v
	}
	return &mp
}

/**
 * 通用泛型查询
 */
func ListData[T any](db *gorm.DB, limitSql string, req any) (*[]T, error) {
	var list = make([]T, 0)
	var err error
	if lv_global.IsDebug {
		db = db.Debug()
	}
	if strings.Contains(limitSql, "@") {
		kvMap, isMap := checkAndExtractMap(req)
		if isMap {
			req = kvMap
		}
		err = db.Raw(limitSql, req).Scan(&list).Error
	} else {
		err = db.Raw(limitSql).Scan(&list).Error
	}

	return &list, err
}

func Count(db *gorm.DB, countSql string, params any) (int64, error) {
	if lv_global.IsDebug {
		db = db.Debug()
	}

	if !strings.Contains(countSql, "count") {
		countSql = " select count(*) from (" + countSql + ") t where 1=1  "
	}
	if !strings.Contains(countSql, "limit") {
		countSql = countSql + "   limit 1  "
	}

	var rows *sql.Rows
	var err error
	if strings.Contains(countSql, "@") {
		kvMap, isMap := checkAndExtractMap(params)
		if isMap {
			params = kvMap
		}
		rows, err = db.Raw(countSql, params).Rows()
	} else {
		rows, err = db.Raw(countSql).Rows()
	}
	if err != nil {
		lv_log.Info(err)
		return 0, err
	}
	//查总数
	var count int64
	if rows != nil {
		for rows.Next() {
			rows.Scan(&count)
		}
	}
	return count, err
}

/**
 * gorm中参数为map指针时，无法正常传参数！！
 * 处理方式：把map的指针转为值类型。
 */
func checkAndExtractMap(value interface{}) (map[string]any, bool) {
	// 判断是否是指针类型
	if ptr, ok := value.(*map[string]any); ok {
		// 指针指向Map类型
		return *ptr, true
	}
	return nil, false
}

// ListMapAny params可是是strtuc指针或map,isCamel key是否按驼峰式命名
func ListMapAny(db *gorm.DB, sqlQuery string, params any, isCamel bool) (*[]map[string]any, error) {
	var rows *sql.Rows
	var err error
	if lv_global.IsDebug {
		db = db.Debug()
	}
	if strings.Contains(sqlQuery, "@") {
		kvMap, isMap := checkAndExtractMap(params)
		if isMap {
			params = kvMap
		}
		rows, err = db.Raw(sqlQuery, params).Rows()
	} else {
		rows, err = db.Raw(sqlQuery).Rows()
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	// 获取列的类型信息
	types, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	// 创建一个切片来存储每行的数据
	var results []map[string]any

	// 遍历每一行
	for rows.Next() {
		rowData := make(map[string]any)
		values := make([]interface{}, len(cols))
		valuePtrs := make([]interface{}, len(cols))
		// 为每列创建一个 interface{} 类型的指针
		for i := range values {
			valuePtrs[i] = &values[i]
		}
		// 扫描当前行的数据
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}
		// 根据列的类型将值转换为更具体的 Go 类型
		var key string
		for i, val := range values {
			if isCamel {
				key = lv_sql.ToCamel(cols[i])
			}
			if val == nil {
				rowData[key] = nil
				continue
			}
			colType := types[i].DatabaseTypeName()
			CastValueType(colType, rowData, key, val)
		}
		// 将处理好的行数据添加到结果切片中
		results = append(results, rowData)
	}

	return &results, err
}

// CastValueType 根据列类型转换值类型
func CastValueType(colType string, rowData map[string]any, key string, val interface{}) {
	switch colType {
	case "VARCHAR", "TEXT":
		rowData[key] = cast.ToString(val)
	case "INT", "INTEGER":
		rowData[key] = cast.ToInt(val)
	case "BIGINT":
		rowData[key] = cast.ToInt64(val)
	case "FLOAT", "DOUBLE":
		rowData[key] = cast.ToFloat64(val)
	case "DATETIME":
		rowData[key] = cast.ToString(val)[:19]
	case "DATE":
		rowData[key] = cast.ToString(val)[:10]
	default:
		// 其他类型，直接存储 interface{}
		rowData[key] = cast.ToString(val)
	}
}

// ListMap sql查询返回map isCamel key是否按驼峰式命名
// Deprecated: 不再使用
func ListMap(db *gorm.DB, sqlQuery string, params any, isCamel bool) (*[]map[string]string, error) {
	return ListMapStr(db, sqlQuery, params, isCamel)
}

// ListMapStr 所有数据转为字符串格式返回，params可是是strtuc指针或map,isCamel key是否按驼峰式命名
func ListMapStr(db *gorm.DB, sqlQuery string, params any, isCamel bool) (*[]map[string]string, error) {
	var rows *sql.Rows
	var err error
	if lv_global.IsDebug {
		db = db.Debug()
	}
	if strings.Contains(sqlQuery, "@") {
		kvMap, isMap := checkAndExtractMap(params)
		if isMap {
			params = kvMap
		}
		rows, err = db.Raw(sqlQuery, params).Rows()
	} else {
		rows, err = db.Raw(sqlQuery).Rows()
	}
	if err != nil {
		lv_log.Info(err)
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		lv_log.Info(err)
		return nil, err
	}
	result := make([]map[string]string, 0)
	values := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			lv_log.Info(err)
			return nil, err
		}
		var value string
		resultC := map[string]string{}
		for i, col := range values {
			if col == nil {
				value = ""
			} else {
				value = string(col)
			}
			colKey := cols[i]
			if isCamel {
				colKey = lv_sql.ToCamel(colKey)
			}
			resultC[colKey] = value
		}
		result = append(result, resultC)
	}
	return &result, err
}

func ListArrStr(db *gorm.DB, sqlQuery string, params any) (*[][]string, error) {
	if lv_global.IsDebug {
		db = db.Debug()
	}
	var rows *sql.Rows
	var err error
	if strings.Contains(sqlQuery, "@") {
		kvMap, isMap := checkAndExtractMap(params)
		if isMap {
			params = kvMap
		}
		rows, err = db.Raw(sqlQuery, params).Rows()
	} else {
		rows, err = db.Raw(sqlQuery).Rows()
	}
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	listRows := make([][]string, 0)
	for rows.Next() {
		row := make([]string, 0)
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		var value string
		for _, col := range values {
			if col == nil {
				value = ""
			} else {
				value = string(col)
			}
			row = append(row, value)
		}
		listRows = append(listRows, row)
	}
	return &listRows, err
}

// ListOneColStr 查询某一列，放到数组中
func ListOneColStr(db *gorm.DB, sqlQuery string, params any) ([]string, error) {
	if lv_global.IsDebug {
		db = db.Debug()
	}
	var rows *sql.Rows
	var err error
	if strings.Contains(sqlQuery, "@") {
		kvMap, isMap := checkAndExtractMap(params)
		if isMap {
			params = kvMap
		}
		rows, err = db.Raw(sqlQuery, params).Rows()
	} else {
		rows, err = db.Raw(sqlQuery).Rows()
	}
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	arr := make([]string, 0)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		for _, col := range values {
			if col != nil {
				arr = append(arr, string(col))
			}
		}
	}
	return arr, err
}
