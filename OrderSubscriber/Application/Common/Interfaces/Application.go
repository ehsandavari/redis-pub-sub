package Interfaces

import (
	"context"
)

type (
	IOrderHandlerCommands interface {
		SubscribeOrderCommand(ctx context.Context) error
	}
)
