package repository

import (
	"github.com/born2ngopi/alterra/mvc/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(data model.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Create(data model.Order) error {
	return r.db.Create(&data).Error
}
