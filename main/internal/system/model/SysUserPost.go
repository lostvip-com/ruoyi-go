// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-01-06 11:47:23 +0800 CST
// 生成人：lv
// ==========================================================================

package model

import (
	"github.com/lostvip-com/lv_framework/db"
)

// SysUserPost 用户与岗位关联
type SysUserPost struct {
	UserId  int64 `gorm:"type:bigint(20);primary_key;auto_increment;用户ID;"     json:"userId"  form:"userId"`
	PostId  int64 `gorm:"type:bigint(20);primary_key;auto_increment;岗位ID;"     json:"postId"  form:"postId"`
	DelFlag int   `gorm:"type:tinyint(1);default:0;comment:删除标记;" column:del_flag; json:"delFlag"`
}

func (e *SysUserPost) TableName() string {
	return "sys_user_post"
}

// 增
func (e *SysUserPost) Save() error {
	return db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysUserPost) FindById() error {
	err := db.GetMasterGorm().Take(e, "user_id=? and post_id=?", e.UserId, e.PostId).Error
	return err
}

// 查第一条
func (e *SysUserPost) FindOne() error {
	err := db.GetMasterGorm().First(e, "user_id=? and post_id=?", e.UserId, e.PostId).Error
	return err
}

// 删
func (e *SysUserPost) Delete() error {
	err := e.FindById()
	if err == nil {
		return db.GetMasterGorm().Delete(e).Error
	}
	return err
}
