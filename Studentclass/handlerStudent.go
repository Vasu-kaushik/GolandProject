package main

import (
	"Studentclass/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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
			"error": "error in binding student",
		})
		return
	}

	if err := database.GetDBInstance().Create(&student).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	name := student.Classname

	var class Class

	if err := database.GetDBInstance().Where("name = ?", name).First(&class).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			class.Name = name
			if err := database.GetDBInstance().Create(&class).Error; err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}
		} else {
			c.JSON(400, gin.H{
				"error": "Error finding class",
			})
			return
		}
	}
	c.JSON(400, student)
}

func UpdateStudent(c *gin.Context) {
	name := c.Param("rollno")
	var student Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": "Failed to decode JSON"})
		return
	}

	var existingStudent Student
	if err := database.GetDBInstance().Where("rollno = ?", name).First(&existingStudent).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			if err := database.GetDBInstance().Create(&student).Error; err != nil {
				c.JSON(500, gin.H{"error": "Failed to create student"})
				return
			}

			var class Class

			if err := database.GetDBInstance().Where("name = ?", student.Classname).First(&class).Error; err != nil {
				if err == gorm.ErrRecordNotFound {

					class.Name = student.Classname
					if err := database.GetDBInstance().Create(&class).Error; err != nil {
						c.JSON(400, gin.H{
							"error": err.Error(),
						})
						return
					}
				} else {
					c.JSON(400, gin.H{
						"error": "Error finding class",
					})
					return
				}
			}

			c.JSON(200, gin.H{"message": "Student created successfully"})
		} else {
			c.JSON(500, gin.H{"error": "Database error"})
		}
		return
	}
	student.ID = existingStudent.ID

	if err := database.GetDBInstance().Model(&existingStudent).Updates(&student).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update student"})
		return
	}
	var class Class

	if err := database.GetDBInstance().Where("name = ?", existingStudent.Classname).First(&class).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			class.Name = existingStudent.Classname
			if err := database.GetDBInstance().Create(&class).Error; err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}
		} else {
			c.JSON(400, gin.H{
				"error": "Error finding class",
			})
			return
		}
	}

	c.JSON(200, gin.H{"message": "Student updated successfully"})
}

func Deletestudent(c *gin.Context) {
	rollno := c.Param("rollno")

	if err := database.GetDBInstance().Where("rollno= ?", rollno).Delete(&Student{}).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.String(200, "deleted person successfully")
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

	query := database.GetDBInstance()

	if queryParams.Name != "" {
		query = query.Where("name = ?", queryParams.Name)
	}
	if queryParams.Rollno != 0 {
		query = query.Where("rollno = ?", queryParams.Rollno)
	}
	if queryParams.Classname != 0 {
		query = query.Where("classname = ?", queryParams.Classname)
	}
	if queryParams.Fathername != "" {
		query = query.Where("fathername = ?", queryParams.Fathername)
	}
	if queryParams.Mothername != "" {
		query = query.Where("mothername = ?", queryParams.Mothername)
	}

	var students []Student
	if err := query.Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}
