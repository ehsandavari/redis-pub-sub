package Interfaces

import (
	DomainEnums "OrderPublisher/Domain/Enums"
	"context"
)

type (
	IRedis interface {
		Publish(ctx context.Context, channelName DomainEnums.RedisQueues, message interface{}) error
		Close()
	}
)
