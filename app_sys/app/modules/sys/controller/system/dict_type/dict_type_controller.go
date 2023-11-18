package dict_type

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/gconv"
	response2 "lostvip.com/utils/response"
	"net/http"
	"robvi/app/modules/sys/model"
	dict_type2 "robvi/app/modules/sys/model/system/dict_type"
	dictTypeService "robvi/app/modules/sys/service/system/dict_type"
)

// 列表页
func List(c *gin.Context) {
	response2.BuildTpl(c, "system/dict/type/list").WriteTpl()
}

// 列表分页数据
func ListAjax(c *gin.Context) {
	var req *dict_type2.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetMsg(err.Error()).Log("字典类型管理", req).WriteJsonExit()
		return
	}
	rows := make([]dict_type2.Entity, 0)
	result, page, err := dictTypeService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response2.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 新增页面
func Add(c *gin.Context) {
	response2.BuildTpl(c, "system/dict/type/add").WriteTpl()
}

// 新增页面保存
func AddSave(c *gin.Context) {
	var req *dict_type2.AddReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
		return
	}

	if dictTypeService.CheckDictTypeUniqueAll(req.DictType) == "1" {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("字典类型已存在").Log("字典管理", req).WriteJsonExit()
		return
	}

	rid, err := dictTypeService.AddSave(req, c)

	if err != nil || rid <= 0 {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).Log("字典管理", req).WriteJsonExit()
		return
	}
	response2.SucessResp(c).SetData(rid).Log("字典管理", req).WriteJsonExit()
}

// 修改页面
func Edit(c *gin.Context) {
	id := gconv.Int64(c.Query("id"))
	if id <= 0 {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典类型错误",
		})
		return
	}

	entity, err := dictTypeService.SelectRecordById(id)

	if err != nil || entity == nil {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典类型不存在",
		})
		return
	}

	response2.BuildTpl(c, "system/dict/type/edit").WriteTpl(gin.H{
		"dict": entity,
	})
}

// 修改页面保存
func EditSave(c *gin.Context) {
	var req *dict_type2.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("字典类型管理", req).WriteJsonExit()
		return
	}

	if dictTypeService.CheckDictTypeUnique(req.DictType, req.DictId) == "1" {
		response2.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg("字典类型已存在").Log("字典类型管理", req).WriteJsonExit()
		return
	}

	rs, err := dictTypeService.EditSave(req, c)

	if err != nil || rs <= 0 {
		response2.ErrorResp(c).SetBtype(model.Buniss_Edit).Log("字典类型管理", req).WriteJsonExit()
		return
	}
	response2.SucessResp(c).Log("字典类型管理", req).WriteJsonExit()
}

// 删除数据
func Remove(c *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
		return
	}

	rs := dictTypeService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response2.SucessResp(c).SetBtype(model.Buniss_Del).Log("字典管理", req).WriteJsonExit()
	} else {
		response2.ErrorResp(c).SetBtype(model.Buniss_Del).Log("字典管理", req).WriteJsonExit()
	}
}

// 数据详情
func Detail(c *gin.Context) {
	dictId := gconv.Int64(c.Query("dictId"))
	if dictId <= 0 {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	dict, _ := dictTypeService.SelectRecordById(dictId)

	if dict == nil {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典类别不存在",
		})
		return
	}

	dictList, _ := dictTypeService.SelectListAll(nil)
	if dictList == nil {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误2",
		})
		return
	}

	response2.BuildTpl(c, "system/dict/data/list").WriteTpl(gin.H{
		"dict":     dict,
		"dictList": dictList,
	})
}

// 选择字典树
func SelectDictTree(c *gin.Context) {
	columnId := gconv.Int64(c.Query("columnId"))
	dictType := c.Query("dictType")
	if columnId <= 0 || dictType == "" {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})

		return
	}

	if dictType == "-" {
		dictType = "-"
	}

	var dict dict_type2.Entity
	rs := dictTypeService.SelectDictTypeByType(dictType)
	if rs != nil {
		dict = *rs
	}

	response2.BuildTpl(c, "system/dict/type/tree").WriteTpl(gin.H{
		"columnId": columnId,
		"dict":     dict,
	})
}

// 导出
func Export(c *gin.Context) {
	var req *dict_type2.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
		return
	}
	url, err := dictTypeService.Export(req)

	if err != nil {
		response2.ErrorResp(c).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
		return
	}

	response2.SucessResp(c).SetMsg(url).Log("导出Excel", req).WriteJsonExit()
}

// 检查字典类型是否唯一不包括本参数
func CheckDictTypeUnique(c *gin.Context) {
	var req *dict_type2.CheckDictTypeReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := dictTypeService.CheckDictTypeUnique(req.DictType, req.DictId)
	c.Writer.WriteString(result)
}

// 检查字典类型是否唯一
func CheckDictTypeUniqueAll(c *gin.Context) {
	var req *dict_type2.CheckDictTypeALLReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := dictTypeService.CheckDictTypeUniqueAll(req.DictType)

	c.Writer.WriteString(result)
}

// 加载部门列表树结构的数据
func TreeData(c *gin.Context) {
	result := dictTypeService.SelectDictTree(nil)
	c.JSON(http.StatusOK, result)
}
