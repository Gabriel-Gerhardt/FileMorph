package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	Files    []File
}
