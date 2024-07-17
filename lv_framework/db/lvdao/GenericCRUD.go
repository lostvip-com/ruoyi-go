package lvdao

import (
	"gorm.io/gorm"
)

// CRUD 接口定义了通用的CRUD操作
type CRUD[T any] interface {
	Create(db *gorm.DB, model *T) error
	Find(db *gorm.DB, out **T, id uint) error
	Update(db *gorm.DB, model *T) error
	Delete(db *gorm.DB, model *T) error
}

// GenericCRUD 是CRUD接口的一个泛型实现
type GenericCRUD[T any] struct {
	db *gorm.DB
}

// NewGenericCRUD 创建一个新的GenericCRUD实例
func NewGenericCRUD[T any](db *gorm.DB) *GenericCRUD[T] {
	return &GenericCRUD[T]{db: db}
}

// Create 创建一条记录
func (g *GenericCRUD[T]) Create(model *T) error {
	return g.db.Create(model).Error
}

// Create 创建一条记录
func (g *GenericCRUD[T]) Save(model *T) error {
	return g.db.Save(model).Error
}

// Find 根据ID查找记录
func (g *GenericCRUD[T]) FindById(out *T, id uint) error {
	return g.db.First(out, id).Error
}

// Find 根据ID查找记录
func (g *GenericCRUD[T]) FindList(list *[]T, start int, pageSize int, condition string, args ...any) error {
	result := g.db.Where(condition, args...).Offset(start).Limit(pageSize).Find(list)
	return result.Error
}

// Find 根据ID查找记录
func (g *GenericCRUD[T]) FindFirst(out *T, condition string, args ...any) error {
	result := g.db.Where(condition, args...).First(out)
	return result.Error
}

// Update 更新记录
func (g *GenericCRUD[T]) Update(model *T) error {
	return g.db.Save(model).Error
}

// Delete 删除记录
func (g *GenericCRUD[T]) Delete(model *T) error {
	return g.db.Delete(model).Error
}
