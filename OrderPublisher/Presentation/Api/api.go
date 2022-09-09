package Api

import (
	ApplicationInterfaces "OrderPublisher/Application/Common/Interfaces"
	"OrderPublisher/Application/Handlers/Order/Commands"
	"OrderPublisher/Infrastructure/config"
	"OrderPublisher/Presentation/Controller"
	"github.com/kataras/iris/v12"
)

type SApplication struct {
	OrderController *Controller.SOrderController
}

func NewApplication(configuration config.SConfiguration, iRedis ApplicationInterfaces.IRedis) *SApplication {
	return &SApplication{
		OrderController: Controller.NewOrderController(Commands.NewPublishOrderCommand(configuration, iRedis)),
	}
}

func (rAppController *SApplication) SetupAPI(app *iris.Application) {
	api := app.Party("/api")
	rAppController.registerOrderAPI(api)
}
