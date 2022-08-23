package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopher-camp/pkg/constants"
	"gopher-camp/pkg/services/env"
	"log"
)

var (
	db *gorm.DB
)

func OpenConnection() {
	//"host=localhost port= user= dbname= password=",
	dbConn, err := gorm.Open(
		"postgres",
		fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v",
			env.GetEnv(constants.DbHost),
			env.GetEnv(constants.DbPort),
			env.GetEnv(constants.DbUser),
			env.GetEnv(constants.DbName),
			env.GetEnv(constants.DbPassword),
		),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	db = dbConn
	err = dbConn.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetDB() *gorm.DB {
	return db
}
