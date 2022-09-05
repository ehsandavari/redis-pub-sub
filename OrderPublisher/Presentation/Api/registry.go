package Api

import (
	"github.com/kataras/iris/v12/core/router"
)

func (rAppController SApplication) registerOrderAPI(router router.Party) {
	router.Post("/order", rAppController.OrderController.Order)
}
