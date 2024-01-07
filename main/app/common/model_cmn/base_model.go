// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2023-11-26 16:09:17 +0800 CST
// 生成人：lv
// ==========================================================================

package model_cmn

import (
	"time"
)

// 数据表映射结构体
type BaseModel struct {
	DelFlag    string    `json:"del_flag" xorm:"comment('删除标识1删除0未删除') char(1)"`
	CreateBy   string    `json:"create_by" xorm:"comment('创建人') varchar(32)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') datetime" time_format:"2006-01-02 15:04:05"`
	UpdateBy   string    `json:"update_by" xorm:"comment('更新者') varchar(32)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') datetime"  time_format:"2006-01-02 15:04:05"`
}
