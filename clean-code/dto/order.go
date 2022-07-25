package dto

type CreateOrderRequest struct {
	ProductID uint `json:"product_id"`
	Qty       uint `json:"qty"`
}
