package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopher-camp/pkg/config/database"
	"gopher-camp/pkg/constants"
	"gopher-camp/pkg/helpers"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/routes"
	"gopher-camp/pkg/services"
	"gopher-camp/pkg/utils"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	port := constants.HostPort()

	db := database.NewDatabase()
	db.OpenConnection("postgres", constants.DatabaseURI())

	models.MigrateAllModels(db)
	logger := utils.NewLogger()

	_ = helpers.SeedDatabaseData(db.GetDB(), logger)

	fileServer := http.FileServer(http.Dir("./static"))

	r.Handle("/", fileServer)

	projectService := services.NewProjectService(*db)
	routes.RegisterProjectRoutes(r, projectService)

	fmt.Printf("Starting server on port " + port + "\n")
	log.Fatal(http.ListenAndServe(port, r))
}
