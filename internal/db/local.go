package db

import (
	"os"
	"path/filepath"

	"github.com/lewissteele/dbat/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var LocalDB *gorm.DB

func History() []string {
	var histories []model.History
	var queries []string

	LocalDB.Order("updated_at").Where("database_id = ?", UserDB.ID).Find(&histories)

	for _, history := range histories {
		queries = append(queries, history.Query)
	}

	return queries
}

func SaveHistory(query string) {
	history := model.History{
		Database: UserDB,
		Query:    query,
	}

	LocalDB.Create(&history)
}

func init() {
	config := os.Getenv("XDG_CONFIG_HOME")

	if len(config) == 0 {
		config = filepath.Join(os.Getenv("HOME"), ".config")
	}

	err := os.MkdirAll(
		filepath.Join(config, "dbat"),
		0700,
	)

	if err != nil {
		panic("cannot create config dir")
	}

	db := filepath.Join(config, "dbat/dbat.db")

	LocalDB, err = gorm.Open(sqlite.Open(db), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("cannot connect to dbat.db")
	}

	LocalDB.AutoMigrate(
		&model.Database{},
		&model.History{},
	)

	dbat := model.Database{
		Driver: "sqlite",
		Name:   "dbat",
		Path:   db,
	}

	LocalDB.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbat)
}
