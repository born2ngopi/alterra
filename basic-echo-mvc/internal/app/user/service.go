package user

import "github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"

type service struct {
}

type Service interface {
}

func NewService(f *factory.Factory) Service {
	return &service{}
}
