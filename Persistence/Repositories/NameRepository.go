package Repositories

import (
	DomainEntities "golangcodebase/Domain/Entities"
	DomainInterfaces "golangcodebase/Domain/Interfaces"
	"golangcodebase/Persistence"
	"gorm.io/gorm"
)

type NameRepository struct {
	DomainInterfaces.IGenericRepository[DomainEntities.NameEntity]
}

func NewNameRepository(db *gorm.DB) DomainInterfaces.INameRepository {
	return NameRepository{
		IGenericRepository: NewGenericRepository[DomainEntities.NameEntity](db),
	}
}

func (nr NameRepository) GetByName(name string) (DomainEntities.NameEntity, error) {
	var a DomainEntities.NameEntity
	Persistence.Database.Where("name = ?", name).First(&a)
	return a, nil
}
