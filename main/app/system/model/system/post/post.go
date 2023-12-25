package post

import (
	"errors"
	"lostvip.com/db"
	"lostvip.com/utils/lv_office"
	"lostvip.com/utils/lv_web"
	userModel "robvi/app/system/model/system"
	"time"
	"xorm.io/builder"
)

// Fill with you ideas below.
// SysPost is the golang structure for table sys_post.
type EntityFlag struct {
	PostId     int64     `json:"post_id" xorm:"not null pk autoincr comment('岗位ID') BIGINT(20)"`
	PostCode   string    `json:"post_code" xorm:"not null comment('岗位编码') VARCHAR(64)"`
	PostName   string    `json:"post_name" xorm:"not null comment('岗位名称') VARCHAR(50)"`
	PostSort   int       `json:"post_sort" xorm:"not null comment('显示顺序') INT(4)"`
	Status     string    `json:"status" xorm:"not null comment('状态（0正常 1停用）') CHAR(1)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
	Flag       bool      `json:"flag"` // 标记
}

// 新增页面请求参数
type AddReq struct {
	PostName string `form:"postName"  binding:"required"`
	PostCode string `form:"postCode"  binding:"required"`
	PostSort int    `form:"postSort"  binding:"required"`
	Status   string `form:"status"    binding:"required"`
	Remark   string `form:"remark"`
}

// 修改页面请求参数
type EditReq struct {
	PostId   int64  `form:"postId" binding:"required"`
	PostName string `form:"postName"  binding:"required"`
	PostCode string `form:"postCode"  binding:"required"`
	PostSort int    `form:"postSort"  binding:"required"`
	Status   string `form:"status"    binding:"required"`
	Remark   string `form:"remark"`
}

// 分页请求参数
type SelectPageReq struct {
	PostCode      string `form:"postCode"`      //岗位编码
	Status        string `form:"status"`        //状态
	PostName      string `form:"postName"`      //岗位名称
	BeginTime     string `form:"beginTime"`     //开始时间
	EndTime       string `form:"endTime"`       //结束时间
	OrderByColumn string `form:"orderByColumn"` //排序字段
	IsAsc         string `form:"isAsc"`         //排序方式
	PageNum       int    `form:"pageNum"`       //当前页码
	PageSize      int    `form:"pageSize"`      //每页数
	Remark        string `form:"remark"`        //每页数
}

// 检查编码请求参数
type CheckPostCodeReq struct {
	PostId   int64  `form:"postId"  binding:"required"`
	PostCode string `form:"postCode"  binding:"required"`
}

// 检查编码请求参数
type CheckPostCodeALLReq struct {
	PostCode string `form:"postCode"  binding:"required"`
}

// 检查名称请求参数
type CheckPostNameReq struct {
	PostId   int64  `form:"postId"  binding:"required"`
	PostName string `form:"postName"  binding:"required"`
}

// 检查名称请求参数
type CheckPostNameALLReq struct {
	PostName string `form:"postName"  binding:"required"`
}

// 根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) ([]SysPost, *lv_web.Paging, error) {
	db := db.GetInstance().Engine()
	p := new(lv_web.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("p")

	if param != nil {
		if param.PostCode != "" {
			model.Where("p.post_code like ?", "%"+param.PostCode+"%")
		}

		if param.Status != "" {
			model.Where("p.status = ", param.Status)
		}

		if param.PostName != "" {
			model.Where("p.post_name like ?", "%"+param.PostName+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(p.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(p.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	totalModel := model.Clone()

	total, err := totalModel.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = lv_web.CreatePaging(param.PageNum, param.PageSize, int(total))

	if param.OrderByColumn != "" {
		model.OrderBy(param.OrderByColumn + " " + param.IsAsc + " ")
	}

	model.Limit(p.Pagesize, p.StartNum)

	var result []SysPost
	err = model.Find(&result)
	return result, p, err
}

// 导出excel
func SelectListExport(param *SelectPageReq, head, col []string) (string, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(TableName(), "t")

	if param != nil {
		if param.PostCode != "" {
			build.Where(builder.Like{"t.post_code", param.PostCode})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"t.status": param.Status})
		}

		if param.PostName != "" {
			build.Where(builder.Like{"t.post_name", param.PostName})
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
func SelectListAll(param *SelectPageReq) ([]EntityFlag, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("p").Select("p.*,false as flag")
	if param != nil {

		if param.PostCode != "" {
			model.Where("p.post_code like ?", "%"+param.PostCode+"%")
		}

		if param.Status != "" {
			model.Where("p.status = ", param.Status)
		}

		if param.PostName != "" {
			model.Where("p.post_name like ?", "%"+param.PostName+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(p.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(p.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []EntityFlag
	err := model.Find(&result)
	return result, err
}

// 根据用户ID查询岗位
func SelectPostsByUserId(userId int64) ([]EntityFlag, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(userModel.TableName()).Alias("u")
	model.Join("LEFT", []string{"sys_user_post", "up"}, "u.user_id = up.user_id")
	model.Join("LEFT", []string{"sys_post", "p"}, "up.post_id = p.post_id")
	model.Where("up.user_id = ?", userId)
	model.Select("p.post_id, p.post_name, p.post_code,false as flag")
	var result []EntityFlag
	err := model.Find(&result)
	return result, err
}

// 校验岗位名称是否唯一
func CheckPostNameUniqueAll(postName string) (*SysPost, error) {
	var entity SysPost
	entity.PostName = postName
	ok, err := entity.FindOne()
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}

// 校验岗位名称是否唯一
func CheckPostCodeUniqueAll(postCode string) (*SysPost, error) {
	var entity SysPost
	entity.PostCode = postCode
	ok, err := entity.FindOne()
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}
