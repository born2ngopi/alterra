package model

import "time"

type Product struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Stock     uint      `json:"stock"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
