package repository

import (
	"github.com/born2ngopi/alterra/mvc/model"
	mock "github.com/stretchr/testify/mock"
)

type mockProductReposory struct {
	mock.Mock
}

func NewMockProductRepo() *mockProductReposory {
	return &mockProductReposory{}
}

func (m *mockProductReposory) Find(productID uint) (model.Product, error) {
	ret := m.Called(productID)
	return ret.Get(0).(model.Product), ret.Error(1)
}
