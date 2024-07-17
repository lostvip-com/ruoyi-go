package dto

import (
	"time"
)

// ///////////////////////////////////////////////////////
// 充血模型的基类
//
// //////////////////////////////////////////////////////
type Model struct {
	CreateBy string `gorm:"varchar(32);comment:创建人;" json:"createBy" form:"createBy"`
	UpdateBy string `gorm:"varchar(32);comment:更新者;" json:"updateBy" form:"updateBy"`
	//Id        int64     `json:"id"         gorm:"column:id;primary_key;auto_increment;comment:创建日期;" form:"id" `
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;type:datetime;comment:创建时间;" time_format:"2006-01-02 15:04:05"  `
	UpdateTime time.Time `json:"updateTime" gorm:"update_time;type:datetime;comment:更新时间;" time_format:"2006-01-02 15:04:05"  `
	DelFlag    int       `json:"delFlag"    gorm:"del_flag;default:0;comment:删除标记;"`
}

//dao 通用的单表增查删改可以写在下面
