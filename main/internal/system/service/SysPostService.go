package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/utils/lv_conv"
	"github.com/lv_framework/utils/lv_office"
	"main/internal/system/dao"
	"main/internal/system/model"
	"main/internal/system/vo"
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
	url, err := lv_office.DownlaodExcelByListMap(&head, &col, result)
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

// 检查角色名是否唯一
func (svc SysPostService) CheckPostNameUniqueAll(postName string) string {
	var d dao.SysPostDao
	post, err := d.CheckPostNameUniqueAll(postName)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}

// 检查岗位名称是否唯一
func (svc SysPostService) CheckPostNameUnique(postName string, postId int64) string {
	var d dao.SysPostDao
	post, err := d.CheckPostNameUniqueAll(postName)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 && post.PostId != postId {
		return "1"
	}
	return "0"
}

// 检查岗位编码是否唯一
func (svc SysPostService) CheckPostCodeUniqueAll(postCode string) string {
	var d dao.SysPostDao
	post, err := d.CheckPostCodeUniqueAll(postCode)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}

// 检查岗位编码是否唯一
func (svc SysPostService) CheckPostCodeUnique(postCode string, postId int64) string {
	var d dao.SysPostDao
	post, err := d.CheckPostCodeUniqueAll(postCode)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 && post.PostId != postId {
		return "1"
	}
	return "0"
}
