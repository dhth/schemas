package db

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	errCouldntCreateDbPool         = errors.New("couldn't create db pool")
	errCouldntConnectToDB          = errors.New("couldn't connect to database")
	errCouldntBuildDbConnectionStr = errors.New("couldn't build database connection string")
)

func CreateDBPool() (*pgxpool.Pool, error) {
	databaseURL, errorStrs := getDBConnectionString()
	if len(errorStrs) > 0 {
		return nil, fmt.Errorf("%w\n- %s", errCouldntBuildDbConnectionStr, strings.Join(errorStrs, "\n- "))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errCouldntCreateDbPool, err.Error())
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errCouldntConnectToDB, err.Error())
	}

	return dbpool, nil
}
