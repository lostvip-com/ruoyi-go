package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
)

func (w HisDataController) ToWizard(c *gin.Context) {
	util.BuildTpl(c, "mywork/his_data_wizard").WriteTpl()
}
