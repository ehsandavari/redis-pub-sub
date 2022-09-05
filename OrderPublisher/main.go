package main

import (
	"OrderPublisher/Infrastructure/Redis"
	"OrderPublisher/Infrastructure/config"
	"OrderPublisher/Presentation/Api"
	"OrderPublisher/Presentation/Common"
	"github.com/kataras/iris/v12"
)

func main() {
	irisApp := iris.New()
	irisApp.Use(Common.ErrorHandlerMiddleware)
	configuration := config.NewConfiguration()
	redis := Redis.NewRedis(configuration.Redis.URL)
	defer redis.Close()
	Api.NewApplication(
		configuration,
		redis,
	).SetupAPI(irisApp)
	irisApp.Logger().Fatal(irisApp.Listen(configuration.Application.Host + ":" + configuration.Application.Port))
}
