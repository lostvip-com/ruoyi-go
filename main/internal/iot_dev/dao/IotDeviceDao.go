// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-07-19 17:09:35 +0800 CST
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
type IotDeviceDao struct{}

// 根据条件分页查询数据
func (d IotDeviceDao) ListMapByPage(req *vo.PageIotDeviceReq) (*[]map[string]string, int64, error) {
	ibatis := lv_batis.NewInstance("iot_dev/iot_device_mapper.sql") //under the mapper directory
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	limitSQL, err := ibatis.GetLimitSql("ListIotDevice", req)
	//查询数据
	rows, err := lv_dao.ListMapByNamedSql(limitSQL, req, true)
	lv_err.HasErrAndPanic(err)
	count, err := lv_dao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_err.HasErrAndPanic(err)
	return rows, count, nil
}

// 根据条件分页查询数据
func (d IotDeviceDao) ListByPage(req *vo.PageIotDeviceReq) (*[]vo.IotDeviceResp, int64, error) {
	ibatis := lv_batis.NewInstance("iot_dev/iot_device_mapper.sql") //under the mapper directory
	// 对应sql文件中的同名tagName
	limitSQL, err := ibatis.GetLimitSql("ListIotDevice", req)
	//查询数据
	rows, err := lv_dao.ListByNamedSql[vo.IotDeviceResp](limitSQL, req)
	lv_err.HasErrAndPanic(err)
	count, err := lv_dao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_err.HasErrAndPanic(err)
	return rows, count, nil
}

// ListAll 导出excel使用
func (d IotDeviceDao) ListAll(req *vo.PageIotDeviceReq, isCamel bool) (*[]map[string]string, error) {
	ibatis := lv_batis.NewInstance("iot_dev/iot_device_mapper.sql")
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	sql, err := ibatis.GetSql("ListIotDevice", req)
	lv_err.HasErrAndPanic(err)

	arr, err := lv_dao.ListMapByNamedSql(sql, req, isCamel)
	return arr, err
}

// FindByWhere 根据条件查询
func (d IotDeviceDao) Find(where, order string) (*[]model.IotDevice, error) {
	var list []model.IotDevice
	err := lv_db.GetMasterGorm().Table("iot_device").Where(where).Order(order).Find(&list).Error
	return &list, err
}

func (d IotDeviceDao) DeleteByIds(ida []int64) int64 {
	db := lv_db.GetMasterGorm().Table("iot_device").Where("id in ? ", ida).Update("del_flag", 1)
	if db.Error != nil {
		panic(db.Error)
	}
	return db.RowsAffected
}
