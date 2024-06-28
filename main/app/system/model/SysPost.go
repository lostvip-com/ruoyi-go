// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-01-02 09:30:46 +0800 CST
// 生成人：lv
// ==========================================================================

package model

import (
	"github.com/lv_framework/db"
	"time"
)

// SysPost 岗位信息
type SysPost struct {
	PostId     int64     `gorm:"type:bigint(20);primary_key;auto_increment;岗位ID;"     json:"postId"  form:"postId"`
	PostCode   string    `gorm:"type:varchar(64);comment:岗位编码;" json:"postCode" form:"postCode"`
	PostName   string    `gorm:"type:varchar(50);comment:岗位名称;" json:"postName" form:"postName"`
	PostSort   int       `gorm:"type:int(11);comment:显示顺序;" json:"postSort" form:"postSort"`
	Status     string    `gorm:"type:char(1);comment:状态（0正常 1停用）;" json:"status" form:"status"`
	UpdateBy   string    `gorm:"type:varchar(64);comment:更新者;" json:"updateBy" form:"updateBy"`
	UpdateTime time.Time `gorm:"type:datetime;comment:更新时间;" json:"updateTime" form:"updateTime" time_format:"2006-01-02 15:04:05"`
	Remark     string    `gorm:"type:varchar(500);comment:备注;" json:"remark" form:"remark"`
	TenantId   int64     `gorm:"type:bigint(20);comment:租户id;" json:"tenantId" form:"tenantId"`
	CreateTime time.Time `gorm:"type:datetime;comment:创建时间;"  column:create_time; time_format:"2006-01-02 15:04:05" json:"createTime"  `
	CreateBy   string    `gorm:"type:varchar(32);comment:创建人;" column:create_by; json:"createBy" form:"createBy"`
	DelFlag    int       `gorm:"type:tinyint(1);default:0;comment:删除标记;" column:del_flag; json:"delFlag"`
	Selected   bool      `gorm:"-" json:"selected"` // 标记
}

func (e *SysPost) TableName() string {
	return "sys_post"
}

// 增
func (e *SysPost) Save() error {
	return db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysPost) FindById() error {
	err := db.GetMasterGorm().Take(e, e.PostId).Error
	return err
}

// 查第一条
func (e *SysPost) FindOne() error {
	tb := db.GetMasterGorm()
	if e.PostId != 0 {
		tb = tb.Where("post_id=?", e.PostId)
	}
	if e.PostCode != "" {
		tb = tb.Where("post_code=?", e.PostCode)
	}
	err := tb.First(e).Error
	return err
}

// 改
func (e *SysPost) Updates() error {
	return db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 删
func (e *SysPost) Delete() error {
	return db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}
