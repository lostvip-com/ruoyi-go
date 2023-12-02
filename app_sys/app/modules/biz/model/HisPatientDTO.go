// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-11-26 16:09:17 +0800 CST
// 生成人：lv
// ==========================================================================
package model

import (
	"errors"
	"time"
    "lostvip.com/db"
	"lostvip.com/utils/lv_office"
	"lostvip.com/utils/lv_web"
    "xorm.io/builder"
)

//新增页面请求参数
type AddHisPatientReq struct { 
	 
	 Name  string   `form:"name" binding:"required"`  
	 Phone  string   `form:"phone" `  
	 DoctorId  int64   `form:"doctorId" `  
	 Idcard  string   `form:"idcard" `  
	 HeadUrl  string   `form:"headUrl" `  
	 IdcardPath  string   `form:"idcardPath" `  
	 BedNo  string   `form:"bedNo" `  
	 OrgId  int64   `form:"orgId" `  
	 OrgAddress  string   `form:"orgAddress" `  
	 OrgEstablish  string   `form:"orgEstablish" `  
	 FamilyId  int64   `form:"familyId" `  
	 Sex  string   `form:"sex" `  
	 Birth  string   `form:"birth" `  
	 Weight  int64   `form:"weight" `  
	 Height  int64   `form:"height" `  
	 Nation  string   `form:"nation" `  
	 NativePlace  string   `form:"nativePlace" `  
	 Address  string   `form:"address" `  
	 Occupation  string   `form:"occupation" `  
	 ContactorPhone  string   `form:"contactorPhone" `  
	 ContactorName  string   `form:"contactorName" binding:"required"`  
	 DelFlag  string   `form:"delFlag" `  
	 CreateBy  string   `form:"createBy" `  
	 CreateTime  time.Time   `form:"createTime" `  
	 UpdateBy  string   `form:"updateBy" `  
	 UpdateTime  time.Time   `form:"updateTime" `  
	 Remark  string   `form:"remark" `  
}

//修改页面请求参数
type EditHisPatientReq struct {
      Id    int64  `form:"id" binding:"required"`    
      Name  string `form:"name" binding:"required"`   
      Phone  string `form:"phone" `   
      DoctorId  int64 `form:"doctorId" `   
      Idcard  string `form:"idcard" `   
      HeadUrl  string `form:"headUrl" `   
      IdcardPath  string `form:"idcardPath" `   
      BedNo  string `form:"bedNo" `   
      OrgId  int64 `form:"orgId" `   
      OrgAddress  string `form:"orgAddress" `   
      OrgEstablish  string `form:"orgEstablish" `   
      FamilyId  int64 `form:"familyId" `   
      Sex  string `form:"sex" `   
      Birth  string `form:"birth" `   
      Weight  int64 `form:"weight" `   
      Height  int64 `form:"height" `   
      Nation  string `form:"nation" `   
      NativePlace  string `form:"nativePlace" `   
      Address  string `form:"address" `   
      Occupation  string `form:"occupation" `   
      ContactorPhone  string `form:"contactorPhone" `   
      ContactorName  string `form:"contactorName" binding:"required"`         
      UpdateBy  string `form:"updateBy" `   
      UpdateTime  time.Time `form:"updateTime" `   
      Remark  string `form:"remark" `  
}

//分页请求参数 
type PageHisPatientReq struct {    
	Name  string `form:"name"` //姓名   
	Phone  string `form:"phone"` //手机号   
	DoctorId  int64 `form:"doctorId"` //责任医生Id   
	Idcard  string `form:"idcard"` //证件号                                                
	BeginTime  string `form:"beginTime"`  //开始时间
	EndTime    string `form:"endTime"`    //结束时间
	PageNum    int    `form:"pageNum"`    //当前页码
	PageSize   int    `form:"pageSize"`   //每页数
}

//根据条件分页查询数据
func (e *HisPatient) SelectListByPage(param *PageHisPatientReq) ([]HisPatient, *lv_web.Paging, error) {
	db := db.Instance().Engine()
    p := new(lv_web.Paging)

	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table("gen_table").Alias("t")

	if param != nil {    
		
		if param.Name != "" {
			model.Where("t.name like ?", "%"+param.Name+"%")
		}    
		
		if param.Phone != "" {
			model.Where("t.phone like ?", "%"+param.Phone+"%")
		}    
		 
		if param.DoctorId != 0 {
			model.Where("t.doctor_id = ?", param.DoctorId)
		}
		    
		
		if param.Idcard != "" {
			model.Where("t.idcard like ?", "%"+param.Idcard+"%")
		}                                                 
		if param.BeginTime != "" {
			model.Where("t.create_time >= ?", param.BeginTime)
		}

		if param.EndTime != "" {
			 model.Where("t.create_time<=?", param.EndTime)
		}
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = lv_web.CreatePaging(param.PageNum, param.PageSize, int(total))
	model.Limit(p.Pagesize, p.StartNum)

	var result []HisPatient
    err = model.Find(&result)
    return result, p, err
}

// 导出excel
func (e *HisPatient) SelectListExport(param *PageHisPatientReq, head, col []string) (string, error) {
	db := db.Instance().Engine()
	build := builder.Select(col...).From("gen_table", "t")

	if param != nil {    
		
		if param.Name != "" {
			build.Where(builder.Like{"t.name", param.Name})
		}    
		
		if param.Phone != "" {
			build.Where(builder.Like{"t.phone", param.Phone})
		}    
		 
		if param.DoctorId != 0 {
			build.Where(builder.Eq{"t.doctor_id": param.DoctorId})
		}
		    
		
		if param.Idcard != "" {
			build.Where(builder.Like{"t.idcard", param.Idcard})
		}                                                 
		if param.BeginTime != "" {
			build.Where(builder.Gte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.BeginTime + "','%y%m%d')"})
		}

		if param.EndTime != "" {
			build.Where(builder.Lte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.EndTime + "','%y%m%d')"})
		}
	}

	sqlStr, _, _ := build.ToSQL()
	arr, err := db.SQL(sqlStr).QuerySliceString()
	path, err := lv_office.DownlaodExcel(head, arr)

	return path, err
}

//获取所有数据
func (e *HisPatient) SelectListAll(param *PageHisPatientReq) ([]HisPatient, error) {
	db := db.Instance().Engine()
	model := db.Table("gen_table").Alias("t")

	if param != nil {    
		
		if param.Name != "" {
			model.Where("t.name like ?", "%"+param.Name+"%")
		}    
		
		if param.Phone != "" {
			model.Where("t.phone like ?", "%"+param.Phone+"%")
		}    
		 
		if param.DoctorId != 0 {
			model.Where("t.doctor_id = ?", param.DoctorId)
		}
		   
		
		if param.Idcard != "" {
			model.Where("t.idcard like ?", "%"+param.Idcard+"%")
		}                                                 
		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []HisPatient
	err := model.Find(&result)
	return result, err
}