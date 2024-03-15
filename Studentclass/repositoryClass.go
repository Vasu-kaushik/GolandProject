package main

import (
	"Studentclass/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateClassRep(class Class, c *gin.Context) {
	if err := database.GetDBInstance().Create(&class).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
}

func UpdateClassRepo(class Class, c *gin.Context, name string) {
	var existing Class
	if err := database.GetDBInstance().Where("name= ?", name).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err = database.GetDBInstance().Create(&class).Error; err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return

			}
		} else {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	class.ID = existing.ID
	if err := database.GetDBInstance().Model(&existing).Save(&class).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func DeleteClassRepo(c *gin.Context, name string) {

	if err := database.GetDBInstance().Where("name= ?", name).Delete(&Class{}).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

	}
	if err := database.GetDBInstance().Where("classname= ?", name).Delete(&Student{}).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

	}

}
