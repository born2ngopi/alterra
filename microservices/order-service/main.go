package main

import (
	"flag"
	"os"

	"github.com/born2ngopi/alterra/basic-echo-mvc/database"
	"github.com/born2ngopi/alterra/basic-echo-mvc/database/migration"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/http"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// load env configuration
func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {

	database.CreateConnection()

	var m string // for check migration

	flag.StringVar(
		&m,
		"migrate",
		"run",
		`this argument for check if user want to migrate table, rollback table, or status migration

to use this flag:
	use -migrate=migrate for migrate table
	use -migrate=rollback for rollback table
	use -migrate=status for get status migration`,
	)
	flag.Parse()

	if m == "migrate" {
		migration.Migrate()
		return
	} else if m == "rollback" {
		migration.Rollback()
		return
	} else if m == "status" {
		migration.Status()
		return
	}

	f := factory.NewFactory()
	e := echo.New()
	middleware.Init(e)
	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))

}
