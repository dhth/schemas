package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDBPool() *pgxpool.Pool {
	databaseUrl, err := getDBConnectionString()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building database connection string: %s\n", err)
		os.Exit(1)
	}

	dbpool, err := pgxpool.New(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return dbpool
}
