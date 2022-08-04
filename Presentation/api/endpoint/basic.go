package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @Summary Get Block
// @ID block_v2
// @Description Get Block information
// @Accept json
// @Produce json
// @Tags Transactions
// @Param coin path string true "the coin name" default(zilliqa)
// @Param address path string true "the query address" default(850321)
// @Failure 500 {object} ErrorResponse
// @Router /v2/{coin}/blocks/{block} [get]
func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"build":  "dev",
		"date":   time.NewTicker,
	})
}
