package factory

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	"gorm.io/gorm"
)

type Factory struct {
	// DB             *gorm.DB
	UserRepository repository.User
}

func NewFactory(db *gorm.DB) *Factory {
	return &Factory{
		UserRepository: repository.NewUser(db),
	}
}
