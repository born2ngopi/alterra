package http

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/app/product"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	product.NewHandler(f).Route(e.Group("/products"))
}
