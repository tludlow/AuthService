package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tludlow/authservice/communications"
	"github.com/tludlow/authservice/database"

	_ "github.com/go-sql-driver/mysql"
)

//Main - Struct holding the references to other parts of the app.
type Main struct {
	router *gin.Engine
	db     *sql.DB
}

func (m *Main) startup() error {
	//Load the .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	//Connect to the database.
	databaseConnection, err := database.ConnectDSN(os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"))
	if err != nil {
		log.Fatal("Error connecting to the database.")
		return err
	}
	m.db = databaseConnection

	//Create a generic base gin router.
	m.router = gin.Default()
	communications.SetupRoutes(m.router)

	return nil
}

func main() {
	main := Main{}

	//Start the server and checks for errors.
	err := main.startup()
	if err != nil {
		log.Fatal("Error starting the service - " + err.Error())
		return
	}

	//Close database connection at the end of the app.
	defer main.db.Close()

	//Start the router
	port := os.Getenv("ROUTER_PORT")
	if port == "" {
		port = ":8080"
		fmt.Println("[DEBUG] No port provided for the router, starting on default port of 8080")
	}
	main.router.Run(port)
}
