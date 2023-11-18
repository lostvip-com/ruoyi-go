package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/response"
)

func (w HisDataController) ToWizard(c *gin.Context) {
	response.BuildTpl(c, "modules/biz/his_data/his_data_wizard").WriteTpl()
}
