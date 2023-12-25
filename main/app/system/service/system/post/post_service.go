package post

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	"robvi/app/system/model/system/post"
	"robvi/app/system/service"
	"time"
)

// 根据主键查询数据
func SelectRecordById(id int64) (*post.SysPost, error) {
	entity := &post.SysPost{PostId: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除数据
func DeleteRecordById(id int64) bool {
	rs, err := (&post.SysPost{PostId: id}).Delete()
	if err == nil {
		if rs > 0 {
			return true
		}
	}
	return false
}

// 批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	ida := lv_conv.ToInt64Array(ids, ",")
	result, err := post.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

// 添加数据
func AddSave(req *post.AddReq, c *gin.Context) (int64, error) {
	var entity post.SysPost
	entity.PostName = req.PostName
	entity.PostCode = req.PostCode
	entity.Status = req.Status
	entity.PostSort = req.PostSort
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()
	entity.CreateBy = ""
	var userService service.UserService
	user := userService.GetProfile(c)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	_, err := entity.Insert()
	return entity.PostId, err
}

// 修改数据
func EditSave(req *post.EditReq, c *gin.Context) (int64, error) {
	entity := &post.SysPost{PostId: req.PostId}
	ok, err := entity.FindOne()
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}

	entity.PostName = req.PostName
	entity.PostCode = req.PostCode
	entity.Status = req.Status
	entity.Remark = req.Remark
	entity.PostSort = req.PostSort
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""
	var userService service.UserService
	user := userService.GetProfile(c)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return entity.Update()
}

// 根据条件分页查询角色数据
func SelectListAll(params *post.SelectPageReq) ([]post.EntityFlag, error) {
	return post.SelectListAll(params)
}

// 根据条件分页查询角色数据
func SelectListByPage(params *post.SelectPageReq) ([]post.SysPost, *lv_web.Paging, error) {
	return post.SelectListByPage(params)
}

// 导出excel
func Export(param *post.SelectPageReq) (string, error) {
	head := []string{"岗位序号", "岗位名称", "岗位编码", "岗位排序", "状态"}
	col := []string{"post_id", "post_name", "post_code", "post_sort", "stat"}
	return post.SelectListExport(param, head, col)
}

// 根据用户ID查询岗位
func SelectPostsByUserId(userId int64) ([]post.EntityFlag, error) {
	var paramsPost *post.SelectPageReq
	postAll, err := post.SelectListAll(paramsPost)

	if err != nil || postAll == nil {
		return nil, errors.New("未查询到岗位数据")
	}
	userPost, err := post.SelectPostsByUserId(userId)

	for i := range postAll {
		if userPost == nil {
			break
		}
		for j := range userPost {
			if userPost[j].PostId == postAll[i].PostId {
				postAll[i].Flag = true
				break
			}
		}
	}

	return postAll, nil
}

// 检查角色名是否唯一
func CheckPostNameUniqueAll(postName string) string {
	post, err := post.CheckPostNameUniqueAll(postName)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}

// 检查岗位名称是否唯一
func CheckPostNameUnique(postName string, postId int64) string {
	post, err := post.CheckPostNameUniqueAll(postName)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 && post.PostId != postId {
		return "1"
	}
	return "0"
}

// 检查岗位编码是否唯一
func CheckPostCodeUniqueAll(postCode string) string {
	post, err := post.CheckPostCodeUniqueAll(postCode)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}

// 检查岗位编码是否唯一
func CheckPostCodeUnique(postCode string, postId int64) string {
	post, err := post.CheckPostCodeUniqueAll(postCode)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 && post.PostId != postId {
		return "1"
	}
	return "0"
}
