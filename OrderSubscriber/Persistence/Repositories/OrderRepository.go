package Repositories

import (
	DomainEntities "OrderSubscriber/Domain/Entities"
	DomainInterfaces "OrderSubscriber/Domain/Interfaces"
	"OrderSubscriber/Persistence"
)

type OrderRepository struct {
	GenericRepository[DomainEntities.OrderEntity]
}

func NewOrderRepository(db *Persistence.Database) DomainInterfaces.IOrderRepository {
	return OrderRepository{
		GenericRepository: NewGenericRepository[DomainEntities.OrderEntity](db),
	}
}
