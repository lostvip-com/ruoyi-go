package cm_model

import (
	"gorm.io/gorm"
	"lostvip.com/utils/lv_time"
	"time"
)

// BaseModel 共享属性
type BaseModel struct {
	DelFlag    int       `gorm:"type:tinyint(1);default:0;comment:删除标记;column:del_flag;" json:"delFlag"`
	CreateTime time.Time `gorm:"type:datetime;comment:创建日期" time_format:"2006-01-02 15:04:05"`
	UpdateTime time.Time `gorm:"type:datetime;comment:更新日期" time_format:"2006-01-02 15:04:05"`
	UpdateBy   string    `gorm:"type:varchar(64);comment:更新者;" json:"updateBy"`
	CreateBy   string    `gorm:"type:varchar(64);comment:创建者;" json:"updateBy"`
	//TenantId   int64     `gorm:"type:bigint(20);comment:租户id;" json:"tenantId" form:"tenantId"`
}

// 实现 BeforeCreate 钩子
func (u *BaseModel) BeforeCreate(db *gorm.DB) error {
	u.UpdateTime = lv_time.GetCurrentTime() // 设置创建时的更新时间
	return nil
}

// 实现 BeforeUpdate 钩子
func (u *BaseModel) BeforeUpdate(db *gorm.DB) error {
	u.CreateTime = lv_time.GetCurrentTime() // 设置更新时的更新时间
	return nil
}
