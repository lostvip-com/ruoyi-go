// ==========================================================================
// LV自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-02-28 14:21:50 +0800 CST
// 生成人：lv
// ==========================================================================
package service

import (
	"fmt"
	"github.com/lv_framework/db"
	"github.com/lv_framework/utils/lv_conv"
	"github.com/lv_framework/utils/lv_err"
	"github.com/lv_framework/utils/lv_office"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"main/internal/mywork/dao"
	"main/internal/mywork/model"
	"main/internal/mywork/vo"
	"time"
)

type AppGenParamsService struct{}

// FindById 根据主键查询数据
func (svc AppGenParamsService) FindById(id int64) (*model.AppGenParams, error) {
	entity := &model.AppGenParams{Id: id}
	err := entity.FindById()
	return entity, err
}

// DeleteById 根据主键删除数据
func (svc AppGenParamsService) DeleteById(id int64) error {
	err := (&model.AppGenParams{Id: id}).Delete()
	return err
}

// DeleteByIds 批量删除数据记录
func (svc AppGenParamsService) DeleteByIds(ids string) int64 {
	ida := lv_conv.ToInt64Array(ids, ",")
	var d dao.AppGenParamsDao
	rowsAffected := d.DeleteByIds(ida)
	return rowsAffected
}

// AddSave 添加数据
func (svc AppGenParamsService) AddSave(req *vo.AddAppGenParamsReq) (int64, error) {
	var entity = new(model.AppGenParams)
	entity.ParamNo = req.ParamNo
	entity.ParamName = req.ParamName
	entity.ParamType = req.ParamType
	entity.Unit = req.Unit
	entity.Remark = req.Remark
	entity.MonitorTypeId = req.MonitorTypeId
	entity.UseFlag = "1"
	entity.CreateTime = time.Now()
	entity.CreateBy = req.CreateBy
	err := entity.Save()
	lv_err.HasErrAndPanic(err)
	return entity.Id, err
}

// EditSave 修改数据
func (svc AppGenParamsService) EditSave(req *vo.EditAppGenParamsReq) error {
	entity := &model.AppGenParams{Id: req.Id}
	err := entity.FindById()
	lv_err.HasErrAndPanic(err)
	if entity.ParamName != "" && entity.UseFlag == "1" && entity.MonitorTypeId > 0 {
		panic("已经启用并绑定业务的参量号不允许修改状态")
	}
	entity.Id = req.Id
	entity.ParamNo = req.ParamNo
	entity.ParamName = req.ParamName
	entity.ParamType = req.ParamType
	entity.Unit = req.Unit
	entity.Remark = req.Remark
	entity.MonitorTypeId = req.MonitorTypeId
	entity.UpdateBy = req.UpdateBy
	entity.UpdateTime = time.Now()
	return entity.Updates()
}

// ListByPage 根据条件分页查询数据
func (svc AppGenParamsService) ListByPage(params *vo.PageAppGenParamsReq) (*[]model.AppGenParams, int64, error) {
	var d dao.AppGenParamsDao
	return d.ListByPage(params)
}

// ExportAll 导出excel
func (svc AppGenParamsService) ExportAll(param *vo.PageAppGenParamsReq) (string, error) {
	var d dao.AppGenParamsDao
	var err error
	var listMap *[]map[string]string
	if param.PageNum > 0 { //分页导出
		listMap, _, err = d.ListMapByPage(param)
	} else { //全部导出
		listMap, err = d.ListAll(param, true)
	}
	lv_err.HasErrAndPanic(err)
	heads := []string{"ID", "是否使用", "参量号", "参量名称", "参量类型", "单位", "备注信息", "监控类型", ""}
	keys := []string{"id", "useFlag", "paramNo", "paramName", "paramType", "unit", "remark", "monitorTypeId", "createTime"}
	url, err := lv_office.DownlaodExcelByListMap(&heads, &keys, listMap)
	return url, err
}

func (svc *AppGenParamsService) GenParamsToDB(baseNUm int, amount int, username string) {
	for i := 0; i < amount; { //第一行是标题，忽略
		baseNUm += 1
		p, err := svc.getParamByNo(baseNUm)
		if err == nil && p != nil {
			fmt.Println("***************已存在的参量号*****************!!!!!!!!!!!!")
			continue //参量号已经存在了
		}
		i++
		app := model.AppGenParams{
			ParamNo:    baseNUm,
			CreateBy:   username,
			UseFlag:    "0",
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		db.GetMasterGorm().Create(&app)
		logrus.Info("=============>baseNUm:" + cast.ToString(baseNUm))
	}
}

func (s *AppGenParamsService) getParamByNo(paramNO int) (*model.AppGenParams, error) {
	list := []model.AppGenParams{}
	err := db.GetMasterGorm().Where("param_no=?", paramNO).Find(&list).Error
	if len(list) > 0 {
		return &list[0], err
	} else {
		return nil, err
	}
}
