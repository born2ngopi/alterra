package auth

import (
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(ctx context.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	Register(ctx context.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error)
}

type service struct {
	Repository repository.User
}

func NewService(f *factory.Factory) *service {
	return &service{f.UserRepository}
}

func (s *service) Login(ctx context.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	var result *dto.AuthLoginResponse

	data, err := s.Repository.FindByEmail(ctx, &payload.Email)
	if data == nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.EmailOrPasswordIncorrect, err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(payload.Password)); err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.EmailOrPasswordIncorrect, err)
	}

	token, err := data.GenerateToken()
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AuthLoginResponse{
		Token: token,
		User:  *data,
	}

	return result, nil
}

func (s *service) Register(ctx context.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error) {
	var result *dto.AuthRegisterResponse
	var data model.User

	// data.User = payload.UserEntity

	// TODO create user

	result = &dto.AuthRegisterResponse{
		User: data,
	}

	return result, nil
}
