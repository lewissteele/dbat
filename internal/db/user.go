package db

import (
	"strings"

	"github.com/lewissteele/dbat/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Driver string

const (
	MariaDB    Driver = "mariadb"
	MySQL      Driver = "mysql"
	PostgreSQL Driver = "postgresql"
	SQLite     Driver = "sqlite"
)

func UserDB(name string) *gorm.DB {
	var userDB model.Database
	LocalDB.Where("name = ?", name).Find(&userDB)

	var dialector gorm.Dialector

	switch Driver(userDB.Driver) {
	case PostgreSQL:
		dialector = postgres.Open(dsn(&userDB))
	default:
		dialector = mysql.Open(dsn(&userDB))
	}

	gormDB, err := gorm.Open(dialector, &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("could not connect")
	}

	return gormDB
}

func UserDBNames() []string {
	databases := []model.Database{}
	LocalDB.Find(&databases)

	var names []string

	for _, db := range databases {
		names = append(names, db.Name)
	}

	return names
}

func Port(driver Driver) string {
	if driver == PostgreSQL {
		return "5432"
	}
	return "3306"
}

func dsn(db *model.Database) string {
	if Driver(db.Driver) == PostgreSQL {
		return strings.Join([]string{
			"host=",
			db.Host,
			" ",
			"user=",
			db.User,
			" ",
			"password=",
			db.Pass,
			" ",
			"port=",
			db.Port,
		}, "")
	}

	return strings.Join([]string{
		db.User,
		":",
		db.Pass,
		"@tcp(",
		db.Host,
		":",
		db.Port,
		")/?parseTime=true",
	}, "")
}
