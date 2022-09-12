package models

import (
	"gopher-camp/pkg/config/database"
	"log"
)

func MigrateProject(database *database.Database) {
	log.Println(".........Migrating Project Table......")
	database.GetDB().AutoMigrate(&Project{})
}

func MigrateCompany(database *database.Database) {
	log.Println(".........Migrating Company Table......")
	database.GetDB().AutoMigrate(&Company{})
}

func MigrateAllModels(database *database.Database) {
	log.Println(".........Migrating all tables......")
	database.GetDB().AutoMigrate(&Founder{}, &Company{}, &Project{})
}
