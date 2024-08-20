// ==========================================================================
// LV自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-07-19 17:09:35 +0800 CST
// 生成人：lv
// ==========================================================================
package service

import (
	model2 "common/model"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_office"
	"github.com/lostvip-com/lv_framework/utils/lv_reflect"
	"main/internal/common/service"
	"main/internal/iot_dev/dao"
	"main/internal/iot_dev/model"
	"main/internal/iot_dev/vo"
	"time"
)

type IotDeviceService struct{}

// FindById 根据主键查询数据
func (svc IotDeviceService) FindById(id int64) (*model.IotDevice, error) {
	entity := &model.IotDevice{Id: id}
	entity, err := entity.FindById(id)
	return entity, err
}

// DeleteById 根据主键删除数据
func (svc IotDeviceService) DeleteById(id int64) error {
	err := (&model.IotDevice{Id: id}).Delete()
	return err
}

// DeleteByIds 批量删除数据记录
func (svc IotDeviceService) DeleteByIds(ids string) int64 {
	ida := lv_conv.ToInt64Array(ids, ",")
	var d dao.IotDeviceDao
	rowsAffected := d.DeleteByIds(ida)
	return rowsAffected
}

// AddSave 添加数据
func (svc IotDeviceService) AddSave(req *vo.AddIotDeviceReq) (int64, error) {
	var entity = new(model.IotDevice)
	lv_reflect.CopyProperties(req, entity)
	entity.CreateTime = time.Now()
	entity.UpdateTime = entity.CreateTime
	entity.CreateBy = req.CreateBy
	err := entity.Save()
	lv_err.HasErrAndPanic(err)
	return entity.Id, err
}

// EditSave 修改数据
func (svc IotDeviceService) EditSave(req *vo.EditIotDeviceReq) error {
	entity := &model.IotDevice{Id: req.Id}
	entity, err := entity.FindById(req.Id)
	lv_err.HasErrAndPanic(err)
	lv_reflect.CopyProperties(req, entity)
	entity.UpdateTime = time.Now()
	entity.UpdateBy = req.UpdateBy
	return entity.Updates()
}

func (svc IotDeviceService) BindGateway(bind string, gwId int64, deviceId int64) error {
	entity := &model.IotDevice{}
	var err error
	if bind == "1" {
		err = lv_db.GetMasterGorm().Model(entity).Where("id=?", deviceId).Update("gateway_id", gwId).Error
	} else {
		err = lv_db.GetMasterGorm().Model(entity).Where("id=?", deviceId).Update("gateway_id", nil).Error
	}
	return err
}

// ListByPage 根据条件分页查询数据
func (svc IotDeviceService) ListByPage(params *vo.PageIotDeviceReq) (*[]vo.IotDeviceResp, int64, error) {
	var d dao.IotDeviceDao
	dept := new(model2.SysDept)
	dept, err := dept.FindById(params.DeptId)
	lv_err.HasErrAndPanic(err)
	params.Ancestors = dept.Ancestors
	listPtr, total, err := d.ListByPage(params)
	listData := *listPtr
	cacheSvc := service.GetDeviceCacheService()
	for i := 0; i < len(listData); i++ {
		it := &listData[i]
		if it.NodeType != "12" {
			continue //非网关子设备
		}
		it.GatewayName = cacheSvc.GetSysDevName(it.GatewayId)
		if it.GatewayName == "" {
			it.GatewayName = "未设置"
		}

	}
	return &listData, total, err
}

// ExportAll 导出excel
func (svc IotDeviceService) ExportAll(param *vo.PageIotDeviceReq) (string, error) {
	var d dao.IotDeviceDao
	var err error
	var listMap *[]map[string]string
	if param.PageNum > 0 { //分页导出
		listMap, _, err = d.ListMapByPage(param)
	} else { //全部导出
		listMap, err = d.ListAll(param, true)
	}
	lv_err.HasErrAndPanic(err)
	heads := []string{"", "产品ID", "云设备ID", "云产品ID", "云实例ID", "驱动实例ID", "名字", "设备状态", "描述", "密钥", "平台名称", "安装地址", "最后一次同步时间", "最后一次在线时间", "删除标记", "创建日期", "更新日期", "更新者", "创建者"}
	keys := []string{"id", "productId", "cloudDeviceId", "cloudProductId", "cloudInstanceId", "driveInstanceId", "name", "status", "description", "secret", "platform", "installLocation", "lastSyncTime", "lastOnlineTime", "delFlag", "createTime", "updateTime", "updateBy", "createBy"}
	url, err := lv_office.DownlaodExcelByListMap(&heads, &keys, listMap)
	return url, err
}
