package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func init() {
	client = resty.New()
}

var client *resty.Client

func InitializeRouter() {
	r := gin.Default()

	fmt.Println("router is running ")

	student := r.Group("/student")
	{
		student.POST("/create", CreateStudent)
		student.DELETE("/delete/:rollno", DeleteStudent)
		student.PUT("/update/:rollno", UpdateStudent)
		student.GET("/get", GetStudentQuery)
	}

	class := r.Group("/class")
	{
		class.POST("/create", CreateClass)
		class.DELETE("/delete/:name", DeleteClass)
		class.PUT("/update/:name", UpdateClass)
	}
	att := r.Group("/att")
	{
		att.POST("/add", Addatt)
	}

	err := r.Run(":5000")
	if err != nil {
		return
	}

}
