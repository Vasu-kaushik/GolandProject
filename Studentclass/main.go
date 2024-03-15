package main

import (
	"Studentclass/database"
	"fmt"
)

func main() {
	database.InitializeSQLDatabase()

	defer database.CloseDB()

	err := database.AutoMigrate(&Class{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = database.AutoMigrate(&Student{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = database.AutoMigrate(&Attendance{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	InitializeRouter()

	database.CloseDB()

}
