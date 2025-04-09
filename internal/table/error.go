package table

import "errors"

var (
	NoTables = errors.New("there are no tables")
	NotFound = errors.New("table not found")
)
