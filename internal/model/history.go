package model

import "gorm.io/gorm"

type History struct {
	gorm.Model
	Database   Database
	DatabaseID int `gorm:"index"`
	Query      string
}
