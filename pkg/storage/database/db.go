package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"profiler/pkg/config"
)

type Database struct {
	db *gorm.DB
}

func (dataBase *Database) OpenPostgresConn(config config.DatabaseConfig) {
	//"host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable",
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
		config.SslMode)

	log.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	dataBase.db = db
}

func (dataBase *Database) GetDB() *gorm.DB {
	return dataBase.db
}

func NewDatabase() *Database {
	return &Database{}
}
