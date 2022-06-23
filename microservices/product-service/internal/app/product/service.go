package product

import (
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	"github.com/born2ngopi/alterra/basic-echo-mvc/pkg/constant"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
)

type service struct {
	ProductRepository repository.Product
}

type Service interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Product], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Product, error)
	Create(ctx context.Context, payload *dto.CreateProductRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateProductRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.Product, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		ProductRepository: f.ProductRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Product], error) {

	Products, info, err := s.ProductRepository.Find(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// var datas []dto.UserResponse

	// for _, user := range users {

	// 	datas = append(datas, dto.UserResponse{
	// 		ID:    user.ID,
	// 		Name:  user.Name,
	// 		Email: user.Email,
	// 	})

	// }

	result := new(dto.SearchGetResponse[model.Product])
	result.Datas = Products
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Product, error) {

	data, err := s.ProductRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateProductRequest) (string, error) {

	var product = model.Product{
		Name:        payload.Name,
		Stock:       payload.Stock,
		Description: payload.Description,
	}

	err := s.ProductRepository.Create(ctx, product)
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

	err := s.ProductRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.Product, error) {

	data, err := s.ProductRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.ProductRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil

}
