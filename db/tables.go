package db

import (
	"context"

	"github.com/dhth/schemas/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetTables(dbpool *pgxpool.Pool) ([]types.Table, error) {
	selectQuery := `SELECT t.table_name, count(c.column_name) as num_cols
FROM information_schema.tables t
         LEFT JOIN information_schema.columns c ON t.table_name = c.table_name
WHERE t.table_schema = 'public'
GROUP BY t.table_name;
`
	rowsResult, err := dbpool.Query(context.Background(), selectQuery)

	if err != nil {
		return nil, err
	}
	defer rowsResult.Close()

	result, err := pgx.CollectRows(rowsResult, pgx.RowToStructByPos[types.Table])
	if err != nil {
		return nil, err
	}
	return result, nil
}
