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

func (table Table) Title() string {
	return fmt.Sprintf("%s", table.Name)
}

func (table Table) Description() string {
	return fmt.Sprintf("%d columns", table.NumColumns)
}

func (table Table) FilterValue() string {
	return table.Name
}
