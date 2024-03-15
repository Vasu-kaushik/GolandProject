package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strconv"
)

type Student struct {
	ID         uint   `gorm:"primarykey"`
	Rollno     int    `gorm:"rollno"`
	Name       string `gorm:"name"`
	Classname  int    `gorm:"classname"`
	Fathername string `gorm:"fathername"`
	Mothername string `gorm:"mothername"`
}

func CreateStudent(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{
			"message": "error in binding json",
		})
		return
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(student).
		Post("http://localhost:8000/student/create")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the response status code

	var user Student

	// Unmarshal JSON data from the variable into the User struct
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		fmt.Println("Error:", err)
		return
	}
	c.JSON(200, user)

}

func DeleteStudent(c *gin.Context) {
	id := c.Param("rollno")
	url := "http://localhost:8000/student/delete/" + id

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Delete(url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.String(200, string(resp.Body()))

}

func UpdateStudent(c *gin.Context) {
	rollno := c.Param("rollno")
	var student Student
	url := "http://localhost:8000/student/update/" + rollno
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{
			"message": "error in binding json",
		})
		return
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(student).
		Put(url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var user map[string]interface{}

	// Unmarshal JSON data from the variable into the User struct
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		fmt.Println("Error:", err)
		return
	}
	c.JSON(200, user)
}

func GetStudentQuery(c *gin.Context) {
	var queryParams struct {
		Name       string `form:"name"`
		Rollno     int    `form:"rollno"`
		Classname  int    `form:"classname"`
		Fathername string `form:"fathername"`
		Mothername string `form:"mothername"`
	}

	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "http://localhost:8000/student/get?"
	if queryParams.Name != "" {
		query += "name=" + queryParams.Name + "&"
	}
	if queryParams.Rollno != 0 {
		query += "rollno=" + strconv.Itoa(queryParams.Rollno) + "&"
	}
	if queryParams.Classname != 0 {
		query += "classname=" + strconv.Itoa(queryParams.Classname) + "&"
	}
	if queryParams.Fathername != "" {
		query += "fathername=" + queryParams.Fathername + "&"
	}
	if queryParams.Mothername != "" {
		query += "mothername=" + queryParams.Mothername
	}

	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		Get(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var users []Student

	if err := json.Unmarshal(resp.Body(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)

}
