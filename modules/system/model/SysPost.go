// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-01-02 09:30:46 +0800 CST
// 生成人：lv
// ==========================================================================

package model

import (
	"common/models"
	"github.com/lostvip-com/lv_framework/lv_db"
)

// SysPost 岗位信息
type SysPost struct {
	PostId   int64  `gorm:"type:bigint(20);primary_key;auto_increment;岗位ID;"     json:"postId"  form:"postId"`
	PostCode string `gorm:"type:varchar(64);comment:岗位编码;" json:"postCode" form:"postCode"`
	PostName string `gorm:"type:varchar(50);comment:岗位名称;" json:"postName" form:"postName"`
	PostSort int    `gorm:"type:int(11);comment:显示顺序;" json:"postSort" form:"postSort"`
	Status   string `gorm:"type:char(1);comment:状态（0正常 1停用）;" json:"status" form:"status"`
	Remark   string `gorm:"type:varchar(500);comment:备注;" json:"remark" form:"remark"`
	models.BaseModel
	Selected bool `gorm:"-" json:"selected"` // 标记
}

func (e *SysPost) TableName() string {
	return "sys_post"
}

// 增
func (e *SysPost) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysPost) FindById() error {
	err := lv_db.GetMasterGorm().Take(e, e.PostId).Error
	return err
}

// 查第一条
func (e *SysPost) FindOne() error {
	tb := lv_db.GetMasterGorm()
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
	return lv_db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 删
func (e *SysPost) Delete() error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}
