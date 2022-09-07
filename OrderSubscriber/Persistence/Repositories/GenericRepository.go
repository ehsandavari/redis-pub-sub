package Repositories

import (
	DomainInterfaces "OrderSubscriber/Domain/Interfaces"
	"OrderSubscriber/Persistence"
)

type GenericRepository[T DomainInterfaces.GenericRepositoryConstraint] struct {
	DataBase *Persistence.Database
}

func NewGenericRepository[T DomainInterfaces.GenericRepositoryConstraint](db *Persistence.Database) GenericRepository[T] {
	return GenericRepository[T]{
		DataBase: db,
	}
}

func (GR GenericRepository[T]) Find() T {
	var find T
	GR.DataBase.Gorm.First(&find)
	return find
}

func (GR GenericRepository[T]) Add(model T) T {
	GR.DataBase.Gorm.Create(&model)
	return model
}
