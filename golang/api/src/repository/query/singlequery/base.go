// Package singlequery provides codes to do querys
package singlequery

import (
	"database/sql"
)

// Query a struct to control a query
type Query struct {
	code   string
	args   []any
	Result sql.Result
}

// New creates a new query empty
func New() *Query {
	return &Query{
		code:   "",
		args:   nil,
		Result: nil,
	}
}
