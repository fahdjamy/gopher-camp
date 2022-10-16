package helpers

import (
	"profiler/pkg/models"
	"profiler/pkg/storage/database"
	"profiler/pkg/types"
)

func MigrateFounder(database *database.Database, logger types.Logger) error {
	logger.Info(".........Migrating Founder Table......")
	err := database.GetDB().AutoMigrate(&models.Founder{})
	return err
}

func MigrateProject(database *database.Database, logger types.Logger) error {
	logger.Info(".........Migrating Project Table......")
	err := database.GetDB().AutoMigrate(&models.Project{})
	return err
}

func MigrateCompany(database *database.Database, logger types.Logger) error {
	logger.Info(".........Migrating Company Table......")
	err := database.GetDB().AutoMigrate(&models.Company{})
	return err
}

func MigrateAllModels(database *database.Database, logger types.Logger) error {
	logger.Info(".........Migrating all tables......")
	err := MigrateFounder(database, logger)
	if err != nil {
		return err
	}
	err = MigrateCompany(database, logger)
	if err != nil {
		return err
	}
	err = MigrateProject(database, logger)

	return err
}
