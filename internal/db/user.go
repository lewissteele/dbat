package db

import "github.com/lewissteele/dbat/internal/model"

func UserDB(name string) *model.Database {
	var userDB model.Database
	LocalDB.Where("name = ?", name).Find(&userDB)
	return &userDB
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

func Port(driver model.Driver) string {
	if driver == model.PostgreSQL {
		return "5432"
	}
	return "3306"
}
