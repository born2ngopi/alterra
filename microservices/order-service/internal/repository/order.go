package repository

import (
	"context"
	"strings"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"gorm.io/gorm"
)

type Order interface {
	Create(ctx context.Context, data model.Order) error
	Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Order, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.Order, error)
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type order struct {
	Db *gorm.DB
}

func NewOrder(db *gorm.DB) *order {
	return &order{
		db,
	}
}

func (p *order) Create(ctx context.Context, data model.Order) error {
	return p.Db.WithContext(ctx).Model(&model.Order{}).Create(&data).Error
}

func (p *order) Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Order, *dto.PaginationInfo, error) {
	var products []model.Order
	var count int64

	query := p.Db.WithContext(ctx).Model(&model.Order{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(customer) LIKE ?  ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&products).Error

	return products, dto.CheckInfoPagination(paginate, count), err
}

func (p *order) FindByID(ctx context.Context, ID uint) (model.Order, error) {

	var data model.Order
	err := p.Db.WithContext(ctx).Model(&data).Where("id = ?", ID).First(&data).Error

	return data, err
}

func (p *order) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := p.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.Order{}).Updates(data).Error
	return err

	// var product model.OrderProduct

	// err := p.Db.WithContext(ctx).Model(&data).Where("id = ?", ID).First(&data).Error

	// if data.Name != nil {
	// product.Name = data.Name
	// }

	// product.Stock = data.Stock
	// product.Description = data.Description

	// err = p.Db.Save(&product).Error

	// return nil
}

func (p *order) Delete(ctx context.Context, ID uint) error {

	err := p.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.Order{}).Error
	return err
}
