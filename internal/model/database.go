package model

import (
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	gorm.Model
	Driver Driver
	Host   string
	Name   string `gorm:"uniqueIndex"`
	Pass   string
	Port   string
	User   string
}

type Driver string

const (
	MariaDB    Driver = "mariadb"
	MySQL      Driver = "mysql"
	PostgreSQL Driver = "postgresql"
	SQLite     Driver = "sqlite"
)

func (d Database) Conn() *gorm.DB {
	var dialector gorm.Dialector

	switch d.Driver {
	case PostgreSQL:
		dialector = postgres.Open(d.dsn())
	default:
		dialector = mysql.Open(d.dsn())
	}

	conn, err := gorm.Open(
		dialector,
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
	if Driver(d.Driver) == PostgreSQL {
		return strings.Join([]string{
			"host=",
			d.Host,
			" ",
			"user=",
			d.User,
			" ",
			"password=",
			d.Pass,
			" ",
			"port=",
			d.Port,
		}, "")
	}

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
