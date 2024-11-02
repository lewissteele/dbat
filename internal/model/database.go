package model

import (
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	gorm.Model
	Driver string
	Host   string
	Name   string `gorm:"uniqueIndex"`
	Pass   string
	Port   string
	User   string
}

func (d Database) Conn() *gorm.DB {
	conn, err := gorm.Open(
		mysql.Open(d.dsn()),
		&gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Silent),
			SkipDefaultTransaction: true,
		})

	if err != nil {
		panic("could not connect")
	}

	return conn
}

func (d Database) dsn() string {
	return strings.Join([]string{
		d.User,
		":",
		d.Pass,
		"@tcp(",
		d.Host,
		":",
		d.Port,
		")/?parseTime=true",
	}, "")
}
