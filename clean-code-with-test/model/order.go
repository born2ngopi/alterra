package model

import "time"

type Order struct {
	ID         uint      `json:"id"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}
