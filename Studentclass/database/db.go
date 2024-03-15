package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitializeSQLDatabase() {
	dsn := "host=localhost user=admin password=admin dbname=crud port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get DB instance: %v", err)
	}
	sqlDB.Close()
}

func GetDBInstance() *gorm.DB {

	return db
}
