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
var Tables []string

func Databases() []string {
	if len(databases) > 0 {
		return databases
	}

	c := NewConn()
	c.Raw("show databases").Scan(&databases)

	for _, database := range databases {
		if strings.Contains(database, "-") {
			continue
		}

		databases = append(
			databases,
			database,
		)
	}

	d, _ := c.DB()
	d.Close()

	return databases
}

func NewConn() *gorm.DB {
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

func cacheObjects() {
	for _, d := range Databases() {
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
}

func cacheTables(database string) {
	c := NewConn()
	c.Exec(fmt.Sprintf("use `%s`", database))

	var tables []string
	c.Raw("show tables").Scan(&tables)

	for _, table := range tables {
		if Selected() == database {
			Tables = append(Tables, table)
			continue
		}

		Tables = append(Tables, strings.Join([]string{database, table}, "."))
	}

	d, _ := c.DB()
	d.Close()
}

func cacheColumns(table string) {
	c := NewConn()

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
