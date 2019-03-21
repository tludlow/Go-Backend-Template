package database

import (
	"database/sql"
	"log"
	"os"

	//Enahncement package ontop of the database/sql one. A driver perhaps
	_ "github.com/go-sql-driver/mysql"
)

//Database - Our connection to the MySQL database.
var Database *sql.DB

//SetupDatabase - Create the connection to the database with DSN paramaters loaded from the .env file.
func SetupDatabase() {
	//We will load the database connection parameters from the .env file and then construct the MySQL DSN string from them.
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")

	//The DSN must follow the pattern "john:pass@/my_db"
	//You can add a domain between @ and / to connect to a remote database
	databaseDSN := username + ":" + password + "@/" + databaseName + "?parseTime=true&charset=utf8mb4"

	db, err := sql.Open("mysql", databaseDSN)

	if err != nil {
		log.Panicln("Database has failed to connect.")
		log.Println(err)
	}
	//Make our database connection public by setting it to the global variable.
	Database = db
	log.Println("[Database] Connection Established")

	//Check if the connection is still alive, will error if not so the connection hasnt worked.
	Database.Ping()
}
