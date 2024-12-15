package db

import (
	"fmt"
	"slices"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Columns []string
var databases []string
var tables []string

func Databases() []string {
	if len(databases) > 0 {
		return databases
	}

	if UserDB.Driver == string(SQLite) {
		return databases
	}

	Conn.Raw("show databases").Scan(&databases)

	for _, database := range databases {
		if strings.Contains(database, "-") {
			continue
		}

		databases = append(
			databases,
			database,
		)
	}

	return databases
}

func Tables() []string {
	if len(tables) != 0 {
		return tables
	}

	if UserDB.Driver == string(SQLite) {
		err := Conn.Raw(
			"select tbl_name from sqlite_master where type = ? and tbl_name != ?",
			"table",
			"sqlite_sequence",
		).Scan(&tables).Error

		if err != nil {
			panic(err)
		}

		return tables
	}

	c := newConn()

	for _, database := range Databases() {
		c.Exec(fmt.Sprintf("use `%s`", database))

		var tables []string
		c.Raw("show tables").Scan(&tables)

		for _, table := range tables {
			if Selected() == database {
				tables = append(tables, table)
				continue
			}

			tables = append(tables, strings.Join([]string{database, table}, "."))
		}
	}

	d, _ := c.DB()
	d.Close()

	return tables
}

func cacheColumns(table string) {
	c := newConn()

	type column struct {
		Field string
	}
	var columns []column

	c.Exec(fmt.Sprintf("use `%s`", Selected()))
	c.Raw("show columns from channels").Scan(&columns)

	for _, column := range columns {
		Columns = append(Columns, column.Field)
	}

	d, _ := c.DB()
	d.Close()
}

func cacheObjects() {
	Databases()

	for _, t := range Tables() {
		if strings.Contains(t, ".") {
			continue
		}
		cacheColumns(t)
	}

	slices.Sort(Columns)
	Columns = slices.Compact(Columns)
}

func newConn() *gorm.DB {
	c, _ := gorm.Open(
		dialector(UserDB),
		&gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Silent),
			SkipDefaultTransaction: true,
		})

	d, _ := c.DB()
	d.SetMaxOpenConns(1)

	return c
}

