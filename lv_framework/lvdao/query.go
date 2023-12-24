package lvdao

import (
	"lostvip.com/db"
	"lostvip.com/db/lvbatis"
	"lostvip.com/namedsql"
	"lostvip.com/utils/lv_logic"
	"lostvip.com/web/dto"
	"robvi/app/common/global"
	"strings"
)

func ListMapByNamedSql(sql string, req any, isCamel bool) (*[]map[string]string, error) {
	d := db.GetMasterGorm()
	return namedsql.ListMap(d, sql, req, isCamel)
}

/**
 * 通用泛型查询
 */
func ListByNamedSql[T any](sql string, req any) (*[]T, error) {
	var list = make([]T, 0)
	db := db.GetMasterGorm()
	if global.GetConfigInstance().IsDebug() {
		db = db.Debug()
	}
	var err error
	if strings.Contains(sql, "@") {
		err = db.Raw(sql, req).Scan(&list).Error
	} else {
		err = db.Raw(sql).Scan(&list).Error
	}

	return &list, err
}

func CountByNamedSql(sqlStr string, params any) (int64, error) {
	db := db.GetMasterGorm()
	if global.GetConfigInstance().IsDebug() {
		db = db.Debug()
	}
	count, err := namedsql.Count(db, sqlStr, params)
	return count, err
}

/**
 * 通用泛型查询
 */
func GetPageByNamedSql[T any](sqlfile string, sqlTag string, req any) dto.RespPage {
	//解析sql
	ibatis := lvbatis.NewInstance(sqlfile)
	sql, err := ibatis.GetLimitSql(sqlTag, req)
	lv_logic.HasErrAndPanic(err)
	//查询数据
	rows, err := ListByNamedSql[T](sql, req)
	lv_logic.HasErrAndPanic(err)
	count, err := CountByNamedSql(sql, req)
	lv_logic.HasErrAndPanic(err)
	return dto.SuccessPage[T](rows, count)
}

//
//func ListMapByNamedSql(sqlStr string, params any, isCamel bool) []map[string]any {
//	list := make([]map[string]any, 0)
//	// 任何命名的占位符参数都将被arg中的字段替换。
//	rows, err := GetMasterGorm().(sqlStr, params)
//	defer rows.Close()
//	lv_logic.HasErrAndPanic(err)
//	for rows.Next() {
//		tempMap := make(map[string]any)
//		// 使用 map 做命名查询
//		err := rows.MapScan(tempMap)
//		//dest, err := rows.SliceScan()
//		if err != nil {
//			fmt.Printf("scan failed, err:%v\n", err)
//			continue
//		}
//		newMap := make(map[string]any)
//		// 将 "name" 字段的值转换为字符串类型
//		for colKey, colVal := range tempMap {
//			if isCamel {
//				colKey = lv_sql.ToCamel(colKey)
//			}
//			newMap[colKey] = cast.ToString(colVal)
//		}
//		list = append(list, newMap)
//	}
//	return list
//}
