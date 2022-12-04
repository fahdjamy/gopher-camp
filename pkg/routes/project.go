package routes

import (
	"profiler/pkg/controllers"
	"profiler/pkg/types"
)

var ProjectRoutes = func(routers types.RestRouters, services types.AllServices) {
	projectController := controllers.NewProjectController(services)

	routers.Get("/projects", projectController.HandleGetAll)
	routers.Post("/projects", projectController.HandleCreate)
	routers.Put("/projects/{projectId}", projectController.HandleUpdate)
	routers.Get("/projects/{projectId}", projectController.HandleGetById)
	routers.Delete("/projects/{projectId}", projectController.HandleDeleteById)
}
