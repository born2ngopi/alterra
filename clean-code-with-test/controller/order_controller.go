package controller

import (
	"net/http"

	"github.com/born2ngopi/alterra/mvc/dto"
	"github.com/born2ngopi/alterra/mvc/usecase"
	"github.com/labstack/echo/v4"
)

type UserController interface{}

type userController struct {
	useCase usecase.OrderUsecase
}

func NewOrderController(orderUsecase usecase.OrderUsecase) *userController {
	return &userController{
		orderUsecase,
	}
}

func (u *userController) Create(c echo.Context) error {

	var payloads []dto.CreateOrderRequest

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	err := u.useCase.Create(payloads)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Success")
}
