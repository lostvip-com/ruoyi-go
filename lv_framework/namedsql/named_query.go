// ///////////////////////////////////////////////////////////////////////////
// 业务逻辑处理类的基类，简单的直接在model中处理即可，不需要service
//
// //////////////////////////////////////////////////////////////////////////
package namedsql

import (
	"database/sql"
	"gorm.io/gorm"
	"lostvip.com/logme"
	"lostvip.com/utils/lv_sql"
	"robvi/app/common/global"
	"strings"
)

/**
 * 通用泛型查询
 */
func ListData[T any](db *gorm.DB, limitSql string, req any) *[]T {
	var list = make([]T, 0)
	if global.GetConfigInstance().IsDebug() {
		db = db.Debug()
	}
	if strings.Contains(limitSql, "@") {
		db.Raw(limitSql, req).Scan(&list)
	} else {
		db.Raw(limitSql).Scan(&list)
	}

	return &list
}

func Count(db *gorm.DB, countSql string, sqlParams any) (int64, error) {
	if !strings.Contains(countSql, "count") {
		countSql = " select count(1) from (" + countSql + ") t where 1=1  "
	}
	if !strings.Contains(countSql, "limit") {
		countSql = countSql + "   limit 1  "
	}
	if global.GetConfigInstance().IsDebug() {
		db = db.Debug()
	}
	var rows *sql.Rows
	var err error
	if strings.Contains(countSql, "@") {
		rows, err = db.Raw(countSql, sqlParams).Rows()
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

// Raw sql查询返回[]map[string]string类型
func ListMap(db *gorm.DB, sqlQuery string, sqlParams any, isCamel bool) (*[]map[string]string, error) {

	var rows *sql.Rows
	var err error
	if strings.Contains(sqlQuery, "@") {
		rows, err = db.Raw(sqlQuery, sqlParams).Rows()
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

func ListArrStr(db *gorm.DB, sqlQuery string, sqlValues ...interface{}) (*[][]string, error) {

	if global.GetConfigInstance().IsDebug() {
		db = db.Debug()
	}
	var rows *sql.Rows
	var err error
	if strings.Contains(sqlQuery, "@") {
		rows, err = db.Raw(sqlQuery, sqlValues).Rows()
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

func CountWhereVal(db *gorm.DB, sql string, whereVal []interface{}) (int, error) {

	if global.GetConfigInstance().IsDebug() {
		db = db.Debug()
	}

	rows, err := db.Raw(sql, whereVal...).Limit(1).Rows()
	if err != nil {
		return 0, err
	}
	//查总数
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	return count, err
}
