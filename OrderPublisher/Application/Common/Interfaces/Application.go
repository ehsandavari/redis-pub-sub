package Interfaces

import (
	"OrderPublisher/Domain/Entities"
	"context"
)

type (
	IOrderHandlerCommands interface {
		PublishOrderCommand(ctx context.Context, orderEntity Entities.OrderEntity)
	}
)
