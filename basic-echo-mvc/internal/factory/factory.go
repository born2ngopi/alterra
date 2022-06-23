package factory

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/database"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
)

type Factory struct {
	// DB             *gorm.DB
	UserRepository repository.User
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		UserRepository: repository.NewUser(db),
	}
}
