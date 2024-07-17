// ///////////////////////////////////////////////////////////////////////////
// 业务逻辑处理类的基类，简单的直接在model中处理即可，不需要service
//
// //////////////////////////////////////////////////////////////////////////
package namedsql

import (
	"database/sql"
	"github.com/lostvip-com/lv_framework/logme"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/utils/lv_sql"
	"gorm.io/gorm"
	"strings"
)

func Exec(db *gorm.DB, dmlSql string, req map[string]any) (int64, error) {
	if lv_global.Config().IsDebug() {
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

/**
 * 通用泛型查询
 */
func ListData[T any](db *gorm.DB, limitSql string, req any) (*[]T, error) {
	var list = make([]T, 0)
	var err error
	if lv_global.Config().IsDebug() {
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
	if lv_global.Config().IsDebug() {
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
		logme.Info(err)
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

/**
 *  注意：
 *  如果是map类型，只能是值类型，不能传地址 ！！！！！！！！！
 *  paramsValues 参数如果是struct类型，可以是指针也可以是值
 *
 */
func ListMap(db *gorm.DB, sqlQuery string, params any, isCamel bool) (*[]map[string]string, error) {
	var rows *sql.Rows
	var err error
	if lv_global.Config().IsDebug() {
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
		logme.Info(err)
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		logme.Info(err)
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
			logme.Info(err)
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
	if lv_global.Config().IsDebug() {
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
	if lv_global.Config().IsDebug() {
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
