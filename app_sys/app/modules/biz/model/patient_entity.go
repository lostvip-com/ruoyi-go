// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2023-10-29 22:05:20 +0800 CST
// 生成人：lv
// ==========================================================================

package model

import (
    "lostvip.com/db"
	"time"
)

// 数据表映射结构体
type HisPatient struct { 
	 Id       int64         `json:"id" xorm:"not null pk autoincr comment('') bigint(32)"`    
	 Name    string         `json:"name" xorm:"comment('姓名') varchar(32)"`    
	 Phone    string         `json:"phone" xorm:"comment('手机号') varchar(11)"`    
	 HeadUrl    string         `json:"head_url" xorm:"comment('照片') varchar(256)"`    
	 IdcardPath    string         `json:"idcard_path" xorm:"comment('') varchar(128)"`    
	 Idcard    string         `json:"idcard" xorm:"comment('证件号') varchar(18)"`    
	 BedNo    string         `json:"bed_no" xorm:"comment('床号') varchar(32)"`    
	 DoctorId    int64         `json:"doctor_id" xorm:"comment('责任医生Id') bigint(20)"`    
	 OrgId    int64         `json:"org_id" xorm:"comment('建档单位') bigint(20)"`    
	 OrgAddress    string         `json:"org_address" xorm:"comment('建档单位地址') varchar(255)"`    
	 OrgEstablish    string         `json:"org_establish" xorm:"comment('建档单位') varchar(32)"`    
	 FamilyId    int64         `json:"family_id" xorm:"comment('') bigint(32)"`    
	 Sex    string         `json:"sex" xorm:"comment('性别') varchar(1)"`    
	 Birth    string         `json:"birth" xorm:"comment('生日') varchar(10)"`    
	 Weight    int64         `json:"weight" xorm:"comment('体重') float(10,0)"`    
	 Height    int64         `json:"height" xorm:"comment('身高') float(10,0)"`    
	 Nation    string         `json:"nation" xorm:"comment('民族') varchar(128)"`    
	 NativePlace    string         `json:"native_place" xorm:"comment('籍贯') varchar(128)"`    
	 Address    string         `json:"address" xorm:"comment('现居地址') varchar(128)"`    
	 Occupation    string         `json:"occupation" xorm:"comment('职业') varchar(32)"`    
	 ContactorPhone    string         `json:"contactor_phone" xorm:"comment('联系人手机号') varchar(11)"`    
	 ContactorName    string         `json:"contactor_name" xorm:"comment('联系人') varchar(128)"`    
	 DelFlag    string         `json:"del_flag" xorm:"comment('') char(1)"`    
	 CreateBy    string         `json:"create_by" xorm:"comment('创建人') varchar(32)"`    
	 CreateTime    time.Time         `json:"create_time" xorm:"comment('创建时间') date"`    
	 UpdateBy    string         `json:"update_by" xorm:"comment('更新者') varchar(32)"`    
	 UpdateTime    time.Time         `json:"update_time" xorm:"comment('更新时间') date"`    
	 Remark    string         `json:"remark" xorm:"comment('备注信息') varchar(32)"`    
	 DeptId    int64         `json:"dept_id" xorm:"comment('建档单位') bigint(20)"`    
}

//映射数据表
func (e *HisPatient) TableName() string {
	return "gen_table"
}

// 插入数据
func (e *HisPatient) Insert() (int64, error) {
	return db.Instance().Engine().Table(e.TableName()).Insert(e)
}

// 更新数据
func (e *HisPatient) Update() (int64, error) {
	return db.Instance().Engine().Table(e.TableName()).ID(e.Id).Update(e)
}

// 删除
func (e *HisPatient) Delete() (int64, error) {
	return db.Instance().Engine().Table(e.TableName()).ID(e.Id).Delete(e)
}

//批量删除
func (e *HisPatient) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(e.TableName()).In("id", ids).Delete(new(HisPatient))
}

// 根据结构体中已有的非空数据来获得单条数据
func (e *HisPatient) FindOne() (bool, error) {
	return db.Instance().Engine().Table(e.TableName()).Get(e)
}

// 根据条件查询
func (e *HisPatient) Find(where, order string) ([]HisPatient, error) {
	var list []HisPatient
	err := db.Instance().Engine().Table(e.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (e *HisPatient) FindIn(column string, args ...interface{}) ([]HisPatient, error) {
	var list []HisPatient
	err := db.Instance().Engine().Table(e.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (e *HisPatient)  FindNotIn(column string, args ...interface{}) ([]HisPatient, error) {
	var list []HisPatient
	err := db.Instance().Engine().Table(e.TableName()).NotIn(column, args).Find(&list)
	return list, err
}