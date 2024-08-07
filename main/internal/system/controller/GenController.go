package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/db"
	"github.com/lostvip-com/lv_framework/db/lvbatis"
	"github.com/lostvip-com/lv_framework/logme"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_file"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	"github.com/lostvip-com/lv_framework/web/dto"
	"html/template"
	"main/internal/common/myconf"
	menuModel "main/internal/system/model"
	tool3 "main/internal/system/model/tool"
	"main/internal/system/service"
	"main/internal/system/service/tool"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type GenController struct{}

// 表单构建
func (w *GenController) Build(c *gin.Context) {
	lv_web.BuildTpl(c, "tool/build").WriteTpl()
}

// swagger文档
func (w *GenController) ExecSqlFile(c *gin.Context) {
	tableId := lv_conv.Int64(c.Query("tableId"))
	//获取参数
	if tableId <= 0 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg("参数错误").Log("执行SQL文件错误", gin.H{"tableId": tableId})
	}
	genTable := tool3.GenTable{}
	po, err := genTable.SelectGenTableById(tableId)
	if err != nil {
		panic(err.Error())
	}

	curDir, _ := os.Getwd()

	//err = db.ExecSqlFile(sqlFile)
	// Loads queries from file
	lvBatis, err := lvbatis.LoadFromFile(curDir + "/tmp/sql/" + po.PackageName + "/" + po.TbName + "_menu.sql")
	// Run queries
	tb := db.GetMasterGorm()
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
	lv_web.SucessData(c, nil)
}

// swagger文档
func (w *GenController) Swagger(c *gin.Context) {
	a := c.Query("a")
	if a == "r" {
		//重新生成文档
		curDir, err := os.Getwd()
		if err != nil {
			lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
				"desc": "参数错误",
			})
			c.Abort()
			return
		}
		genPath := curDir + "/static/swagger"
		err = w.generateSwaggerFiles(genPath)
		if err != nil {
			lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
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
	lv_web.BuildTpl(c, "tool/gen/list").WriteTpl()
}

func (w *GenController) GenList(c *gin.Context) {
	var req *tool3.SelectPageReq
	tableService := tool.TableService{}
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("生成代码", req).WriteJsonExit()
		return
	}
	rows := make([]tool3.GenTable, 0)
	result, page, err := tableService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	lv_web.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 导入数据表
func (w *GenController) ImportTable(c *gin.Context) {
	lv_web.BuildTpl(c, "tool/gen/importTable").WriteTpl()
}

// 删除数据
func (w *GenController) Remove(c *gin.Context) {
	var req *dto.IdsReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).Log("生成代码", req).WriteJsonExit()
		return
	}
	tableService := tool.TableService{}
	rs := tableService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(dto.Buniss_Del).Log("生成代码", req).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Del).Log("生成代码", req).WriteJsonExit()
	}
}

// 修改数据
func (w *GenController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	tableService := tool.TableService{}
	entity, err := tableService.SelectRecordById(id)
	if err != nil || entity == nil {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数不存在",
		})
		return
	}
	goTypeTpl := tableService.GoTypeTpl()
	queryTypeTpl := tableService.QueryTypeTpl()
	htmlTypeTpl := tableService.HtmlTypeTpl()

	lv_web.BuildTpl(c, "tool/gen/edit").WriteTpl(gin.H{
		"table":        entity,
		"goTypeTpl":    template.HTML(goTypeTpl),
		"queryTypeTpl": template.HTML(queryTypeTpl),
		"htmlTypeTpl":  template.HTML(htmlTypeTpl),
	})
}

// 修改数据保存
func (w *GenController) EditSave(c *gin.Context) {
	var req tool3.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).SetBtype(dto.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
		return
	}
	tableService := tool.TableService{}
	_, err := tableService.SaveEdit(&req, c)
	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).SetBtype(dto.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetBtype(dto.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
}

// 预览代码
func (w *GenController) Preview(c *gin.Context) {
	tableId := lv_conv.Int64(c.Query("tableId"))
	if tableId <= 0 {
		c.JSON(http.StatusOK, dto.CommonRes{
			Code:  500,
			Btype: dto.Buniss_Other,
			Msg:   "参数错误",
		})
	}
	tableService := tool.TableService{}
	entity, err := tableService.SelectRecordById(tableId)

	if err != nil || entity == nil {
		c.JSON(http.StatusOK, dto.CommonRes{
			Code:  500,
			Btype: dto.Buniss_Other,
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
	listValue, _ = tableService.LoadTemplate(listTmp, gin.H{"table": entity})

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
		c.JSON(http.StatusOK, dto.CommonRes{
			Code:  200,
			Btype: dto.Buniss_Other,
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
		c.JSON(http.StatusOK, dto.CommonRes{
			Code:  200,
			Btype: dto.Buniss_Other,
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
	if tableId <= 0 {
		c.JSON(http.StatusOK, dto.CommonRes{
			Code:  500,
			Btype: dto.Buniss_Other,
			Msg:   "参数错误",
		})
	}
	tableService := tool.TableService{}
	entity, err := tableService.SelectRecordById(tableId)

	if err != nil || entity == nil {
		c.JSON(http.StatusOK, dto.CommonRes{
			Code:  500,
			Btype: dto.Buniss_Other,
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
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("生成代码", gin.H{"tableId": tableId}).WriteJsonExit()
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
	lv_web.SucessResp(c).Log("生成代码", gin.H{"tableId": tableId}).WriteJsonExit()
}

func canGenIt(overwrite bool, file string) bool {
	if overwrite { //允许覆盖
		logme.Warn("--------->您配置了 overwrite 开关的值为true，旧文件会被覆盖！！！！ ")
		return true
	} else { // 不允许覆盖
		if lv_file.Exists(file) { //文件已经存在，不允许重新生成
			logme.Warn("=======> 文件已经存在，本次将不会生成新文件！！！！！！！！！！！！ ")
			return false
		} else { //文件不存在，允许重新生成
			return true
		}
	}
}

// 查询数据库列表
func (w *GenController) DataList(c *gin.Context) {
	var req *tool3.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("生成代码", req).WriteJsonExit()
	}
	tableService := tool.TableService{}
	rows := make([]tool3.GenTable, 0)
	result, page, err := tableService.SelectDbTableList(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	c.JSON(http.StatusOK, dto.TableDataInfo{
		Code:  200,
		Msg:   "操作成功",
		Total: page.Total,
		Rows:  rows,
	})
}

// 导入表结构（保存）
func (w *GenController) ImportTableSave(c *gin.Context) {
	tables := c.PostForm("tables")
	if tables == "" {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg("参数错误tables未选中").Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
	}
	var userService service.UserService
	user := userService.GetProfile(c)
	if user == nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg("登录超时").Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
	}

	operName := user.LoginName
	tableService := tool.TableService{}
	tableArr := strings.Split(tables, ",")
	tableList, err := tableService.SelectDbTableListByNames(tableArr)
	if err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg(err.Error()).WriteJsonExit()
		return
	}

	if tableList == nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg("请选择需要导入的表").WriteJsonExit()
		return
	}
	tableService.ImportGenTable(&tableList, operName)
	lv_web.SucessResp(c).WriteJsonExit()
}

// 根据table_id查询表列数据
func (w *GenController) ColumnList(c *gin.Context) {
	tableId := lv_conv.Int64(c.PostForm("tableId"))
	//获取参数
	if tableId <= 0 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg("参数错误").Log("生成代码", gin.H{"tableId": tableId})
	}
	rows := make([]tool3.Entity, 0)
	tableService := tool.TableService{}
	result, err := tableService.SelectGenTableColumnListByTableId(tableId)

	if err == nil && len(result) > 0 {
		rows = result
	}

	c.JSON(http.StatusOK, dto.TableDataInfo{
		Code:  200,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}
