// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-02-17 14:03:51
// 生成路径: app/model_cmn/online/online.go
// 生成人：yunjie
// ==========================================================================
package online

import (
	"errors"
	"github.com/lostvip-com/lv_framework/db"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	"time"
)

// 新增页面请求参数
type AddReq struct {
	LoginName      string    `form:"loginName" binding:"required"`
	DeptName       string    `form:"deptName" binding:"required"`
	Ipaddr         string    `form:"ipaddr" `
	LoginLocation  string    `form:"loginLocation" `
	Browser        string    `form:"browser" `
	Os             string    `form:"os" `
	Status         string    `form:"status" binding:"required"`
	StartTimestamp time.Time `form:"startTimestamp" `
	LastAccessTime time.Time `form:"lastAccessTime" `
	ExpireTime     int       `form:"expireTime" `
}

// 修改页面请求参数
type EditReq struct {
	SessionId      string    `form:"sessionId" binding:"required"`
	LoginName      string    `form:"loginName" binding:"required"`
	DeptName       string    `form:"deptName" binding:"required"`
	Ipaddr         string    `form:"ipaddr" `
	LoginLocation  string    `form:"loginLocation" `
	Browser        string    `form:"browser" `
	Os             string    `form:"os" `
	Status         string    `form:"status" binding:"required"`
	StartTimestamp time.Time `form:"startTimestamp" `
	LastAccessTime time.Time `form:"lastAccessTime" `
	ExpireTime     int       `form:"expireTime" `
}

// 分页请求参数
type SelectPageReq struct {
	SessionId      string    `form:"sessionId"`      //用户会话id
	LoginName      string    `form:"loginName"`      //登录账号
	DeptName       string    `form:"deptName"`       //部门名称
	Ipaddr         string    `form:"ipaddr"`         //登录IP地址
	LoginLocation  string    `form:"loginLocation"`  //登录地点
	Browser        string    `form:"browser"`        //浏览器类型
	Os             string    `form:"os"`             //操作系统
	Status         string    `form:"status"`         //在线状态on_line在线off_line离线
	StartTimestamp time.Time `form:"startTimestamp"` //session创建时间
	LastAccessTime time.Time `form:"lastAccessTime"` //session最后访问时间
	ExpireTime     int       `form:"expireTime"`     //超时时间，单位为分钟
	BeginTime      string    `form:"beginTime"`      //开始时间
	EndTime        string    `form:"endTime"`        //结束时间
	PageNum        int       `form:"pageNum"`        //当前页码
	PageSize       int       `form:"pageSize"`       //每页数
	SortName       string    `form:"sortName"`       //排序字段
	SortOrder      string    `form:"sortOrder"`      //排序方式
}

// 根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) ([]UserOnline, *lv_web.Paging, error) {
	db := db.GetInstance().Engine()
	p := new(lv_web.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table("sys_user_online").Alias("t")

	if param != nil {

		if param.SessionId != "" {
			model.Where("t.session_id = ?", param.SessionId)
		}

		if param.LoginName != "" {
			model.Where("t.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.DeptName != "" {
			model.Where("t.dept_name like ?", "%"+param.DeptName+"%")
		}

		if param.Ipaddr != "" {
			model.Where("t.ipaddr = ?", param.Ipaddr)
		}

		if param.LoginLocation != "" {
			model.Where("t.login_location = ?", param.LoginLocation)
		}

		if param.Browser != "" {
			model.Where("t.browser = ?", param.Browser)
		}

		if param.Os != "" {
			model.Where("t.os = ?", param.Os)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
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

	p = lv_web.CreatePaging(param.PageNum, param.PageSize, int64(total))

	if param.SortName != "" {
		model.OrderBy(param.SortName + " " + param.SortOrder + " ")
	}

	model.Limit(p.PageSize, p.StartNum)

	var result []UserOnline
	err = model.Find(&result)

	return result, p, err
}

// 获取所有数据
func SelectListAll(param *SelectPageReq) ([]UserOnline, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table("sys_user_online").Alias("t")

	if param != nil {

		if param.SessionId != "" {
			model.Where("t.session_id = ?", param.SessionId)
		}

		if param.LoginName != "" {
			model.Where("t.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.DeptName != "" {
			model.Where("t.dept_name like ?", "%"+param.DeptName+"%")
		}

		if param.Ipaddr != "" {
			model.Where("t.ipaddr = ?", param.Ipaddr)
		}

		if param.LoginLocation != "" {
			model.Where("t.login_location = ?", param.LoginLocation)
		}

		if param.Browser != "" {
			model.Where("t.browser = ?", param.Browser)
		}

		if param.Os != "" {
			model.Where("t.os = ?", param.Os)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []UserOnline
	err := model.Find(&result)
	return result, err
}

// 批量删除除参数以外的数据
func DeleteNotIn(ids ...string) (int64, error) {
	return db.GetInstance().Engine().Table("sys_user_online").NotIn("session_id", ids).Delete(new(UserOnline))
}
