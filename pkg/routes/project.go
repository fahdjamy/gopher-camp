package routes

import (
	"gopher-camp/pkg/controllers"
	"gopher-camp/pkg/types"
)

var RegisterProjectRoutes = func(routers types.RestRouters, services types.AllServices) {
	projectController := controllers.NewProjectController(services)

	routers.Get("/projects", projectController.GetProjects)
	routers.Post("/projects", projectController.CreateProject)
	routers.Get("/projects/{projectId}", projectController.GetProject)
	routers.Put("/projects/{projectId}", projectController.UpdateProject)
	routers.Delete("/projects/{projectId}", projectController.DeleteProjectById)
}
