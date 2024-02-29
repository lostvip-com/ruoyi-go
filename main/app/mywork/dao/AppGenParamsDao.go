// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-02-28 14:21:50 +0800 CST
// 生成人：lv
// ==========================================================================
package dao

import (
	"lostvip.com/db"
	"lostvip.com/db/lvbatis"
	"lostvip.com/db/lvdao"
	"lostvip.com/utils/lv_err"
	"main/app/mywork/model"
	"main/app/mywork/vo"
)

// 新增页面请求参数
type AppGenParamsDao struct{}

// 根据条件分页查询数据
func (d AppGenParamsDao) ListMapByPage(req *vo.PageAppGenParamsReq) (*[]map[string]string, int64, error) {
	ibatis := lvbatis.NewInstance("mywork/app_gen_params_mapper.tpl") //under the mapper directory
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	limitSQL, err := ibatis.GetLimitSql("ListAppGenParams", req)
	//查询数据
	rows, err := lvdao.ListMapByNamedSql(limitSQL, req, true)
	lv_err.HasErrAndPanic(err)
	count, err := lvdao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_err.HasErrAndPanic(err)
	return rows, count, nil
}

// 根据条件分页查询数据
func (d AppGenParamsDao) ListByPage(req *vo.PageAppGenParamsReq) (*[]model.AppGenParams, int64, error) {
	ibatis := lvbatis.NewInstance("mywork/app_gen_params_mapper.tpl") //under the mapper directory
	// 对应sql文件中的同名tagName
	limitSQL, err := ibatis.GetLimitSql("ListAppGenParams", req)
	//查询数据
	rows, err := lvdao.ListByNamedSql[model.AppGenParams](limitSQL, req)
	lv_err.HasErrAndPanic(err)
	count, err := lvdao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_err.HasErrAndPanic(err)
	return rows, count, nil
}

// ListAll 导出excel使用
func (d AppGenParamsDao) ListAll(req *vo.PageAppGenParamsReq, isCamel bool) (*[]map[string]string, error) {
	ibatis := lvbatis.NewInstance("mywork/app_gen_params_mapper.tpl")
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	sql, err := ibatis.GetSql("ListAppGenParams", req)
	lv_err.HasErrAndPanic(err)

	arr, err := lvdao.ListMapByNamedSql(sql, req, isCamel)
	return arr, err
}

// FindByWhere 根据条件查询
func (d AppGenParamsDao) Find(where, order string) (*[]model.AppGenParams, error) {
	var list []model.AppGenParams
	err := db.GetMasterGorm().Table("app_gen_params").Where(where).Order(order).Find(&list).Error
	return &list, err
}

func (d AppGenParamsDao) DeleteByIds(ida []int64) int64 {
	db := db.GetMasterGorm().Table("app_gen_params").Where("id in ? ", ida).Update("del_flag", 1)
	if db.Error != nil {
		panic(db.Error)
	}
	return db.RowsAffected
}
