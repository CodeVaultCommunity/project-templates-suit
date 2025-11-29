// Package query provides codes to do querys
package query

import "mod_name/repository/shared"

func NewQueryWithCode(code string, args ...any) *Query {
	query := NewQuery()

	query.code = code
	query.args = args
	return query
}

func NewQueryFromSQLFile(filepath string, args ...any) *Query {
	query := NewQuery()

	query.code = shared.ReadSQLFile(filepath)
	query.args = args
	return query
}
