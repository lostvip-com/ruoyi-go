package system

import (
	"github.com/lostvip-com/lv_framework/db"
	"github.com/lostvip-com/lv_framework/logme"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"main/internal/common/myconf"
	"main/internal/system/model"
	"main/internal/system/model/monitor/online"
)

//自动建表

func init() {

	//自动同步表结构
	cfg := myconf.GetConfigInstance()
	migrate := cfg.GetAutoMigrate()
	if migrate == "create" || migrate == "update" || migrate == "true" {
		logme.Warn("######### 开始同步表结构: ############## migrate" + migrate)
		err := db.GetMasterGorm().AutoMigrate(
			model.SysDept{}, model.SysPost{}, model.SysUser{},
			model.SysMenu{}, model.SysRole{}, online.UserOnline{})
		lv_err.HasErrAndPanic(err)
	} else {
		logme.Warn("========== 已关闭表结构同步功能========== migrate" + migrate)
	}
}
