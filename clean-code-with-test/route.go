package main

import (
	"github.com/born2ngopi/alterra/mvc/controller"
	"github.com/born2ngopi/alterra/mvc/repository"
	"github.com/born2ngopi/alterra/mvc/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {

	orderRepository := repository.NewOrderRepository(db)
	productRepository := repository.NewProductRepository(db)

	orderService := usecase.NewOrderUsecase(orderRepository, productRepository)

	orderController := controller.NewOrderController(orderService)

	e.POST("/order", orderController.Create)
}
