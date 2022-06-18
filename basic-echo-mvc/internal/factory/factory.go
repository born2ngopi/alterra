package factory

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/database"
	"gorm.io/gorm"
)

type Factory struct {
	DB *gorm.DB
}

func NewFactory() *Factory {
	return &Factory{
		DB: database.GetConnection(),
	}
}
