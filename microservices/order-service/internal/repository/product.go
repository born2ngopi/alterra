package repository

import (
	"context"
	"strings"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"gorm.io/gorm"
)

type Product interface {
	Create(ctx context.Context, data model.Product) error
	Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Product, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.Product, error)
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type product struct {
	Db *gorm.DB
}

func NewProduct(db *gorm.DB) *product {
	return &product{
		db,
	}
}

func (p *product) Create(ctx context.Context, data model.Product) error {
	return p.Db.WithContext(ctx).Model(&model.Product{}).Create(&data).Error
}

func (p *product) Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Product, *dto.PaginationInfo, error) {
	var products []model.Product
	var count int64

	query := p.Db.WithContext(ctx).Model(&model.Product{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?  ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&products).Error

	return products, dto.CheckInfoPagination(paginate, count), err
}

func (p *product) FindByID(ctx context.Context, ID uint) (model.Product, error) {

	var data model.Product
	err := p.Db.WithContext(ctx).Model(&data).Where("id = ?", ID).First(&data).Error

	return data, err
}

func (p *product) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := p.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.Product{}).Updates(data).Error
	return err

	// var product model.ProductProduct

	// err := p.Db.WithContext(ctx).Model(&data).Where("id = ?", ID).First(&data).Error

	// if data.Name != nil {
	// product.Name = data.Name
	// }

	// product.Stock = data.Stock
	// product.Description = data.Description

	// err = p.Db.Save(&product).Error

	// return nil
}

func (p *product) Delete(ctx context.Context, ID uint) error {

	err := p.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.Product{}).Error
	return err
}
