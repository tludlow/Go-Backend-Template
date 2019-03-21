package main

import (
	"billy-maths/database"
	"billy-maths/router"

	"github.com/subosito/gotenv"
)

//Init functions, across all packages run before the main func in this package (main package) keep that in mind.
func init() {
	//Load our .env file from the root of the backend.
	gotenv.Load()
}

func main() {
	//Establish a database connection.
	database.SetupDatabase()
	defer database.Database.Close()

	//Start our HTTP router.
	router.CreateRouter()

}
