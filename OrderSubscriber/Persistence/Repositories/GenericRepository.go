package Repositories

import (
	DomainInterfaces "OrderSubscriber/Domain/Interfaces"
	"OrderSubscriber/Persistence"
)

type GenericRepository[T DomainInterfaces.GenericRepositoryConstraint] struct {
	DataBase Persistence.Database
}

func NewGenericRepository[T DomainInterfaces.GenericRepositoryConstraint](db Persistence.Database) GenericRepository[T] {
	return GenericRepository[T]{
		DataBase: db,
	}
}

func (GR GenericRepository[T]) Find() (T, error) {
	var a T
	GR.DataBase.Gorm.First(&a)
	return a, nil
}

func (GR GenericRepository[T]) Add(model *T) (T, error) {
	GR.DataBase.Gorm.Create(&model)
	return *model, nil
}
