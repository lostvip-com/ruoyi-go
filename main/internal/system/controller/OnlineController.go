package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"main/internal/system/model"
	"main/internal/system/service"
	"main/internal/system/vo"
	"strings"
)

type OnlineController struct {
}

// 列表页
func (w *OnlineController) List(c *gin.Context) {
	util.WriteTpl(c, "monitor/list_online")
}

// 列表分页数据
func (w *OnlineController) ListAjax(c *gin.Context) {
	var param vo.OnlinePageReq
	//获取参数
	err := c.ShouldBind(&param)
	lv_err.HasErrAndPanic(err)
	rows := make([]model.SysUserOnline, 0)
	db := lv_db.GetMasterGorm()
	tb := db.Table("sys_user_online t")
	if param.SessionId != "" {
		tb.Where("t.session_id = ?", param.SessionId)
	}

	if param.LoginName != "" {
		tb.Where("t.login_name like ?", "%"+param.LoginName+"%")
	}

	if param.DeptName != "" {
		tb.Where("t.dept_name like ?", "%"+param.DeptName+"%")
	}

	if param.Ipaddr != "" {
		tb.Where("t.ipaddr = ?", param.Ipaddr)
	}

	if param.LoginLocation != "" {
		tb.Where("t.login_location = ?", param.LoginLocation)
	}

	if param.Browser != "" {
		tb.Where("t.browser = ?", param.Browser)
	}

	if param.Os != "" {
		tb.Where("t.os = ?", param.Os)
	}

	if param.Status != "" {
		tb.Where("t.status = ?", param.Status)
	}

	if param.BeginTime != "" {
		tb.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
	}

	if param.EndTime != "" {
		tb.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
	}
	var total int64
	tb = tb.Count(&total)
	tb.Order("t.login_time desc").Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize).Find(&rows)
	err = tb.Error
	lv_err.HasErrAndPanic(err)
	util.SucessPage(c, rows, total)
}

// 用户强退
func (w *OnlineController) ForceLogout(c *gin.Context) {
	sessionId := c.PostForm("sessionId")
	var userService service.SessionService
	err := userService.ForceLogout(sessionId)
	lv_err.HasErrAndPanic(err)
	util.Success(c, gin.H{"sessionId": sessionId}, "success")
}

// 批量强退
func (w *OnlineController) BatchForceLogout(c *gin.Context) {
	var userService service.SessionService
	ids := c.Query("ids")
	if ids == "" {
		util.ErrorResp(c).SetMsg("参数错误").Log("批量强退", gin.H{"ids": ids}).WriteJsonExit()
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
	util.SucessResp(c).Log("批量强退", gin.H{"ids": ids}).WriteJsonExit()
}
