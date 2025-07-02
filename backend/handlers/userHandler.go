package handlers

import (
	"backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func newUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}

}

func (u *UserController) GetUsers(c *gin.Context) {
	var users []models.User
	if err := u.DB.Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (u *UserController) PostUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"Error in parsing user to json": err.Error()})
	}

	if err := u.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"Error in writing user to database": err.Error()})
		return
	}
	c.JSON(200, user)
}
