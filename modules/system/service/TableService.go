package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_reflect"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"os"
	"strings"
	"system/dao"
	"system/model"
	"system/vo"
	"text/template"
	"time"
)

type TableService struct {
}

// SelectRecordById 根据主键查询数据
func (svc TableService) SelectRecordById(id int64) (*vo.GenTableVO, error) {
	var dao dao.GenTableDao
	vo, err := dao.ListColumn(id)
	if err != nil {
		return nil, err
	}
	if vo.Options != "" {
		p := make(map[string]interface{}, 2)
		if e := json.Unmarshal([]byte(vo.Options), &p); e == nil {
			treeCode := p["treeCode"].(string)
			treeParentCode := p["treeParentCode"].(string)
			treeName := p["treeName"].(string)
			vo.TreeCode = treeCode
			vo.TreeParentCode = treeParentCode
			vo.TreeName = treeName
		}
	}
	//

	return vo, err
}

// DeleteRecordById 根据主键删除数据
func (svc TableService) DeleteRecordById(id int64) bool {
	err := (&model.GenTable{TableId: id}).Delete()
	if err != nil {
		return true
	}
	return false
}

// 批量删除数据记录
func (svc TableService) DeleteRecordByIds(ids string) error {
	idarr := lv_conv.ToInt64Array(ids, ",")
	err := lv_db.GetMasterGorm().Exec("delete from gen_table where table_id in (?)", idarr).Error
	err = lv_db.GetMasterGorm().Exec("delete from gen_table_column where table_id in (?)", idarr).Error
	return err
}

// 保存修改数据
func (svc TableService) SaveEdit(req *vo.GenTableEditReq, c *gin.Context) error {
	table := new(model.GenTable)
	table, err := table.FindById(req.TableId)
	if err != nil {
		return errors.New("数据不存在")
	}
	lv_reflect.CopyProperties(req, table)
	table.UpdateTime = time.Now()
	var userService UserService
	user := userService.GetProfile(c)
	if user != nil {
		table.UpdateBy = user.LoginName
	}
	err = lv_db.GetMasterGorm().Transaction(func(tx *gorm.DB) error {
		var err error
		err = table.Updates()
		//保存列数据
		if req.Columns == "" {
			return err
		}
		var columnList []model.GenTableColumn
		err = json.Unmarshal([]byte(req.Columns), &columnList)
		lv_err.HasError1(err)
		if columnList == nil {
			return err
		}
		for _, column := range columnList {
			po := new(model.GenTableColumn)
			if column.ColumnId == 0 {
				continue
			}
			po, err = po.FindById(column.ColumnId)
			lv_err.HasErrAndPanic(err)
			lv_reflect.CopyProperties(column, po)
			err = po.Updates()
			lv_err.HasErrAndPanic(err)
		}
		return err
	})
	return err
}

// 设置主键列信息
func (svc TableService) SetPkColumn(table *vo.GenTableVO, columns []model.GenTableColumn) {
	for _, column := range columns {
		//是否存在 需要编辑的时间字段 time.Time类型
		if column.GoType == "time.Time" && column.IsEdit == "1" {
			table.HasEditTime = "1"
		}

		if column.IsPk == "1" {
			table.PkColumn = column
		}
	} //end for
	if &(table.PkColumn) == nil {
		table.PkColumn = columns[0]
	}
}

// 根据条件分页查询数据
func (svc TableService) SelectListByPage(param *vo.GenTablePageReq) ([]model.GenTable, int64, error) {
	var table dao.GenTableDao
	return table.SelectListByPage(param)
}

// 查询据库列表
func (svc TableService) SelectDbTableList(param *vo.GenTablePageReq) ([]model.GenTable, int64, error) {
	var dao dao.GenTableDao
	return dao.SelectDbTableList(param)
}

// 查询据库列表
func (svc TableService) SelectDbTableListByNames(tableNames []string) ([]model.GenTable, error) {
	var dao dao.GenTableDao
	return dao.SelectDbTableListByNames(tableNames)
}

// 查询据库列表
func (svc TableService) SelectTableByName(tableName string) (*model.GenTable, error) {
	var dao dao.GenTableDao
	return dao.SelectTableByName(tableName)
}

// SelectGenTableById 查询表ID业务信息
func (svc TableService) SelectGenTableById(tableId int64) (*model.GenTable, error) {
	var table model.GenTable
	return table.FindById(tableId)
}

// SelectGenTableByName 查询表名称业务信息
func (svc TableService) SelectGenTableByName(tbName string) (*model.GenTable, error) {
	var table = model.GenTable{TbName: tbName}
	return table.FindOne()
}

// ImportGenTable 导入表结构
func (svc TableService) ImportGenTable(tableList *[]model.GenTable, operName string) error {
	var err error
	if tableList != nil && operName != "" {
		err = lv_db.GetMasterGorm().Transaction(func(tx *gorm.DB) error {
			var err error
			for _, table := range *tableList {
				tableName := table.TbName
				svc.InitTable(&table, operName)
				err = tx.Table(table.TableName()).Save(&table).Error
				if err != nil {
					return err
				}
				if err != nil || table.TableId <= 0 {
					tx.Rollback()
					return errors.New("保存数据失败")
				}
				colSvc := TableColumnService{}
				genTableColumns, err := colSvc.SelectDbTableColumnsByName(tableName)

				if err != nil || len(genTableColumns) <= 0 {
					tx.Rollback()
					return errors.New("获取列数据失败")
				}

				for index, column := range genTableColumns {
					if index > 0 && index < 5 { // 默认前 前4个字段为查询用
						column.IsQuery = "1"
					}
					svc.InitColumnField(&column, &table)
					err = tx.Save(&column).Error
					if err != nil {
						tx.Rollback()
						return errors.New("保存列数据失败")
					}
				}
			}
			return err
		})
	}
	return err
}

// InitTable 初始化表信息
func (svc TableService) InitTable(table *model.GenTable, createBy string) {
	table.ClassName = svc.ConvertClassName(table.TbName)
	table.BusinessName = svc.GetBusinessName(table.TbName)
	table.FunctionName = table.BusinessName
	table.PackageName = lv_global.Config().GetVipperCfg().GetString("gen.packageName")
	table.ModuleName = lv_global.Config().GetVipperCfg().GetString("gen.moduleName")
	table.FunctionAuthor = lv_global.Config().GetVipperCfg().GetString("gen.author")
	table.CreateBy = createBy
	table.TplCategory = "crud"
	table.CreateTime = time.Now()
	table.UpdateTime = table.CreateTime
	table.TableComment = strings.ReplaceAll(table.TableComment, "表", "")
	if table.TableComment == "" {
		table.TableComment = table.ClassName
	}
}

// InitColumnField 初始化列属性字段
func (svc TableService) InitColumnField(column *model.GenTableColumn, table *model.GenTable) {
	dataType := svc.GetDbType(column.ColumnType)
	columnName := column.ColumnName
	column.TableId = table.TableId
	column.CreateBy = table.CreateBy
	//设置字段名
	column.GoField = svc.ConvertToCamelCase(columnName)
	column.HtmlField = svc.ConvertToCamelCase1(columnName)

	if svc.IsStringObject(dataType) {
		//字段为字符串类型
		column.GoType = "string"
		if strings.EqualFold(dataType, "text") || strings.EqualFold(dataType, "tinytext") || strings.EqualFold(dataType, "mediumtext") || strings.EqualFold(dataType, "longtext") {
			column.HtmlType = "textarea"
		} else {
			columnLength := svc.GetColumnLength(column.ColumnType)
			column.ColumnSize = columnLength
			if columnLength >= 255 {
				column.HtmlType = "textarea"
			} else {
				column.HtmlType = "input"
			}
		}
	} else if svc.IsTimeObject(dataType) {
		//字段为时间类型
		column.GoType = "time.Time"
		column.HtmlType = "datetime"
	} else if svc.IsNumberObject(dataType) {
		//字段为数字类型
		column.HtmlType = "input"
		// 如果是浮点型
		tmp := column.ColumnType
		if strings.HasPrefix(tmp, "float") {
			column.GoType = "float"
		} else if strings.Contains(tmp, "double") || strings.HasPrefix(tmp, "decimal") {
			column.GoType = "float64"
		} else if strings.Contains(tmp, "bigint") {
			column.GoType = "int64"
		} else if strings.Contains(tmp, "int") {
			column.GoType = "int"
		} else {
			start := strings.Index(tmp, "(")
			end := strings.Index(tmp, ")")
			result := tmp[start+1 : end]
			arr := strings.Split(result, ",")
			if len(arr) == 2 && cast.ToInt(arr[1]) > 0 {
				column.GoType = "float64"
			} else if len(arr) == 1 && cast.ToInt(arr[0]) <= 10 {
				column.GoType = "int"
			} else {
				column.GoType = "int64"
			}
		}

	}
	//新增字段，不生成到界面上哐不需要人工编辑
	if columnName == "create_by" || columnName == "create_time" || columnName == "update_by" || columnName == "update_time" {
		column.IsRequired = "0"
		column.IsInsert = "0"
	} else { //需要生成到界面
		if column.IsPk == "1" {
			column.IsInsert = "0"
		} else {
			column.IsInsert = "1"
		}
		column.IsRequired = "0"
		if strings.Index(columnName, "name") >= 0 || strings.Index(columnName, "status") >= 0 {
			column.IsRequired = "1"
		}
	}

	// 编辑字段
	if svc.IsNotEdit(columnName) {
		if column.IsPk == "1" {
			column.IsEdit = "0"
		} else {
			column.IsEdit = "1"
		}
	} else {
		column.IsEdit = "0"
	}
	// 列表字段
	if svc.IsNotList(columnName) {
		column.IsList = "1"
	} else {
		column.IsList = "0"
	}
	// 黑名单
	if column.IsQuery == "" {
		column.IsQuery = "0" //默认非查询
	} else { //检查黑名单
		if svc.IsNotQuery(columnName) {
			column.IsQuery = "0"
		}
	}

	// 查询字段类型
	if svc.IsStringType(column.ColumnType) {
		column.QueryType = "LIKE"
	} else {
		column.QueryType = "EQ"
	}

	// 状态字段设置单选框
	if svc.CheckStatusColumn(columnName) {
		column.HtmlType = "radio"
	} else if svc.CheckTypeColumn(columnName) || svc.CheckSexColumn(columnName) {
		// 类型&性别字段设置下拉框
		column.HtmlType = "select"
	}
}

// 检查字段名后3位是否是sex
func (svc TableService) CheckSexColumn(columnName string) bool {
	if len(columnName) >= 3 {
		end := len(columnName)
		start := end - 3

		if start <= 0 {
			start = 0
		}

		if columnName[start:end] == "sex" {
			return true
		}
	}
	return false
}

// 检查字段名后4位是否是type
func (svc TableService) IsStringType(columnType string) bool {
	if strings.HasPrefix(columnType, "char") {
		return true
	}
	if strings.HasPrefix(columnType, "varchar") {
		return true
	}
	if strings.HasPrefix(columnType, "text") {
		return true
	}
	if strings.HasPrefix(columnType, "string") {
		return true
	}
	return false
}

// 检查字段名后4位是否是type
func (svc TableService) CheckTypeColumn(columnName string) bool {
	if len(columnName) >= 4 {
		end := len(columnName)
		start := end - 4

		if start <= 0 {
			start = 0
		}

		if columnName[start:end] == "type" {
			return true
		}
	}
	return false
}

// CheckNameColumn 检查字段名后4位是否是name
func (svc TableService) CheckNameColumn(columnName string) bool {
	if len(columnName) >= 4 {
		end := len(columnName)
		start := end - 4

		if start <= 0 {
			start = 0
		}

		tmp := columnName[start:end]

		if tmp == "name" {
			return true
		}
	}
	return false
}

// CheckStatusColumn 检查字段名后6位是否是status
func (svc TableService) CheckStatusColumn(columnName string) bool {
	if len(columnName) >= 6 {
		end := len(columnName)
		start := end - 6

		if start <= 0 {
			start = 0
		}
		tmp := columnName[start:end]

		if tmp == "status" {
			return true
		}
	}

	return false
}

// GetDbType  获取数据库类型字段
func (svc TableService) GetDbType(columnType string) string {
	if strings.Index(columnType, "(") > 0 {
		return columnType[0:strings.Index(columnType, "(")]
	} else {
		return columnType
	}
}

// ConvertClassName 表名转换成类名
func (svc TableService) ConvertClassName(tableName string) string {
	return svc.ConvertToCamelCase(tableName)
}

// GetBusinessName 获取业务名
func (svc TableService) GetBusinessName(tableName string) string {
	lastIndex := strings.LastIndex(tableName, "_")
	nameLength := len(tableName)
	businessName := tableName[lastIndex+1 : nameLength]
	return businessName
}

// ConvertToCamelCase 将下划线大写方式命名的字符串转换为驼峰式。如果转换前的下划线大写方式命名的字符串为空，则返回空字符串。 例如：HELLO_WORLD->HelloWorld
func (svc TableService) ConvertToCamelCase(name string) string {
	if name == "" {
		return ""
	} else if !strings.Contains(name, "_") {
		// 不含下划线，仅将首字母大写
		return strings.ToUpper(name[0:1]) + name[1:len(name)]
	}
	var result string = ""
	camels := strings.Split(name, "_")
	for index := range camels {
		if camels[index] == "" {
			continue
		}
		camel := camels[index]
		result = result + strings.ToUpper(camel[0:1]) + strings.ToLower(camel[1:len(camel)])
	}
	return result
}

// ConvertToCamelCase1 将下划线大写方式命名的字符串转换为驼峰式,首字母小写。如果转换前的下划线大写方式命名的字符串为空，则返回空字符串。 例如：HELLO_WORLD->helloWorld
func (svc TableService) ConvertToCamelCase1(name string) string {
	if name == "" {
		return ""
	} else if !strings.Contains(name, "_") {
		// 不含下划线，原值返回
		return name
	}
	var result string = ""
	camels := strings.Split(name, "_")
	for index := range camels {
		if camels[index] == "" {
			continue
		}
		camel := camels[index]
		if result == "" {
			result = strings.ToLower(camel[0:1]) + strings.ToLower(camel[1:len(camel)])
		} else {
			result = result + strings.ToUpper(camel[0:1]) + strings.ToLower(camel[1:len(camel)])
		}
	}
	return result
}

// 获取字段长度
func (svc TableService) GetColumnLength(columnType string) int {
	start := strings.Index(columnType, "(")
	end := strings.Index(columnType, ")")
	result := columnType[start+1 : end-1]
	return cast.ToInt(result)
}

// 获取Go类别下拉框
func (svc TableService) GoTypeTpl() string {
	return `<script id="goTypeTpl" type="text/x-jquery-tmpl">
<div>
<select class='form-control' name='columns[${index}].goType'>
    <option value="int64" {{if goType==="int64"}}selected{{/if}}>int64</option>
    <option value="int" {{if goType==="int"}}selected{{/if}}>int</option>
    <option value="string" {{if goType==="string"}}selected{{/if}}>string</option>
    <option value="time.Time" {{if goType==="time.Time"}}selected{{/if}}>time.Time</option>
    <option value="float64" {{if goType==="float64"}}selected{{/if}}>float64</option>
    <option value="byte" {{if goType==="byte"}}selected{{/if}}>byte</option>
</select>
</div>
</script>`
}

// 获取查询方式下拉框
func (svc TableService) QueryTypeTpl() string {
	return `<script id="queryTypeTpl" type="text/x-jquery-tmpl">
<div>
<select class='form-control' name='columns[${index}].queryType'>
    <option value="EQ" {{if queryType==="EQ"}}selected{{/if}}>=</option>
    <option value="NE" {{if queryType==="NE"}}selected{{/if}}>!=</option>
    <option value="GT" {{if queryType==="GT"}}selected{{/if}}>></option>
    <option value="GTE" {{if queryType==="GTE"}}selected{{/if}}>>=</option>
    <option value="LT" {{if queryType==="LT"}}selected{{/if}}><</option>
    <option value="LTE" {{if queryType==="LTE"}}selected{{/if}}><=</option>
    <option value="LIKE" {{if queryType==="LIKE"}}selected{{/if}}>Like</option>
    <option value="BETWEEN" {{if queryType==="BETWEEN"}}selected{{/if}}>Between</option>
</select>
</div>
</script>`
}

// 获取显示类型下拉框
func (svc TableService) HtmlTypeTpl() string {
	return `<script id="htmlTypeTpl" type="text/x-jquery-tmpl">
<div>
<select class='form-control' name='columns[${index}].htmlType'>
    <option value="input" {{if htmlType==="input"}}selected{{/if}}>文本框</option>
    <option value="textarea" {{if htmlType==="textarea"}}selected{{/if}}>文本域</option>
    <option value="select" {{if htmlType==="select"}}selected{{/if}}>下拉框</option>
    <option value="radio" {{if htmlType==="radio"}}selected{{/if}}>单选框</option>
    <option value="checkbox" {{if htmlType==="checkbox"}}selected{{/if}}>复选框</option>
    <option value="datetime" {{if htmlType==="datetime"}}selected{{/if}}>日期控件</option>
</select>
</div>
</script>`
}

// 读取模板
func (svc TableService) LoadTemplate(templateName string, data interface{}) (string, error) {
	cur, err := os.Getwd()
	if err != nil {
		return "", err
	}
	b, err := os.ReadFile(cur + "/template/" + templateName)
	if err != nil {
		return "", err
	}
	templateStr := string(b)

	// 创建函数映射
	funcs := template.FuncMap{
		"contains": Contains,
	}

	tmpl, err := template.New(templateName).Funcs(funcs).Parse(templateStr) //建立一个模板，内容是"hello, {{OssUrl}}"
	if err != nil {
		return "", err
	}
	buffer := bytes.NewBufferString("")
	err = tmpl.Execute(buffer, data) //将string与模板合成，变量name的内容会替换掉{{OssUrl}}
	return buffer.String(), err
}

// IsExistInArray 判断string 是否存在在数组中
func (dao TableService) IsExistInArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// 判断是否是数据库字符串类型
func (dao TableService) IsStringObject(value string) bool {
	return dao.IsExistInArray(value, COLUMNTYPE_STR)
}

// 判断是否是数据库时间类型
func (dao TableService) IsTimeObject(value string) bool {
	return dao.IsExistInArray(value, COLUMNTYPE_TIME)
}

// IsNumberObject 判断是否是数据库数字类型
func (dao TableService) IsNumberObject(value string) bool {
	return dao.IsExistInArray(value, COLUMNTYPE_NUMBER)
}

// IsNotEdit 页面不需要编辑字段
func (dao TableService) IsNotEdit(value string) bool {
	return !dao.IsExistInArray(value, COLUMNNAME_NOT_EDIT)
}

// IsNotList 页面不需要显示的列表字段
func (dao TableService) IsNotList(value string) bool {
	return !dao.IsExistInArray(value, COLUMNNAME_NOT_LIST)
}

// IsNotQuery 页面不需要查询字段
func (dao TableService) IsNotQuery(value string) bool {
	return dao.IsExistInArray(value, COLUMNNAME_NOT_QUERY)
}

// COLUMNTYPE_STR 数据库字符串类型
var COLUMNTYPE_STR = []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext"}

// 数据库时间类型
var COLUMNTYPE_TIME = []string{"datetime", "time", "date", "timestamp"}

// 数据库数字类型
var COLUMNTYPE_NUMBER = []string{"tinyint", "smallint", "mediumint", "int", "number", "integer", "bigint", "float", "float", "double", "decimal"}

// 页面不需要编辑字段
var COLUMNNAME_NOT_EDIT = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time"}

// 页面不需要显示的列表字段
var COLUMNNAME_NOT_LIST = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time"}

// 页面不需要查询字段
var COLUMNNAME_NOT_QUERY = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time", "remark"}

func Contains(str, subStr string) bool {
	return strings.Contains(str, subStr)
}
