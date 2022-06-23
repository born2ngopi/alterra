package dto

import "github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"

type OrderResponse struct {
	model.Order
	Product model.Product
}

type CreateOrderRequest struct {
	ProductID uint   `json:"product_id" validate:"required"`
	Total     uint   `json:"total" validate:"required,max=10"`
	Customer  string `json:"customer" validate:"required"`
}
