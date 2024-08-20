package service

import (
	"common/common_vo"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"github.com/spf13/cast"
	dao2 "main/internal/system/dao"
	"main/internal/system/model"
	"time"
)

type DictTypeService struct {
}

// 根据主键查询数据
func (svc *DictTypeService) SelectRecordById(id int64) (*model.SysDictType, error) {
	entity := &model.SysDictType{DictId: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除数据
func (svc *DictTypeService) DeleteRecordById(id int64) bool {
	err := (&model.SysDictType{DictId: id}).Delete()
	if err == nil {
		return true
	}
	return false
}

// 批量删除数据记录
func (svc *DictTypeService) DeleteRecordByIds(ids string) error {
	ida := lv_conv.ToInt64Array(ids, ",")
	data := new(dao2.DictDataDao)
	//data.DeleteBatch()
	tp := new(model.SysDictType)
	for _, id := range ida {
		tp.DictId = cast.ToInt64(id)
		_, err := tp.FindOne()
		lv_err.HasErrAndPanic(err)
		//
		data.DeleteByType(tp.DictType)
		err = tp.Delete()
		lv_err.HasErrAndPanic(err)
	}
	return nil
}

// 添加数据
func (svc *DictTypeService) AddSave(req *common_vo.AddDictTypeReq, c *gin.Context) (int64, error) {
	var entity model.SysDictType
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
	err := entity.Save()
	return entity.DictId, err
}

// 修改数据
func (svc *DictTypeService) EditSave(req *common_vo.EditDictTypeReq, c *gin.Context) (int64, error) {
	entity := &model.SysDictType{DictId: req.DictId}
	entity, err := entity.FindOne()
	if err != nil {
		return 0, err
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
	err = entity.Updates()
	return entity.DictId, err
}

// 根据条件分页查询角色数据
func (svc *DictTypeService) SelectListAll(params *common_vo.DictTypePageReq) ([]model.SysDictType, error) {
	var dao dao2.DictTypeDao
	return dao.SelectListAll(params)
}

// 根据条件分页查询角色数据
func (svc *DictTypeService) SelectListByPage(params *common_vo.DictTypePageReq) ([]model.SysDictType, int64, error) {
	var dao dao2.DictTypeDao
	return dao.SelectListByPage(params)
}

// 根据字典类型查询信息
func (svc *DictTypeService) SelectDictTypeByType(dictType string) *model.SysDictType {
	entity := &model.SysDictType{DictType: dictType}
	entity, err := entity.FindOne()
	if err != nil {
		return nil
	}
	return entity
}

// 导出excel
func (svc *DictTypeService) Export(param *common_vo.DictTypePageReq) (string, error) {
	head := []string{"字典主键", "字典名称", "字典类型", "状态", "创建者", "创建时间", "更新者", "更新时间", "备注"}
	col := []string{"dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark"}
	var dao dao2.DictTypeDao
	return dao.SelectListExport(param, head, col)
}

// 检查字典类型是否唯一
func (svc *DictTypeService) CheckDictTypeUniqueAll(configKey string) (bool, error) {
	var entity model.SysDictType
	entity.DictType = configKey
	count, err := entity.Count()

	return count > 0, err
}

// 检查字典类型是否唯一
func (svc *DictTypeService) IsDictTypeExist(configKey string) bool {
	var entity model.SysDictType
	entity.DictType = configKey
	count, err := entity.Count()
	lv_err.HasErrAndPanic(err)
	return count > 0
}

// 查询字典类型树
func (svc *DictTypeService) SelectDictTree(params *common_vo.DictTypePageReq) *[]lv_dto.Ztree {
	var result []lv_dto.Ztree
	var dao dao2.DictTypeDao
	dictList, err := dao.SelectListAll(params)
	if err == nil && dictList != nil {
		for _, item := range dictList {
			var tmp lv_dto.Ztree
			tmp.Id = item.DictId
			tmp.Name = svc.transDictName(item)
			tmp.Title = item.DictType
			result = append(result, tmp)
		}
	}
	return &result
}

func (svc *DictTypeService) transDictName(entity model.SysDictType) string {
	return `(` + entity.DictName + `)&nbsp;&nbsp;&nbsp;` + entity.DictType
}
