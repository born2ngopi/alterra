package user

import (
	"encoding/json"
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/mock"
	"github.com/labstack/echo/v4"
)

func TestGetByID(t *testing.T) {
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

	// setup context
	e := echo.New()
	f := factory.NewFactory(db)
	h := NewHandler(f)
	h.Route(e.Group("/users"))

	echoMock := mock.HttpMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/users")

	h.Get(c)

	var data map[string]interface{}

	err := json.Unmarshal(rec.Body, &data)
	if err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("got status code %d want 200", rec.Code)
	}

	if len(data) != 2 {
		t.Fatalf("got %d data, want 2 data", len(data))
	}

}
