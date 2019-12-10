package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Database global database variable to access the database...
var Database *sql.DB

//ConnectDSN - Connects to the database.
func ConnectDSN(username, password, databaseName string) (*sql.DB, error) {
	//Connect to the database
	databaseDSN := username + ":" + password + "@/" + databaseName + "?parseTime=true&charset=utf8mb4"
	db, err := sql.Open("mysql", databaseDSN)
	if err != nil {
		log.Fatal("Error connecting to the database")
		return nil, err
	}
	Database = db
	return db, nil
}
