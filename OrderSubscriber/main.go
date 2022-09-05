package main

import (
	"OrderSubscriber/Application/Handlers/Order/Commands"
	"OrderSubscriber/Infrastructure/Redis"
	"OrderSubscriber/Infrastructure/config"
	"OrderSubscriber/Persistence"
	PersistenceRepositories "OrderSubscriber/Persistence/Repositories"
	"context"
)

func main() {
	configuration := config.NewConfiguration()
	database := Persistence.NewDatabase(configuration.MySql.URL)
	defer database.Close()
	redis := Redis.NewRedis(configuration.Redis.URL)
	defer redis.Close()

	Commands.NewSubscribeOrderCommand(
		configuration,
		PersistenceRepositories.NewUnitOfWork(database),
		redis,
	).SubscribeOrderCommand(context.Background())
}
