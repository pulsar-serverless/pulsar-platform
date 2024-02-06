package main

import (
	"log"
	"os"
	_ "pulsar/docs"
	"pulsar/internal/adapters/primary"
	"pulsar/internal/adapters/secondary/postgres"

	"github.com/joho/godotenv"
)

// @title						Pulsar API
// @version					1.0
// @description				This is a server for  pulsar (serverless web platform) server.
//
// @host						localhost:1323
// @BasePath					/
// @SecurityDefinitions.apiKey	Bearer
// @in							header
// @name						Authorization
func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbConn := postgres.SetupDB(os.Getenv("POSTGRES_CONNECTION"))
	projectRepo := postgres.NewProjectRepo(dbConn)

	if err != nil {
		panic("Unable to setup authentication")
	}

	primary.StartServer(&projectRepo)
}

// swag init -g cmd/main.go -o docs/
