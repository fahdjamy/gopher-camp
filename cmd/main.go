package main

import (
	"fmt"
	"gopher-camp/pkg/config/database"
	"gopher-camp/pkg/constants"
	"gopher-camp/pkg/helpers"
	muxServer "gopher-camp/pkg/http/rest/mux"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/routes"
	"gopher-camp/pkg/services"
	"gopher-camp/pkg/utils"
	"log"
	"net/http"
	"time"
)

func main() {
	port := constants.HostPort()
	logger := utils.NewCustomLogger()
	address := fmt.Sprintf("127.0.0.1%v", port)
	muxSrv, srv := muxServer.NewMuxServer(address, 15*time.Second, 15*time.Second)

	db := database.NewDatabase()
	db.OpenConnection("postgres", constants.DatabaseURI())

	models.MigrateAllModels(db)

	_ = helpers.SeedDatabaseData(db.GetDB(), logger)

	fileServer := http.FileServer(http.Dir("./static"))
	companyService := services.NewCompanyService(*db, logger)
	projectService := services.NewProjectService(*db, logger, companyService)

	muxSrv.Router.Handle("/", fileServer)
	routes.RegisterProjectRoutes(muxSrv, projectService)

	fmt.Printf("Starting server on port " + port + "\n")
	log.Fatal(srv.ListenAndServe())
}
