package db

import (
	"strings"

	"github.com/lewissteele/dbat/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetUserDB(host string) *gorm.DB {
	userDB, err := gorm.Open(mysql.Open(
		getDSN(host)),
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

func GetUserDBList() []string {
	databases := []model.Database{}
	LocalDB.Find(&databases)

	var names []string

	for _, db := range databases {
		names = append(names, db.Host)
	}

	return names
}

func getDSN(host string) string {
	db := model.Database{}

	LocalDB.Where("host = ?", host).First(&db)

	return strings.Join([]string{
		db.Username,
		":",
		db.Password,
		"@tcp(",
		db.Host,
		":3306",
		")/",
	}, "")
}
