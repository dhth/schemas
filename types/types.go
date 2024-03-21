package types

import "fmt"

type Table struct {
	Name       string `db:"table_name"`
	NumColumns int    `db:"num_cols"`
}

type ColumnDetails struct {
	Name       string `db:"column_name"`
	DataType   string `db:"data_type"`
	IsNullable string `db:"is_nullable"`
}

type TableConstraint struct {
	Name        string  `db:"constraint_name"`
	Type        string  `db:"constraint_type"`
	CheckClause *string `db:"check_clause"`
}

func (table Table) Title() string {
	return fmt.Sprintf("%s", table.Name)
}

func (table Table) Description() string {
	return fmt.Sprintf("%d columns", table.NumColumns)
}

func (table Table) FilterValue() string {
	return table.Name
}
