package model

import "gorm.io/gorm"

type Database struct {
	gorm.Model
	Database string
	Driver   string
	Host     string
	Name     string `gorm:"uniqueIndex"`
	Pass     string
	Path     string
	Port     string
	User     string
}
