package handlers

import (
	"backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type fileController struct {
	db *gorm.DB
}

func NewFileController(db *gorm.DB) *fileController {
	return &fileController{db: db}

}

func (u *fileController) PostFile(c *gin.Context) {
	var file models.File

	if err := c.BindJSON(&file); err != nil {
		c.JSON(400, gin.H{"Error in parsing file to json": err.Error()})
	}

	if err := u.db.Create(&file).Error; err != nil {
		c.JSON(500, gin.H{"Error in writing file to database": err.Error()})
		return
	}
	c.JSON(200, file)
}
