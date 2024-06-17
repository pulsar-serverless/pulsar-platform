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
)

func (server *Server) DefineRoutes(jwtSecrete string) {

	server.echo.GET("/api/projects/plans", billing.GetPricingPlans(server.billingService))

	apiController := server.echo.Group("/api")

	apiController.Use(auth.IsAuthenticated)
	apiController.Use(auth.AuthorizeStatus(server.userService))

	userController := apiController.Group("/users")

	userController.GET("", users.GetUsers(server.userService))
	userController.GET("/status", users.GetAccountStatus(server.userService))
	userController.PUT("/:id/", users.ChangeAccountStatus(server.userService))
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
			projectController.PUT("/site/:projectId", projects.UploadAssets(server.projectService))
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
			projectController.POST("/:projectId/plan", billing.SetProjectPricing(server.billingService))
			projectController.POST("/:projectId/invoice", billing.GenerateInvoice(server.billingService))
		}
	}

	server.echo.POST("/app/status", apps.Status(server.containerService))
	server.echo.GET("/static", apps.ServeSite(server.projectService), auth.IsAuthorized(server.projectService, jwtSecrete))
	server.echo.Any("*",
		apps.ExecuteFunction(server.containerService, server.projectService, server.analyticsService, server.billingService),
		auth.IsAuthorized(server.projectService, jwtSecrete),
	)
}
