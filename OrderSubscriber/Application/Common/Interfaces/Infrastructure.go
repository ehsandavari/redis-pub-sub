package Interfaces

import (
	DomainEnums "OrderSubscriber/Domain/Enums"
	"context"
)

//go:generate mockgen -destination=../../Mocks/MockIRedis.go -package=MockIRedis OrderSubscriber/Application/Common/Interfaces IRedis

type (
	IRedis interface {
		Subscribe(ctx context.Context, channelName DomainEnums.RedisQueues, channel chan<- string)
		Close()
	}
)
