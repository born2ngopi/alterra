package seeder

import "github.com/born2ngopi/alterra/basic-echo-mvc/database"

func Seed() {

	conn := database.GetConnection()

	userTableSeeder(conn)
	// otherTableSeeder(conn)
}
