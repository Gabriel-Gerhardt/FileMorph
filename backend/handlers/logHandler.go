package handlers

import (
	"backend/models"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type logController struct {
	db *gorm.DB
}

var (
	logCrtl *logController
	once    sync.Once
)

func GetLogController(db *gorm.DB) *logController {
	once.Do(func() {
		logCrtl = &logController{db: db}
	})
	return logCrtl
}
func (l *logController) GetLog(c *gin.Context) {
	var log models.Log
	if err := l.db.First(&log).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, log)
}
