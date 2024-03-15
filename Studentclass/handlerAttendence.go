package main

import (
	"Studentclass/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Attendance struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"name"`
	Rollno    int    `gorm:"rollno"`
	Classname int    `gorm:"classname"`
	Day       int    `gorm:"day"`
	Status    string `gorm:"status"`
}

func AddAttendence(c *gin.Context) {
	var att Attendance

	if err := c.ShouldBindJSON(&att); err != nil {
		c.JSON(400, gin.H{
			"message": "error in binding attendance",
		})
	}
	var stu Student
	if err := database.GetDBInstance().Where("rollno= ?", att.Rollno).Where("classname= ?", att.Classname).First(&stu).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(400, gin.H{
				"error": "student does'nt exist",
			})
			return
		} else {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	if err := database.GetDBInstance().Create(&att).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

	}
	c.JSON(400, att)
}

func UpdateAtt(c *gin.Context) {
	var attq Attendance
	if err := c.ShouldBindJSON(&attq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var existingAttendance Attendance
	if err := database.GetDBInstance().Where("day = ?", attq.Day).
		Where("classname = ?", attq.Classname).
		Where("rollno = ?", attq.Rollno).
		First(&existingAttendance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Attendance record not found",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	existingAttendance.Name = attq.Name
	existingAttendance.Status = attq.Status

	if err := database.GetDBInstance().Model(&existingAttendance).Save(&existingAttendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update Attendance record",
		})
		return
	}

	c.JSON(http.StatusOK, existingAttendance)
}

func TotalAtt(c *gin.Context) {
	rollno := c.Query("rollno")
	classname := c.Query("classname")

	var att []Attendance
	if err := database.GetDBInstance().Where("rollno= ?", rollno).Where("classname= ?", classname).Find(&att).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "student doesn't exist",
		})
	}
	c.JSON(400, len(att))
}

func Getatt(c *gin.Context) {
	rollno := c.Query("rollno")
	classname := c.Query("classname")
	day := c.Query("day")

	var att Attendance
	if err := database.GetDBInstance().Where("rollno= ?", rollno).Where("classname= ?", classname).Where("day= ?", day).First(&att).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "student doesn't exist",
		})
	}
	c.JSON(400, att.Status)
}
