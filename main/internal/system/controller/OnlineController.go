package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	"main/internal/system/model/monitor/online"
	"main/internal/system/service"
	onlineService "main/internal/system/service/monitor/online"
	"strings"
)

type OnlineController struct {
}

// 列表页
func (w *OnlineController) List(c *gin.Context) {
	sessinIdArr := make([]string, 0)

	//global.SessionList.Range(func(k, v interface{}) bool {
	//	return true
	//})
	if len(sessinIdArr) > 0 {
		onlineService.DeleteRecordNotInIds(sessinIdArr)
	}

	lv_web.BuildTpl(c, "monitor/online/list").WriteTpl()
}

// 列表分页数据
func (w *OnlineController) ListAjax(c *gin.Context) {
	var req *online.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	rows := make([]online.UserOnline, 0)
	result, page, err := onlineService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	lv_web.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 用户强退
func (w *OnlineController) ForceLogout(c *gin.Context) {
	sessionId := c.PostForm("sessionId")
	if sessionId == "" {
		lv_web.ErrorResp(c).SetMsg("参数错误").Log("用户强退", gin.H{"sessionId": sessionId}).WriteJsonExit()
		return
	}
	var userService service.SessionService
	err := userService.ForceLogout(sessionId)
	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("用户强退", gin.H{"sessionId": sessionId}).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).Log("用户强退", gin.H{"sessionId": sessionId}).WriteJsonExit()
}

// 批量强退
func (w *OnlineController) BatchForceLogout(c *gin.Context) {
	var userService service.SessionService
	ids := c.Query("ids")
	if ids == "" {
		lv_web.ErrorResp(c).SetMsg("参数错误").Log("批量强退", gin.H{"ids": ids}).WriteJsonExit()
		return
	}
	ids = strings.ReplaceAll(ids, "[", "")
	ids = strings.ReplaceAll(ids, "]", "")
	ids = strings.ReplaceAll(ids, `"`, "")
	idarr := strings.Split(ids, ",")
	if len(idarr) > 0 {
		for _, sessionId := range idarr {
			if sessionId != "" {
				userService.ForceLogout(sessionId)
			}
		}
	}
	lv_web.SucessResp(c).Log("批量强退", gin.H{"ids": ids}).WriteJsonExit()
}
