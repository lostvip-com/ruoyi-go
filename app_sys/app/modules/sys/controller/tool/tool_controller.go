package tool

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/response"
	"net/http"
	"os"
	"os/exec"
	"robvi/app/modules/sys/model"
)

// 表单构建
func Build(c *gin.Context) {
	response.BuildTpl(c, "tool/build").WriteTpl()
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

		genPath := curDir + "/public/swagger"
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
