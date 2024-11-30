package db

import (
	"fmt"
	"slices"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Columns []string
var Databases []string
var Tables []string

var cacheConn *gorm.DB

func cacheObjects() {
	cacheConn, _ = gorm.Open(
		dialector(UserDB),
		&gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Silent),
			SkipDefaultTransaction: true,
	})

	db, _ := cacheConn.DB()
	db.SetMaxOpenConns(1)

	cacheDatabases()

	for _, d := range Databases {
		cacheTables(d)
	}

	for _, t := range Tables {
		if strings.Contains(t, ".") {
			continue
		}
		cacheColumns(t)
	}

	slices.Sort(Columns)
	Columns = slices.Compact(Columns)

	db.Close()
}

func cacheDatabases() {
	var databases []string
	cacheConn.Raw("show databases").Scan(&databases)

	for _, database := range databases {
		if strings.Contains(database, "-") {
			continue
		}

		Databases = append(
			Databases,
			database,
		)
	}
}

func cacheTables(database string) {
	cacheConn.Exec(fmt.Sprintf("use `%s`", database))

	var tables []string
	cacheConn.Raw("show tables").Scan(&tables)

	for _, table := range tables {
		if Selected() == database {
			Tables = append(Tables, table)
			continue
		}

		Tables = append(Tables, strings.Join([]string{database, table}, "."))
	}
}

func cacheColumns(table string) {
	type column struct {
		Field string
	}
	var columns []column

	cacheConn.Exec(fmt.Sprintf("use `%s`", Selected()))
	cacheConn.Raw("show columns from channels").Scan(&columns)

	for _, c := range columns {
		Columns = append(Columns, c.Field)
	}
}
