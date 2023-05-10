package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	host     string
	port     int
	username string
	password string
	dbname   string
}

func (m MySQL) Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.username, m.password, m.host, m.port, m.dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (p MySQL) CreateUsers(user User) (string, error) {
	query := "INSERT INTO `User` (`username`, `email`) VALUES (?, ?)"
	insertResult, err := db.ExecContext(context.Background(), query, user.Username, user.Email)
	if err != nil {
		panic(err.Error())
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		return "", err
	}
	return "The User has been inserted successfully! Id: " + fmt.Sprint(id), nil
}

func (p MySQL) GetAll() ([]User, error) {
	query := "Select * from User"
	result, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var users []User
	for result.Next() {
		var id int
		var username string
		var email string

		err = result.Scan(&id, &username, &email)
		if err != nil {
			return nil, err
		}

		users = append(users, User{Username: username, Email: email})
	}
	return users, nil
}

func (p MySQL) DeleteUsers() (string, error) {
	query := "DELETE FROM User"
	result, err := db.Exec(query)
	if err != nil {
		panic(err.Error())
	}

	row, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	return "All Users have been deleted successfully!  Row Deleted: " + fmt.Sprint(row), nil
}

func (p MySQL) DeleteUser(id string) (string, error) {
	query := "DELETE FROM User where id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		panic(err.Error())
	}

	row, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	return "All Users have been deleted successfully!  Row Deleted: " + fmt.Sprint(row), nil
}
