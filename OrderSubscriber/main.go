package main

import (
	"OrderSubscriber/Application/Handlers/Order/Commands"
	"OrderSubscriber/Infrastructure/Redis"
	"OrderSubscriber/Infrastructure/config"
	"OrderSubscriber/Persistence"
	PersistenceRepositories "OrderSubscriber/Persistence/Repositories"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	stop()
}

func stop() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	defer signal.Stop(signals)

	<-signals
	go func() {
		<-signals
		os.Exit(1)
	}()
}

func run() error {
	configuration, err := config.NewConfiguration()
	if err != nil {
		return err
	}
	database, err := Persistence.NewDatabase(configuration.MySql.URL)
	if err != nil {
		return err
	}
	redis := Redis.NewRedis(configuration.Redis.URL)

	err = Commands.NewSubscribeOrderCommand(
		configuration,
		PersistenceRepositories.NewUnitOfWork(database),
		redis,
	).SubscribeOrderCommand(context.Background())
	if err != nil {
		return err
	}
	return nil
}
