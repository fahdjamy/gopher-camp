package routes

import (
	"profiler/pkg/controllers"
	"profiler/pkg/types"
)

var ProjectRoutes = func(routers types.RestRouters, services types.AllServices) {
	projectController := controllers.NewProjectController(services)

	routers.Get("/projects", projectController.GetProjects)
	routers.Post("/projects", projectController.CreateProject)
	routers.Get("/projects/{projectId}", projectController.GetProject)
	routers.Put("/projects/{projectId}", projectController.UpdateProject)
	routers.Delete("/projects/{projectId}", projectController.DeleteProjectById)
}
