package web

import (
	"pulsar/internal/adapters/primary/web/apps"
	"pulsar/internal/adapters/primary/web/auth"
	"pulsar/internal/adapters/primary/web/projects"

	"github.com/labstack/echo/v4"
)

func (server *Server) DefineRoutes() {
	apiController := server.echo.Group("/api")

	apiController.Use(auth.IsAuthenticated)

	projectController := apiController.Group("/projects")
	{
		projectController.POST("", projects.CreateProject(server.projectService))
		projectController.DELETE("/:id", projects.DeleteProject(server.projectService))
		projectController.GET("", projects.GetProjects(server.projectService))
		projectController.GET("/:id", projects.GetProject(server.projectService))
		projectController.PUT("/:id", projects.UpdateProjects(server.projectService))

		{
			projectController.GET("/code/:projectId", projects.DownloadSourceCode(server.projectService))
		}
	}

	server.echo.POST("/app/status", apps.Status(server.containerService))
	server.echo.Any("*",
		echo.WrapHandler(apps.NewProxy()),
		apps.ExecuteFunction(server.containerService, server.projectService))
}
