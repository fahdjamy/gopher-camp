package routes

import (
	"gopher-camp/pkg/controllers"
	"gopher-camp/pkg/dto"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/storage"
	"gopher-camp/pkg/types"
)

var RegisterProjectRoutes = func(routers types.RestRouters, service storage.Storage[models.Project, dto.ProjectReqDTO, dto.ProjectResponseDTO]) {
	projectController := controllers.NewProjectController(service)

	routers.Get("/projects/", projectController.GetProjects)
	routers.Post("/projects/", projectController.CreateProject)
	routers.Get("/projects/{projectId}/", projectController.GetProject)
	routers.Put("/projects/{projectId}/", projectController.UpdateProject)
	routers.Delete("/projects/{projectId}/", projectController.DeleteProjects)
}
