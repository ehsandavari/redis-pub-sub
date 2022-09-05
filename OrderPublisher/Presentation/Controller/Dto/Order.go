package Dto

import DomainEntities "OrderPublisher/Domain/Entities"

type CreateOrderRequest struct {
	Id    uint64 `json:"order_id"`
	Price uint   `json:"price"`
	Title string `json:"title"`
}

type CreateOrderResponse struct {
	User DomainEntities.OrderEntity `json:"user"`
}
