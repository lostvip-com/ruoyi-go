package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/utils/lv_conv"
	"github.com/lv_framework/utils/lv_web"
	"github.com/lv_framework/web/dto"
	"main/internal/system/model"
	"main/internal/system/service"
	"main/internal/system/vo"
	"net/http"
)

type DictTypeController struct {
}

// 列表页
func (w *DictTypeController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "system/dict/type/list").WriteTpl()
}

// 列表分页数据
func (w *DictTypeController) ListAjax(c *gin.Context) {
	var req *vo.SelectDictTypePageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("字典类型管理", req).WriteJsonExit()
		return
	}
	rows := make([]model.DictType, 0)
	var dictTypeService service.DictTypeService
	result, page, err := dictTypeService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	lv_web.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 新增页面
func (w *DictTypeController) Add(c *gin.Context) {
	lv_web.BuildTpl(c, "system/dict/type/add").WriteTpl()
}

// 新增页面保存
func (w *DictTypeController) AddSave(c *gin.Context) {
	var req *model.DictType
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
		return
	}
	var dictTypeService service.DictTypeService
	if dictTypeService.CheckDictTypeUniqueAll(req.DictType) == "1" {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg("字典类型已存在").Log("字典管理", req).WriteJsonExit()
		return
	}

	rid, err := dictTypeService.AddSave(req, c)

	if err != nil || rid <= 0 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).Log("字典管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(rid).Log("字典管理", req).WriteJsonExit()
}

// 修改页面
func (w *DictTypeController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典类型错误",
		})
		return
	}
	var dictTypeService service.DictTypeService
	entity, err := dictTypeService.SelectRecordById(id)

	if err != nil || entity == nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典类型不存在",
		})
		return
	}

	lv_web.BuildTpl(c, "system/dict/type/edit").WriteTpl(gin.H{
		"dict": entity,
	})
}

// 修改页面保存
func (w *DictTypeController) EditSave(c *gin.Context) {
	var req *vo.EditDictTypeReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).SetMsg(err.Error()).Log("字典类型管理", req).WriteJsonExit()
		return
	}
	var dictTypeService service.DictTypeService
	if dictTypeService.CheckDictTypeUnique(req.DictType, req.DictId) == "1" {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).SetMsg("字典类型已存在").Log("字典类型管理", req).WriteJsonExit()
		return
	}

	rs, err := dictTypeService.EditSave(req, c)

	if err != nil || rs <= 0 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).Log("字典类型管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).Log("字典类型管理", req).WriteJsonExit()
}

// 删除数据
func (w *DictTypeController) Remove(c *gin.Context) {
	var req *dto.IdsReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
		return
	}
	var dictTypeService service.DictTypeService
	rs := dictTypeService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(dto.Buniss_Del).Log("字典管理", req).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).Log("字典管理", req).WriteJsonExit()
	}
}

// 数据详情
func (w *DictTypeController) Detail(c *gin.Context) {
	dictId := lv_conv.Int64(c.Query("dictId"))
	if dictId <= 0 {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	var dictTypeService service.DictTypeService
	dict, _ := dictTypeService.SelectRecordById(dictId)

	if dict == nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典类别不存在",
		})
		return
	}

	dictList, _ := dictTypeService.SelectListAll(nil)
	if dictList == nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误2",
		})
		return
	}

	lv_web.BuildTpl(c, "system/dict/data/list").WriteTpl(gin.H{
		"dict":     dict,
		"dictList": dictList,
	})
}

// 选择字典树
func (w *DictTypeController) SelectDictTree(c *gin.Context) {
	columnId := lv_conv.Int64(c.Query("columnId"))
	dictType := c.Query("dictType")
	if columnId <= 0 || dictType == "" {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})

		return
	}

	if dictType == "-" {
		dictType = "-"
	}
	var dictTypeService service.DictTypeService
	var dict model.DictType
	rs := dictTypeService.SelectDictTypeByType(dictType)
	if rs != nil {
		dict = *rs
	}

	lv_web.BuildTpl(c, "system/dict/type/tree").WriteTpl(gin.H{
		"columnId": columnId,
		"dict":     dict,
	})
}

// 导出
func (w *DictTypeController) Export(c *gin.Context) {
	var req *vo.SelectDictTypePageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
		return
	}
	var dictTypeService service.DictTypeService
	url, err := dictTypeService.Export(req)
	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
		return
	}

	lv_web.SucessResp(c).SetMsg(url).Log("导出Excel", req).WriteJsonExit()
}

// 检查字典类型是否唯一不包括本参数
func (w *DictTypeController) CheckDictTypeUnique(c *gin.Context) {
	var req *vo.CheckDictTypeReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}
	var dictTypeService service.DictTypeService
	result := dictTypeService.CheckDictTypeUnique(req.DictType, req.DictId)
	c.Writer.WriteString(result)
}

// 检查字典类型是否唯一
func (w *DictTypeController) CheckDictTypeUniqueAll(c *gin.Context) {
	var req *vo.CheckDictTypeALLReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}
	var dictTypeService service.DictTypeService
	result := dictTypeService.CheckDictTypeUniqueAll(req.DictType)

	c.Writer.WriteString(result)
}

// 加载部门列表树结构的数据
func (w *DictTypeController) TreeData(c *gin.Context) {
	var dictTypeService service.DictTypeService
	result := dictTypeService.SelectDictTree(nil)
	c.JSON(http.StatusOK, result)
}
