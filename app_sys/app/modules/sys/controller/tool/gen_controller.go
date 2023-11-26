package tool

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"lostvip.com/utils/gconv"
	"lostvip.com/utils/lib_file"
	response2 "lostvip.com/utils/response"
	"net/http"
	"os"
	"robvi/app/modules/sys/model"
	tool2 "robvi/app/modules/sys/model/tool"
	"robvi/app/modules/sys/service"
	"robvi/app/modules/sys/service/tool"
	"strings"
)

// 生成代码列表页面
func Gen(c *gin.Context) {
	response2.BuildTpl(c, "tool/gen/list").WriteTpl()
}

func GenList(c *gin.Context) {
	var req *tool2.SelectPageReq
	tableService := tool.TableService{}
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetMsg(err.Error()).Log("生成代码", req).WriteJsonExit()
		return
	}
	rows := make([]tool2.GenTable, 0)
	result, page, err := tableService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response2.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 导入数据表
func ImportTable(c *gin.Context) {
	response2.BuildTpl(c, "tool/gen/importTable").WriteTpl()
}

// 删除数据
func Remove(c *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Del).Log("生成代码", req).WriteJsonExit()
		return
	}
	tableService := tool.TableService{}
	rs := tableService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response2.SucessResp(c).SetBtype(model.Buniss_Del).Log("生成代码", req).WriteJsonExit()
	} else {
		response2.ErrorResp(c).SetBtype(model.Buniss_Del).Log("生成代码", req).WriteJsonExit()
	}
}

// 修改数据
func Edit(c *gin.Context) {
	id := gconv.Int64(c.Query("id"))
	if id <= 0 {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	tableService := tool.TableService{}
	entity, err := tableService.SelectRecordById(id)
	if err != nil || entity == nil {
		response2.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数不存在",
		})
		return
	}
	goTypeTpl := tableService.GoTypeTpl()
	queryTypeTpl := tableService.QueryTypeTpl()
	htmlTypeTpl := tableService.HtmlTypeTpl()

	response2.BuildTpl(c, "tool/gen/edit").WriteTpl(gin.H{
		"table":        entity,
		"goTypeTpl":    template.HTML(goTypeTpl),
		"queryTypeTpl": template.HTML(queryTypeTpl),
		"htmlTypeTpl":  template.HTML(htmlTypeTpl),
	})
}

// 修改数据保存
func EditSave(c *gin.Context) {
	var req tool2.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetMsg(err.Error()).SetBtype(model.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
		return
	}
	tableService := tool.TableService{}
	_, err := tableService.SaveEdit(&req, c)
	if err != nil {
		response2.ErrorResp(c).SetMsg(err.Error()).SetBtype(model.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
		return
	}
	response2.SucessResp(c).SetBtype(model.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
}

// 预览代码
func Preview(c *gin.Context) {
	tableId := gconv.Int64(c.Query("tableId"))
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

	sqlKey := "vm/sql/menu.sql.vm"
	sqlValue := ""
	entityKey := "vm/go/" + entity.BusinessName + "_entity.go.vm"
	entityValue := ""
	extendKey := "vm/go/" + entity.BusinessName + ".go.vm"
	extendValue := ""
	serviceKey := "vm/go/" + entity.BusinessName + "_service.go.vm"
	serviceValue := ""
	routerKey := "vm/go/" + entity.BusinessName + "_router.go.vm"
	routerValue := ""
	controllerKey := "vm/go/" + entity.BusinessName + "_controller.go.vm"
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
func GenCode(c *gin.Context) {
	tableId := gconv.Int64(c.Query("tableId"))
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
		response2.ErrorResp(c).SetMsg(err.Error()).Log("生成代码", gin.H{"tableId": tableId}).WriteJsonExit()
	}
	var htmlMoudlePath = curDir + "/template/modules/" + entity.ModuleName
	//add模板
	if tmp, err := tableService.LoadTemplate("vm/html/add.txt", gin.H{"table": entity}); err == nil {
		fileName := htmlMoudlePath + "/" + entity.BusinessName + "/add.html"

		//if !file.Exists(fileName) { //改为直接覆盖
		f, err := lib_file.Create(fileName)
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
		f, err := lib_file.Create(fileName)
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
		f, err := lib_file.Create(fileName)
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
			f, err := lib_file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
			//}
		}
	}
	var goModulePath = curDir + "/app/modules/" + entity.ModuleName
	//entity模板
	if tmp, err := tableService.LoadTemplate("vm/go/entity.txt", gin.H{"table": entity}); err == nil {
		fileName := goModulePath + "/model/" + entity.ClassName + ".go"
		if lib_file.Exists(fileName) {
			os.RemoveAll(fileName)
		}

		f, err := lib_file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
	}

	//extend模板
	if tmp, err := tableService.LoadTemplate("vm/go/extend.txt", gin.H{"table": entity}); err == nil {
		fileName := goModulePath + "/model/" + entity.ClassName + "DTO.go"
		//if !file.Exists(fileName) {//改为直接覆盖
		f, err := lib_file.Create(fileName)
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
		f, err := lib_file.Create(fileName)
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
		f, err := lib_file.Create(fileName)
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
		f, err := lib_file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}

	//sql模板
	if tmp, err := tableService.LoadTemplate("vm/sql/sql.txt", gin.H{"table": entity}); err == nil {
		tmpPath := curDir + "/tmp/sql"
		if !lib_file.Exists(tmpPath) {
			lib_file.Mkdir(tmpPath)
		}
		fileName := strings.Join([]string{tmpPath, "/", entity.ModuleName, "/", entity.TbName, "_menu.sql"}, "")

		//if !file.Exists(fileName) {
		f, err := lib_file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}
	response2.SucessResp(c).Log("生成代码", gin.H{"tableId": tableId}).WriteJsonExit()
}

// 查询数据库列表
func DataList(c *gin.Context) {
	var req *tool2.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response2.ErrorResp(c).SetMsg(err.Error()).Log("生成代码", req).WriteJsonExit()
	}
	tableService := tool.TableService{}
	rows := make([]tool2.GenTable, 0)
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
func ImportTableSave(c *gin.Context) {
	tables := c.PostForm("tables")
	if tables == "" {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("参数错误tables未选中").Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
	}
	var userService service.UserService
	user := userService.GetProfile(c)
	if user == nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("登陆超时").Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
	}

	operName := user.LoginName
	tableService := tool.TableService{}
	tableArr := strings.Split(tables, ",")
	tableList, err := tableService.SelectDbTableListByNames(tableArr)
	if err != nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
		return
	}

	if tableList == nil {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("请选择需要导入的表").Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
		return
	}

	tableService.ImportGenTable(&tableList, operName)
	response2.SucessResp(c).Log("导入表结构", gin.H{"tables": tables}).WriteJsonExit()
}

// 根据table_id查询表列数据
func ColumnList(c *gin.Context) {
	tableId := gconv.Int64(c.PostForm("tableId"))
	//获取参数
	if tableId <= 0 {
		response2.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("参数错误").Log("生成代码", gin.H{"tableId": tableId})
	}
	rows := make([]tool2.Entity, 0)
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
