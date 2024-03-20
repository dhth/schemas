package ui

import "github.com/dhth/schemas/types"

type TablesFetchedMsg struct {
	tables []types.Table
	err    error
}

type TableChosenMsg struct {
	tableName string
}

type TablesDetailsFetchedMsg struct {
	tableName string
	columns   []types.ColumnDetails
	err       error
}
