package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type Database struct {
	db *gorm.DB
}

func (dataBase *Database) OpenConnection(dialect string, dbURI string) {
	//"host=localhost port= user= dbname= password=",

	dbConn, err := gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
		return
	}
	dataBase.db = dbConn
}

func (dataBase *Database) GetDB() *gorm.DB {
	return dataBase.db
}

func NewDatabase() *Database {
	return &Database{}
}
