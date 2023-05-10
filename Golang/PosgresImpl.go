package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type PostgreSQL struct {
	host     string
	port     int
	username string
	password string
	dbname   string
}

func (p PostgreSQL) Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.host, p.port, p.username, p.password, p.dbname)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (p PostgreSQL) CreateUsers(user User) (string, error) {
	query := `INSERT INTO public.User (username, email) VALUES ($1, $2) RETURNING id`
	id := 0
	err = db.QueryRow(query, user.Username, user.Email).Scan(&id)
	if err != nil {
		return "", err
	}
	return "The User has been inserted successfully! Id: " + fmt.Sprint(id), nil
}

func (p PostgreSQL) GetAll() ([]User, error) {
	query := "Select * from public.user"
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

func (p PostgreSQL) DeleteUsers() (string, error) {
	query := "DELETE FROM public.user"
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

func (p PostgreSQL) DeleteUser(id string) (string, error) {
	query := "DELETE FROM public.User where id = $1"
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
