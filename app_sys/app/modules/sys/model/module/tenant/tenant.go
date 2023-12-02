// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-06-29 22:21:21 +0800 CST
// 生成路径: app/model/modules/tenant/tenant.go
// 生成人：robnote
// ==========================================================================
package tenant

import (
	"errors"
	"lostvip.com/db"
	"lostvip.com/utils/lv_office"
	"lostvip.com/utils/lv_web"
	"xorm.io/builder"
)

// 新增页面请求参数
type AddReq struct {
	DelFlag   string `form:"delFlag" `
	Name      string `form:"name" binding:"required"`
	Address   string `form:"address" `
	Manager   string `form:"manager" `
	Phone     string `form:"phone" `
	Remark    string `form:"remark" `
	StartTime string `form:"startTime" `
	EndTime   string `form:"endTime" `
	Email     string `form:"email" `
}

// 修改页面请求参数
type EditReq struct {
	Id        int64  `form:"id" binding:"required"`
	Name      string `form:"name" binding:"required"`
	Address   string `form:"address" `
	Manager   string `form:"manager" `
	Phone     string `form:"phone" `
	Remark    string `form:"remark" `
	StartTime string `form:"startTime" `
	EndTime   string `form:"endTime" `
	Email     string `form:"email" `
}

// 分页请求参数
type SelectPageReq struct {
	Name      string `form:"name"`      //商户名称
	Address   string `form:"address"`   //联系地址
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	PageNum   int    `form:"pageNum"`   //当前页码
	PageSize  int    `form:"pageSize"`  //每页数
}

// 根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) ([]SysTenant, *lv_web.Paging, error) {
	db := db.Instance().Engine()
	paging := new(lv_web.Paging)

	if db == nil {
		return nil, paging, errors.New("获取数据库连接失败")
	}

	model := db.Table("sys_tenant").Alias("t")

	if param != nil {

		if param.Name != "" {
			model.Where("t.name like ?", "%"+param.Name+"%")
		}

		if param.Address != "" {
			model.Where("t.address = ?", param.Address)
		}
		if param.BeginTime != "" {
			model.Where("t.create_time >= ?", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("t.create_time<=?", param.EndTime)
		}
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, paging, errors.New("读取行数失败")
	}

	paging = lv_web.CreatePaging(param.PageNum, param.PageSize, int(total))
	model.Limit(paging.Pagesize, paging.StartNum)

	var result []SysTenant
	err = model.Find(&result)
	return result, paging, err
}

// 导出excel
func SelectListExport(param *SelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From("sys_tenant", "t")

	if param != nil {

		if param.Name != "" {
			build.Where(builder.Like{"t.name", param.Name})
		}

		if param.Address != "" {
			build.Where(builder.Eq{"t.address": param.Address})
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

// 获取所有数据
func SelectListAll(param *SelectPageReq) ([]SysTenant, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table("sys_tenant").Alias("t")

	if param != nil {

		if param.Name != "" {
			model.Where("t.name like ?", "%"+param.Name+"%")
		}

		if param.Address != "" {
			model.Where("t.address = ?", param.Address)
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []SysTenant
	err := model.Find(&result)
	return result, err
}
