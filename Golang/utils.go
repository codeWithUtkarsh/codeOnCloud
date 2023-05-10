package main

import (
	"errors"
	"os"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func validate() (bool, error) {

	datasource_name := getenv("DATA_SOURCE", "")
	if datasource_name == "" {
		panic(errors.New("datasource not defined, availaible options are mysql, postgres"))
	}
	return true, nil
}
