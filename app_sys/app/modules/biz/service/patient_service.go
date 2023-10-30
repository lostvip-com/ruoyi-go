// ==========================================================================
// LV自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-10-29 22:05:20 +0800 CST
// 生成人：lv
// ==========================================================================
package service

import (
		"errors"
    	"github.com/gin-gonic/gin"
    	"lostvip.com/utils/convert"
    	"lostvip.com/utils/page"
    	"robvi/app/modules/biz/model"
    	"robvi/app/modules/sys/service"
    	"time"
)
type HisPatientService struct{}
//根据主键查询数据
func (svc HisPatientService) SelectRecordById(id int64) (*model.HisPatient, error) {
	entity := &model.HisPatient{ Id: id}
	_, err := entity.FindOne()
	return entity, err
}

//根据主键删除数据
func (svc HisPatientService) DeleteRecordById(id int64) bool {
	rs, err := (&model.HisPatient{ Id: id}).Delete()
	if err == nil {
		if rs > 0 {
			return true
		}
	}
	return false
}

//批量删除数据记录
func (svc HisPatientService) DeleteRecordByIds(ids string) int64 {
	ida := convert.ToInt64Array(ids, ",")
    var entity *model.HisPatient
	result, err := entity.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

//添加数据
func (svc HisPatientService) AddSave(req *model.AddHisPatientReq, c *gin.Context) (int64, error) {
	var entity *model.HisPatient

	
	   
	  entity.Name = req.Name  
	  entity.Phone = req.Phone  
	  entity.HeadUrl = req.HeadUrl  
	  entity.IdcardPath = req.IdcardPath  
	  entity.Idcard = req.Idcard  
	  entity.BedNo = req.BedNo  
	  entity.DoctorId = req.DoctorId  
	  entity.OrgId = req.OrgId  
	  entity.OrgAddress = req.OrgAddress  
	  entity.OrgEstablish = req.OrgEstablish  
	  entity.FamilyId = req.FamilyId  
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
	 
	 
	 
	 
	  entity.Remark = req.Remark  
	  entity.DeptId = req.DeptId  

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

//修改数据
func (svc HisPatientService) EditSave(req *model.EditHisPatientReq, c *gin.Context) (int64, error) {
	entity := &model.HisPatient{ Id: req.Id }
	ok, err := entity.FindOne()

	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}

	   
	entity.Name = req.Name  
	entity.Phone = req.Phone  
	entity.HeadUrl = req.HeadUrl  
	entity.IdcardPath = req.IdcardPath  
	entity.Idcard = req.Idcard  
	entity.BedNo = req.BedNo  
	entity.DoctorId = req.DoctorId  
	entity.OrgId = req.OrgId  
	entity.OrgAddress = req.OrgAddress  
	entity.OrgEstablish = req.OrgEstablish  
	entity.FamilyId = req.FamilyId  
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
	entity.Remark = req.Remark  
	entity.DeptId = req.DeptId 
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""

	var userService service.UserService
	user := userService.GetProfile(c)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return entity.Update()
}

//根据条件查询数据
func (svc HisPatientService) SelectListAll(params *model.PageHisPatientReq) ([]model.HisPatient, error) {
	entity := model.HisPatient{}
	return entity.SelectListAll(params)
}

//根据条件分页查询数据
func (svc HisPatientService) SelectListByPage(params *model.PageHisPatientReq) ([]model.HisPatient, *page.Paging, error) {
	entity := model.HisPatient{}
	return entity.SelectListByPage(params)
}

// 导出excel
func (svc HisPatientService) Export(param *model.PageHisPatientReq) (string, error) {
	head := []string{  "" ,"姓名" ,"手机号" ,"照片" ,"" ,"证件号" ,"床号" ,"责任医生Id" ,"建档单位" ,"建档单位地址" ,"建档单位" ,"" ,"性别" ,"生日" ,"体重" ,"身高" ,"民族" ,"籍贯" ,"现居地址" ,"职业" ,"联系人手机号" ,"联系人" ,"" ,"创建人" ,"创建时间" ,"更新者" ,"更新时间" ,"备注信息" ,"建档单位"}
	col := []string{  "id" ,"name" ,"phone" ,"head_url" ,"idcard_path" ,"idcard" ,"bed_no" ,"doctor_id" ,"org_id" ,"org_address" ,"org_establish" ,"family_id" ,"sex" ,"birth" ,"weight" ,"height" ,"nation" ,"native_place" ,"address" ,"occupation" ,"contactor_phone" ,"contactor_name" ,"del_flag" ,"create_by" ,"create_time" ,"update_by" ,"update_time" ,"remark" ,"dept_id"}
	entity := model.HisPatient{}
	return entity.SelectListExport(param, head, col)
}