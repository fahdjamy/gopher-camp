package routes

import (
	"github.com/gorilla/mux"
	"gopher-camp/pkg/config/database"
	"gopher-camp/pkg/controllers"
)

var RegisterProjectRoutes = func(router *mux.Router, db database.Database) {
	projectController := controllers.NewProjectController(db)
	router.HandleFunc("/projects/", projectController.GetProjects).Methods("GET")
	router.HandleFunc("/projects/", projectController.CreateProject).Methods("POST")
	router.HandleFunc("/projects/{projectId}", projectController.GetProject).Methods("GET")
	router.HandleFunc("/projects/{projectId}", projectController.UpdateProject).Methods("POST")
	router.HandleFunc("/projects/{projectId}", projectController.DeleteProjects).Methods("DELETE")
}
