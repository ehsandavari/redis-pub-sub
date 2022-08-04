package Repositories

import (
	DomainInterfaces "golangcodebase/Domain/Interfaces"
	"gorm.io/gorm"
)

type GenericRepository[T any] struct {
	DataBase *gorm.DB
}

func NewGenericRepository[T any](db *gorm.DB) DomainInterfaces.IGenericRepository[T] {
	return GenericRepository[T]{
		DataBase: db,
	}
}

func (GR GenericRepository[T]) Find() (T, error) {
	var a T
	GR.DataBase.First(&a)
	return a, nil
}

func (GR GenericRepository[T]) Add(model *T) (T, error) {
	GR.DataBase.Create(&model)
	return *model, nil
}
