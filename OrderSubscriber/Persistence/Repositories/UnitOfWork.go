package Repositories

import (
	DomainInterfaces "OrderSubscriber/Domain/Interfaces"
	"OrderSubscriber/Persistence"
	"gorm.io/gorm"
)

type sUnitOfWork struct {
	conn            Persistence.Database
	orderRepository DomainInterfaces.IOrderRepository
}

func (SUnitOfWork sUnitOfWork) OrderRepository() DomainInterfaces.IOrderRepository {
	return SUnitOfWork.orderRepository
}

func NewUnitOfWork(db Persistence.Database) DomainInterfaces.IUnitOfWork {
	return &sUnitOfWork{
		conn:            db,
		orderRepository: NewOrderRepository(db),
	}
}

func (SUnitOfWork sUnitOfWork) Do(fn DomainInterfaces.UnitOfWorkBlock) error {
	return SUnitOfWork.conn.Gorm.Transaction(func(tx *gorm.DB) error {
		SUnitOfWork.conn.Gorm = tx
		SUnitOfWork.orderRepository = NewOrderRepository(SUnitOfWork.conn)
		return fn(SUnitOfWork)
	})
}
