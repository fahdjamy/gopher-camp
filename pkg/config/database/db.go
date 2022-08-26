package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopher-camp/pkg/constants"
	"log"
)

var (
	db *gorm.DB
)

func OpenConnection() {
	//"host=localhost port= user= dbname= password=",

	dbConn, err := gorm.Open("postgres", constants.DatabaseURI())
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
