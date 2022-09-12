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
	logger := utils.NewCustomLogger()

	db := database.NewDatabase()
	db.OpenConnection("postgres", constants.DatabaseURI())

	models.MigrateAllModels(db)

	_ = helpers.SeedDatabaseData(db.GetDB(), logger)

	fileServer := http.FileServer(http.Dir("./static"))
	companyService := services.NewCompanyService(*db, logger)
	projectService := services.NewProjectService(*db, logger, companyService)

	r.Handle("/", fileServer)
	routes.RegisterProjectRoutes(r, projectService)

	fmt.Printf("Starting server on port " + port + "\n")
	log.Fatal(http.ListenAndServe(port, r))
}
