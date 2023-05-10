package main

import (
	"database/sql"
	"fmt"
	"strconv"
)

var db *sql.DB
var err error

type Database interface {
	Connect() (*sql.DB, error)
	GetAll() ([]User, error)
	CreateUsers(User) (string, error)
	DeleteUsers() (string, error)
	DeleteUser(string) (string, error)
}

type DatabaseFactory struct{}

func (df DatabaseFactory) CreateDatabase() (Database, error) {

	datasource := getenv("DATA_SOURCE", "")

	switch datasource {
	case "mysql":
		port, _ := strconv.Atoi(getenv("DB_PORT", "3306"))
		return MySQL{
			host:     getenv("DB_HOST", "localhost"),
			port:     port,
			username: getenv("DB_USERNAME", "root"),
			password: getenv("DB_PASSWORD", "password"),
			dbname:   getenv("DB_NAME", "testdb"),
		}, nil
	case "postgres":
		port, _ := strconv.Atoi(getenv("DB_PORT", "3306"))
		return PostgreSQL{
			host:     getenv("DB_HOST", "localhost"),
			port:     port,
			username: getenv("DB_USERNAME", "root"),
			password: getenv("DB_PASSWORD", "password"),
			dbname:   getenv("DB_NAME", "testdb"),
		}, nil
	default:
		return nil, fmt.Errorf("unsupported database type")
	}
}
