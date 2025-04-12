package vo

import (
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"system/model"
)

// Fill with you ideas below.

// GenTable is the golang structure for table gen_table.
type GenTableVO struct {
	model.GenTable
	TreeCode       string                 `gorm:"-"` // 树编码字段
	TreeParentCode string                 `gorm:"-"` // 树父编码字段
	TreeName       string                 `gorm:"-"` // 树名称字段
	Columns        []model.GenTableColumn `gorm:"-"` // 表列信息
	PkColumn       model.GenTableColumn   `gorm:"-"` // 表列信息
}

type GenTableParams struct {
	TreeCode       string `form:"treeCode"`
	TreeParentCode string `form:"treeParentCode"`
	TreeName       string `form:"treeName"`
}

// 修改页面请求参数
type GenTableEditReq struct {
	TableId        int64  `form:"tableId" binding:"required"`
	TableName      string `form:"tableName"  binding:"required"`
	TableComment   string `form:"tableComment"  binding:"required"`
	ClassName      string `form:"className" binding:"required"`
	FunctionAuthor string `form:"functionAuthor"  binding:"required"`
	TplCategory    string `form:"tplCategory"`
	PackageName    string `form:"packageName" binding:"required"`
	ModuleName     string `form:"moduleName" binding:"required"`
	BusinessName   string `form:"businessName" binding:"required"`
	FunctionName   string `form:"functionName" binding:"required"`
	Remark         string `form:"remark"`
	Params         string `form:"params"`
	Columns        string `form:"columns"`
}

// 分页请求参数
type GenTablePageReq struct {
	TableName    string `form:"tableName"`    //表名称
	TableComment string `form:"tableComment"` //表描述
	BeginTime    string `form:"beginTime"`    //开始时间
	EndTime      string `form:"endTime"`      //结束时间
	lv_dto.Paging
}
