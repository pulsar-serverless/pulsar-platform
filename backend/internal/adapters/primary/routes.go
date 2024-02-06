package primary

import "pulsar/internal/adapters/primary/projects"

func (server *Server) DefineRoutes() {

	projectController := server.echo.Group("/projects")
	{
		projectController.POST("", projects.CreateProject(server.projectService))
		projectController.DELETE("/:id", projects.DeleteProject(server.projectService))
		projectController.GET("", projects.GetProjects(server.projectService))
		projectController.GET("/:id", projects.GetProject(server.projectService))
		projectController.PUT("/:id", projects.UpdateProjects(server.projectService))
	}
}
