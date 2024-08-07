package service

import (
	"common/cm_vo"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	"github.com/lostvip-com/lv_framework/web/dto"
	"github.com/spf13/cast"
	dao2 "main/internal/system/dao"
	"main/internal/system/model"
	"time"
)

type DictTypeService struct {
}

// 根据主键查询数据
func (svc *DictTypeService) SelectRecordById(id int64) (*model.DictType, error) {
	entity := &model.DictType{DictId: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除数据
func (svc *DictTypeService) DeleteRecordById(id int64) bool {
	rs, err := (&model.DictType{DictId: id}).Delete()
	if err == nil {
		if rs > 0 {
			return true
		}
	}
	return false
}

// 批量删除数据记录
func (svc *DictTypeService) DeleteRecordByIds(ids string) error {
	ida := lv_conv.ToInt64Array(ids, ",")
	data := new(model.SysDictData)
	//data.DeleteBatch()
	tp := new(model.DictType)
	for _, id := range ida {
		tp.DictId = cast.ToInt64(id)
		_, err := tp.FindOne()
		lv_err.HasErrAndPanic(err)
		//
		data.DeleteByDictType(tp.DictType)
		_, err = tp.Delete()
		lv_err.HasErrAndPanic(err)
	}
	return nil
}

// 添加数据
func (svc *DictTypeService) AddSave(req *cm_vo.AddDictTypeReq, c *gin.Context) (int64, error) {
	var entity model.DictType
	entity.Status = req.Status
	entity.DictType = req.DictType
	entity.DictName = req.DictName
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()
	entity.CreateBy = ""
	var userService UserService
	user := userService.GetProfile(c)
	if user != nil {
		entity.CreateBy = user.LoginName
	}
	_, err := entity.Insert()
	return entity.DictId, err
}

// 修改数据
func (svc *DictTypeService) EditSave(req *cm_vo.EditDictTypeReq, c *gin.Context) (int64, error) {
	entity := &model.DictType{DictId: req.DictId}
	ok, err := entity.FindOne()

	if err != nil || !ok {
		return 0, err
	}

	if entity == nil {
		return 0, errors.New("数据不存在")
	}
	entity.Status = req.Status
	entity.DictType = req.DictType
	entity.DictName = req.DictName
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
func (svc *DictTypeService) SelectListAll(params *cm_vo.DictTypePageReq) ([]model.DictType, error) {
	var dao dao2.DictTypeDao
	return dao.SelectListAll(params)
}

// 根据条件分页查询角色数据
func (svc *DictTypeService) SelectListByPage(params *cm_vo.DictTypePageReq) ([]model.DictType, *lv_web.Paging, error) {
	var dao dao2.DictTypeDao
	return dao.SelectListByPage(params)
}

// 根据字典类型查询信息
func (svc *DictTypeService) SelectDictTypeByType(dictType string) *model.DictType {
	entity := &model.DictType{DictType: dictType}
	ok, err := entity.FindOne()
	if err != nil || !ok {
		return nil
	}
	return entity
}

// 导出excel
func (svc *DictTypeService) Export(param *cm_vo.DictTypePageReq) (string, error) {
	head := []string{"字典主键", "字典名称", "字典类型", "状态", "创建者", "创建时间", "更新者", "更新时间", "备注"}
	col := []string{"dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark"}
	var dao dao2.DictTypeDao
	return dao.SelectListExport(param, head, col)
}

// 检查字典类型是否唯一
func (svc *DictTypeService) CheckDictTypeUniqueAll(configKey string) (bool, error) {
	var dao dao2.DictTypeDao
	exist, err := dao.CheckDictTypeUniqueAll(configKey)

	return exist, err
}

// 检查字典类型是否唯一
func (svc *DictTypeService) IsDictTypeExist(configKey string) bool {
	var dao dao2.DictTypeDao
	exist, err := dao.CheckDictTypeUniqueAll(configKey)
	lv_err.HasErrAndPanic(err)
	return exist
}

// 查询字典类型树
func (svc *DictTypeService) SelectDictTree(params *cm_vo.DictTypePageReq) *[]dto.Ztree {
	var result []dto.Ztree
	var dao dao2.DictTypeDao
	dictList, err := dao.SelectListAll(params)
	if err == nil && dictList != nil {
		for _, item := range dictList {
			var tmp dto.Ztree
			tmp.Id = item.DictId
			tmp.Name = svc.transDictName(item)
			tmp.Title = item.DictType
			result = append(result, tmp)
		}
	}
	return &result
}

func (svc *DictTypeService) transDictName(entity model.DictType) string {
	return `(` + entity.DictName + `)&nbsp;&nbsp;&nbsp;` + entity.DictType
}
