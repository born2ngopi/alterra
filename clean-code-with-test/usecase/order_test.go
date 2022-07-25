package usecase

import (
	"testing"

	"github.com/born2ngopi/alterra/mvc/dto"
	"github.com/born2ngopi/alterra/mvc/model"
	"github.com/born2ngopi/alterra/mvc/repository"
)

func TestCreateOrder(t *testing.T) {

	data := []dto.CreateOrderRequest{
		{ProductID: 1, Qty: 6},
	}

	mockProduct := model.Product{
		ID:    1,
		Name:  "product-mock",
		Price: 4500,
		Stock: 10,
	}

	mockProductRepository := repository.NewMockProductRepo()
	mockProductRepository.On("Find", uint(1)).Return(mockProduct, nil)

	mockOrder := model.Order{
		TotalPrice: 24300,
		OrderItems: []model.OrderItem{
			{
				ProductID: 1,
				Qty:       6,
				Price:     24300,
			},
		},
	}

	mockOrderRepository := repository.NewMockOrderRepo()
	mockOrderRepository.On("Create", mockOrder).Return(nil)

	service := NewOrderUsecase(mockOrderRepository, mockProductRepository)

	if err := service.Create(data); err != nil {
		t.Errorf("Got Error %v", err)
	}

}
