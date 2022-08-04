package Interfaces

import (
	"golangcodebase/Domain/Entities"
)

type INameRepository interface {
	IGenericRepository[Entities.NameEntity]
	GetByName(name string) (Entities.NameEntity, error)
}
