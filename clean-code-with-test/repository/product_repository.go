package repository

import (
	"github.com/born2ngopi/alterra/mvc/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Find(productID uint) (model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) Find(productID uint) (model.Product, error) {
	var product model.Product

	if err := r.db.Model(&product).Where("id = ?", productID).First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}
