// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2023-12-24 16:29:05 +0800 CST
// 生成人：lv
// ==========================================================================

package model

import (
    "robvi/app/common/model_cmn"
    "lostvip.com/db"
	"time"
)

// 数据表映射结构体
type DpcTaskItem struct {
    Id  int64  `gorm:"primary_key;auto_increment;ID;"     json:"id"  form:"id"`
    TaskId  int64  `gorm:"bigint(20);comment:任务ID;" json:"taskId" form:"taskId"`
    Action  string  `gorm:"varchar(16);comment:click dbclick write write_enter;" json:"action" form:"action"`
    Content  string  `gorm:"varchar(16);comment:内容;" json:"content" form:"content"`
    IdxXpath  string  `gorm:"varchar(64);comment:html xpath;" json:"idxXpath" form:"idxXpath"`
    IdxImg1  string  `gorm:"varchar(128);comment:图片1;" json:"idxImg1" form:"idxImg1"`
    IdxImg2  string  `gorm:"varchar(128);comment:图片2;" json:"idxImg2" form:"idxImg2"`
    IdxImg3  string  `gorm:"varchar(128);comment:图片3;" json:"idxImg3" form:"idxImg3"`
    Status  string  `gorm:"varchar(16);comment:操作状态，ready running end;" json:"status" form:"status"`
    Sort  int64  `gorm:"int(11);comment:排序，大的优先;" json:"sort" form:"sort"`
    model_cmn.Model
}

//映射数据表
func (e *DpcTaskItem) TableName() string {
	return "dpc_task_item"
}

// 增
func (e *DpcTaskItem) Save() error {
	return db.GetMasterGorm().Save(e).Error
}
// 查
func (e *DpcTaskItem) FindById() error {
	err := db.GetMasterGorm().Find(e, e.Id).Error
	return err
}
// 改
func (e *DpcTaskItem) Updates() error {
	return db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}
// 删
func (e *DpcTaskItem) Delete() error {
	return db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}

