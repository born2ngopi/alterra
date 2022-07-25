package usecase

import (
	"github.com/born2ngopi/alterra/mvc/dto"
	"github.com/born2ngopi/alterra/mvc/model"
	"github.com/born2ngopi/alterra/mvc/repository"
)

type OrderUsecase interface {
	Create(payloads []dto.CreateOrderRequest) error
}

type orderUsecase struct {
	orderRepository   repository.OrderRepository
	productRepository repository.ProductRepository
}

func NewOrderUsecase(orderRepo repository.OrderRepository, productRepo repository.ProductRepository) *orderUsecase {
	return &orderUsecase{orderRepository: orderRepo, productRepository: productRepo}
}

func (s *orderUsecase) Create(payloads []dto.CreateOrderRequest) error {
	var (
		orderData  model.Order
		orderItems []model.OrderItem
	)

	for _, payload := range payloads {

		product, err := s.productRepository.Find(payload.ProductID)
		if err != nil {
			return err
		}

		price := product.Price * float64(payload.Qty)

		// add discount 10% if qty > 5
		if payload.Qty > 5 {
			price = price - (float64(10) / float64(100) * price)
		}

		orderData.TotalPrice += price

		orderItems = append(orderItems, model.OrderItem{
			ProductID: payload.ProductID,
			Qty:       payload.Qty,
			Price:     price,
		})
	}

	orderData.OrderItems = orderItems

	if err := s.orderRepository.Create(orderData); err != nil {
		return err
	}

	return nil

}
