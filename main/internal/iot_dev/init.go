package iot_dev

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"main/internal/common/myconf"
	models "main/internal/iot_dev/model"
)

func init() {
	//自动同步表结构
	cfg := myconf.GetConfigInstance()
	migrate := cfg.GetAutoMigrate()
	if migrate == "create" || migrate == "update" || migrate == "true" {
		lv_log.Warn("######### 开始同步表结构: ############## migrate" + migrate)
		err := lv_db.GetMasterGorm().AutoMigrate(
			models.IotProduct{},
			models.IotDevice{},
			models.IotPrdProperty{},
			models.IotPrdAction{},
			models.IotPrdEvent{},
			//// 告警规则
			//models.RuleEngine{},
			//models.AlertRule{},
			//models.AlertContent{},
		)
		lv_err.HasErrAndPanic(err)
	} else {
		lv_log.Warn("========== 已关闭表结构同步功能========== migrate" + migrate)
	}
}
