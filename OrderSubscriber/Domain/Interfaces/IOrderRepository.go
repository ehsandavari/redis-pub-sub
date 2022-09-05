package Interfaces

import (
	"OrderSubscriber/Domain/Entities"
)

//go:generate mockgen -destination=../Mocks/MockIOrderRepository.go -package=Mock  OrderSubscriber/Domain/Interfaces IOrderRepository

type IOrderRepository interface {
	IGenericRepository[Entities.OrderEntity]
}
