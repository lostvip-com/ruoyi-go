package service

import (
	"common/common_vo"
	"common/model"
	"common/session"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_reflect"
	dao2 "main/internal/system/dao"
	"time"
)

type DictDataService struct {
}

// 根据主键查询数据
func (svc *DictDataService) SelectRecordById(id int64) (*model.SysDictData, error) {
	entity := &model.SysDictData{DictCode: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除数据
func (svc *DictDataService) DeleteRecordById(id int64) bool {
	err := (&model.SysDictData{DictCode: id}).Delete()
	if err == nil {
		return true
	}
	return false
}

// 批量删除数据记录
func (svc *DictDataService) DeleteRecordByIds(ids string) error {
	ida := lv_conv.ToInt64Array(ids, ",")
	data := new(dao2.DictDataDao)
	err := data.DeleteBatch(ida...)
	return err
}

// 添加数据
func (svc *DictDataService) AddSave(req *common_vo.AddDictDataReq, c *gin.Context) (int64, error) {
	var entity model.SysDictData
	entity.DictType = req.DictType
	entity.Status = req.Status
	entity.DictLabel = req.DictLabel
	entity.CssClass = req.CssClass
	entity.DictSort = req.DictSort
	entity.DictValue = req.DictValue
	entity.IsDefault = req.IsDefault
	entity.ListClass = req.ListClass
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()
	entity.CreateBy = ""
	var userService UserService
	user := userService.GetProfile(c)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	err := entity.Save()

	return entity.DictCode, err
}

// 修改数据
func (svc *DictDataService) EditSave(req *common_vo.EditDictDataReq, c *gin.Context) error {
	entity := &model.SysDictData{DictCode: req.DictCode}
	po, err := entity.FindOne()
	if err != nil {
		return errors.New("数据不存在")
	}
	lv_reflect.CopyProperties(req, po)
	po.UpdateTime = time.Now()
	po.UpdateBy = session.GetLoginInfo(c).Username
	var userService UserService
	user := userService.GetProfile(c)
	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return entity.Updates()
}

// 根据条件分页查询角色数据
func (svc *DictDataService) SelectListAll(params *common_vo.SelectDictDataPageReq) ([]model.SysDictData, error) {
	var dao dao2.DictDataDao
	return dao.SelectListAll(params)
}

// 根据条件分页查询角色数据
func (svc *DictDataService) SelectListByPage(params *common_vo.SelectDictDataPageReq) (*[]model.SysDictData, int64, error) {
	var dao dao2.DictDataDao
	return dao.SelectListByPage(params)
}

// 导出excel
func (svc *DictDataService) Export(param *common_vo.SelectDictDataPageReq) (string, error) {
	head := []string{"字典编码", "字典排序", "字典标签", "字典键值", "字典类型", "样式属性", "表格回显样式", "是否默认", "状态", "创建者", "创建时间", "更新者", "更新时间", "备注"}
	col := []string{"dict_code", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "status", "create_by", "create_time", "update_by", "update_time", "remark"}
	var dao dao2.DictDataDao
	return dao.SelectListExport(param, head, col)
}
