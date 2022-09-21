package main

import (
	"fmt"
	"gopher-camp/pkg/config"
	"gopher-camp/pkg/constants"
	"gopher-camp/pkg/env"
	"gopher-camp/pkg/helpers"
	muxServer "gopher-camp/pkg/http/rest/mux"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/routes"
	"gopher-camp/pkg/services"
	"gopher-camp/pkg/storage/database"
	"gopher-camp/pkg/types"
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
	dbConfig := config.DatabaseConfig{
		Host:     env.GetEnv(constants.DBHost),
		Name:     env.GetEnv(constants.DBName),
		Port:     env.GetEnv(constants.DBPort),
		User:     env.GetEnv(constants.DBUser),
		Password: env.GetEnv(constants.DBPassword),
		SslMode:  env.GetEnvOrDefault(constants.DBSSLMode, "disable"),
	}
	db.OpenPostgresConn(dbConfig)

	err := models.MigrateAllModels(db)
	if err != nil {
		logger.LogError(types.CustomError{
			Err:      err,
			DateTime: time.Now(),
			Source:   "tables.MigrateAllModels",
		})
		return
	}

	fileServer := http.FileServer(http.Dir("./static"))
	companyService := services.NewCompanyService(*db, logger)
	founderService := services.NewFounderService(*db, logger)
	projectService := services.NewProjectService(*db, logger, companyService)

	allServices := types.AllServices{
		ProjectService: projectService,
		CompanyService: companyService,
		FounderService: founderService,
	}
	err = helpers.SeedDatabaseData(logger, allServices)
	if err != nil {
		logger.LogError(types.CustomError{
			Err:      err,
			DateTime: time.Now(),
			Source:   "helpers.SeedDatabaseData",
		})
	}

	muxSrv.Router.Handle("/", fileServer)
	routes.RegisterProjectRoutes(muxSrv, allServices)

	fmt.Printf("Starting server on port " + port + "\n")
	log.Fatal(srv.ListenAndServe())
}
