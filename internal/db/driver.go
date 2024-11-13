package db

type Driver string

const (
	MariaDB    Driver = "mariadb"
	MySQL      Driver = "mysql"
	PostgreSQL Driver = "postgresql"
	SQLite     Driver = "sqlite"
)
