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
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "no file received", "details": err.Error()})
		return
	}

	fileData, err := fileHeader.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot open file", "details": err.Error()})
		return
	}
	defer fileData.Close()

	buf := make([]byte, fileHeader.Size)
	_, err = fileData.Read(buf)
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot read file", "details": err.Error()})
		return
	}

	mimeType := fileHeader.Header.Get("Content-Type")

	newFile := models.File{
		Name:     fileHeader.Filename,
		Data:     buf,
		MimeType: mimeType,
	}

	if err := u.db.Create(&newFile).Error; err != nil {
		c.JSON(500, gin.H{"error": "cannot save file", "details": err.Error()})
		return
	}

	c.JSON(200, newFile)
}
