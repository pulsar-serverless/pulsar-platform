package web

import (
	"os"
	"pulsar/internal/adapters/secondary/postgres"
	"pulsar/internal/core/services/analytics"
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
	analyticsService analytics.IAnalyticsService
	resourceService  analytics.IResourceService
}

func StartServer(db *postgres.Database, mq ports.IMessageQueue, containerMan ports.IContainerManager, fileRepo ports.IFileRepository) {
	logService := log.NewLogService(mq, db)
	resourceService := analytics.NewResourceService(db)
	containerService := container.NewContainerService(containerMan, logService, fileRepo, db, resourceService)
	projectService := project.NewProjectService(db, containerService, fileRepo)
	envService := envs.NewEnvService(db, *projectService)
	analyticsService := analytics.NewAnalyticsService(db, mq)

	server := &Server{
		echo:             echo.New(),
		projectService:   projectService, // inject project service inject authentication service
		containerService: containerService,
		envService:       envService,
		logService:       logService,
		analyticsService: analyticsService,
		resourceService:  resourceService,
	}

	server.echo.Use(middleware.CORS())
	server.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	// setup API routes
	server.DefineRoutes()
	// start server
	server.echo.Logger.Fatal(server.echo.Start(":" + os.Getenv("PORT")))
}
