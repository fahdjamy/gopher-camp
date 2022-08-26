package routes

import (
	"github.com/gorilla/mux"
	"gopher-camp/pkg/controllers"
)

var RegisterProjectRoutes = func(router *mux.Router) {
	router.HandleFunc("/projects/", controllers.GetProjects).Methods("GET")
	router.HandleFunc("/projects/", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/projects/{projectId}", controllers.GetProject).Methods("GET")
	router.HandleFunc("/projects/{projectId}", controllers.UpdateProject).Methods("POST")
	router.HandleFunc("/projects/{projectId}", controllers.DeleteProjects).Methods("DELETE")
}
