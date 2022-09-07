package Commands

import (
	ApplicationInterfaces "OrderSubscriber/Application/Common/Interfaces"
	DomainEntities "OrderSubscriber/Domain/Entities"
	DomainEnums "OrderSubscriber/Domain/Enums"
	DomainInterfaces "OrderSubscriber/Domain/Interfaces"
	"OrderSubscriber/Infrastructure/config"
	"context"
	"encoding/json"
	"fmt"
)

type SSubscribeOrderCommand struct {
	sConfiguration config.SConfiguration
	iUnitOfWork    DomainInterfaces.IUnitOfWork
	iRedis         ApplicationInterfaces.IRedis
}

func NewSubscribeOrderCommand(sConfiguration config.SConfiguration, iUnitOfWork DomainInterfaces.IUnitOfWork, iRedis ApplicationInterfaces.IRedis) ApplicationInterfaces.IOrderHandlerCommands {
	return SSubscribeOrderCommand{
		sConfiguration: sConfiguration,
		iUnitOfWork:    iUnitOfWork,
		iRedis:         iRedis,
	}
}

func (rSOrderCommand SSubscribeOrderCommand) SubscribeOrderCommand(ctx context.Context) error {
	channel := make(chan string)
	go rSOrderCommand.iRedis.Subscribe(ctx, rSOrderCommand.sConfiguration.Redis.Queues[DomainEnums.ORDERS], channel)
	go func() {
		orderEntity := DomainEntities.OrderEntity{}
		for {
			select {
			case channelData := <-channel:
				if err := json.Unmarshal([]byte(channelData), &orderEntity); err != nil {
					panic(err)
				}
				add := rSOrderCommand.iUnitOfWork.OrderRepository().Add(orderEntity)
				fmt.Println(add)
			}
		}
	}()
	return nil
}
