package routes

import (
	"github.com/gorilla/mux"
	"gopher-camp/pkg/controllers"
	"gopher-camp/pkg/dto"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/storage"
)

var RegisterProjectRoutes = func(router *mux.Router, service storage.Storage[models.Project, dto.ProjectDTO]) {
	projectController := controllers.NewProjectController(service)
	router.HandleFunc("/projects/", projectController.GetProjects).Methods("GET")
	router.HandleFunc("/projects/", projectController.CreateProject).Methods("POST")
	router.HandleFunc("/projects/{projectId}", projectController.GetProject).Methods("GET")
	router.HandleFunc("/projects/{projectId}", projectController.UpdateProject).Methods("POST")
	router.HandleFunc("/projects/{projectId}", projectController.DeleteProjects).Methods("DELETE")
}
