package web

import (
	"os"
	"pulsar/internal/adapters/secondary/postgres"
	"pulsar/internal/core/services/container"
	"pulsar/internal/core/services/envs"
	"pulsar/internal/core/services/log"
	"pulsar/internal/core/services/project"
	"pulsar/internal/ports"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	echo             *echo.Echo
	projectService   project.IProjectService
	containerService container.IContainerService
	envService       envs.IEnvService
	logService       log.ILogService
}

func StartServer(db *postgres.Database, mq ports.IMessageQueue, containerMan ports.IContainerManager, fileRepo ports.IFileRepository) {
	logService := log.NewLogService(mq, db)
	containerService := container.NewContainerService(containerMan, logService, fileRepo, db)
	projectService := project.NewProjectService(db, containerService, fileRepo)
	envService := envs.NewEnvService(db, *projectService)

	server := &Server{
		echo:             echo.New(),
		projectService:   projectService, // inject project service inject authentication service
		containerService: containerService,
		envService:       envService,
		logService:       logService,
	}

	server.echo.Use(middleware.CORS())
	server.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	// setup API routes
	server.DefineRoutes()
	// start server
	server.echo.Logger.Fatal(server.echo.Start(":" + os.Getenv("PORT")))
}
