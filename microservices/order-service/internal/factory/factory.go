package factory

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/database"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
)

type Factory struct {
	OrderRepository repository.Order
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		OrderRepository: repository.NewOrder(db),
	}
}
