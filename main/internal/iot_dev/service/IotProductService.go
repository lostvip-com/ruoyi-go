// ==========================================================================
// LV自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-07-19 17:16:13 +0800 CST
// 生成人：lv
// ==========================================================================
package service

import (
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_office"
	"github.com/lostvip-com/lv_framework/utils/lv_reflect"
	"main/internal/iot_dev/dao"
	"main/internal/iot_dev/model"
	"main/internal/iot_dev/vo"
	"main/internal/system/service"
	"time"
)

type IotProductService struct{}

// FindById 根据主键查询数据
func (svc IotProductService) FindById(id int64) (*model.IotProduct, error) {
	entity := &model.IotProduct{Id: id}
	err := entity.FindById()
	// nodeType
	var dictSvc service.DictService
	label := dictSvc.GetDictLabel("iot_node_type", entity.NodeType)
	entity.NodeTypeName = label

	protoName := dictSvc.GetDictLabel("iot_protocol_type", entity.Protocol)
	entity.ProtocolName = protoName
	return entity, err
}

// DeleteById 根据主键删除数据
func (svc IotProductService) DeleteById(id int64) error {
	err := (&model.IotProduct{Id: id}).Delete()
	return err
}

// DeleteByIds 批量删除数据记录
func (svc IotProductService) DeleteByIds(ids string) int64 {
	ida := lv_conv.ToInt64Array(ids, ",")
	var d dao.IotProductDao
	rowsAffected := d.DeleteByIds(ida)
	return rowsAffected
}

// AddSave 添加数据
func (svc IotProductService) AddSave(req *vo.AddIotProductReq) (int64, error) {
	var entity = new(model.IotProduct)
	lv_reflect.CopyProperties(req, entity)

	entity.CreateTime = time.Now()
	entity.UpdateTime = entity.CreateTime
	entity.CreateBy = req.CreateBy
	err := entity.Save()
	lv_err.HasErrAndPanic(err)
	return entity.Id, err
}

// EditSave 修改数据
func (svc IotProductService) EditSave(req *vo.EditIotProductReq) error {
	entity := &model.IotProduct{Id: req.Id}
	err := entity.FindById()
	lv_err.HasErrAndPanic(err)

	lv_reflect.CopyProperties(req, entity)
	entity.UpdateTime = time.Now()
	entity.UpdateBy = req.UpdateBy
	return entity.Updates()
}

// ListByPage 根据条件分页查询数据
func (svc IotProductService) ListByPage(params *vo.PageIotProductReq) (*[]model.IotProduct, int64, error) {
	var d dao.IotProductDao
	return d.ListByPage(params)
}

// ExportAll 导出excel
func (svc IotProductService) ExportAll(param *vo.PageIotProductReq) (string, error) {
	var d dao.IotProductDao
	var err error
	var listMap *[]map[string]string
	if param.PageNum > 0 { //分页导出
		listMap, _, err = d.ListMapByPage(param)
	} else { //全部导出
		listMap, err = d.ListAll(param, true)
	}
	lv_err.HasErrAndPanic(err)
	heads := []string{"主键", "名字", "产品标识", "云产品ID", "云实例ID", "平台", "协议", "节点类型", "网络类型", "数据类型", "最后一次同步时间", "工厂名称", "描述", "产品状态", "扩展字段", "删除标记", "创建日期", "更新日期", "更新者", "创建者"}
	keys := []string{"id", "name", "key", "cloudProductId", "cloudInstanceId", "platform", "protocol", "nodeType", "netType", "dataFormat", "lastSyncTime", "factory", "description", "status", "extra", "delFlag", "createTime", "updateTime", "updateBy", "createBy"}
	url, err := lv_office.DownlaodExcelByListMap(&heads, &keys, listMap)
	return url, err
}
