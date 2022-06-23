package factory

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/database"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
)

type Factory struct {
	ProductRepository repository.Product
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		ProductRepository: repository.NewProduct(db),
	}
}
