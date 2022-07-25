package repository

import (
	"github.com/born2ngopi/alterra/mvc/model"
	mock "github.com/stretchr/testify/mock"
)

type mockOrderReposory struct {
	mock.Mock
}

func NewMockOrderRepo() *mockOrderReposory {
	return &mockOrderReposory{}
}

func (m *mockOrderReposory) Create(data model.Order) error {
	ret := m.Called(data)
	return ret.Error(0)
}
