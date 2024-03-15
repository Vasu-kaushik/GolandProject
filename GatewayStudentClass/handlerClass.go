package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Class struct {
	ID             uint   `gorm:"primarykey"`
	Name           int    `gorm:"name"`
	Incharge       string `gorm:"Incharge"`
	Representative string `gorm:"representative"`
}

func CreateClass(c *gin.Context) {
	var class Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(400, gin.H{
			"message": "error in binding json",
		})
		return
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(class).
		Post("http://localhost:8000/class/create")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the response status code

	var user Class

	// Unmarshal JSON data from the variable into the User struct
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		fmt.Println("Error:", err)
		return
	}
	c.JSON(200, user)
}

func DeleteClass(c *gin.Context) {
	id := c.Param("name")
	url := "http://localhost:8000/class/delete/" + id

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Delete(url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var ee map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &ee); err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.JSON(200, ee)
}

func UpdateClass(c *gin.Context) {

	var class Class
	name := c.Param("name")
	url := "http://localhost:8000/class/update/" + name
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(400, gin.H{
			"message": "error in binding json",
		})
		return
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(class).
		Put(url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var user Class

	// Unmarshal JSON data from the variable into the User struct
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		fmt.Println("Error:", err)
		return
	}
	c.JSON(200, user)
}
