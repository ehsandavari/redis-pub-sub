package Interfaces

type UnitOfWorkBlock func(IUnitOfWork) error
type IUnitOfWork interface {
	NameRepositories() INameRepository
	Name1Repositories() IName1Repository
	Do(UnitOfWorkBlock) error
}
