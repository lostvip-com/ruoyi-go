package model

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/namedsql"
	"github.com/lostvip-com/lv_framework/utils/lv_time"
	"gorm.io/gorm"
	"time"
)

type GenTableColumn struct {
	ColumnId      int64  `gorm:"type:bigint(20);primary_key;auto_increment;编号;" json:"columnId"`
	TableId       int64  `gorm:"type:bigint(20);comment:归属表编号;" json:"tableId"`
	ColumnName    string `gorm:"type:varchar(200);comment:列名称;" json:"columnName"`
	ColumnComment string `gorm:"type:varchar(500);comment:列描述;" json:"columnComment"`
	ColumnType    string `gorm:"type:varchar(100);comment:列类型;" json:"columnType"`
	GoType        string `gorm:"type:varchar(500);comment:Go类型;" json:"goType"`
	GoField       string `gorm:"type:varchar(200);comment:Go字段名;" json:"goField"`
	HtmlField     string `gorm:"type:varchar(200);comment:html字段名;" json:"htmlField"`
	IsPk          string `gorm:"type:char(1);comment:是否主键（1是）;" json:"isPk"`
	IsIncrement   string `gorm:"type:char(1);comment:是否自增（1是）;" json:"isIncrement"`
	IsRequired    string `gorm:"type:char(1);comment:是否必填（1是）;" json:"isRequired"`
	IsInsert      string `gorm:"type:char(1);comment:是否为插入字段（1是）;" json:"isInsert"`
	IsEdit        string `gorm:"type:char(1);comment:是否编辑字段（1是）;" json:"isEdit"`
	IsList        string `gorm:"type:char(1);comment:是否列表字段（1是）;" json:"isList"`
	IsQuery       string `gorm:"type:char(1);comment:是否查询字段（1是）;" json:"isQuery"`
	QueryType     string `gorm:"type:varchar(200);comment:查询方式（等于、不等于、大于、小于、范围）;" json:"queryType"`
	HtmlType      string `gorm:"type:varchar(200);comment:显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）;" json:"htmlType"`
	DictType      string `gorm:"type:varchar(200);comment:字典类型;" json:"dictType"`

	CreateTime time.Time `gorm:"type:datetime;comment:创建日期" time_format:"2006-01-02 15:04:05" json:"createTime"`
	UpdateTime time.Time `gorm:"type:datetime;comment:更新日期" time_format:"2006-01-02 15:04:05" json:"updateTime"`
	UpdateBy   string    `gorm:"type:varchar(64);comment:更新者;" json:"updateBy"`
	CreateBy   string    `gorm:"type:varchar(64);comment:创建者;" json:"createBy"`
}

// 映射数据表
func (GenTableColumn) TableName() string {
	return "gen_table_column"
}

// BeforeCreate 实现钩子
func (u *GenTableColumn) BeforeCreate(db *gorm.DB) error {
	u.CreateTime = lv_time.GetCurrentTime() // 设置创建时的更新时间
	u.UpdateTime = u.CreateTime             // 设置创建时的更新时间
	return nil
}

// BeforeUpdate 实现 BeforeUpdate 钩子
func (u *GenTableColumn) BeforeUpdate(db *gorm.DB) error {
	u.CreateTime = lv_time.GetCurrentTime() // 设置更新时的更新时间
	return nil
}

// 增
func (e *GenTableColumn) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// 查
func (e *GenTableColumn) FindById(id int64) (*GenTableColumn, error) {
	err := lv_db.GetMasterGorm().Take(e, id).Error
	return e, err
}

// 查第一条
func (e *GenTableColumn) FindOne() (*GenTableColumn, error) {
	tb := lv_db.GetMasterGorm().Table(e.TableName())

	if e.TableId != 0 {
		tb = tb.Where("table_id=?", e.TableId)
	}
	if e.ColumnName != "" {
		tb = tb.Where("column_name=?", e.ColumnName)
	}
	if e.ColumnComment != "" {
		tb = tb.Where("column_comment=?", e.ColumnComment)
	}
	if e.ColumnType != "" {
		tb = tb.Where("column_type=?", e.ColumnType)
	}
	err := tb.First(e).Error
	return e, err
}

// 改
func (e *GenTableColumn) Updates() error {
	return lv_db.GetMasterGorm().Updates(e).Error
}

// 删
func (e *GenTableColumn) Delete() error {
	return lv_db.GetMasterGorm().Delete(e).Error
}

func (e *GenTableColumn) Count() (int64, error) {
	sql := " select count(*) from gen_table_column where del_flag = 0 "

	if e.TableId != 0 {
		sql += " and table_id = @TableId "
	}
	if e.ColumnName != "" {
		sql += " and column_name = @ColumnName "
	}
	if e.ColumnComment != "" {
		sql += " and column_comment = @ColumnComment "
	}
	if e.ColumnType != "" {
		sql += " and column_type = @ColumnType "
	}

	return namedsql.Count(lv_db.GetMasterGorm(), sql, e)
}
