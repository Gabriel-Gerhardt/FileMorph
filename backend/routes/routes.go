package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReturnRoutes(r *gin.Engine, db *gorm.DB) {
	userController := handlers.NewUserController(db)
	fileController := handlers.NewFileController(db)
	r.GET("", handlers.Test)
	r.GET("/users", userController.GetUsers)
	r.POST("/users", userController.PostUser)
	r.GET("/users/id/:id", userController.GetUserById)
	r.POST("/file", fileController.PostFile)
	r.GET("/users/name/:name", userController.GetUserByName)

}
