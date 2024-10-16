package db

import (
	"strconv"
	"strings"

	"github.com/lewissteele/dbat/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Drivers = []string{
	"mariadb",
	"mysql",
	"sqlite",
}

func UserDB(host string) *gorm.DB {
	userDB, err := gorm.Open(mysql.Open(
		dsn(host)),
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

func dsn(host string) string {
	db := model.Database{}

	LocalDB.Where("host = ?", host).First(&db)

	return strings.Join([]string{
		db.User,
		":",
		db.Pass,
		"@tcp(",
		db.Host,
		":",
		strconv.FormatUint(uint64(db.Port), 10),
		")/",
	}, "")
}
