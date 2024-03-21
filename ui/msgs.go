package ui

import "github.com/dhth/schemas/types"

type HideHelpMsg struct{}

type TablesFetchedMsg struct {
	tables []types.Table
	err    error
}

type TableChosenMsg struct {
	tableName string
}

type TableDetailsFetchedMsg struct {
	tableName string
	columns   []types.ColumnDetails
	err       error
}

type TableConstraintsFetchedMsg struct {
	tableName   string
	constraints []types.TableConstraint
	err         error
}
