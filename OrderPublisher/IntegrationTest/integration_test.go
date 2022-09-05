package integration_test

import (
	. "github.com/Eun/go-hit"
	"net/http"
	"testing"
)

const (
	host     = "localhost:8000"
	basePath = host + "/api"
)

// todo: not complate
func TestOrderApi(t *testing.T) {
	body := `{
				"order_id": 100,
				"price": 12000,
				"title": "test"
			 }`
	Test(t,
		Description("Send Order Success"),
		Post(basePath+"/order"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
	)
}
