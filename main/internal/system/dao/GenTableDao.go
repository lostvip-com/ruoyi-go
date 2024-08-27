package dao

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"main/internal/system/model"
	"main/internal/system/vo"
)

type GenTableDao struct {
}

// 根据ID获取记录
func (r *GenTableDao) ListColumn(tableId int64) (*vo.GenTableVO, error) {
	db := lv_db.GetMasterGorm()
	var result vo.GenTableVO
	tb := db.Table("gen_table").Where("table_id=?", tableId)
	err := tb.Find(&result).Error
	if err != nil {
		return nil, err
	}
	//表数据列
	columModel := db.Table("gen_table_column").Where("table_id=?", tableId)
	var columList []model.GenTableColumn
	err = columModel.Find(&columList).Error

	if err != nil {
		return nil, err
	}
	result.Columns = columList
	return &result, nil
}

// 根据条件分页查询数据
func (r *GenTableDao) SelectListByPage(param *vo.GenTablePageReq) ([]model.GenTable, int64, error) {
	db := lv_db.GetMasterGorm()
	tb := db.Table(" gen_table t")

	if param != nil {
		if param.TableName != "" {
			tb.Where("t.table_name like ?", "%"+param.TableName+"%")
		}

		if param.TableComment != "" {
			tb.Where("t.table_comment like ?", "%"+param.TableComment+"%")
		}

		if param.BeginTime != "" {
			tb.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			tb.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	var total int64
	tb = tb.Count(&total)
	tb.Limit(param.GetPageSize()).Offset(param.GetStartNum()).Order(param.SortName + " " + param.SortOrder)
	var result []model.GenTable
	err := tb.Find(&result).Error

	return result, total, err
}

// 查询据库列表
func (r *GenTableDao) SelectDbTableList(req *vo.GenTablePageReq) ([]model.GenTable, int64, error) {
	db := lv_db.GetMasterGorm()
	tb := db.Table("information_schema.tables")
	tb.Where("table_schema = (select database())")
	tb.Where("table_name NOT LIKE 'qrtz_%' AND table_name NOT LIKE 'gen_%' ")
	tb.Where("table_name NOT IN (select table_name from gen_table)")
	if req != nil {
		if req.TableName != "" {
			tb.Where("lower(table_name) like lower(?)", "%"+req.TableName+"%")
		}

		if req.TableComment != "" {
			tb.Where("lower(table_comment) like lower(?)", "%"+req.TableComment+"%")
		}

		if req.BeginTime != "" {
			tb.Where("date_format(create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", req.BeginTime)
		}

		if req.EndTime != "" {
			tb.Where("date_format(create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", req.EndTime)
		}
	}
	var total int64
	tb.Count(&total)
	tb.Select("table_name TbName, table_comment TableComment, create_time CreateTime,update_time UpdateTime")
	tb.Limit(req.GetPageSize()).Offset(req.GetStartNum())
	var result []model.GenTable
	err := tb.Find(&result).Error
	return result, total, err
}

// 查询据库列表
func (r *GenTableDao) SelectDbTableListByNames(tableNames []string) ([]model.GenTable, error) {
	db := lv_db.GetMasterGorm()
	tb := db.Table("information_schema.tables")
	tb.Select("0 as TableId, table_name as TbName,ifnull(table_comment,table_name) as TableComment")
	tb.Where("table_name NOT LIKE 'qrtz_%'")
	tb.Where("table_name NOT LIKE 'gen_%'")
	tb.Where("table_schema = (select database())")
	if len(tableNames) > 0 {
		tb.Where("table_name in ? ", tableNames)
	}

	var result []model.GenTable
	err := tb.Find(&result).Error
	return result, err
}

// 查询据库列表
func (r *GenTableDao) SelectTableByName(tableName string) (*model.GenTable, error) {
	db := lv_db.GetMasterGorm()
	tb := db.Table("information_schema.tables")
	tb.Select("0 as TableId, table_name TbName, table_comment TableComment")
	tb.Where("table_comment <> ''")
	tb.Where("table_schema = (select database())")
	if tableName != "" {
		tb.Where("table_name = ?", tableName)
	}

	var result model.GenTable
	err := tb.First(&result).Error
	return &result, err
}
