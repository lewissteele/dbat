package db

import (
	"strings"

	"github.com/lewissteele/dbat/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB
var Databases []string
var UserDB model.Database

func Connect(name string) (*gorm.DB, *model.Database) {
	LocalDB.Where("name = ?", name).Find(&UserDB)

	var dialector gorm.Dialector

	switch Driver(UserDB.Driver) {
	case PostgreSQL:
		dialector = postgres.Open(dsn(UserDB))
	default:
		dialector = mysql.Open(dsn(UserDB))
	}

	var err error

	Conn, err = gorm.Open(
		dialector,
		&gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Silent),
			SkipDefaultTransaction: true,
		})

	if err != nil {
		panic("could not connect")
	}

	go populateDatabases()

	return Conn, &UserDB
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

func Port(d Driver) string {
	if d == PostgreSQL {
		return "5432"
	}
	return "3306"
}

func populateDatabases() {
	rows, err := Conn.Raw("show databases").Rows()

	if err != nil {
		panic("could not get databases")
	}

	var results []map[string]interface{}

	rows.Next()
	err = Conn.ScanRows(rows, &results)

	for _, val := range results {
		Databases = append(
			Databases,
			val["Database"].(string),
		)
	}
}

func dsn(d model.Database) string {
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
