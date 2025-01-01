package model

import "gorm.io/gorm"

type History struct {
	gorm.Model
	Database   Database
	DatabaseID int    `gorm:"uniqueIndex:idx_histories"`
	Query      string `gorm:"uniqueIndex:idx_histories"`
}
