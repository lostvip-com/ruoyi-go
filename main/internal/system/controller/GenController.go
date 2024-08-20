package controller

import (
	util2 "common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/lv_batis"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_file"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"html/template"
	"main/internal/common/myconf"
	menuModel "main/internal/system/model"
	"main/internal/system/service"
	"main/internal/system/vo"
	"net/http"
	"os"
	"os/exec"
	"strings"
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

	curDir, _ := os.Getwd()

	//err = lv_db.ExecSqlFile(sqlFile)
	// Loads queries from file
	lvBatis, err := lv_batis.LoadFromFile(curDir + "/tmp/sql/" + po.PackageName + "/" + po.TbName + "_menu.sql")
	// Run queries
	tb := lv_db.GetMasterGorm()
	//cfg := global.GetConfigInstance()
	lvBatis.Exec(tb, "menu")
	menuName := po.FunctionName
	sysmenu := menuModel.SysMenu{}
	sysmenu.MenuName = menuName
	err = sysmenu.FindLastOne()
	lv_err.HasErrAndPanic(err)
	pmenuId := sysmenu.MenuId
	_, err = lvBatis.Exec(tb, "menu_button_create", pmenuId)
	_, err = lvBatis.Exec(tb, "menu_button_retrieve", pmenuId)
	_, err = lvBatis.Exec(tb, "menu_button_update", pmenuId)
	_, err = lvBatis.Exec(tb, "menu_button_delete", pmenuId)
	_, err = lvBatis.Exec(tb, "menu_button_export", pmenuId)
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
	util2.WriteTpl(c, "tool/gen_import_able")
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

	addKey := "vm/html/add.html.vm"
	addValue := ""
	editKey := "vm/html/edit.html.vm"
	editValue := ""

	listKey := "vm/html/list.html.vm"
	listValue := ""
	listTmp := "vm/html/list.txt"

	treeKey := "vm/html/tree.html.vm"
	treeValue := ""

	if entity.TplCategory == "tree" {
		listTmp = "vm/html/list-tree.txt"
	}

	mapperKey := "vm/mapper/" + entity.TbName + "_mapper.tpl.vm"
	mapperValue := ""
	sqlKey := "vm/sql/" + entity.TbName + "_menu.sql.vm"
	sqlValue := ""
	entityKey := "vm/go/" + entity.ClassName + ".go.vm"
	entityValue := ""
	extendKey := "vm/go/" + entity.ClassName + "VO.go.vm"
	extendValue := ""
	daoKey := "vm/go/" + entity.ClassName + "Dao.go.vm"
	daoValue := ""
	serviceKey := "vm/go/" + entity.ClassName + "Service.go.vm"
	serviceValue := ""
	routerKey := "vm/go/" + entity.BusinessName + "_router.go.vm"
	routerValue := ""
	controllerKey := "vm/go/" + entity.ClassName + "Controller.go.vm"
	controllerValue := ""
	apiKey := "vm/go/" + entity.ClassName + "Api.go.vm"
	apiValue := ""

	//新增页面模板
	addValue, _ = tableService.LoadTemplate("vm/html/add.txt", gin.H{"table": entity})
	//修改页面模板
	editValue, _ = tableService.LoadTemplate("vm/html/edit.txt", gin.H{"table": entity})

	//列表页面模板
	listValue, err = tableService.LoadTemplate(listTmp, gin.H{"table": entity})
	if err != nil {
		lv_log.Error("XXXXXXXXXX"+listTmp, err)
	}
	if entity.TplCategory == "tree" {
		//选择树页面模板
		treeValue, _ = tableService.LoadTemplate("vm/html/tree.txt", gin.H{"table": entity})
	}

	//entity模板
	entityValue, _ = tableService.LoadTemplate("vm/go/model.txt", gin.H{"table": entity})

	//extend模板
	extendValue, _ = tableService.LoadTemplate("vm/go/vo.txt", gin.H{"table": entity})

	//service模板
	serviceValue, _ = tableService.LoadTemplate("vm/go/service.txt", gin.H{"table": entity})
	//dao模板
	daoValue, _ = tableService.LoadTemplate("vm/go/dao.txt", gin.H{"table": entity})
	//router模板
	routerValue, _ = tableService.LoadTemplate("vm/go/router.txt", gin.H{"table": entity})

	//controller模板
	controllerValue, _ = tableService.LoadTemplate("vm/go/controller.txt", gin.H{"table": entity})
	//controller模板
	apiValue, _ = tableService.LoadTemplate("vm/go/api.txt", gin.H{"table": entity})
	//sql模板
	sqlValue, _ = tableService.LoadTemplate("vm/sql/sql.txt", gin.H{"table": entity})
	//mapper模板
	mapperValue, _ = tableService.LoadTemplate("vm/mapper/mapper.sql", gin.H{"table": entity})

	if entity.TplCategory == "tree" {
		c.JSON(http.StatusOK, lv_dto.CommonRes{
			Code:  200,
			Btype: lv_dto.Buniss_Other,
			Data: gin.H{
				addKey:        addValue,
				editKey:       editValue,
				listKey:       listValue,
				treeKey:       treeValue,
				sqlKey:        sqlValue,
				entityKey:     entityValue,
				extendKey:     extendValue,
				serviceKey:    serviceValue,
				daoKey:        daoValue,
				routerKey:     routerValue,
				controllerKey: controllerValue,
				apiKey:        apiValue,
				mapperKey:     mapperValue,
			},
		})
	} else {
		c.JSON(http.StatusOK, lv_dto.CommonRes{
			Code:  200,
			Btype: lv_dto.Buniss_Other,
			Data: gin.H{
				addKey:        addValue,
				editKey:       editValue,
				listKey:       listValue,
				sqlKey:        sqlValue,
				entityKey:     entityValue,
				extendKey:     extendValue,
				serviceKey:    serviceValue,
				daoKey:        daoValue,
				routerKey:     routerValue,
				controllerKey: controllerValue,
				apiKey:        apiValue,
				mapperKey:     mapperValue,
			},
		})
	}

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

	listTmp := "vm/html/list.txt"
	if entity.TplCategory == "tree" {
		listTmp = "vm/html/list-tree.txt"
	}

	//获取当前运行时目录
	curDir, err := os.Getwd()

	if err != nil {
		util2.ErrorResp(c).SetMsg(err.Error()).Log("生成代码", gin.H{"tableId": tableId}).WriteJsonExit()
	}
	var htmlPath = curDir + "/template/" + entity.PackageName
	//add模板
	if tmp, err := tableService.LoadTemplate("vm/html/add.txt", gin.H{"table": entity}); err == nil {
		fileName := htmlPath + "/" + entity.BusinessName + "/add.html"

		if canGenIt(overwrite, fileName) {
			f, err := lv_file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}

	//edit模板
	if tmp, err := tableService.LoadTemplate("vm/html/edit.txt", gin.H{"table": entity}); err == nil {
		fileName := htmlPath + "/" + entity.BusinessName + "/edit.html"
		if canGenIt(overwrite, fileName) {
			f, err := lv_file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}

	//list模板
	if tmp, err := tableService.LoadTemplate(listTmp, gin.H{"table": entity}); err == nil {
		fileName := htmlPath + "/" + entity.BusinessName + "/list.html"
		if canGenIt(overwrite, fileName) {
			f, err := lv_file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}

	if entity.TplCategory == "tree" {
		//tree模板
		if tmp, err := tableService.LoadTemplate("vm/html/tree.txt", gin.H{"table": entity}); err == nil {
			fileName := htmlPath + "/" + entity.BusinessName + "/tree.html"
			if canGenIt(overwrite, fileName) {
				f, err := lv_file.Create(fileName)
				if err == nil {
					f.WriteString(tmp)
				}
				f.Close()
			}
		}
	}
	var packagePath = curDir + "/internal/" + entity.PackageName
	//entity模板
	if tmp, err := tableService.LoadTemplate("vm/go/model.txt", gin.H{"table": entity}); err == nil {
		fileName := packagePath + "/model/" + entity.ClassName + ".go"
		if canGenIt(overwrite, fileName) {
			f, err := lv_file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}

	//dao模板
	if tmp, err := tableService.LoadTemplate("vm/go/dao.txt", gin.H{"table": entity}); err == nil {
		fileName := packagePath + "/dao/" + entity.ClassName + "Dao.go"
		if canGenIt(overwrite, fileName) {
			f, err := lv_file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}
	//vo模板
	if tmp, err := tableService.LoadTemplate("vm/go/vo.txt", gin.H{"table": entity}); err == nil {
		fileName := packagePath + "/vo/" + entity.ClassName + "VO.go"
		if canGenIt(overwrite, fileName) {
			f, err := lv_file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}
	//service模板
	if tmp, err := tableService.LoadTemplate("vm/go/service.txt", gin.H{"table": entity}); err == nil {
		fileName := packagePath + "/service/" + entity.ClassName + "Service.go"
		if canGenIt(overwrite, fileName) {
			f, err := lv_file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}
	//controller模板
	if tmp, err := tableService.LoadTemplate("vm/go/controller.txt", gin.H{"table": entity}); err == nil {
		fileName := packagePath + "/controller/" + entity.ClassName + "Controller.go"
		if canGenIt(overwrite, fileName) {
			f, err := lv_file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}
	//controller模板
	if tmp, err := tableService.LoadTemplate("vm/go/api.txt", gin.H{"table": entity}); err == nil {
		fileName := packagePath + "/controller/" + entity.ClassName + "Api.go"
		if canGenIt(overwrite, fileName) {
			f, err := lv_file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}

	//router模板
	if tmp, err := tableService.LoadTemplate("vm/go/router.txt", gin.H{"table": entity}); err == nil {
		fileName := packagePath + "/" + entity.TbName + "_router.go"
		if canGenIt(overwrite, fileName) {
			f, err := lv_file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}
	//ibatis sql文件
	if mapper, err := tableService.LoadTemplate("vm/mapper/mapper.sql", gin.H{"table": entity}); err == nil {
		mapperPath := curDir + "/mapper"
		fileName := strings.Join([]string{mapperPath, "/", entity.PackageName, "/", entity.TbName, "_mapper.sql"}, "")
		if canGenIt(overwrite, fileName) {
			f, err := lv_file.Create(fileName)
			if err == nil {
				f.WriteString(mapper)
			}
			f.Close()
		}
	}
	//sql模板
	if tmp, err := tableService.LoadTemplate("vm/sql/sql.txt", gin.H{"table": entity}); err == nil {
		tmpPath := curDir + "/tmp/sql"
		fileName := strings.Join([]string{tmpPath, "/", entity.PackageName, "/", entity.TbName, "_menu.sql"}, "")
		if canGenIt(overwrite, fileName) {
			f, err := lv_file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}
	util2.SucessResp(c).Log("生成代码", gin.H{"tableId": tableId}).WriteJsonExit()
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
