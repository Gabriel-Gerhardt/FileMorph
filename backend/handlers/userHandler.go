package handlers

import (
	"backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *userController {
	return &userController{db: db}

}

func (u *userController) GetUsers(c *gin.Context) {
	var users []models.User
	if err := u.db.Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (u *userController) PostUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"Error in parsing user to json": err.Error()})
	}

	if err := u.db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"Error in writing user to database": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (u *userController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		c.JSON(500, gin.H{"Error in finding user": err.Error()})
		return
	}
	c.JSON(200, user)
}
