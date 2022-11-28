package main

import (
	"fmt"
	"net/http"
	"profiler/pkg/config"
	"profiler/pkg/constants"
	"profiler/pkg/env"
	"profiler/pkg/helpers"
	muxServer "profiler/pkg/http/rest/mux"
	"profiler/pkg/routes"
	"profiler/pkg/services"
	"profiler/pkg/storage/database"
	"profiler/pkg/types"
	"profiler/pkg/utils"
	"time"
)

func main() {
	port := constants.HostPort()
	logger := utils.NewCustomLogger(constants.InfoLevel)
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

	err := helpers.MigrateAllModels(db, logger)
	if err != nil {
		logger.CustomError(types.CustomError{
			Err:      err,
			DateTime: time.Now(),
			Source:   "tables.MigrateAllModels",
		})
		return
	}

	fileServer := http.FileServer(http.Dir("./static"))
	profileService := services.NewProfileService(*db, logger)
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
		logger.CustomError(types.CustomError{
			Err:      err,
			DateTime: time.Now(),
			Source:   "helpers.SeedDatabaseData",
		})
	}

	muxSrv.Router.Handle("/", fileServer)

	routes.ProjectRoutes(muxSrv, allServices)
	routes.ProfileRoutes(muxSrv, profileService)

	logger.Info("Starting server on port " + port + "\n")

	logger.Fatal(srv.ListenAndServe())
}
