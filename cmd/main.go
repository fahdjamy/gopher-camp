package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopher-camp/pkg/config/database"
	"gopher-camp/pkg/constants"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/routes"
	"gopher-camp/pkg/services/storage"
	"log"
	"net/http"
)

func main() {
	port := ":8008"
	r := mux.NewRouter()

	db := database.NewDatabase()
	db.OpenConnection("postgres", constants.DatabaseURI())
	models.MigrateAllModels(db)

	fileServer := http.FileServer(http.Dir("./static"))

	r.Handle("/", fileServer)

	projectService := storage.NewProjectService(*db)
	routes.RegisterProjectRoutes(r, projectService)

	fmt.Printf("Starting server on port " + port + "\n")
	log.Fatal(http.ListenAndServe(port, r))
}
