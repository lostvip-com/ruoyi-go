package db

import (
	"fmt"
	"github.com/spf13/cast"
	"lostvip.com/db/sqlx"
	"lostvip.com/utils/lv_sql"
	"lostvip.com/web/dto"
	"strings"
)

/**
 * 通用泛型查询
 */
func ListByNamedSql[T any](sql string, req any) []T {
	rows, err := GetMasterSqlX().NamedQuery(sql, req)
	if err != nil {
		panic(err)
	}
	var list = make([]T, 0)
	for rows.Next() {
		var obj T
		err := rows.StructScan(&obj)
		if err == nil {
			list = append(list, obj)
		}
	}
	return list
}

/**
 * 查询总记录
 */
func CountByNamedSql(sql string, req any) int64 {
	count, err := GetInstance().CountByNamedSql(sql, req)
	if err != nil {
		panic(err)
	}
	return count
}

/**
 * 通用泛型查询
 */
func GetPageByNamedSql[T any](sqlfile string, sqlTag string, req any) dto.RespPage {
	//解析sql
	ibatis := NewIBatis(sqlfile)
	sql := ibatis.GetLimitSql(sqlTag, req)
	//查询数据
	rows := ListByNamedSql[T](sql, req)
	count := CountByNamedSql(sql, req)
	return dto.SuccessPage[T](rows, count)
}

func (db *dbEngine) ListByNamedSql(sqlStr string, params any, isCamel bool) ([]map[string]any, error) {
	list := make([]map[string]any, 0)
	// 任何命名的占位符参数都将被arg中的字段替换。
	sqlx.NameMapper = strings.ToTitle
	rows, err := db.sqlxMaster.NamedQuery(sqlStr, params)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return list, err
	}
	defer rows.Close()
	for rows.Next() {
		tempMap := make(map[string]any)
		// 使用 map 做命名查询
		err := rows.MapScan(tempMap)
		//dest, err := rows.SliceScan()
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		newMap := make(map[string]any)
		// 将 "name" 字段的值转换为字符串类型
		for colKey, colVal := range tempMap {
			if isCamel {
				colKey = lv_sql.ToCamel(colKey)
			}
			newMap[colKey] = cast.ToString(colVal)
		}

		list = append(list, newMap)
	}
	return list, err
}

func (db *dbEngine) CountByNamedSql(sqlStr string, params any) (int64, error) {
	list, err := db.ListByNamedSql(sqlStr, params, false)
	if err == nil && list != nil && len(list) > 0 {
		rowMap := list[0]
		for _, value := range rowMap {
			return cast.ToInt64(value), err
		}
	}
	return 0, err
}
