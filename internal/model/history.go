package model

import "gorm.io/gorm"

type History struct {
	gorm.Model
	DatabaseID int
	Query string
}
