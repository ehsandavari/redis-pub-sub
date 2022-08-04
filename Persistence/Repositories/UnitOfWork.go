package Repositories

import (
	DomainInterfaces "golangcodebase/Domain/Interfaces"
	"gorm.io/gorm"
)

type sUnitOfWork struct {
	conn            *gorm.DB
	nameRepository  DomainInterfaces.INameRepository
	name1Repository DomainInterfaces.IName1Repository
}

func (SUnitOfWork sUnitOfWork) NameRepositories() DomainInterfaces.INameRepository {
	return SUnitOfWork.nameRepository
}

func (SUnitOfWork sUnitOfWork) Name1Repositories() DomainInterfaces.IName1Repository {
	return SUnitOfWork.name1Repository
}

func NewUnitOfWork(db *gorm.DB) DomainInterfaces.IUnitOfWork {
	return &sUnitOfWork{
		conn:            db,
		nameRepository:  NewNameRepository(db),
		name1Repository: NewName1Repository(db),
	}
}

func (SUnitOfWork sUnitOfWork) Do(fn DomainInterfaces.UnitOfWorkBlock) error {
	return SUnitOfWork.conn.Transaction(func(tx *gorm.DB) error {
		SUnitOfWork.conn = tx
		SUnitOfWork.nameRepository = NewNameRepository(tx)
		SUnitOfWork.name1Repository = NewName1Repository(tx)
		return fn(SUnitOfWork)
	})
}
