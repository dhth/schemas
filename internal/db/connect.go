package db

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

const (
	envVarDatabaseAddress  = "DATABASE_ADDRESS"
	envVarDatabasePort     = "DATABASE_PORT"
	envVarDatabaseUsername = "DATABASE_USERNAME"
	envVarDatabasePassword = "DATABASE_PASSWORD"
	envVarDatabaseDbName   = "DATABASE_DBNAME"
)

func getDBConnectionString() (string, []string) {
	address := os.Getenv(envVarDatabaseAddress)
	var errors []string
	if address == "" {
		errors = append(errors, fmt.Sprintf("Environment variable %s is not set", envVarDatabaseAddress))
	}

	portStr := os.Getenv(envVarDatabasePort)
	if portStr == "" {
		errors = append(errors, fmt.Sprintf("Environment variable %s is not set", envVarDatabasePort))
	}

	port, err := strconv.Atoi(portStr)
	if portStr != "" && err != nil {
		errors = append(errors, fmt.Sprintf("%s is not a valid number", envVarDatabasePort))
	}

	username := os.Getenv(envVarDatabaseUsername)
	if username == "" {
		errors = append(errors, fmt.Sprintf("Environment variable %s is not set", envVarDatabaseUsername))
	}

	password := os.Getenv(envVarDatabasePassword)
	if password == "" {
		errors = append(errors, fmt.Sprintf("Environment variable %s is not set", envVarDatabasePassword))
	}
	passwordEncoded := url.QueryEscape(password)

	dbname := os.Getenv(envVarDatabaseDbName)
	if dbname == "" {
		errors = append(errors, fmt.Sprintf("Environment variable %s is not set", envVarDatabaseDbName))
	}

	if len(errors) > 0 {
		return "", errors
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", username, passwordEncoded, address, port, dbname), nil
}
