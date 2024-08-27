// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-02-28 14:21:50 +0800 CST
// 生成人：lv
// ==========================================================================
package dao

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/lv_batis"
	"github.com/lostvip-com/lv_framework/lv_db/lv_dao"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"main/internal/mywork/model"
	"main/internal/mywork/vo"
)

// 新增页面请求参数
type AppGenParamsDao struct{}

// 根据条件分页查询数据
func (d AppGenParamsDao) ListMapByPage(req *vo.PageAppGenParamsReq) (*[]map[string]string, int64, error) {
	ibatis := lv_batis.NewInstance("mywork/app_gen_params_mapper.sql") //under the mapper directory
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	limitSQL, err := ibatis.GetLimitSql("ListAppGenParams", req)
	//查询数据
	rows, err := lv_dao.ListMapStrByNamedSql(limitSQL, req, true)
	lv_err.HasErrAndPanic(err)
	count, err := lv_dao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_err.HasErrAndPanic(err)
	return rows, count, nil
}

// 根据条件分页查询数据
func (d AppGenParamsDao) ListByPage(req *vo.PageAppGenParamsReq) (*[]model.AppGenParams, int64, error) {
	ibatis := lv_batis.NewInstance("mywork/app_gen_params_mapper.sql") //under the mapper directory
	// 对应sql文件中的同名tagName
	limitSQL, err := ibatis.GetLimitSql("ListAppGenParams", req)
	//查询数据
	rows, err := lv_dao.ListByNamedSql[model.AppGenParams](limitSQL, req)
	lv_err.HasErrAndPanic(err)
	count, err := lv_dao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_err.HasErrAndPanic(err)
	return rows, count, nil
}

// ListAll 导出excel使用
func (d AppGenParamsDao) ListAll(req *vo.PageAppGenParamsReq, isCamel bool) (*[]map[string]string, error) {
	ibatis := lv_batis.NewInstance("mywork/app_gen_params_mapper.sql")
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	sql, err := ibatis.GetSql("ListAppGenParams", req)
	lv_err.HasErrAndPanic(err)

	arr, err := lv_dao.ListMapStrByNamedSql(sql, req, isCamel)
	return arr, err
}

// FindByWhere 根据条件查询
func (d AppGenParamsDao) Find(where, order string) (*[]model.AppGenParams, error) {
	var list []model.AppGenParams
	err := lv_db.GetMasterGorm().Table("app_gen_params").Where(where).Order(order).Find(&list).Error
	return &list, err
}

func (d AppGenParamsDao) DeleteByIds(ida []int64) int64 {
	db := lv_db.GetMasterGorm().Table("app_gen_params").Where("id in ? ", ida).Update("del_flag", 1)
	if db.Error != nil {
		panic(db.Error)
	}
	return db.RowsAffected
}
