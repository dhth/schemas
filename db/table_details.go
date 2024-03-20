package db

import (
	"context"
	"github.com/dhth/schemas/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetColumnDetails(dbpool *pgxpool.Pool, tableName string) ([]types.ColumnDetails, error) {
	selectQuery := `SELECT column_name, data_type, CAST(is_nullable as TEXT)
FROM information_schema.columns
WHERE table_schema = 'public'
AND table_name = $1
ORDER BY column_name;
`
	rowsResult, err := dbpool.Query(context.Background(), selectQuery, tableName)

	if err != nil {
		return nil, err
	}
	defer rowsResult.Close()

	result, err := pgx.CollectRows(rowsResult, pgx.RowToStructByPos[types.ColumnDetails])
	if err != nil {
		return nil, err
	}
	return result, nil
}
