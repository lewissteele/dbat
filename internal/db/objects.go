package db

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Databases []string
var Tables []string

func cacheObjects() {
	conn, _ := gorm.Open(
		dialector(UserDB),
		&gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Silent),
			SkipDefaultTransaction: true,
	})

	db, _ := conn.DB()
	db.SetMaxOpenConns(1)

	cacheDatabases(conn)

	for _, d := range Databases {
		cacheTables(d, conn)
	}

	db.Close()
}

func cacheDatabases(conn *gorm.DB) {
	var databases []string
	Conn.Raw("show databases").Scan(&databases)

	for _, database := range databases {
		Databases = append(
			Databases,
			database,
		)
	}
}

func cacheTables(database string, conn *gorm.DB) {
	conn.Exec(fmt.Sprintf("use `%s`", database))

	var tables []string
	conn.Raw("show tables").Scan(&tables)

	for _, table := range tables {
		if Selected() == database {
			Tables = append(Tables, table)
			continue
		}

		Tables = append(Tables, strings.Join([]string{database, table}, "."))
	}
}
