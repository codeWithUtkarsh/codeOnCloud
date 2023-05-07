package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/go-swagno/swagno"
	"github.com/go-swagno/swagno-http/swagger"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var datasource Database
var users []User

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	datasource_name := getenv("datasource", "")
	if datasource_name == "" {
		panic(errors.New("The datasource not defined. Availaible options are mysql, postgres"))
	}

	response, err := datasource.GetAll()
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	var response string
	if user.Username == "" || user.Email == "" {
		response = "You are missing username or email parameter."
	} else {

		fmt.Println("Inserting new User with username: " + user.Username + " and email: " + user.Email)

		datasource_name := getenv("datasource", "")
		if datasource_name == "" {
			panic(errors.New("The datasource not defined. Availaible options are mysql, postgres"))
		}

		response, err = datasource.CreateUsers(User{Username: user.Username, Email: user.Email})
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	var response string
	if id == "" {
		response = "You are missing id parameter."
	} else {

		datasource_name := getenv("datasource", "")
		if datasource_name == "" {
			panic(errors.New("The datasource not defined. Availaible options are mysql, postgres"))
		}

		response, err = datasource.DeleteUser(id)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {

	var response string
	datasource_name := getenv("datasource", "")
	if datasource_name == "" {
		panic(errors.New("The datasource not defined. Availaible options are mysql, postgres"))
	}

	response, err = datasource.DeleteUsers()
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(response)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/users", GetUsers).Methods("GET")
	myRouter.HandleFunc("/users", CreateUser).Methods("POST")
	myRouter.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users", DeleteUsers).Methods("DELETE")

	SetSwagger(myRouter)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	os.Setenv("datasource", "mysql")
	os.Setenv("HOST", "127.0.0.1")
	os.Setenv("PORT", "3306")
	os.Setenv("USERNAME", "user")
	os.Setenv("PASSWORD", "password")
	os.Setenv("DB_NAME", "db")

	// os.Setenv("datasource", "postgres")
	// os.Setenv("HOST", "127.0.0.1")
	// os.Setenv("PORT", "5438")
	// os.Setenv("USERNAME", "postgres")
	// os.Setenv("PASSWORD", "postgres")
	// os.Setenv("DB_NAME", "db")

	dbFactory := DatabaseFactory{}
	datasource, err = dbFactory.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}

	db, err = datasource.Connect()
	defer db.Close()

	handleRequests()
}

func SetSwagger(myRouter *mux.Router) {
	endpoints := []Endpoint{
		EndPoint(GET, "/users", "Users", Params(), nil, []User{}, err, "Get all users", nil),
		EndPoint(POST, "/users", "Users", Params(), User{}, nil, err, "", nil),
		EndPoint(DELETE, "/users/{id}", "Users", Params(IntParam("id", true, "")), nil, nil, err, "", nil),
		EndPoint(DELETE, "/users", "Users", Params(), nil, nil, err, "", nil),
	}

	sw := CreateNewSwagger("Swagger API", "1.0")
	AddEndpoints(endpoints)

	myRouter.PathPrefix("/swagger").Handler(swagger.SwaggerHandler(sw.GenerateDocs()))
	// if you want to export your swagger definition to a file
	// sw.ExportSwaggerDocs("docs/swagger.json") // optional
}
