// ==========================================================================
// LV自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-11-26 16:09:17 +0800 CST
// 生成人：lv
// ==========================================================================
package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	model2 "robvi/app/biz/model"
	"robvi/app/system/service"
	"time"
)

type HisPatientService struct{}

// 根据主键查询数据
func (svc HisPatientService) SelectRecordById(id int64) (*model2.HisPatient, error) {
	entity := &model2.HisPatient{Id: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除数据
func (svc HisPatientService) DeleteRecordById(id int64) bool {
	rs, err := (&model2.HisPatient{Id: id}).Delete()
	if err == nil {
		if rs > 0 {
			return true
		}
	}
	return false
}

// 批量删除数据记录
func (svc HisPatientService) DeleteRecordByIds(ids string) int64 {
	ida := lv_conv.ToInt64Array(ids, ",")
	var entity *model2.HisPatient
	result, err := entity.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

// 添加数据
func (svc HisPatientService) AddSave(req *model2.AddHisPatientReq, c *gin.Context) (int64, error) {
	var entity *model2.HisPatient

	entity.Name = req.Name
	entity.Phone = req.Phone
	entity.DoctorId = req.DoctorId
	entity.Idcard = req.Idcard
	entity.HeadUrl = req.HeadUrl
	entity.IdcardPath = req.IdcardPath
	entity.BedNo = req.BedNo
	entity.OrgId = req.OrgId
	entity.OrgAddress = req.OrgAddress
	entity.OrgEstablish = req.OrgEstablish
	entity.Sex = req.Sex
	entity.Birth = req.Birth
	entity.Weight = req.Weight
	entity.Height = req.Height
	entity.Nation = req.Nation
	entity.NativePlace = req.NativePlace
	entity.Address = req.Address
	entity.Occupation = req.Occupation
	entity.ContactorPhone = req.ContactorPhone
	entity.ContactorName = req.ContactorName
	entity.DelFlag = req.DelFlag
	entity.CreateBy = req.CreateBy
	entity.CreateTime = req.CreateTime
	entity.UpdateBy = req.UpdateBy
	entity.UpdateTime = req.UpdateTime
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()
	entity.CreateBy = ""

	var userService service.UserService
	user := userService.GetProfile(c)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	_, err := entity.Insert()
	return entity.Id, err
}

// 修改数据
func (svc HisPatientService) EditSave(req *model2.EditHisPatientReq, c *gin.Context) (int64, error) {
	entity := &model2.HisPatient{Id: req.Id}
	ok, err := entity.FindOne()

	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}

	entity.Name = req.Name
	entity.Phone = req.Phone
	entity.DoctorId = req.DoctorId
	entity.Idcard = req.Idcard
	entity.HeadUrl = req.HeadUrl
	entity.IdcardPath = req.IdcardPath
	entity.BedNo = req.BedNo
	entity.OrgId = req.OrgId
	entity.OrgAddress = req.OrgAddress
	entity.OrgEstablish = req.OrgEstablish
	entity.Sex = req.Sex
	entity.Birth = req.Birth
	entity.Weight = req.Weight
	entity.Height = req.Height
	entity.Nation = req.Nation
	entity.NativePlace = req.NativePlace
	entity.Address = req.Address
	entity.Occupation = req.Occupation
	entity.ContactorPhone = req.ContactorPhone
	entity.ContactorName = req.ContactorName
	entity.UpdateBy = req.UpdateBy
	entity.UpdateTime = req.UpdateTime
	entity.Remark = req.Remark
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""

	var userService service.UserService
	user := userService.GetProfile(c)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return entity.Update()
}

// 根据条件查询数据
func (svc HisPatientService) SelectListAll(params *model2.PageHisPatientReq) ([]model2.HisPatient, error) {
	entity := model2.HisPatient{}
	return entity.SelectListAll(params)
}

// 根据条件分页查询数据
func (svc HisPatientService) SelectListByPage(params *model2.PageHisPatientReq) ([]model2.HisPatient, *lv_web.Paging, error) {
	entity := model2.HisPatient{}
	return entity.SelectListByPage(params)
}

// 导出excel
func (svc HisPatientService) Export(param *model2.PageHisPatientReq) (string, error) {
	head := []string{"", "姓名", "手机号", "责任医生Id", "证件号", "照片", "身份证照片", "床号", "建档单位", "建档单位地址", "建档单位", "家庭ID", "性别", "生日", "体重", "身高", "民族", "籍贯", "现居地址", "职业", "联系人手机号", "联系人", "删除标识1删除0未删除", "创建人", "创建时间", "更新者", "更新时间", "备注信息"}
	col := []string{"id", "name", "phone", "doctor_id", "idcard", "head_url", "idcard_path", "bed_no", "org_id", "org_address", "org_establish", "family_id", "sex", "birth", "weight", "height", "nation", "native_place", "address", "occupation", "contactor_phone", "contactor_name", "del_flag", "create_by", "create_time", "update_by", "update_time", "remark"}
	entity := model2.HisPatient{}
	return entity.SelectListExport(param, head, col)
}
