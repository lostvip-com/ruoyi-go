package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_web"
)

func (w HisDataController) ToWizard(c *gin.Context) {
	lv_web.BuildTpl(c, "mywork/his_data_wizard").WriteTpl()
}
