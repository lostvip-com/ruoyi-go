package model

import "robvi/app/common/model_cmn"

type SysPost struct {
	PostId   int64  `json:"postId" gorm:"primary_key;auto_increment" ` //表编码
	PostCode string `json:"postCode" gorm:"not null comment:岗位编码;VARCHAR(64)"`
	PostName string `json:"postName" gorm:"not null comment:岗位名称;VARCHAR(50)"`
	PostSort int    `json:"postSort" gorm:"not null comment:显示顺序;INT(4)"`
	Status   string `json:"status" gorm:"not null comment:状态（0正常 1停用）;CHAR(1)"`
	CreateBy string `json:"createBy" gorm:"default '' comment:创建者;VARCHAR(64)"`
	Remark   string `json:"remark" gorm:"comment:备注;VARCHAR(500)"`
	TenantId int64  `json:"tenantId" gorm:"default 0 comment:租户id;BIGINT(20)"`
	model_cmn.Model
}

func (SysPost) TableName() string {
	return "sys_post"
}
