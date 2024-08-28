// //////////////////////////////////////////////////////////////////////////////////////
// from Dotsql
// It is not an ORM, it is not a query builder.
// Dotsql is a library that helps you keep sql files in one place and use it with ease.
// For more usage examples see https://github.com/qustavo/dotsql
// /////////////////////////////////////////////////////////////////////////////////////
package lv_batis

import (
	"bufio"
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_file"
	"github.com/lostvip-com/lv_framework/utils/lv_tpl"
	"github.com/morrisxyang/xreflect"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"io"
	"os"
	"reflect"
	"strings"
)

// Execer is an interface used by Exec.
type Execer interface {
	Exec(query string, args ...interface{}) *gorm.DB
}

// ExecerContext is an interface used by ExecContext.
type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type LvBatis struct {
	Queries     map[string]string
	TplFile     string
	CurrBaseSql string
}

/**
 * 从mapper目录解析sql文件
 */
func NewInstance(relativePath string) (string *LvBatis) {
	basePath, _ := os.Getwd()
	absolutePath := basePath + "/mapper" //为了方便管理，必须把映射文件放到mapper目录
	if strings.HasPrefix(relativePath, "/") {
		absolutePath = absolutePath + relativePath
	} else {
		absolutePath = absolutePath + "/" + relativePath
	}
	dot, err := LoadFromFile(absolutePath)
	dot.TplFile = relativePath
	if err != nil {
		panic(err)
	}
	return dot
}

func (e *LvBatis) GetSql(tagName string, params interface{}) (string, error) {
	query, err := e.LookupQuery(tagName)
	if err != nil || query == "" {
		panic("tpl文件格式错误!")
	}
	//动态解析
	sql, err := lv_tpl.ParseTemplateStr(query, params)
	if sql == "" || err != nil {
		lv_log.Error(err)
		panic(e.getTplFile() + " 可能存在错误：<p/>1.使用了参数对象中不存在的属性<p/>2.template语法错误！")
	}
	e.CurrBaseSql = sql //缓存当前正在执行的分页sql
	return sql, err
}

/**
 * 从mapper目录解析sql文件
 */
func (e *LvBatis) GetLimitSql(tagName string, params interface{}) (string, error) {
	sql, err := e.GetSql(tagName, params)
	if err != nil {
		panic(err)
	}
	paramType := reflect.TypeOf(params).Kind()
	if paramType == reflect.Map {
		paramMap := params.(map[string]interface{})
		pageNum := paramMap["pageNum"]
		pageSize := paramMap["pageSize"]
		if pageSize == nil || pageNum == nil {
			panic("???????????分页信息错误" + sql)
		} else {
			start := cast.ToInt64(pageSize) * (cast.ToInt64(pageNum) - 1)
			sql = sql + " limit  " + cast.ToString(start) + "," + cast.ToString(pageSize)
		}
	} else {
		pageNum, err1 := xreflect.FieldValue(params, "PageNum")
		pageSize, err2 := xreflect.FieldValue(params, "PageSize")
		if pageSize == 0 || pageNum == 0 || err1 != nil || err2 != nil {
			lv_log.Info("XXXX=====>pageSize:", pageSize, " pageNum:", pageNum)
			panic("分页参数错误，请检查分页参数pageSize和pageNum的值")
		} else {
			start := cast.ToInt64(pageSize) * (cast.ToInt64(pageNum) - 1)
			sql = sql + " limit  " + cast.ToString(start) + "," + cast.ToString(pageSize)
		}
	}

	return sql, err
}

func (e *LvBatis) GetCountSql() string {
	if e.CurrBaseSql == "" {
		panic("未初始化过ibatis对象,未传入过sql参数！" + e.getTplFile())
	}
	sql := " select count(1)  from (" + e.CurrBaseSql + ") t "
	return sql
}

func (d LvBatis) LookupQuery(name string) (query string, err error) {
	query, ok := d.Queries[name]
	if !ok {
		err = fmt.Errorf("dotsql: '%s' could not be found", name)
	}

	return
}

// Exec is a wrapper for database/sql's Exec(), using dotsql named query.
func (d LvBatis) Exec(db Execer, name string, args ...interface{}) (*gorm.DB, error) {
	query, err := d.LookupQuery(name)
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...), err
}

// ExecContext is a wrapper for database/sql's ExecContext(), using dotsql named query.
func (d LvBatis) ExecContext(ctx context.Context, db ExecerContext, name string, args ...interface{}) (sql.Result, error) {
	query, err := d.LookupQuery(name)
	if err != nil {
		return nil, err
	}

	return db.ExecContext(ctx, query, args...)
}

// GetRawSql returns the query, everything after the --name tag
func (d LvBatis) GetRawSql(name string) (string, error) {
	return d.LookupQuery(name)
}

// GetQueryMap returns a map[string]string of loaded Queries
func (d LvBatis) GetQueryMap() map[string]string {
	return d.Queries
}

func (e *LvBatis) getTplFile() string {
	return e.TplFile
}

// Load imports sql Queries from any io.Reader.
func Load(r io.Reader) (*LvBatis, error) {
	scanner := &Scanner{}
	queries := scanner.Run(bufio.NewScanner(r))

	dotsql := &LvBatis{
		Queries: queries,
	}

	return dotsql, nil
}

// LoadFromFile imports SQL Queries from the file.
func LoadFromFile(sqlFile string) (*LvBatis, error) {
	if !lv_file.IsFileExist(sqlFile) {
		panic("生成代码后再执行此操作!未发现文件 " + sqlFile)
	}
	f, err := os.Open(sqlFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Load(f)
}

// LoadFromString imports SQL Queries from the string.
func LoadFromString(sql string) (*LvBatis, error) {
	buf := bytes.NewBufferString(sql)
	return Load(buf)
}

// Merge takes one or more *LvBatis and merge its Queries
// It's in-order, so the last source will override Queries with the same name
// in the previous arguments if any.
func Merge(dots ...*LvBatis) *LvBatis {
	queries := make(map[string]string)

	for _, dot := range dots {
		for k, v := range dot.GetQueryMap() {
			queries[k] = v
		}
	}

	return &LvBatis{
		Queries: queries,
	}
}
