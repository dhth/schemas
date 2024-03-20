package db

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type DBConfigError string

func (e DBConfigError) Error() string {
	return string(e)
}

func getDBConnectionString() (string, error) {
	address := os.Getenv("DATABASE_ADDRESS")
	var errorStr string
	if address == "" {
		errorStr += "\n\tEnvironment variable DATABASE_ADDRESS not set"
	}
	portStr := os.Getenv("DATABASE_PORT")
	if portStr == "" {
		errorStr += "\n\tEnvironment variable DATABASE_PORT not set"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		errorStr += "\n\tDATABASE_PORT value is not a valid number"
	}
	username := os.Getenv("DATABASE_USERNAME")
	if username == "" {
		errorStr += "\n\tEnvironment variable DATABASE_USERNAME not set"
	}
	password := os.Getenv("DATABASE_PASSWORD")
	if password == "" {
		errorStr += "\n\tEnvironment variable DATABASE_PASSWORD not set"
	}
	passwordEncoded := url.QueryEscape(password)

	dbname := os.Getenv("DATABASE_DBNAME")
	if dbname == "" {
		errorStr += "\n\tEnvironment variable DATABASE_DBNAME not set"
	}
	if errorStr != "" {
		return "", DBConfigError(errorStr)
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", username, passwordEncoded, address, port, dbname), nil

}
