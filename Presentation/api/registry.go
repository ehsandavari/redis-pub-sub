package api

import (
	"github.com/gin-gonic/gin"
	"golangcodebase/Presentation/api/endpoint"
)

func RegisterBasicAPI(router gin.IRouter) {
	router.GET("/", endpoint.GetStatus)
}
