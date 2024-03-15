package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	r := gin.Default()

	r.POST("/class/create", CreateClass)
	r.PUT("/class/update/:name", UpdateClass)
	r.DELETE("/class/delete/:name", DeleteClass)

	r.POST("/student/create", CreateStudent)
	r.PUT("/student/update/:rollno", UpdateStudent)
	r.DELETE("/student/delete/:rollno", Deletestudent)
	r.GET("/student/get", GetStudentQuery)

	r.POST("/att/add", AddAttendence)
	r.PUT("/att/update", UpdateAtt)
	r.GET("/att/total", TotalAtt)
	r.GET("/att/status", Getatt)

	fmt.Println("router is running ")

	err := r.Run(":8000")
	if err != nil {
		return
	}

}
