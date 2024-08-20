package lv_dao

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/lv_batis"
	"github.com/lostvip-com/lv_framework/lv_db/namedsql"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
)

func CountCol(table, column, value string) (int64, error) {
	var total int64
	err := lv_db.GetMasterGorm().Table(table).Where("del_flag=0 and "+column+"=?", value).Count(&total).Error
	return total, err
}

func ListMapByNamedSql(sql string, req any, isCamel bool) (*[]map[string]string, error) {
	d := lv_db.GetMasterGorm()
	return namedsql.ListMap(d, sql, req, isCamel)
}

/**
 * 通用泛型查询
 */
func ListByNamedSql[T any](sql string, req any) (*[]T, error) {
	return namedsql.ListData[T](lv_db.GetMasterGorm(), sql, req)
}

func CountByNamedSql(sql string, params any) (int64, error) {
	return namedsql.Count(lv_db.GetMasterGorm(), sql, params)
}

/**
 * 通用泛型查询
 */
func GetPageByNamedSql[T any](sqlfile string, sqlTag string, req any) lv_dto.RespPage {
	//解析sql
	ibatis := lv_batis.NewInstance(sqlfile)
	sql, err := ibatis.GetLimitSql(sqlTag, req)
	lv_err.HasErrAndPanic(err)
	//查询数据
	rows, err := namedsql.ListData[T](lv_db.GetMasterGorm(), sql, req)
	lv_err.HasErrAndPanic(err)
	count, err := namedsql.Count(lv_db.GetMasterGorm(), sql, req)
	lv_err.HasErrAndPanic(err)
	return lv_dto.SuccessPage[T](rows, count)
}

func DeleteByIds(tableName string, ids []int64) error {
	delSql := "delete from " + tableName + " where id in (?)"
	err := lv_db.GetMasterGorm().Exec(delSql, ids).Error
	return err
}

func DeleteSoftByIds(tableName string, ids []int64) error {
	delSql := "update " + tableName + "set del_flag=? where id in (?)"
	err := lv_db.GetMasterGorm().Exec(delSql, lv_global.FLAG_DEL_YES, ids).Error
	return err
}
