package primary

import (
	"os"
	"pulsar/internal/core/services/project"
	"pulsar/internal/ports"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	echo           *echo.Echo
	projectService project.IProjectService
}

func StartServer(pr ports.IProjectRepo) {
	projectService := project.NewProjectService(pr)

	server := &Server{
		echo:           echo.New(),
		projectService: projectService, // inject project service inject authentication service
	}

	server.echo.Use(middleware.CORS())
	server.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	// setup API routes
	server.DefineRoutes()
	// start server
	server.echo.Logger.Fatal(server.echo.Start(":" + os.Getenv("PORT")))
}
