package db

import (
	"context"

	"github.com/dhth/schemas/internal/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetTableConstraints(dbpool *pgxpool.Pool, tableName string) ([]types.TableConstraint, error) {
	selectQuery := `
SELECT c.constraint_name,
       c.constraint_type,
       cc.check_clause
FROM information_schema.table_constraints c
         LEFT JOIN information_schema.check_constraints cc
                   ON c.constraint_name = cc.constraint_name
WHERE table_schema = 'public'
  AND table_name = $1
ORDER BY c.constraint_type;
`
	rowsResult, err := dbpool.Query(context.Background(), selectQuery, tableName)
	if err != nil {
		return nil, err
	}
	defer rowsResult.Close()

	result, err := pgx.CollectRows(rowsResult, pgx.RowToStructByPos[types.TableConstraint])
	if err != nil {
		return nil, err
	}
	return result, nil
}
