package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/gconv"
	"robvi/app/modules/sys/model"
	dept2 "robvi/app/modules/sys/model/system/dept"
	"strings"
	"time"
)

type DeptService struct{}

// 新增保存信息
func (svc *DeptService) AddSave(req *dept2.AddReq, c *gin.Context) (int64, error) {
	dept0 := new(dept2.SysDept)
	parent := &dept2.SysDept{DeptId: req.ParentId}
	ok, err := parent.FindOne()
	if err == nil && ok {
		if parent.Status != "0" {
			return 0, errors.New("部门停用，不允许新增")
		}
	} else {
		return 0, errors.New("父部门不能为空")
	}
	var userService UserService
	dept0.DeptName = req.DeptName
	dept0.Status = req.Status
	dept0.ParentId = req.ParentId
	dept0.DelFlag = "0"
	dept0.Email = req.Email
	dept0.Leader = req.Leader
	dept0.Phone = req.Phone
	dept0.OrderNum = req.OrderNum
	dept0.TenantId = req.TenantId
	user := userService.GetProfile(c)

	if user != nil && user.UserId > 0 {
		dept0.CreateBy = user.LoginName
	}

	dept0.CreateTime = time.Now()

	_, err = dept0.Insert()
	//这里跟原版不一样了，多加了一级自己的ID，以方便数据权限控制
	dept0.Ancestors = parent.Ancestors + "," + gconv.String(dept0.DeptId)
	return dept0.DeptId, err
}

// 修改保存信息
func (svc *DeptService) EditSave(req *dept2.EditReq, c *gin.Context) (int64, error) {
	dept0 := &dept2.SysDept{DeptId: req.DeptId}
	ok, err := dept0.FindOne()
	if err != nil || !ok {
		return 0, errors.New("数据不存在")
	}
	pdept := &dept2.SysDept{DeptId: req.ParentId}

	ok, err = pdept.FindOne()
	if pdept != nil {
		if pdept.Status != "0" {
			return 0, errors.New("部门停用，不允许新增")
		} else {
			if !strings.HasSuffix(pdept.Ancestors, gconv.String(pdept.DeptId)) { //修复数据用，
				pdept.Ancestors = pdept.Ancestors + "," + gconv.String(pdept.DeptId)
			}
			newAncestors := pdept.Ancestors + "," + gconv.String(dept0.DeptId) //上级的+自己的
			dept2.UpdateDeptChildren(dept0.DeptId, newAncestors, dept0.Ancestors)

			dept0.DeptName = req.DeptName
			dept0.Status = req.Status
			dept0.ParentId = req.ParentId
			dept0.DelFlag = "0"
			dept0.Email = req.Email
			dept0.Leader = req.Leader
			dept0.Phone = req.Phone
			dept0.OrderNum = req.OrderNum

			var userService UserService
			user := userService.GetProfile(c)

			if user != nil && user.UserId > 0 {
				dept0.UpdateBy = user.LoginName
			}

			dept0.UpdateTime = time.Now()

			dept0.Update()
			return 1, nil
		}

	} else {
		return 0, errors.New("父部门不能为空")
	}
}

// 根据分页查询部门管理数据
func (svc *DeptService) SelectListAll(param *dept2.SelectPageReq) ([]dept2.SysDept, error) {
	if param == nil {
		return svc.SelectDeptList(0, "", "", param.TenantId)
	} else {
		return svc.SelectDeptList(param.ParentId, param.DeptName, param.Status, param.TenantId)
	}
}

// 删除部门管理信息
func (svc *DeptService) DeleteDeptById(deptId int64) int64 {
	return dept2.DeleteDeptById(deptId)
}

// 根据部门ID查询信息
func (svc *DeptService) SelectDeptById(deptId int64) *dept2.SysDeptExtend {
	deptEntity, err := dept2.SelectDeptById(deptId)
	if err != nil {
		return nil
	}

	return deptEntity
}

// 根据ID查询所有子部门
func (svc *DeptService) SelectChildrenDeptById(deptId int64) []*dept2.SysDept {
	return dept2.SelectChildrenDeptById(deptId)
}

// 加载部门列表树
func (svc *DeptService) SelectDeptTree(parentId int64, deptName, status string, tenantId int64) (*[]model.Ztree, error) {
	list, err := dept2.SelectDeptList(parentId, deptName, status, tenantId)
	if err != nil {
		return nil, err
	}

	return svc.InitZtree(&list, nil), nil

}

// 查询部门管理数据
func (svc *DeptService) SelectDeptList(parentId int64, deptName, status string, tenantId int64) ([]dept2.SysDept, error) {
	return dept2.SelectDeptList(parentId, deptName, status, tenantId)
}

// 根据角色ID查询部门（数据权限）
func (svc *DeptService) RoleDeptTreeData(roleId int64, tenantId int64) (*[]model.Ztree, error) {
	var result *[]model.Ztree
	deptList, err := dept2.SelectDeptList(0, "", "", tenantId)
	if err != nil {
		return nil, err
	}

	if roleId > 0 {
		roleDeptList, err := dept2.SelectRoleDeptTree(roleId)
		if err != nil || roleDeptList == nil {
			result = svc.InitZtree(&deptList, nil)
		} else {
			result = svc.InitZtree(&deptList, &roleDeptList)
		}
	} else {
		result = svc.InitZtree(&deptList, nil)
	}
	return result, nil
}

// 对象转部门树
func (svc *DeptService) InitZtree(deptList *[]dept2.SysDept, roleDeptList *[]string) *[]model.Ztree {
	var result []model.Ztree
	isCheck := false
	if roleDeptList != nil && len(*roleDeptList) > 0 {
		isCheck = true
	}

	for i := range *deptList {
		if (*deptList)[i].Status == "0" {
			var ztree model.Ztree
			ztree.Id = (*deptList)[i].DeptId
			ztree.Pid = (*deptList)[i].ParentId
			ztree.Name = (*deptList)[i].DeptName
			ztree.Title = (*deptList)[i].DeptName
			if isCheck {
				tmp := gconv.String((*deptList)[i].DeptId) + (*deptList)[i].DeptName
				tmpcheck := false
				for j := range *roleDeptList {
					if strings.EqualFold((*roleDeptList)[j], tmp) {
						tmpcheck = true
						break
					}
				}
				ztree.Checked = tmpcheck
			}
			result = append(result, ztree)
		}
	}
	return &result
}

// 查询部门是否存在用户
func (svc *DeptService) CheckDeptExistUser(deptId int64) bool {
	return dept2.CheckDeptExistUser(deptId)
}

// 查询部门人数
func (svc *DeptService) SelectDeptCount(deptId, parentId int64) int64 {
	return dept2.SelectDeptCount(deptId, parentId)
}

// 校验部门名称是否唯一
func (svc *DeptService) CheckDeptNameUniqueAll(deptName string, parentId int64) string {
	dept, err := dept2.CheckDeptNameUniqueAll(deptName, parentId)
	if err != nil {
		return "1"
	}
	if dept != nil && dept.DeptId > 0 {
		return "1"
	} else {
		return "0"
	}
}

// 校验部门名称是否唯一
func (svc *DeptService) CheckDeptNameUnique(deptName string, deptId, parentId int64) string {
	dept, err := dept2.CheckDeptNameUniqueAll(deptName, parentId)

	if err != nil {
		return "1"
	}
	if dept != nil && dept.DeptId > 0 && dept.DeptId != deptId {
		return "1"
	}
	return "0"
}
