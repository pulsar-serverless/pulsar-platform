package main

import (
	"log"
	"os"
	_ "pulsar/docs"
	"pulsar/internal/adapters/primary"
	"pulsar/internal/adapters/secondary/docker"
	"pulsar/internal/adapters/secondary/fs"
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

	containerManager := docker.NewContainerManager()
	defer containerManager.Close()

	fileRepository := fs.NewProjectFileRepository(
		os.Getenv("PROJECT_STORAGE_PATH"),
		os.Getenv("DOCKERFILE_TEMPLATE_PATH"),
		os.Getenv("STARTER_CODE_PATH"),
	)

	if err != nil {
		panic("Unable to setup authentication")
	}

	primary.StartServer(&projectRepo, containerManager, fileRepository)
}

// swag init -g cmd/main.go -o docs/
