package routes

import (
	"github.com/gorilla/mux"
	"gopher-camp/pkg/controllers"
)

var RegisterProjectRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.GetProjects).Methods("GET")
	router.HandleFunc("/", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/{projectId}", controllers.GetProject).Methods("GET")
	router.HandleFunc("/{projectId}", controllers.UpdateProject).Methods("POST")
	router.HandleFunc("/{projectId}", controllers.DeleteProjects).Methods("DELETE")
}
