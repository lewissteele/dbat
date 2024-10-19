package db

import (
	"os"
	"path/filepath"

	"github.com/lewissteele/dbat/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var LocalDB *gorm.DB

func init() {
	configHome := os.Getenv("XDG_CONFIG_HOME")

	if len(configHome) == 0 {
		configHome = filepath.Join(os.Getenv("HOME"), ".config")
	}

	var err error
	path := filepath.Join(configHome, "dbat/dbat.db")

	LocalDB, err = gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("could not connect to sqlite db")
	}

	LocalDB.AutoMigrate(
		&model.Database{},
		&model.History{},
	)
}
