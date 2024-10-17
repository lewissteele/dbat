package db

import (
	"strings"

	"github.com/lewissteele/dbat/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Driver string

const (
	MariaDB    Driver = "mariadb"
	MySQL      Driver = "mysql"
	PostgreSQL Driver = "postgresql"
	SQLite     Driver = "sqlite"
)

func UserDB(name string) *gorm.DB {
	userDB, err := gorm.Open(mysql.Open(
		dsn(name)),
		&gorm.Config{
			//Logger: logger.Default.LogMode(logger.Silent),
			SkipDefaultTransaction: true,
		},
	)

	if err != nil {
		panic("could not connect")
	}

	return userDB
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

func dsn(name string) string {
	db := model.Database{}

	LocalDB.Where("name = ?", name).First(&db)

	return strings.Join([]string{
		db.User,
		":",
		db.Pass,
		"@tcp(",
		db.Host,
		":",
		db.Port,
		")/",
	}, "")
}
