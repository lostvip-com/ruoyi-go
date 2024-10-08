// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-01-03 21:50:54 +0800 CST
// 生成人：lv
// ==========================================================================
package dao

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/lv_batis"
	"github.com/lostvip-com/lv_framework/lv_db/lv_dao"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_reflect"
	"main/internal/mywork/model"
	"main/internal/mywork/vo"
)

// 新增页面请求参数
type DpcTaskDao struct{}

// 根据条件分页查询数据
func (e DpcTaskDao) ListByPage(req *vo.PageDpcTaskReq) (*[]model.DpcTask, int64, error) {
	ibatis := lv_batis.NewInstance("mywork/dpc_task_mapper.sql") //under the directory mapper
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	tagName := lv_reflect.GetMethodName()
	limitSQL, err := ibatis.GetLimitSql(tagName, req)
	//查询数据
	rows, err := lv_dao.ListByNamedSql[model.DpcTask](limitSQL, req)
	lv_err.HasErrAndPanic(err)
	count, err := lv_dao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_err.HasErrAndPanic(err)
	return rows, count, nil
}

// ListAll 导出excel使用
func (e DpcTaskDao) ListAll(req *vo.PageDpcTaskReq, isCamel bool) (*[]map[string]string, error) {
	ibatis := lv_batis.NewInstance("mywork/dpc_task_mapper.sql") //under the directory mapper
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	sql, err := ibatis.GetLimitSql(lv_reflect.GetMethodName(), req)
	lv_err.HasErrAndPanic(err)

	arr, err := lv_dao.ListMapStrByNamedSql(sql, req, isCamel)
	return arr, err
}

// FindByWhere 根据条件查询
func (e DpcTaskDao) Find(where, order string) (*[]model.DpcTask, error) {
	var list []model.DpcTask
	err := lv_db.GetMasterGorm().Table("dpc_task").Where(where).Order(order).Find(&list).Error
	return &list, err
}

func (e DpcTaskDao) DeleteByIds(ida []int64) int64 {
	db := lv_db.GetMasterGorm().Table("dpc_task").Where("id in ? ", ida).Update("del_flag", 1)
	if db.Error != nil {
		panic(db.Error)
	}
	return db.RowsAffected
}
