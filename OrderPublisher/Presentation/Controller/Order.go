package Controller

import (
	ApplicationInterfaces "OrderPublisher/Application/Common/Interfaces"
	DomainEntities "OrderPublisher/Domain/Entities"
	"OrderPublisher/Presentation/Common"
	"OrderPublisher/Presentation/Controller/Dto"
	"github.com/kataras/iris/v12"
	"net/http"
)

type SOrderController struct {
	iOrderCommand ApplicationInterfaces.IOrderHandlerCommands
}

func NewOrderController(iOrderCommand ApplicationInterfaces.IOrderHandlerCommands) *SOrderController {
	return &SOrderController{
		iOrderCommand: iOrderCommand,
	}
}

// Order todo: use Validator for request body
func (sOrderController SOrderController) Order(ctx iris.Context) {
	params := &Dto.CreateOrderRequest{}
	Common.ReadJson(ctx, &params)
	sOrderController.iOrderCommand.PublishOrderCommand(ctx.Request().Context(), DomainEntities.OrderEntity{
		Id:    params.Id,
		Price: params.Price,
		Title: params.Title,
	})
	ctx.StopWithStatus(http.StatusOK)
}
