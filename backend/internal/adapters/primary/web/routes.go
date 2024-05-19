package web

import (
	"pulsar/internal/adapters/primary/web/analytics"
	"pulsar/internal/adapters/primary/web/apps"
	"pulsar/internal/adapters/primary/web/auth"
	"pulsar/internal/adapters/primary/web/billing"
	"pulsar/internal/adapters/primary/web/envs"
	"pulsar/internal/adapters/primary/web/log"
	"pulsar/internal/adapters/primary/web/projects"
	"pulsar/internal/adapters/primary/web/resources"
	"pulsar/internal/adapters/primary/web/users"

	"github.com/labstack/echo/v4"
)

func (server *Server) DefineRoutes(jwtSecrete string) {
	apiController := server.echo.Group("/api")

	apiController.Use(auth.IsAuthenticated)

	userController := apiController.Group("/users")

	userController.GET("", users.GetUsers(server.userService))
	userController.DELETE("/:id/projects", projects.DeleteAllProjects(server.projectService))

	projectController := apiController.Group("/projects")
	{
		projectController.POST("", projects.CreateProject(server.projectService))
		projectController.DELETE("/:id", projects.DeleteProject(server.projectService))
		projectController.GET("", projects.GetProjects(server.projectService))
		projectController.GET("/:id", projects.GetProject(server.projectService))
		projectController.PUT("/:id", projects.UpdateProjects(server.projectService))
		projectController.PUT("/:id/api-token", projects.GenerateAPIToken(server.projectService))
		projectController.DELETE("/:id/api-token", projects.RemoveAPIKey(server.projectService))

		{
			projectController.GET("/code/:projectId", projects.DownloadSourceCode(server.projectService))
			projectController.PUT("/code/:projectId", projects.UploadProjectCode(server.projectService))
		}

		{
			projectController.POST("/envs/:projectId", envs.OverwriteEnvVariables(server.envService))
			projectController.GET("/envs/:projectId", envs.GetEnvVariables(server.envService))
		}

		{
			projectController.GET("/logs/:projectId", log.GetProjectLogs(server.logService))
			projectController.DELETE("/logs/:projectId", log.DeleteProjectLogs(server.logService))
		}

		{
			projectController.GET("/:projectId/analytics/hourly", analytics.GetProjectHourlyInvocations(server.analyticsService))
			projectController.GET("/:projectId/analytics/monthly", analytics.GetProjectMonthlyInvocations(server.analyticsService))
			projectController.GET("/:projectId/analytics/weekly", analytics.GetProjectWeeklyInvocations(server.analyticsService))
		}

		{
			projectController.GET("/:projectId/resources/monthly", resources.GetProjectMonthlyConsumption(server.resourceService))
			projectController.GET("/:projectId/resources", resources.GetResourceUtilList(server.resourceService))
			projectController.GET("/:projectId/resources/total", resources.GetProjectTotalUtil(server.resourceService))
		}

		{
			projectController.GET("/plans", billing.GetPricingPlans(server.billingService))
			projectController.POST("/:projectId/plan", billing.SetProjectPricing(server.billingService))
		}
	}

	server.echo.POST("/app/status", apps.Status(server.containerService))
	server.echo.Any("*",
		echo.WrapHandler(apps.NewProxy()),
		auth.IsAuthorized(server.projectService, jwtSecrete),
		apps.ExecuteFunction(server.containerService, server.projectService, server.analyticsService),
	)
}
