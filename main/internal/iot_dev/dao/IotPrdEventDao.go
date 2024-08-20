// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-07-30 21:59:36 +0800 CST
// 生成人：lv
// ==========================================================================
package dao

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/lv_batis"
	"github.com/lostvip-com/lv_framework/lv_db/lv_dao"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"main/internal/iot_dev/model"
	"main/internal/iot_dev/vo"
)

// 新增页面请求参数
type IotPrdEventDao struct{}

// 根据条件分页查询数据
func (d IotPrdEventDao) ListMapByPage(req *vo.PageIotPrdEventReq) (*[]map[string]string, int64, error) {
	ibatis := lv_batis.NewInstance("iot_dev/iot_prd_event_mapper.sql") //under the mapper directory
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	limitSQL, err := ibatis.GetLimitSql("ListIotPrdEvent", req)
	//查询数据
	rows, err := lv_dao.ListMapByNamedSql(limitSQL, req, true)
	lv_err.HasErrAndPanic(err)
	count, err := lv_dao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_err.HasErrAndPanic(err)
	return rows, count, nil
}

// 根据条件分页查询数据
func (d IotPrdEventDao) ListByPage(req *vo.PageIotPrdEventReq) (*[]vo.IotPrdEventResp, int64, error) {
	ibatis := lv_batis.NewInstance("iot_dev/iot_prd_event_mapper.sql") //under the mapper directory
	// 对应sql文件中的同名tagName
	limitSQL, err := ibatis.GetLimitSql("ListIotPrdEvent", req)
	//查询数据
	rows, err := lv_dao.ListByNamedSql[vo.IotPrdEventResp](limitSQL, req)
	lv_err.HasErrAndPanic(err)
	count, err := lv_dao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_err.HasErrAndPanic(err)
	return rows, count, nil
}

// ListAll 导出excel使用
func (d IotPrdEventDao) ListAll(req *vo.PageIotPrdEventReq, isCamel bool) (*[]map[string]string, error) {
	ibatis := lv_batis.NewInstance("iot_dev/iot_prd_event_mapper.sql")
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	sql, err := ibatis.GetSql("ListIotPrdEvent", req)
	lv_err.HasErrAndPanic(err)

	arr, err := lv_dao.ListMapByNamedSql(sql, req, isCamel)
	return arr, err
}

// FindByWhere 根据条件查询
func (d IotPrdEventDao) Find(where, order string) (*[]model.IotPrdEvent, error) {
	var list []model.IotPrdEvent
	err := lv_db.GetMasterGorm().Table("iot_prd_event").Where(where).Order(order).Find(&list).Error
	return &list, err
}

func (d IotPrdEventDao) DeleteByIds(ida []int64) int64 {
	db := lv_db.GetMasterGorm().Table("iot_prd_event").Where("id in ? ", ida).Update("del_flag", 1)
	if db.Error != nil {
		panic(db.Error)
	}
	return db.RowsAffected
}
