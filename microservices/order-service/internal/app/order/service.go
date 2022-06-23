package order

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	"github.com/born2ngopi/alterra/basic-echo-mvc/pkg/constant"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
)

type service struct {
	OrderRepository repository.Order
}

type Service interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Order], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*dto.OrderResponse, error)
	Create(ctx context.Context, payload *dto.CreateOrderRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateProductRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.Order, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		OrderRepository: f.OrderRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Order], error) {

	Products, info, err := s.OrderRepository.Find(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.Order])
	result.Datas = Products
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*dto.OrderResponse, error) {

	data, err := s.OrderRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	resp, err := http.Get(fmt.Sprintf("http://localhost:8000/products/%d", data.ProductID))
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var product model.Product

	if err := json.NewDecoder(resp.Body).Decode(product); err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := dto.OrderResponse{
		Order:   data,
		Product: product,
	}

	return &result, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateOrderRequest) (string, error) {

	resp, err := http.Get(fmt.Sprintf("http://localhost:8000/products/%d", payload.ProductID))
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var product model.Product

	if err := json.NewDecoder(resp.Body).Decode(product); err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	if product.Stock < int(payload.Total) {
		return "", res.ErrorBuilder(&res.ErrorConstant.StockTidakAda, err)
	}

	product.Stock -= int(payload.Total)
	// TODO

	var order = model.Order{
		Customer:  payload.Customer,
		TotalItem: payload.Total,
		ProductID: product.ID,
	}

	err = s.OrderRepository.Create(ctx, order)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdateProductRequest) (string, error) {

	var data = make(map[string]interface{})

	if payload.Name != nil {
		data["name"] = payload.Name
	}
	if payload.Stock != nil {
		data["stock"] = payload.Stock
	}
	if payload.Description != nil {
		data["description"] = payload.Description
	}

	err := s.OrderRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.Order, error) {

	data, err := s.OrderRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.OrderRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil

}
