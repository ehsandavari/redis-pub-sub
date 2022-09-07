package Repositories

import (
	DomainInterfaces "OrderSubscriber/Domain/Interfaces"
	"OrderSubscriber/Persistence"
	"gorm.io/gorm"
)

type sUnitOfWork struct {
	conn            *Persistence.Database
	orderRepository DomainInterfaces.IOrderRepository
}

func NewUnitOfWork(db *Persistence.Database) DomainInterfaces.IUnitOfWork {
	return &sUnitOfWork{
		conn:            db,
		orderRepository: NewOrderRepository(db),
	}
}

func (SUnitOfWork sUnitOfWork) OrderRepository() DomainInterfaces.IOrderRepository {
	return SUnitOfWork.orderRepository
}

func (SUnitOfWork sUnitOfWork) Do(unitOfWorkBlock DomainInterfaces.UnitOfWorkBlock) error {
	return SUnitOfWork.conn.Gorm.Transaction(func(transaction *gorm.DB) error {
		SUnitOfWork.conn.Gorm = transaction
		SUnitOfWork.orderRepository = NewOrderRepository(SUnitOfWork.conn)
		return unitOfWorkBlock(SUnitOfWork)
	})
}
