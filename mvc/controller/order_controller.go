package controller

import (
	"net/http"

	"github.com/born2ngopi/alterra/mvc/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController interface{}

type userController struct {
	db *gorm.DB
}

func NewOrderController(db *gorm.DB) *userController {
	return &userController{
		db,
	}
}

type CreateOrderRequest struct {
	ProductID uint `json:"product_id"`
	Qty       uint `json:"qty"`
}

func (u *userController) Create(c echo.Context) error {

	var payloads []CreateOrderRequest

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	var (
		orderData  model.Order
		orderItems []model.OrderItem
	)

	for _, payload := range payloads {

		var product model.Product

		if err := u.db.Model(&product).Where("id = ?", payload.ProductID).First(&product).Error; err != nil {
			return err
		}

		price := product.Price * float64(payload.Qty)

		// add discount 10% if qty > 5
		if payload.Qty > 5 {
			price = price - (float64(10) / float64(100) * price)
		}

		orderData.TotalPrice += price

		orderItems = append(orderItems, model.OrderItem{
			ProductID: payload.ProductID,
			Qty:       payload.Qty,
			Price:     price,
		})
	}

	orderData.OrderItems = orderItems

	if err := u.db.Create(&orderData).Error; err != nil {
		return err
	}

	return c.String(http.StatusOK, "Success")
}
