// Package singlequery provides codes to do querys
package singlequery

import reposhared "mod_name/repository/shared"

// NewWithCode creates a new query with provided code and args
func NewWithCode(code string, args ...any) *Query {
	query := New()

	query.code = code
	query.args = args
	return query
}

// NewFromSQLFile creates a new query based on sql file path and passed args
func NewFromSQLFile(filepath string, args ...any) *Query {
	query := New()

	query.code = reposhared.ReadSQLFile(filepath)
	query.args = args
	return query
}
