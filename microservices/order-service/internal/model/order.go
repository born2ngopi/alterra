package model

type Order struct {
	Model

	Customer  string `json:"customer"   gorm:"size:200;not null"`
	ProductID uint   `json:"product_id" gorm:"not null"`
	TotalItem uint   `json:"total_item" gorm:"not null"`
}
