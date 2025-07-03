package models

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	UserID   uint
	Name     string
	MimeType string
	Data     []byte `gorm:"type:bytea"`
}
