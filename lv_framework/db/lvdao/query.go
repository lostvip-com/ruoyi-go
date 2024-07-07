package lvdao

import (
	"github.com/lv_framework/db"
	"github.com/lv_framework/db/lvbatis"
	"github.com/lv_framework/db/namedsql"
	"github.com/lv_framework/utils/lv_err"
	"github.com/lv_framework/web/dto"
)

func ListMapByNamedSql(sql string, req any, isCamel bool) (*[]map[string]string, error) {
	d := db.GetMasterGorm()
	return namedsql.ListMap(d, sql, req, isCamel)
}

/**
 * 通用泛型查询
 */
func ListByNamedSql[T any](sql string, req any) (*[]T, error) {
	return namedsql.ListData[T](db.GetMasterGorm(), sql, req)
}

func CountByNamedSql(sql string, params any) (int64, error) {
	return namedsql.Count(db.GetMasterGorm(), sql, params)
}

/**
 * 通用泛型查询
 */
func GetPageByNamedSql[T any](sqlfile string, sqlTag string, req any) dto.RespPage {
	//解析sql
	ibatis := lvbatis.NewInstance(sqlfile)
	sql, err := ibatis.GetLimitSql(sqlTag, req)
	lv_err.HasErrAndPanic(err)
	//查询数据
	rows, err := namedsql.ListData[T](db.GetMasterGorm(), sql, req)
	lv_err.HasErrAndPanic(err)
	count, err := namedsql.Count(db.GetMasterGorm(), sql, req)
	lv_err.HasErrAndPanic(err)
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
