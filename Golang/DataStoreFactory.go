package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type Database interface {
	Connect() (*sql.DB, error)
	GetAll() ([]User, error)
	CreateUsers(User) (string, error)
	DeleteUsers() (string, error)
	DeleteUser(string) (string, error)
}

type DatabaseFactory struct{}

func (df DatabaseFactory) CreateDatabase() (Database, error) {

	datasource := getenv("datasource", "")
	if datasource == "" {
		return nil,
			errors.New("The datasource not defined. Availaible options are mysql, postgresql")
	}

	switch datasource {
	case "mysql":
		port, _ := strconv.Atoi(getenv("PORT", "3306"))
		return MySQL{
			host:     getenv("HOST", "localhost"),
			port:     port,
			username: getenv("USERNAME", "root"),
			password: getenv("PASSWORD", "password"),
			dbname:   getenv("DB_NAME", "testdb"),
		}, nil
	case "postgres":
		port, _ := strconv.Atoi(getenv("PORT", "3306"))
		return PostgreSQL{
			host:     getenv("HOST", "localhost"),
			port:     port,
			username: getenv("USERNAME", "root"),
			password: getenv("PASSWORD", "password"),
			dbname:   getenv("DB_NAME", "testdb"),
		}, nil
	default:
		return nil, fmt.Errorf("Unsupported database type")
	}
}
