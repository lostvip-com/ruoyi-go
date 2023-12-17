// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-12-17 19:23:47 +0800 CST
// 生成人：lv
// ==========================================================================
package dao

import (
	"errors"
	 "xorm.io/builder"
    "lostvip.com/db"
    "lostvip.com/utils/lv_office"
    "lostvip.com/utils/lv_web"
    "robvi/app/mywork/dto"
    "robvi/app/mywork/model"
)

//新增页面请求参数
type DpcTaskDao struct { }

//根据条件分页查询数据
func (e DpcTaskDao) SelectListByPage(param *dto.PageDpcTaskReq) ([]model.DpcTask, *lv_web.Paging, error) {
	db := db.Instance().Engine()
    p := new(lv_web.Paging)

	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	query := db.Table("dpc_task").Alias("t")

	if param != nil {    
		
		if param.Username != "" {
			query.Where("t.username like ?", "%"+param.Username+"%")
		}    
		
		if param.Password != "" {
			query.Where("t.password like ?", "%"+param.Password+"%")
		}    
		
		if param.PrjCode != "" {
			query.Where("t.prj_code like ?", "%"+param.PrjCode+"%")
		}    
		
		if param.TaskContent != "" {
			query.Where("t.task_content like ?", "%"+param.TaskContent+"%")
		}                         
		if param.BeginTime != "" {
			query.Where("t.create_time >= ?", param.BeginTime)
		}

		if param.EndTime != "" {
			 query.Where("t.create_time<=?", param.EndTime)
		}
	}

	total, err := query.Clone().Count()
	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = lv_web.CreatePaging(param.PageNum, param.PageSize, int(total))
	query.Limit(p.Pagesize, p.StartNum)

	var result []model.DpcTask
    err = query.Find(&result)
    return result, p, err
}

// 导出excel
func (e DpcTaskDao) SelectListExport(param *dto.PageDpcTaskReq, head, col []string) (string, error) {
	db := db.Instance().Engine()
	build := builder.Select(col...).From("dpc_task", "t")

	if param != nil {    
		
		if param.Username != "" {
			build.Where(builder.Like{"t.username", param.Username})
		}    
		
		if param.Password != "" {
			build.Where(builder.Like{"t.password", param.Password})
		}    
		
		if param.PrjCode != "" {
			build.Where(builder.Like{"t.prj_code", param.PrjCode})
		}    
		
		if param.TaskContent != "" {
			build.Where(builder.Like{"t.task_content", param.TaskContent})
		}                         
		if param.BeginTime != "" {
			build.Where(builder.Gte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.BeginTime + "','%y%m%d')"})
		}

		if param.EndTime != "" {
			build.Where(builder.Lte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.EndTime + "','%y%m%d')"})
		}
	}

	sqlStr, _, _ := build.ToSQL()
	arr, err := db.SQL(sqlStr).QuerySliceString()
	path, err := lv_office.DownlaodExcel(head, arr)

	return path, err
}

//获取所有数据
func (e DpcTaskDao) SelectListAll(param *dto.PageDpcTaskReq) ([]model.DpcTask, error) {
	db := db.Instance().Engine()
	query := db.Table("dpc_task").Alias("t")

	if param != nil {    
		
		if param.Username != "" {
			query.Where("t.username like ?", "%"+param.Username+"%")
		}    
		
		if param.Password != "" {
			query.Where("t.password like ?", "%"+param.Password+"%")
		}    
		
		if param.PrjCode != "" {
			query.Where("t.prj_code like ?", "%"+param.PrjCode+"%")
		}    
		
		if param.TaskContent != "" {
			query.Where("t.task_content like ?", "%"+param.TaskContent+"%")
		}                         
		if param.BeginTime != "" {
			query.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			query.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []model.DpcTask
	err := query.Find(&result)
	return result, err
}



//批量删除
func (e DpcTaskDao)DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table("dpc_task").In("id", ids).Delete(new(model.DpcTask))
}

// 根据条件查询
func (e DpcTaskDao)Find(where, order string) ([]model.DpcTask, error) {
	var list []model.DpcTask
	err := db.Instance().Engine().Table("dpc_task").Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (e DpcTaskDao)FindIn(column string, args ...interface{}) ([]model.DpcTask, error) {
	var list []model.DpcTask
	err := db.Instance().Engine().Table("dpc_task").In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (e DpcTaskDao)  FindNotIn(column string, args ...interface{}) ([]model.DpcTask, error) {
	var list []model.DpcTask
	err := db.Instance().Engine().Table("dpc_task").NotIn(column, args).Find(&list)
	return list, err
}