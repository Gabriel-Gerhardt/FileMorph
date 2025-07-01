package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func ReturnRoutes(r *gin.Engine) {
	r.GET("", handlers.Test)
}
