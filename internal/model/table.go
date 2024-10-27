package model

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	Database   Database
	DatabaseID int `gorm:"index"`
	Name       string
}
