package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_db/lv_dao"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_office"
	"system/dao"
	"system/model"
	"system/vo"
	"time"
)

type SysPostService struct {
}

// 根据主键查询数据
func (svc SysPostService) SelectRecordById(id int64) (*model.SysPost, error) {
	entity := &model.SysPost{PostId: id}
	err := entity.FindOne()
	return entity, err
}

// 根据主键删除数据
func (svc SysPostService) DeleteRecordById(id int64) bool {
	err := (&model.SysPost{PostId: id}).Delete()
	if err == nil {
		return true
	}
	return false
}

// 批量删除数据记录
func (svc SysPostService) DeleteRecordByIds(ids string) int64 {
	ida := lv_conv.ToInt64Array(ids, ",")
	var d dao.SysPostDao
	num, _ := d.DeleteByIds(ida)
	return num
}

// 添加数据
func (svc SysPostService) AddSave(req *vo.AddPostReq, c *gin.Context) (int64, error) {
	var entity model.SysPost
	entity.PostName = req.PostName
	entity.PostCode = req.PostCode
	entity.Status = req.Status
	entity.PostSort = req.PostSort
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()
	entity.CreateBy = ""
	var userService UserService
	user := userService.GetProfile(c)
	if user != nil {
		entity.CreateBy = user.LoginName
	}

	err := entity.Save()
	return entity.PostId, err
}

// 修改数据
func (svc SysPostService) EditSave(req *vo.EditSysPostReq, c *gin.Context) error {
	entity := &model.SysPost{PostId: req.PostId}
	err := entity.FindOne()
	if err != nil {
		return err
	}
	entity.PostName = req.PostName
	entity.PostCode = req.PostCode
	entity.Status = req.Status
	entity.Remark = req.Remark
	entity.PostSort = req.PostSort
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""
	var userService UserService
	user := userService.GetProfile(c)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return entity.Updates()
}

// 根据条件分页查询角色数据
func (svc SysPostService) SelectListAll(params *vo.SelectPostPageReq) (*[]model.SysPost, error) {
	var d dao.SysPostDao
	ret, err := d.ListAll(params)
	return ret, err
}

// 根据条件分页查询角色数据
func (svc SysPostService) SelectListByPage(params *vo.SelectPostPageReq) (*[]map[string]string, int64, error) {
	var d dao.SysPostDao
	return d.SelectPageList(params)
}

// 导出excel
func (svc SysPostService) Export(param *vo.SelectPostPageReq) (string, error) {
	head := []string{"岗位序号", "岗位名称", "岗位编码", "岗位排序", "状态"}
	col := []string{"post_id", "post_name", "post_code", "post_sort", "status"}
	var d dao.SysPostDao
	result, err := d.ListAllMap(param, false)
	url, err := lv_office.DownlaodExcelByListMapStr(&head, &col, result)
	return url, err
}

// 根据用户ID查询岗位
func (svc SysPostService) SelectPostsByUserId(userId int64) (*[]model.SysPost, error) {
	var paramsPost *vo.SelectPostPageReq
	var d dao.SysPostDao
	postAll, err := d.ListAll(paramsPost)

	if err != nil || postAll == nil {
		return nil, errors.New("未查询到岗位数据")
	}
	userPost, err := d.SelectPostsByUserId(userId)

	for i := range *postAll {
		if userPost == nil {
			break
		}
		for j := range *userPost {
			if (*userPost)[j].PostId == (*postAll)[i].PostId {
				(*postAll)[i].Selected = true
				break
			}
		}
	}

	return postAll, err
}

// IsPostCodeExist 检查岗位编码是否唯一
func (svc SysPostService) IsPostCodeExist(postCode string) (exist bool) {
	//total, err := d.CountCol("post_code", postCode)
	total, err := lv_dao.CountCol("sys_post", "post_code", postCode)
	lv_err.HasErrAndPanic(err)
	if total > 0 {
		exist = true
	}
	return
}
