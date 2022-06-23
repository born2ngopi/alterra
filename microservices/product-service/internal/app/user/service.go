package user

import (
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	"github.com/born2ngopi/alterra/basic-echo-mvc/pkg/constant"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[dto.UserResponse], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*dto.UserResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[dto.UserResponse], error) {

	users, info, err := s.UserRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var datas []dto.UserResponse

	for _, user := range users {

		datas = append(datas, dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})

	}

	result := new(dto.SearchGetResponse[dto.UserResponse])
	result.Datas = datas
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*dto.UserResponse, error) {
	var result *dto.UserResponse

	data, err := s.UserRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.UserResponse{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
	}

	return result, nil
}
