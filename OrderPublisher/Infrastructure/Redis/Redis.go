package Redis

import (
	ApplicationInterfaces "OrderPublisher/Application/Common/Interfaces"
	DomainEnums "OrderPublisher/Domain/Enums"
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type SRedis struct {
	Client *redis.Client
}

func NewRedis(host string) ApplicationInterfaces.IRedis {
	return &SRedis{
		Client: redis.NewClient(&redis.Options{
			Addr: host,
		}),
	}
}

func (rRedis SRedis) Publish(ctx context.Context, channelName DomainEnums.RedisQueues, message interface{}) error {
	return rRedis.Client.Publish(ctx, string(channelName), message).Err()
}

func (rRedis SRedis) Close() {
	err := rRedis.Client.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
