package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

/**
CREATE TABLE User (
    id SERIAL,
    username varchar(50) NOT NULL,
    email varchar(50) NOT NULL,
    PRIMARY KEY (id)
)

INSERT INTO User (
    username,
    email
)
VALUES
    ('Utkarsh', 'utkarshkviim@gmail.com'),
    ('Ravi', 'ravi009988@gmail.com'),
    ('Kavi', 'Kavi@gmail.com');

*/

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dbFactory := DatabaseFactory{}
	datasource, err = dbFactory.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}

	db, err = datasource.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handleRequests()
}
