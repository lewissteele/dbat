package model

import "gorm.io/gorm"

type Database struct {
	gorm.Model
	Driver string
	Host   string
	Name   string
	Pass   string
	Port   string
	User   string
}
