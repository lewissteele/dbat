package model

import "gorm.io/gorm"

type Database struct {
	gorm.Model
	Host string
	Username string
	Password string
}
