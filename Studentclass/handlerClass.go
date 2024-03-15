package main

import (
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

	err := c.ShouldBindJSON(&class)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error in binding student ",
		})
	}
	CreateClassRep(class, c)
	//
	//if err = database.GetDBInstance()s.Create(&class).Error; err != nil {
	//	c.JSON(400, gin.H{
	//		"error": err.Error(),
	//	})
	//}
	c.JSON(200, class)
}

func UpdateClass(c *gin.Context) {
	name := c.Param("name")

	var class Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(400, gin.H{
			"message": "error in binding class",
		})
	}
	//var existing Class
	//if err := database.GetDBInstance().Where("name= ?", name).First(&existing).Error; err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		if err = database.GetDBInstance().Create(&class).Error; err != nil {
	//			c.JSON(400, gin.H{
	//				"error": err.Error(),
	//			})
	//			return
	//
	//		}
	//	} else {
	//		c.JSON(400, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}
	//}
	//class.ID = existing.ID
	//if err := database.GetDBInstance().Model(&existing).Save(&class).Error; err != nil {
	//	c.JSON(400, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	UpdateClassRepo(class, c, name)
	c.JSON(200, class)

}

func DeleteClass(c *gin.Context) {
	name := c.Param("name")
	//
	//if err := database.GetDBInstance().Where("name= ?", name).Delete(&Class{}).Error; err != nil {
	//	c.JSON(400, gin.H{
	//		"error": err.Error(),
	//	})
	//
	//}
	//if err := database.GetDBInstance().Where("classname= ?", name).Delete(&Student{}).Error; err != nil {
	//	c.JSON(400, gin.H{
	//		"error": err.Error(),
	//	})
	//
	//}
	DeleteClassRepo(c, name)
	c.JSON(200, gin.H{
		"message": "deleted Successfully",
	})

}
