package database

import "log"

func AutoMigrate(data interface{}) (err error) {
	// auto migrate create the table based on the struct if it doesn't exit
	err = db.AutoMigrate(data)
	if err != nil {
		log.Fatal("faied to auto migrate table: %v", err)
	}
	return
}
