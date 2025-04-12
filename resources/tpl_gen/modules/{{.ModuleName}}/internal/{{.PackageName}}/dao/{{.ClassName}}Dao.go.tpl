// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：{{.CreateTime}}
// 生成人：{{.FunctionAuthor}}
// ==========================================================================
package dao

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/lv_batis"
	"github.com/lostvip-com/lv_framework/lv_db/lv_dao"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
    "{{.ModuleName}}/internal/{{.PackageName}}/vo"
    "{{.ModuleName}}/internal/{{.PackageName}}/model"
)

type {{.ClassName}}Dao struct { }

// ListMapByPage 根据条件分页查询数据
func (d {{.ClassName}}Dao) ListMapByPage(req *vo.{{.ClassName}}Req) (*[]map[string]any, int64, error) {
	ibatis := lv_batis.NewInstance("{{.PackageName}}/{{.TbName}}_mapper.sql") //under the mapper directory
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	limitSQL, err := ibatis.GetLimitSql("List{{.ClassName}}", req)
	//查询数据
	rows, err := lv_dao.ListMapByNamedSql(limitSQL, req,true)
	lv_err.HasErrAndPanic(err)
	count, err := lv_dao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_err.HasErrAndPanic(err)
	return rows, count, nil
}

// ListByPage 根据条件分页查询数据
func (d {{.ClassName}}Dao) ListByPage(req *vo.{{.ClassName}}Req) (*[]vo.{{.ClassName}}Resp, int64, error) {
	ibatis := lv_batis.NewInstance("{{.PackageName}}/{{.TbName}}_mapper.sql") //under the mapper directory
	// 对应sql文件中的同名tagName
	limitSQL, err := ibatis.GetLimitSql("List{{.ClassName}}", req)
	//查询数据
	rows, err := lv_dao.ListByNamedSql[vo.{{.ClassName}}Resp](limitSQL, req)
	lv_err.HasErrAndPanic(err)
	count, err := lv_dao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_err.HasErrAndPanic(err)
	return rows, count, nil
}

// ListAll 导出excel使用
func (d {{.ClassName}}Dao) ListAll(req *vo.{{.ClassName}}Req, isCamel bool) (*[]map[string]any, error) {
	ibatis := lv_batis.NewInstance("{{.PackageName}}/{{.TbName}}_mapper.sql")
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	sql, err := ibatis.GetSql("List{{.ClassName}}", req)
	lv_err.HasErrAndPanic(err)

	arr, err := lv_dao.ListMapByNamedSql(sql, req, isCamel)
	return arr, err
}

// Find 根据条件查询
func (d {{.ClassName}}Dao) Find(where, order string) (*[]model.{{.ClassName}}, error) {
	var list []model.{{.ClassName}}
	err := lv_db.GetMasterGorm().Table("{{.TbName}}").Where(where).Order(order).Find(&list).Error
	return &list, err
}

// Find 通过主键批量删除
func (d {{.ClassName}}Dao) DeleteByIds(ida []int64) int64 {
	db := lv_db.GetMasterGorm().Table("{{.TbName}}").Where("{{.PkColumn.ColumnName}} in ? ", ida).Update("del_flag", 1)
    if db.Error != nil {
        panic(db.Error)
    }
    return db.RowsAffected
}
