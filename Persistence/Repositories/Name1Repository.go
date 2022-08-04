package Repositories

import (
	DomainEntities "golangcodebase/Domain/Entities"
	DomainInterfaces "golangcodebase/Domain/Interfaces"
	"golangcodebase/Persistence"
	"gorm.io/gorm"
)

type Name1Repository struct {
	DomainInterfaces.IGenericRepository[DomainEntities.Name1Entity]
}

func NewName1Repository(db *gorm.DB) DomainInterfaces.IName1Repository {
	return Name1Repository{
		IGenericRepository: NewGenericRepository[DomainEntities.Name1Entity](db),
	}
}

func (nr Name1Repository) GetByName(name string) (DomainEntities.Name1Entity, error) {
	var a DomainEntities.Name1Entity
	Persistence.Database.Where("name = ?", name).First(&a)
	return a, nil
}
