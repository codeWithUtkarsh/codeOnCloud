package main

import (
	_ "github.com/go-sql-driver/mysql"
)

// func main() {
// 	dbFactory := DatabaseFactory{}

// 	mysqlDB, err := dbFactory.CreateDatabase("mysql")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	mysqlDB.Connect()

// 	defer mysqlDB.Close()

// 	postgresDB, err := dbFactory.CreateDatabase("postgresql")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	postgresDB.Connect()

// 	defer postgresDB.Close()

// 	fmt.Println("Successfully connected to databases!")
// }
