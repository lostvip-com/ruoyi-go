package model

import (
	"common/model"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/namedsql"
)

type GenTable struct {
	TableId        int64  `gorm:"type:bigint(20);primary_key;auto_increment;编号;" json:"tableId"`
	TbName         string `gorm:"type:varchar(200);comment:表名称;column:table_name" json:"tableName"`
	TableComment   string `gorm:"type:varchar(500);comment:表描述;" json:"tableComment"`
	ClassName      string `gorm:"type:varchar(100);comment:实体类名称;" json:"className"`
	TplCategory    string `gorm:"type:varchar(200);comment:使用的模板（crud单表操作 tree树表操作）;" json:"tplCategory"`
	PackageName    string `gorm:"type:varchar(100);comment:生成包路径;" json:"packageName"`
	ModuleName     string `gorm:"type:varchar(30);comment:生成模块名;" json:"moduleName"`
	BusinessName   string `gorm:"type:varchar(30);comment:生成业务名;" json:"businessName"`
	FunctionName   string `gorm:"type:varchar(50);comment:生成功能名;" json:"functionName"`
	FunctionAuthor string `gorm:"type:varchar(50);comment:生成功能作者;" json:"functionAuthor"`
	Options        string `gorm:"type:varchar(1000);comment:其它生成选项;" json:"options"`
	Remark         string `gorm:"type:varchar(500);comment:备注;" json:"remark"`
	model.BaseModel

	HasEditTime string `gorm:"-"` //1需要导入time.Time 0 不需要
}

// TableName 映射数据表
func (r *GenTable) TableName() string {
	return "gen_table"
}

// Save 增
func (e *GenTable) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// FindById 查
func (e *GenTable) FindById(id int64) (*GenTable, error) {
	err := lv_db.GetMasterGorm().Take(e, id).Error
	return e, err
}

// FindOne 查第一条
func (e *GenTable) FindOne() (*GenTable, error) {
	tb := lv_db.GetMasterGorm().Table(e.TableName())

	if e.TbName != "" {
		tb = tb.Where("table_name=?", e.TableName)
	}
	if e.TableComment != "" {
		tb = tb.Where("table_comment=?", e.TableComment)
	}
	if e.ClassName != "" {
		tb = tb.Where("class_name=?", e.ClassName)
	}
	if e.TplCategory != "" {
		tb = tb.Where("tpl_category=?", e.TplCategory)
	}
	err := tb.First(e).Error
	return e, err
}

// Updates 改
func (e *GenTable) Updates() error {
	return lv_db.GetMasterGorm().Updates(e).Error
}

// Delete 删
func (e *GenTable) Delete() error {
	return lv_db.GetMasterGorm().Delete(e).Error
}

func (e *GenTable) Count() (int64, error) {
	sql := " select count(*) from gen_table where del_flag = 0 "

	if e.TbName != "" {
		sql += " and table_name = @TbName "
	}
	if e.TableComment != "" {
		sql += " and table_comment = @TableComment "
	}
	if e.ClassName != "" {
		sql += " and class_name = @ClassName "
	}
	if e.TplCategory != "" {
		sql += " and tpl_category = @TplCategory "
	}

	return namedsql.Count(lv_db.GetMasterGorm(), sql, e)
}