package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Attendance struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"name"`
	Rollno    int    `gorm:"rollno"`
	Classname int    `gorm:"classname"`
	Day       int    `gorm:"day"`
	Status    string `gorm:"status"`
}

func Addatt(c *gin.Context) {
	var att Attendance
	if err := c.ShouldBindJSON(&att); err != nil {
		c.JSON(400, gin.H{
			"message": "error in binding json",
		})
		return
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(att).
		Post("http://localhost:8000/att/add")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var user Attendance

	// Unmarshal JSON data from the variable into the User struct
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		fmt.Println("Error:", err)
		return
	}
	c.JSON(200, user)
}
