package post

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/convert"
	"lostvip.com/utils/page"
	post2 "robvi/app/modules/sys/model/system/post"
	"robvi/app/modules/sys/service"
	"time"
)

// 根据主键查询数据
func SelectRecordById(id int64) (*post2.SysPost, error) {
	entity := &post2.SysPost{PostId: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据主键删除数据
func DeleteRecordById(id int64) bool {
	rs, err := (&post2.SysPost{PostId: id}).Delete()
	if err == nil {
		if rs > 0 {
			return true
		}
	}
	return false
}

// 批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	ida := convert.ToInt64Array(ids, ",")
	result, err := post2.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

// 添加数据
func AddSave(req *post2.AddReq, c *gin.Context) (int64, error) {
	var entity post2.SysPost
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
func EditSave(req *post2.EditReq, c *gin.Context) (int64, error) {
	entity := &post2.SysPost{PostId: req.PostId}
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
func SelectListAll(params *post2.SelectPageReq) ([]post2.EntityFlag, error) {
	return post2.SelectListAll(params)
}

// 根据条件分页查询角色数据
func SelectListByPage(params *post2.SelectPageReq) ([]post2.SysPost, *page.Paging, error) {
	return post2.SelectListByPage(params)
}

// 导出excel
func Export(param *post2.SelectPageReq) (string, error) {
	head := []string{"岗位序号", "岗位名称", "岗位编码", "岗位排序", "状态"}
	col := []string{"post_id", "post_name", "post_code", "post_sort", "stat"}
	return post2.SelectListExport(param, head, col)
}

// 根据用户ID查询岗位
func SelectPostsByUserId(userId int64) ([]post2.EntityFlag, error) {
	var paramsPost *post2.SelectPageReq
	postAll, err := post2.SelectListAll(paramsPost)

	if err != nil || postAll == nil {
		return nil, errors.New("未查询到岗位数据")
	}

	userPost, err := post2.SelectPostsByUserId(userId)

	if err != nil || userPost == nil {
		return nil, errors.New("未查询到用户岗位数据")
	} else {
		for i := range postAll {
			for j := range userPost {
				if userPost[j].PostId == postAll[i].PostId {
					postAll[i].Flag = true
					break
				}
			}
		}
	}

	return postAll, nil
}

// 检查角色名是否唯一
func CheckPostNameUniqueAll(postName string) string {
	post, err := post2.CheckPostNameUniqueAll(postName)
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
	post, err := post2.CheckPostNameUniqueAll(postName)
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
	post, err := post2.CheckPostCodeUniqueAll(postCode)
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
	post, err := post2.CheckPostCodeUniqueAll(postCode)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 && post.PostId != postId {
		return "1"
	}
	return "0"
}
