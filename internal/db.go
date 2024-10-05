package db

import (
	"os"
	"path/filepath"

	"github.com/lewissteele/dbat/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var localDB *gorm.DB

func SaveConnection(host string, username string, password string) {
	localDB.Create(&model.Database{Host: "localhost", Username: "root", Password: ""})
}

func init() {
	configHome := os.Getenv("XDG_CONFIG_HOME")

	if len(configHome) == 0 {
		configHome = filepath.Join(os.Getenv("HOME"), ".config")
	}

	path := filepath.Join(configHome, "dbat/dbat.db")
	conn, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("could not connect to sqlite db")
	}

	conn.AutoMigrate(&model.Database{})

	localDB = conn
}
