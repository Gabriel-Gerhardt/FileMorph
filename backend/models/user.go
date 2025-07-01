package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Files []File // use "BLOB" para MySQL, "bytea" para PostgreSQL
}
