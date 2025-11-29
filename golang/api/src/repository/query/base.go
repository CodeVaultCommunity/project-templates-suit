// Package query provides codes to do querys
package query

import (
	"database/sql"
)

type Query struct {
	code   string
	args   []any
	Result *sql.Result
}

func NewQuery() *Query {
	return &Query{
		code:   "",
		args:   nil,
		Result: nil,
	}
}
