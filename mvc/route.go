package main

import (
	"github.com/born2ngopi/alterra/mvc/controller"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {

	orderController := controller.NewOrderController(db)

	e.POST("/order", orderController.Create)
}
