// ==========================================================================
// LV自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-07-30 21:59:38 +0800 CST
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
	"time"
)

type IotPrdPropertyService struct{}

// FindById 根据主键查询数据
func (svc IotPrdPropertyService) FindById(id int64) (*model.IotPrdProperty, error) {
	entity := &model.IotPrdProperty{Id: id}
	err := entity.FindById()
	return entity, err
}

// DeleteById 根据主键删除数据
func (svc IotPrdPropertyService) DeleteById(id int64) error {
	err := (&model.IotPrdProperty{Id: id}).Delete()
	return err
}

// DeleteByIds 批量删除数据记录
func (svc IotPrdPropertyService) DeleteByIds(ids string) int64 {
	ida := lv_conv.ToInt64Array(ids, ",")
	var d dao.IotPrdPropertyDao
	rowsAffected := d.DeleteByIds(ida)
	return rowsAffected
}

// AddSave 添加数据
func (svc IotPrdPropertyService) AddSave(req *vo.AddIotPrdPropertyReq) (int64, error) {
	var entity = new(model.IotPrdProperty)
	lv_reflect.CopyProperties(req, entity)
	entity.CreateTime = time.Now()
	entity.UpdateTime = entity.CreateTime
	entity.CreateBy = req.CreateBy
	err := entity.Save()
	lv_err.HasErrAndPanic(err)
	return entity.Id, err
}

// EditSave 修改数据
func (svc IotPrdPropertyService) EditSave(req *vo.EditIotPrdPropertyReq) error {
	entity := &model.IotPrdProperty{Id: req.Id}
	err := entity.FindById()
	lv_err.HasErrAndPanic(err)
	lv_reflect.CopyProperties(req, entity)
	entity.UpdateTime = time.Now()
	entity.UpdateBy = req.UpdateBy
	return entity.Updates()
}

// ListByPage 根据条件分页查询数据
func (svc IotPrdPropertyService) ListByPage(params *vo.PageIotPrdPropertyReq) (*[]model.IotPrdProperty, int64, error) {
	var d dao.IotPrdPropertyDao
	return d.ListByPage(params)
}

// ExportAll 导出excel
func (svc IotPrdPropertyService) ExportAll(param *vo.PageIotPrdPropertyReq) (string, error) {
	var d dao.IotPrdPropertyDao
	var err error
	var listMap *[]map[string]string
	if param.PageNum > 0 { //分页导出
		listMap, _, err = d.ListMapByPage(param)
	} else { //全部导出
		listMap, err = d.ListAll(param, true)
	}
	lv_err.HasErrAndPanic(err)
	heads := []string{"主键", "产品ID", "名字", "标识符", "标签", "读写模型", "数据类型", "单位", "描述", "删除标记", "创建日期", "更新日期", "更新者", "创建者", "租户id"}
	keys := []string{"id", "productId", "name", "code", "tag", "accessMode", "type", "unit", "remark", "delFlag", "createTime", "updateTime", "updateBy", "createBy", "tenantId"}
	url, err := lv_office.DownlaodExcelByListMap(&heads, &keys, listMap)
	return url, err
}
