package dict_type

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/convert"
	"lostvip.com/utils/page"
	"robvi/app/modules/sys/model"
	dict_type2 "robvi/app/modules/sys/model/system/dict_type"
	"robvi/app/modules/sys/service"
	"time"
)

// 根据主键查询数据
func SelectRecordById(id int64) (*dict_type2.Entity, error) {
	entity := &dict_type2.Entity{DictId: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除数据
func DeleteRecordById(id int64) bool {
	rs, err := (&dict_type2.Entity{DictId: id}).Delete()
	if err == nil {
		if rs > 0 {
			return true
		}
	}
	return false
}

// 批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	ida := convert.ToInt64Array(ids, ",")
	result, err := dict_type2.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

// 添加数据
func AddSave(req *dict_type2.AddReq, c *gin.Context) (int64, error) {
	var entity dict_type2.Entity
	entity.Status = req.Status
	entity.DictType = req.DictType
	entity.DictName = req.DictName
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()
	entity.CreateBy = ""
	var userService service.UserService
	user := userService.GetProfile(c)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	_, err := entity.Insert()

	return entity.DictId, err
}

// 修改数据
func EditSave(req *dict_type2.EditReq, c *gin.Context) (int64, error) {
	entity := &dict_type2.Entity{DictId: req.DictId}
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
	var userService service.UserService
	user := userService.GetProfile(c)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return entity.Update()
}

// 根据条件分页查询角色数据
func SelectListAll(params *dict_type2.SelectPageReq) ([]dict_type2.Entity, error) {
	return dict_type2.SelectListAll(params)
}

// 根据条件分页查询角色数据
func SelectListByPage(params *dict_type2.SelectPageReq) ([]dict_type2.Entity, *page.Paging, error) {
	return dict_type2.SelectListByPage(params)
}

// 根据字典类型查询信息
func SelectDictTypeByType(dictType string) *dict_type2.Entity {
	entity := &dict_type2.Entity{DictType: dictType}
	ok, err := entity.FindOne()
	if err != nil || !ok {
		return nil
	}
	return entity
}

// 导出excel
func Export(param *dict_type2.SelectPageReq) (string, error) {
	head := []string{"字典主键", "字典名称", "字典类型", "状态", "创建者", "创建时间", "更新者", "更新时间", "备注"}
	col := []string{"dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark"}
	return dict_type2.SelectListExport(param, head, col)
}

// 检查字典类型是否唯一
func CheckDictTypeUniqueAll(configKey string) string {
	entity, err := dict_type2.CheckDictTypeUniqueAll(configKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.DictId > 0 {
		return "1"
	}
	return "0"
}

// 检查字典类型是否唯一
func CheckDictTypeUnique(configKey string, dictId int64) string {
	entity, err := dict_type2.CheckDictTypeUniqueAll(configKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.DictId > 0 && entity.DictId != dictId {
		return "1"
	}
	return "0"
}

// 查询字典类型树
func SelectDictTree(params *dict_type2.SelectPageReq) *[]model.Ztree {
	var result []model.Ztree
	dictList, err := dict_type2.SelectListAll(params)
	if err == nil && dictList != nil {
		for _, item := range dictList {
			var tmp model.Ztree
			tmp.Id = item.DictId
			tmp.Name = transDictName(item)
			tmp.Title = item.DictType
			result = append(result, tmp)
		}
	}
	return &result
}

func transDictName(entity dict_type2.Entity) string {
	return `(` + entity.DictName + `)&nbsp;&nbsp;&nbsp;` + entity.DictType
}
