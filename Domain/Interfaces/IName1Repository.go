package Interfaces

import (
	"golangcodebase/Domain/Entities"
)

type IName1Repository interface {
	IGenericRepository[Entities.Name1Entity]
	GetByName(name string) (Entities.Name1Entity, error)
}
