package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	dao2 "main/internal/system/dao"
	"main/internal/system/model"
	"main/internal/system/vo"
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
	rs, _ := (&model.SysDictData{DictCode: id}).Delete()
	if rs > 0 {
		return true
	}
	return false
}

// 批量删除数据记录
func (svc *DictDataService) DeleteRecordByIds(ids string) int64 {
	ida := lv_conv.ToInt64Array(ids, ",")
	result, err := model.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

// 添加数据
func (svc *DictDataService) AddSave(req *vo.AddDictDataReq, c *gin.Context) (int64, error) {
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

	_, err := entity.Insert()

	return entity.DictCode, err
}

// 修改数据
func (svc *DictDataService) EditSave(req *vo.EditDictDataReq, c *gin.Context) (int64, error) {
	entity := &model.SysDictData{DictCode: req.DictCode}
	ok, err := entity.FindOne()

	if err != nil || !ok {
		return 0, err
	}

	if entity == nil {
		return 0, errors.New("数据不存在")
	}

	entity.DictType = req.DictType
	entity.Status = req.Status
	entity.DictLabel = req.DictLabel
	entity.CssClass = req.CssClass
	entity.DictSort = req.DictSort
	entity.DictValue = req.DictValue
	entity.IsDefault = req.IsDefault
	entity.ListClass = req.ListClass
	entity.Remark = req.Remark
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""
	var userService UserService
	user := userService.GetProfile(c)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return entity.Update()
}

// 根据条件分页查询角色数据
func (svc *DictDataService) SelectListAll(params *vo.SelectDictDataPageReq) ([]model.SysDictData, error) {
	var dao dao2.DictDataDao
	return dao.SelectListAll(params)
}

// 根据条件分页查询角色数据
func (svc *DictDataService) SelectListByPage(params *vo.SelectDictDataPageReq) (*[]model.SysDictData, *lv_web.Paging, error) {
	var dao dao2.DictDataDao
	return dao.SelectListByPage(params)
}

// 导出excel
func (svc *DictDataService) Export(param *vo.SelectDictDataPageReq) (string, error) {
	head := []string{"字典编码", "字典排序", "字典标签", "字典键值", "字典类型", "样式属性", "表格回显样式", "是否默认", "状态", "创建者", "创建时间", "更新者", "更新时间", "备注"}
	col := []string{"dictCode", "dictSort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "status", "create_by", "create_time", "update_by", "update_time", "remark"}
	var dao dao2.DictDataDao
	return dao.SelectListExport(param, head, col)
}
