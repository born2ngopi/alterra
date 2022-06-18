package user

import "github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}
