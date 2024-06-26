package web

import (
	"os"
	"pulsar/internal/adapters/secondary/postgres"
	"pulsar/internal/core/services/analytics"
	"pulsar/internal/core/services/billing"
	"pulsar/internal/core/services/container"
	"pulsar/internal/core/services/envs"
	"pulsar/internal/core/services/log"
	"pulsar/internal/core/services/project"
	"pulsar/internal/core/services/user"
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
	billingService   billing.IBillingService
	userService      user.IUserService
}

func StartServer(db *postgres.Database, mq ports.IMessageQueue, containerMan ports.IContainerManager, fileRepo ports.IFileRepository, jwtSecrete string) {
	logService := log.NewLogService(mq, db)
	resourceService := analytics.NewResourceService(db)
	containerService := container.NewContainerService(containerMan, logService, fileRepo, db, resourceService)
	projectService := project.NewProjectService(db, containerService, fileRepo, jwtSecrete, db)
	envService := envs.NewEnvService(db, *projectService)
	analyticsService := analytics.NewAnalyticsService(db, mq)
	billingService := billing.NewBillingService(db, fileRepo, projectService, analyticsService, resourceService)
	userService := user.NewUserService(db)

	server := &Server{
		echo:             echo.New(),
		projectService:   projectService, // inject project service inject authentication service
		containerService: containerService,
		envService:       envService,
		logService:       logService,
		analyticsService: analyticsService,
		resourceService:  resourceService,
		billingService:   billingService,
		userService:      userService,
	}

	server.echo.Use(middleware.CORS())
	server.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	// setup API routes
	server.DefineRoutes(jwtSecrete)
	// start server
	server.echo.Logger.Fatal(server.echo.Start(":" + os.Getenv("PORT")))
}
