package tool

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/db"
	"lostvip.com/db/ibatis"
	"lostvip.com/utils/gconv"
	"lostvip.com/utils/lib_file"
	"lostvip.com/utils/response"
	"net/http"
	"os"
	"os/exec"
	"robvi/app/modules/sys/model"
	menuModel "robvi/app/modules/sys/model/system/menu"
	tool2 "robvi/app/modules/sys/model/tool"
)

// 表单构建
func Build(c *gin.Context) {
	response.BuildTpl(c, "tool/build").WriteTpl()
}

// swagger文档
func ExecSqlFile(c *gin.Context) {
	tableId := gconv.Int64(c.Query("tableId"))
	//获取参数
	if tableId <= 0 {
		response.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("参数错误").Log("执行SQL文件错误", gin.H{"tableId": tableId})
	}
	tb := tool2.GenTable{}
	po, err := tb.SelectGenTableById(tableId)
	if err != nil {
		panic(err.Error())
	}

	curDir, _ := os.Getwd()
	sqlFile := curDir + "/tmp/sql/" + po.ModuleName + "/" + po.TbName + "_menu.sql"

	if !lib_file.FileExist(sqlFile) {
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
	response.Sucess(c, nil)
}

// swagger文档
func Swagger(c *gin.Context) {
	a := c.Query("a")
	if a == "r" {
		//重新生成文档
		curDir, err := os.Getwd()
		if err != nil {
			response.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
				"desc": "参数错误",
			})
			c.Abort()
			return
		}
		genPath := curDir + "/static/swagger"
		err = generateSwaggerFiles(genPath)
		if err != nil {
			response.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
				"desc": "参数错误",
			})
			c.Abort()
			return
		}
	}
	c.Redirect(http.StatusFound, "/static/swagger/index.html")
}

// 自动生成文档 swag init -o /Volumes/File/WorkSpaces/app-yjzx/public/swagger
func generateSwaggerFiles(output string) error {

	cmd := exec.Command("swag", "init -o "+output)
	// 保证关闭输出流
	if err := cmd.Start(); err != nil { // 运行命令
		return err
	}

	return nil
}
