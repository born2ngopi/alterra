package http

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/app/order"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	order.NewHandler(f).Route(e.Group("/orders"))
}
