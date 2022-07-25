package model

type OrderItem struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Qty       uint    `json:"qty"`
	Price     float64 `json:"price"`
	OrderID   uint    `json:"order_id"`
}
