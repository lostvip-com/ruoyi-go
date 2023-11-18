package tool

import (
	"errors"
	"lostvip.com/db"
	"lostvip.com/utils/page"
)

// Fill with you ideas below.

// GenTable is the golang structure for table gen_table.
type EntityExtend struct {
	GenTable       `xorm:"extends"`
	TreeCode       string   `xorm:"-"` // 树编码字段
	TreeParentCode string   `xorm:"-"` // 树父编码字段
	TreeName       string   `xorm:"-"` // 树名称字段
	Columns        []Entity `xorm:"-"` // 表列信息
	PkColumn       Entity   `xorm:"-"` // 表列信息
}

type Params struct {
	TreeCode       string `form:"treeCode"`
	TreeParentCode string `form:"treeParentCode"`
	TreeName       string `form:"treeName"`
}

// 修改页面请求参数
type EditReq struct {
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
type SelectPageReq struct {
	TableName    string `form:"tableName"`    //表名称
	TableComment string `form:"tableComment"` //表描述
	BeginTime    string `form:"beginTime"`    //开始时间
	EndTime      string `form:"endTime"`      //结束时间
	PageNum      int    `form:"pageNum"`      //当前页码
	PageSize     int    `form:"pageSize"`     //每页数
}

// 根据ID获取记录
func (r *GenTable) SelectRecordById(id int64) (*EntityExtend, error) {
	db := db.Instance().Engine()
	var result EntityExtend
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(r.TableName()).Where("table_id=?", id)
	ok, err := model.Get(&result)
	if !ok {
		return nil, err
	}

	//表数据列
	columModel := db.Table("gen_table_column").Where("table_id=?", id)

	var columList []Entity
	err = columModel.Find(&columList)

	if err != nil {
		return nil, err
	}
	result.Columns = columList
	return &result, nil
}

// 根据条件分页查询数据
func (r *GenTable) SelectListByPage(param *SelectPageReq) ([]GenTable, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table(r.TableName()).Alias("t")

	if param != nil {
		if param.TableName != "" {
			model.Where("t.table_name like ?", "%"+param.TableName+"%")
		}

		if param.TableComment != "" {
			model.Where("t.table_comment like ?", "%"+param.TableComment+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	totalModel := model.Clone()

	total, err := totalModel.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

	model.Limit(p.Pagesize, p.StartNum)
	var result []GenTable
	err = model.Find(&result)

	return result, p, err
}

// 查询据库列表
func SelectDbTableList(param *SelectPageReq) ([]GenTable, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table("information_schema.tables")
	model.Where("table_schema = (select database())")
	model.Where("table_name NOT LIKE 'qrtz_%' AND table_name NOT LIKE 'gen_%'")
	model.Where("table_name NOT IN (select table_name from gen_table)")
	if param != nil {
		if param.TableName != "" {
			model.Where("lower(table_name) like lower(?)", "%"+param.TableName+"%")
		}

		if param.TableComment != "" {
			model.Where("lower(table_comment) like lower(?)", "%"+param.TableComment+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	totalModel := model.Clone()

	total, err := totalModel.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

	model.Select("table_name, table_comment, create_time, update_time")
	model.Limit(p.Pagesize, p.StartNum)

	var result []GenTable
	err = model.Find(&result)
	return result, p, err
}

// 查询据库列表
func SelectDbTableListByNames(tableNames []string) ([]GenTable, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table("information_schema.tables")
	model.Select("0 as table_id, table_name, table_comment,'' as class_name,'' as tpl_category,'' as package_name,'' as module_name,'' as business_name,'' as function_name,'' as function_author,'' as options,'' as create_by, create_time,'' as update_by, update_time,'' as remark")
	model.Where("table_name NOT LIKE 'qrtz_%'")
	model.Where("table_name NOT LIKE 'gen_%'")
	model.Where("table_schema = (select database())")
	if len(tableNames) > 0 {
		model.In("table_name", tableNames)
	}

	var result []GenTable
	err := model.Find(&result)
	return result, err
}

// 查询据库列表
func SelectTableByName(tableName string) (*GenTable, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table("information_schema.tables")
	model.Select("0 as table_id, table_name, table_comment,'' as class_name,'' as tpl_category,'' as package_name,'' as module_name,'' as business_name,'' as function_name,'' as function_author,'' as options,'' as create_by, create_time,'' as update_by, update_time,'' as remark")
	model.Where("table_comment <> ''")
	model.Where("table_schema = (select database())")
	if tableName != "" {
		model.Where("table_name = ?", tableName)
	}

	var result GenTable
	_, err := model.Get(&result)
	return &result, err
}

// 查询表ID业务信息
func (r *GenTable) SelectGenTableById(tableId int64) (*GenTable, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(r.TableName()).Alias("t")
	model.Join("LEFT", []string{"gen_table_column", "c"}, "t.table_id = c.table_id")
	model.Where("t.table_id = ?", tableId)
	model.Select("t.table_id, t.table_name, t.table_comment, t.class_name, t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.options, t.remark,c.column_id, c.column_name, c.column_comment, c.column_type, c.is_pk, c.is_increment, c.is_required, c.is_insert, c.is_edit, c.is_list, c.is_query, c.query_type, c.html_type, c.dict_type, c.sort")

	var result GenTable
	_, err := model.Get(&result)
	return &result, err
}

func (r *GenTable) SelectableExtendById(tableId int64) (*EntityExtend, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(r.TableName()).Alias("t")
	model.Where("t.table_id = ?", tableId)

	var result EntityExtend
	_, err := model.Get(&result)
	return &result, err
}

// 查询表名称业务信息
func (r *GenTable) SelectGenTableByName(tableName string) (*GenTable, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(r.TableName()).Alias("t")
	model.Join("LEFT", []string{"gen_table_column", "c"}, "t.table_id = c.table_id")
	model.Where("t.table_name = ?", tableName)
	model.Select("t.table_id, t.table_name, t.table_comment, t.class_name, t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.options, t.remark,c.column_id, c.column_name, c.column_comment, c.column_type, c.java_type, c.java_field, c.is_pk, c.is_increment, c.is_required, c.is_insert, c.is_edit, c.is_list, c.is_query, c.query_type, c.html_type, c.dict_type, c.sort")

	var result GenTable
	_, err := model.Get(&result)
	return &result, err
}
