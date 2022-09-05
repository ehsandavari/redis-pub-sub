package Commands

import (
	"OrderPublisher/Application/Common"
	ApplicationInterfaces "OrderPublisher/Application/Common/Interfaces"
	"OrderPublisher/Domain/Entities"
	"OrderPublisher/Domain/Enums"
	"OrderPublisher/Infrastructure/config"
	"context"
)

type SPublishOrderCommand struct {
	Configuration config.SConfiguration
	iRedis        ApplicationInterfaces.IRedis
}

func NewPublishOrderCommand(configuration config.SConfiguration, iRedis ApplicationInterfaces.IRedis) ApplicationInterfaces.IOrderHandlerCommands {
	return &SPublishOrderCommand{
		Configuration: configuration,
		iRedis:        iRedis,
	}
}

func (rSOrderCommand *SPublishOrderCommand) PublishOrderCommand(ctx context.Context, orderEntity Entities.OrderEntity) {
	payload := Common.MarshalJson(orderEntity)
	err := rSOrderCommand.iRedis.Publish(ctx, rSOrderCommand.Configuration.Redis.Queues[Enums.ORDERS], payload)
	if err != nil {
		panic(err)
	}
}
