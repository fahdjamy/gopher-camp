package models

import (
	"gopher-camp/pkg/storage/database"
	"log"
)

func MigrateFounder(database *database.Database) error {
	log.Println(".........Migrating Founder Table......")
	err := database.GetDB().AutoMigrate(&Founder{})
	return err
}

func MigrateProject(database *database.Database) error {
	log.Println(".........Migrating Project Table......")
	err := database.GetDB().AutoMigrate(&Project{})
	return err
}

func MigrateCompany(database *database.Database) error {
	log.Println(".........Migrating Company Table......")
	err := database.GetDB().AutoMigrate(&Company{})
	return err
}

func MigrateAllModels(database *database.Database) error {
	log.Println(".........Migrating all tables......")
	err := MigrateFounder(database)
	if err != nil {
		return err
	}
	err = MigrateCompany(database)
	if err != nil {
		return err
	}
	err = MigrateProject(database)

	return err
}
