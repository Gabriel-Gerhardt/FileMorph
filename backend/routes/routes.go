package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReturnRoutes(r *gin.Engine, db *gorm.DB) {
	controller := handlers.NewUserController(db)
	fileController := handlers.NewFileController(db)
	r.GET("", handlers.Test)
	r.GET("/users", controller.GetUsers)
	r.POST("/users", controller.PostUser)
	r.GET("/users/:id", controller.GetUserById)
	r.POST("/file", fileController.PostFile)

}
