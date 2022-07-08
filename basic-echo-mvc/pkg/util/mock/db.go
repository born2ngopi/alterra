package mock

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, sqlmock.Sqlmock) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	psql := postgres.New(postgres.Config{
		Conn:       dbMock,
		DriverName: "postgres",
	})

	db, err := gorm.Open(psql, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, mock

}
