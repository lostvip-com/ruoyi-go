package controller

import (
	"common/myconf"
	util2 "common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/lv_batis"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_file"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"strings"
	menuModel "system/model"
	"system/service"
	"system/vo"
)

type GenController struct{}

// 表单构建
func (w *GenController) Build(c *gin.Context) {
	util2.BuildTpl(c, "tool/build").WriteTpl()
}

// swagger文档
func (w *GenController) ExecSqlFile(c *gin.Context) {
	tableId := lv_conv.Int64(c.Query("tableId"))
	//获取参数
	if tableId <= 0 {
		util2.ErrorResp(c).SetBtype(lv_dto.Buniss_Add).SetMsg("参数错误").Log("执行SQL文件错误", gin.H{"tableId": tableId})
	}
	genTable := menuModel.GenTable{}
	po, err := genTable.FindById(tableId)
	if err != nil {
		panic(err.Error())
	}

	//err = lv_db.ExecSqlFile(sqlFile)
	// Loads queries from file
	batis, err := lv_batis.LoadFromFile(lv_global.Config().GetTmpPath() + "/" + po.TbName + "_menu.sql")
	// Run queries
	tb := lv_db.GetMasterGorm()
	//cfg := global.GetConfigInstance()
	batis.Exec(tb, "menu")
	menuName := po.FunctionName
	sysmenu := menuModel.SysMenu{}
	sysmenu.MenuName = menuName
	err = sysmenu.FindLastOne()
	lv_err.HasErrAndPanic(err)
	pmenuId := sysmenu.MenuId
	_, err = batis.Exec(tb, "menu_button_create", pmenuId)
	_, err = batis.Exec(tb, "menu_button_retrieve", pmenuId)
	_, err = batis.Exec(tb, "menu_button_update", pmenuId)
	_, err = batis.Exec(tb, "menu_button_delete", pmenuId)
	_, err = batis.Exec(tb, "menu_button_export", pmenuId)
	if err != nil {
		panic(err)
	}
	util2.SucessData(c, nil)
}

// swagger文档
func (w *GenController) Swagger(c *gin.Context) {
	a := c.Query("a")
	if a == "r" {
		//重新生成文档
		curDir, err := os.Getwd()
		if err != nil {
			util2.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
				"desc": "参数错误",
			})
			c.Abort()
			return
		}
		genPath := curDir + "/static/swagger"
		err = w.generateSwaggerFiles(genPath)
		if err != nil {
			util2.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
				"desc": "参数错误",
			})
			c.Abort()
			return
		}
	}
	c.Redirect(http.StatusFound, "/static/swagger/index.html")
}

// 自动生成文档 swag init -o static/swagger
func (w *GenController) generateSwaggerFiles(output string) error {

	cmd := exec.Command("swag", "init -o "+output)
	// 保证关闭输出流
	if err := cmd.Start(); err != nil { // 运行命令
		return err
	}

	return nil
}

// 生成代码列表页面
func (w *GenController) Gen(c *gin.Context) {
	util2.BuildTpl(c, "tool/gen_list_tables").WriteTpl()
}

func (w *GenController) GenList(c *gin.Context) {
	var req *vo.GenTablePageReq
	tableService := service.TableService{}
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util2.ErrorResp(c).SetMsg(err.Error()).Log("生成代码", req).WriteJsonExit()
		return
	}
	rows := make([]menuModel.GenTable, 0)
	result, total, err := tableService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	util2.SucessPage(c, rows, total)
}

// 导入数据表
func (w *GenController) ImportTable(c *gin.Context) {
	util2.WriteTpl(c, "tool/gen_import_table")
}

// 删除数据
func (w *GenController) Remove(c *gin.Context) {
	var req *lv_dto.IdsReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util2.ErrorResp(c).SetBtype(lv_dto.Buniss_Del).Log("生成代码", req).WriteJsonExit()
		return
	}
	tableService := service.TableService{}
	err := tableService.DeleteRecordByIds(req.Ids)

	if err == nil {
		util2.SucessResp(c).SetBtype(lv_dto.Buniss_Del).Log("生成代码", req).WriteJsonExit()
	} else {
		util2.ErrorResp(c).SetBtype(lv_dto.Buniss_Del).Log("生成代码", req).WriteJsonExit()
	}
}

// Edit 修改数据
func (w *GenController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	tableService := service.TableService{}
	entity, err := tableService.SelectRecordById(id)
	lv_err.HasErrAndPanic(err)
	goTypeTpl := tableService.GoTypeTpl()
	queryTypeTpl := tableService.QueryTypeTpl()
	htmlTypeTpl := tableService.HtmlTypeTpl()

	util2.WriteTpl(c, "tool/gen_edit_table", gin.H{
		"table":        entity,
		"goTypeTpl":    template.HTML(goTypeTpl),
		"queryTypeTpl": template.HTML(queryTypeTpl),
		"htmlTypeTpl":  template.HTML(htmlTypeTpl),
	})
}

// EditSave 修改数据保存
func (w *GenController) EditSave(c *gin.Context) {
	var req vo.GenTableEditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util2.ErrorResp(c).SetMsg(err.Error()).SetBtype(lv_dto.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
		return
	}
	tableService := service.TableService{}
	err := tableService.SaveEdit(&req, c)
	if err != nil {
		util2.ErrorResp(c).SetMsg(err.Error()).SetBtype(lv_dto.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
		return
	}
	util2.SucessResp(c).SetBtype(lv_dto.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
}

// Preview 预览代码
func (w *GenController) Preview(c *gin.Context) {
	tableId := lv_conv.Int64(c.Query("tableId"))
	if tableId <= 0 {
		c.JSON(http.StatusOK, lv_dto.CommonRes{
			Code:  500,
			Btype: lv_dto.Buniss_Other,
			Msg:   "参数错误",
		})
	}
	tableService := service.TableService{}
	entity, err := tableService.SelectRecordById(tableId)

	if err != nil || entity == nil {
		c.JSON(http.StatusOK, lv_dto.CommonRes{
			Code:  500,
			Btype: lv_dto.Buniss_Other,
			Msg:   "数据不存在",
		})
	}
	tableService.SetPkColumn(entity, entity.Columns)
	var codeGenService service.CodeGenService
	dataMap := codeGenService.PreviewCode(entity)
	retMap := make(map[string]string)
	for _, mp := range dataMap {
		for k, v := range mp {
			retMap[k] = v
		}
	}
	c.JSON(http.StatusOK, lv_dto.CommonRes{
		Code: 200,
		Data: retMap,
	})
}

// 生成代码
func (w *GenController) GenCode(c *gin.Context) {
	overwrite := myconf.GetConfigInstance().GetBool("gen.overwrite")
	tableId := lv_conv.Int64(c.Query("tableId"))
	tableService := service.TableService{}
	entity, err := tableService.SelectRecordById(tableId)
	if err != nil || entity == nil {
		c.JSON(http.StatusOK, lv_dto.CommonRes{
			Code:  500,
			Btype: lv_dto.Buniss_Other,
			Msg:   "数据不存在",
		})
	}
	tableService.SetPkColumn(entity, entity.Columns)
	var codeGenService service.CodeGenService
	codeGenService.GenCode(entity, overwrite)
	//(genService)
	util2.Success(c, gin.H{"tableId": tableId}, "生成代码")
}

func canGenIt(overwrite bool, file string) bool {
	if overwrite { //允许覆盖
		lv_log.Warn("--------->您配置了 overwrite 开关的值为true，旧文件会被覆盖！！！！ ")
		return true
	} else { // 不允许覆盖
		if lv_file.Exists(file) { //文件已经存在，不允许重新生成
			lv_log.Warn("=======> 文件已经存在，本次将不会生成新文件！！！！！！！！！！！！ ")
			return false
		} else { //文件不存在，允许重新生成
			return true
		}
	}
}

// 查询数据库列表
func (w *GenController) DataList(c *gin.Context) {
	var req *vo.GenTablePageReq
	//获取参数
	err := c.ShouldBind(&req)
	lv_err.HasErrAndPanic(err)
	tableService := service.TableService{}
	rows := make([]menuModel.GenTable, 0)
	result, total, err := tableService.SelectDbTableList(req)
	if err == nil && len(result) > 0 {
		rows = result
	}

	c.JSON(http.StatusOK, lv_dto.TableDataInfo{
		Code:  200,
		Msg:   "操作成功",
		Total: total,
		Rows:  rows,
	})
}

// 导入表结构（保存）
func (w *GenController) ImportTableSave(c *gin.Context) {
	tables := c.PostForm("tables")
	if tables == "" {
		util2.ErrorResp(c).SetBtype(lv_dto.Buniss_Add).SetMsg("参数错误tables未选中").Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
	}
	var userService service.UserService
	user := userService.GetProfile(c)
	if user == nil {
		util2.ErrorResp(c).SetBtype(lv_dto.Buniss_Add).SetMsg("登录超时").Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
	}

	operName := user.LoginName
	tableService := service.TableService{}
	tableArr := strings.Split(tables, ",")
	tableList, err := tableService.SelectDbTableListByNames(tableArr)
	if err != nil {
		util2.ErrorResp(c).SetBtype(lv_dto.Buniss_Add).SetMsg(err.Error()).WriteJsonExit()
		return
	}

	if tableList == nil {
		util2.ErrorResp(c).SetBtype(lv_dto.Buniss_Add).SetMsg("请选择需要导入的表").WriteJsonExit()
		return
	}
	tableService.ImportGenTable(&tableList, operName)
	util2.SucessResp(c).WriteJsonExit()
}

// 根据table_id查询表列数据
func (w *GenController) ColumnList(c *gin.Context) {
	tableId := lv_conv.Int64(c.PostForm("tableId"))
	rows := make([]menuModel.GenTableColumn, 0)
	tableService := service.TableColumnService{}
	result, err := tableService.SelectGenTableColumnListByTableId(tableId)

	if err == nil && len(result) > 0 {
		rows = result
	}

	c.JSON(http.StatusOK, lv_dto.TableDataInfo{
		Code:  200,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}
