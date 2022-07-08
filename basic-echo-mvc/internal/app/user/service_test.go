package user

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/mock"
)

func TestFind(t *testing.T) {

	// setup database
	db, mock := mock.DBConnection()

	users := []model.User{
		{
			Model: model.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:     "test",
			Email:    "test@test.com",
			Password: "test123",
		},
		{
			Model: model.Model{
				ID:        2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:     "test2",
			Email:    "test2@test.com",
			Password: "test123",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(users[0].ID, users[0].Name, users[0].Email, users[0].Password, users[0].CreatedAt, users[0].UpdatedAt, users[0].DeletedAt).
		AddRow(users[1].ID, users[1].Name, users[0].Email, users[0].Password, users[1].CreatedAt, users[1].UpdatedAt, users[1].DeletedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL`)).WillReturnRows(rows)

	f := factory.NewFactory(db)
	service := NewService(f)

	dto, err := service.Find(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(dto.Datas) != 2 {
		t.Fatalf("got %d data, want 2 data", len(dto.Datas))
	}

}
