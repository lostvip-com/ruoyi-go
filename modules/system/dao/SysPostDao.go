package dao

import (
	"errors"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/namedsql"
	"github.com/spf13/cast"
	"system/model"
	"system/vo"
)

type SysPostDao struct {
}

func (e SysPostDao) DeleteByIds(ida []int64) (int64, error) {
	db := lv_db.GetMasterGorm().Table("sys_post").Where("post_id in ? ", ida).Update("del_flag", 1)
	return db.RowsAffected, db.Error
}

// 根据条件分页查询用户列表
func (d SysPostDao) SelectPageList(param *vo.SelectPostPageReq) (*[]map[string]string, int64, error) {
	db := lv_db.GetMasterGorm()
	sqlParams, sql := d.GetSql(param)
	limitSql := sql + " order by u.post_id desc "
	limitSql += "  limit " + cast.ToString(param.GetStartNum()) + "," + cast.ToString(param.GetPageSize())
	result, err := namedsql.ListMapStr(db, limitSql, sqlParams, true)
	if err != nil {
		return nil, 0, err
	}
	countSql := "select count(*) from (" + sql + ") t "
	total, err := namedsql.Count(db, countSql, sqlParams)
	return result, total, err
}

func (d SysPostDao) GetSql(param *vo.SelectPostPageReq) (map[string]interface{}, string) {
	sqlParams := make(map[string]interface{})
	sql := `
            select * from sys_post u where u.del_flag =0 
           `
	if param != nil {
		if param.PostName != "" {
			sql += " and  u.post_name like @PostName "
			sqlParams["PostName"] = "%" + param.PostName + "%"
		}

		if param.Status != "" {
			sql += " and  u.status = @Status "
			sqlParams["Status"] = param.Status
		}

		if param.BeginTime != "" {
			sql += " and  u.create_time >= @BeginTime "
			sqlParams["BeginTime"] = param.BeginTime
		}
		if param.EndTime != "" {
			sql += " and  u.create_time <= @EndTime "
			sqlParams["EndTime"] = param.EndTime
		}

	}
	return sqlParams, sql
}

// 导出excel
func (d SysPostDao) ListAll(param *vo.SelectPostPageReq) (*[]model.SysPost, error) {
	db := lv_db.GetMasterGorm()
	sqlParams, sql := d.GetSql(param)
	allSql := sql + " order by u.post_id desc "
	result, err := namedsql.ListData[model.SysPost](db, allSql, &sqlParams)
	return result, err
}

// 导出excel
func (d SysPostDao) ListAllMap(param *vo.SelectPostPageReq, camel bool) (*[]map[string]string, error) {
	db := lv_db.GetMasterGorm()
	sqlParams, sql := d.GetSql(param)
	allSql := sql + " order by u.post_id desc "
	result, err := namedsql.ListMapStr(db, allSql, &sqlParams, camel)
	return result, err
}

// 根据用户ID查询岗位
func (dao SysPostDao) SelectPostsByUserId(userId int64) (*[]model.SysPost, error) {
	db := lv_db.GetMasterGorm()
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	sql := `
            select p.post_id, p.post_name, p.post_code,false as selected 
            from sys_post p,sys_user_post up 
            where p.del_flag =0 and p.post_id = up.post_id
                  and up.user_id = @UserId
           `
	sqlParams := map[string]any{"UserId": 2}
	//sqlParams := model.SysUserPost{UserId: 2}
	result, err := namedsql.ListData[model.SysPost](db, sql, sqlParams)
	return result, err
}

// CountCol 按字段值统计数量
func (dao SysPostDao) CountCol(column, value string) (total int64, err error) {
	err = lv_db.GetMasterGorm().Table("sys_post").Where("del_flag=0 and "+column+"=?", value).Count(&total).Error
	return
}
