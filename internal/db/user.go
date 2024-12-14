package db

import (
	"fmt"
	"strings"

	"github.com/lewissteele/dbat/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB
var UserDB model.Database

func Connect(name string) {
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

	go cacheObjects()

	d, _ := Conn.DB()
	d.SetMaxOpenConns(1)

	if len(UserDB.Database) > 0 {
		Select(UserDB.Database)
	}
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

func Selected() string {
	var database string
	Conn.Raw("select database()").Scan(&database)
	return database
}

func Select(d string) {
	Conn.Exec(fmt.Sprintf("use `%s`", d))
	go updateSelected()
}

func Query(q string) ([]map[string]interface{}, error) {
	if Conn == nil {
		panic("no connection")
	}

	var results []map[string]interface{}
	err := Conn.Raw(q).Scan(&results).Error

	go updateSelected()

	return results, err
}

func Port(d Driver) string {
	if d == PostgreSQL {
		return "5432"
	}
	return "3306"
}

func updateSelected() {
	UserDB.Database = Selected()
	LocalDB.Save(UserDB)
}

func dialector(u model.Database) gorm.Dialector {
	var dialector gorm.Dialector

	switch Driver(UserDB.Driver) {
	case PostgreSQL:
		dialector = postgres.Open(dsn(UserDB))
	default:
		dialector = mysql.Open(dsn(UserDB))
	}

	return dialector
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
