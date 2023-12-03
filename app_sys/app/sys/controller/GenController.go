package controller

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"lostvip.com/db"
	"lostvip.com/db/ibatis"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_file"
	"lostvip.com/utils/lv_web"
	"net/http"
	"os"
	"os/exec"
	"robvi/app/sys/model"
	menuModel "robvi/app/sys/model/system/menu"
	tool3 "robvi/app/sys/model/tool"
	"robvi/app/sys/service"
	"robvi/app/sys/service/tool"
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
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("参数错误").Log("执行SQL文件错误", gin.H{"tableId": tableId})
	}
	tb := tool3.GenTable{}
	po, err := tb.SelectGenTableById(tableId)
	if err != nil {
		panic(err.Error())
	}

	curDir, _ := os.Getwd()
	sqlFile := curDir + "/tmp/sql/" + po.ModuleName + "/" + po.TbName + "_menu.sql"

	if !lv_file.IsFileExist(sqlFile) {
		panic("生成代码后再执行此操作")
	}
	//err = db.ExecSqlFile(sqlFile)
	// Loads queries from file
	dot, err := ibatis.LoadFromFile(sqlFile)
	// Run queries
	_, err = dot.Exec(db.GetInstance().Engine().DB(), "menu")
	menuName := po.FunctionName
	sysmenu := menuModel.SysMenu{}
	sysmenu.MenuName = menuName
	_, err = sysmenu.FindOne()
	pmenuId := sysmenu.MenuId
	_, err = dot.Exec(db.GetInstance().Engine().DB(), "menu_button_create", pmenuId)
	_, err = dot.Exec(db.GetInstance().Engine().DB(), "menu_button_retrieve", pmenuId)
	_, err = dot.Exec(db.GetInstance().Engine().DB(), "menu_button_update", pmenuId)
	_, err = dot.Exec(db.GetInstance().Engine().DB(), "menu_button_delete", pmenuId)
	_, err = dot.Exec(db.GetInstance().Engine().DB(), "menu_button_export", pmenuId)
	if err != nil {
		panic(err)
	}
	lv_web.Sucess(c, nil)
}

// swagger文档
func (w *GenController) Swagger(c *gin.Context) {
	a := c.Query("a")
	if a == "r" {
		//重新生成文档
		curDir, err := os.Getwd()
		if err != nil {
			lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
				"desc": "参数错误",
			})
			c.Abort()
			return
		}
		genPath := curDir + "/static/swagger"
		err = w.generateSwaggerFiles(genPath)
		if err != nil {
			lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
				"desc": "参数错误",
			})
			c.Abort()
			return
		}
	}
	c.Redirect(http.StatusFound, "/static/swagger/index.html")
}

// 自动生成文档 swag init -o /Volumes/File/WorkSpaces/app-yjzx/public/swagger
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
	var req *model.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Del).Log("生成代码", req).WriteJsonExit()
		return
	}
	tableService := tool.TableService{}
	rs := tableService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(model.Buniss_Del).Log("生成代码", req).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Del).Log("生成代码", req).WriteJsonExit()
	}
}

// 修改数据
func (w *GenController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	tableService := tool.TableService{}
	entity, err := tableService.SelectRecordById(id)
	if err != nil || entity == nil {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
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
		lv_web.ErrorResp(c).SetMsg(err.Error()).SetBtype(model.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
		return
	}
	tableService := tool.TableService{}
	_, err := tableService.SaveEdit(&req, c)
	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).SetBtype(model.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetBtype(model.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
}

// 预览代码
func (w *GenController) Preview(c *gin.Context) {
	tableId := lv_conv.Int64(c.Query("tableId"))
	if tableId <= 0 {
		c.JSON(http.StatusOK, model.CommonRes{
			Code:  500,
			Btype: model.Buniss_Other,
			Msg:   "参数错误",
		})
	}
	tableService := tool.TableService{}
	entity, err := tableService.SelectRecordById(tableId)

	if err != nil || entity == nil {
		c.JSON(http.StatusOK, model.CommonRes{
			Code:  500,
			Btype: model.Buniss_Other,
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

	sqlKey := "vm/sql/" + entity.TbName + "_menu.sql.vm"
	sqlValue := ""
	entityKey := "vm/go/" + entity.ClassName + ".go.vm"
	entityValue := ""
	extendKey := "vm/go/" + entity.ClassName + "DTO.go.vm"
	extendValue := ""
	serviceKey := "vm/go/" + entity.ClassName + "Service.go.vm"
	serviceValue := ""
	routerKey := "vm/go/" + entity.BusinessName + "_router.go.vm"
	routerValue := ""
	controllerKey := "vm/go/" + entity.ClassName + "Controller.go.vm"
	controllerValue := ""

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
	entityValue, _ = tableService.LoadTemplate("vm/go/entity.txt", gin.H{"table": entity})

	//extend模板
	extendValue, _ = tableService.LoadTemplate("vm/go/extend.txt", gin.H{"table": entity})

	//service模板
	serviceValue, _ = tableService.LoadTemplate("vm/go/service.txt", gin.H{"table": entity})

	//router模板
	routerValue, _ = tableService.LoadTemplate("vm/go/router.txt", gin.H{"table": entity})

	//controller模板
	controllerValue, _ = tableService.LoadTemplate("vm/go/controller.txt", gin.H{"table": entity})

	//sql模板
	sqlValue, _ = tableService.LoadTemplate("vm/sql/sql.txt", gin.H{"table": entity})

	if entity.TplCategory == "tree" {
		c.JSON(http.StatusOK, model.CommonRes{
			Code:  200,
			Btype: model.Buniss_Other,
			Data: gin.H{
				addKey:        addValue,
				editKey:       editValue,
				listKey:       listValue,
				treeKey:       treeValue,
				sqlKey:        sqlValue,
				entityKey:     entityValue,
				extendKey:     extendValue,
				serviceKey:    serviceValue,
				routerKey:     routerValue,
				controllerKey: controllerValue,
			},
		})
	} else {
		c.JSON(http.StatusOK, model.CommonRes{
			Code:  200,
			Btype: model.Buniss_Other,
			Data: gin.H{
				addKey:        addValue,
				editKey:       editValue,
				listKey:       listValue,
				sqlKey:        sqlValue,
				entityKey:     entityValue,
				extendKey:     extendValue,
				serviceKey:    serviceValue,
				routerKey:     routerValue,
				controllerKey: controllerValue,
			},
		})
	}

}

// 生成代码
func (w *GenController) GenCode(c *gin.Context) {
	tableId := lv_conv.Int64(c.Query("tableId"))
	if tableId <= 0 {
		c.JSON(http.StatusOK, model.CommonRes{
			Code:  500,
			Btype: model.Buniss_Other,
			Msg:   "参数错误",
		})
	}
	tableService := tool.TableService{}
	entity, err := tableService.SelectRecordById(tableId)

	if err != nil || entity == nil {
		c.JSON(http.StatusOK, model.CommonRes{
			Code:  500,
			Btype: model.Buniss_Other,
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
	var htmlMoudlePath = curDir + "/template/" + entity.ModuleName
	//add模板
	if tmp, err := tableService.LoadTemplate("vm/html/add.txt", gin.H{"table": entity}); err == nil {
		fileName := htmlMoudlePath + "/" + entity.BusinessName + "/add.html"

		//if !file.Exists(fileName) { //改为直接覆盖
		f, err := lv_file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}

	//edit模板
	if tmp, err := tableService.LoadTemplate("vm/html/edit.txt", gin.H{"table": entity}); err == nil {
		fileName := htmlMoudlePath + "/" + entity.BusinessName + "/edit.html"
		//if !file.Exists(fileName) { //改为直接覆盖
		f, err := lv_file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}

	//list模板
	if tmp, err := tableService.LoadTemplate(listTmp, gin.H{"table": entity}); err == nil {
		fileName := htmlMoudlePath + "/" + entity.BusinessName + "/list.html"

		//if !file.Exists(fileName) {//改为直接覆盖
		f, err := lv_file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}

	if entity.TplCategory == "tree" {
		//tree模板
		if tmp, err := tableService.LoadTemplate("vm/html/tree.txt", gin.H{"table": entity}); err == nil {
			fileName := htmlMoudlePath + "/" + entity.BusinessName + "/tree.html"

			//if !file.Exists(fileName) {
			f, err := lv_file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
			//}
		}
	}
	var goModulePath = curDir + "/app/" + entity.ModuleName
	//entity模板
	if tmp, err := tableService.LoadTemplate("vm/go/entity.txt", gin.H{"table": entity}); err == nil {
		fileName := goModulePath + "/model/" + entity.ClassName + ".go"
		if lv_file.Exists(fileName) {
			os.RemoveAll(fileName)
		}

		f, err := lv_file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
	}

	//extend模板
	if tmp, err := tableService.LoadTemplate("vm/go/extend.txt", gin.H{"table": entity}); err == nil {
		fileName := goModulePath + "/model/" + entity.ClassName + "DTO.go"
		//if !file.Exists(fileName) {//改为直接覆盖
		f, err := lv_file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}

	//service模板
	if tmp, err := tableService.LoadTemplate("vm/go/service.txt", gin.H{"table": entity}); err == nil {
		fileName := goModulePath + "/service/" + entity.ClassName + "Service.go"

		//if !file.Exists(fileName) {
		f, err := lv_file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}
	//controller模板
	if tmp, err := tableService.LoadTemplate("vm/go/controller.txt", gin.H{"table": entity}); err == nil {
		fileName := goModulePath + "/controller/" + entity.ClassName + "Controller.go"

		//if !file.Exists(fileName) {
		f, err := lv_file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}
	//router模板
	if tmp, err := tableService.LoadTemplate("vm/go/router.txt", gin.H{"table": entity}); err == nil {
		fileName := goModulePath + "/" + entity.TbName + "_router.go"

		//if !file.Exists(fileName) {
		f, err := lv_file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}

	//sql模板
	if tmp, err := tableService.LoadTemplate("vm/sql/sql.txt", gin.H{"table": entity}); err == nil {
		tmpPath := curDir + "/tmp/sql"
		if !lv_file.Exists(tmpPath) {
			lv_file.Mkdir(tmpPath)
		}
		fileName := strings.Join([]string{tmpPath, "/", entity.ModuleName, "/", entity.TbName, "_menu.sql"}, "")

		//if !file.Exists(fileName) {
		f, err := lv_file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}
	lv_web.SucessResp(c).Log("生成代码", gin.H{"tableId": tableId}).WriteJsonExit()
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

	c.JSON(http.StatusOK, model.TableDataInfo{
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
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("参数错误tables未选中").Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
	}
	var userService service.UserService
	user := userService.GetProfile(c)
	if user == nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("登陆超时").Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
	}

	operName := user.LoginName
	tableService := tool.TableService{}
	tableArr := strings.Split(tables, ",")
	tableList, err := tableService.SelectDbTableListByNames(tableArr)
	if err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
		return
	}

	if tableList == nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("请选择需要导入的表").Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
		return
	}

	tableService.ImportGenTable(&tableList, operName)
	lv_web.SucessResp(c).Log("导入表结构", gin.H{"tables": tables}).WriteJsonExit()
}

// 根据table_id查询表列数据
func (w *GenController) ColumnList(c *gin.Context) {
	tableId := lv_conv.Int64(c.PostForm("tableId"))
	//获取参数
	if tableId <= 0 {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("参数错误").Log("生成代码", gin.H{"tableId": tableId})
	}
	rows := make([]tool3.Entity, 0)
	tableService := tool.TableService{}
	result, err := tableService.SelectGenTableColumnListByTableId(tableId)

	if err == nil && len(result) > 0 {
		rows = result
	}

	c.JSON(http.StatusOK, model.TableDataInfo{
		Code:  200,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}
