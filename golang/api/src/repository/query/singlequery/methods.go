// Package singlequery provides codes to do querys
package singlequery

import (
	errorsapp "mod_name/error"
	reposhared "mod_name/repository/shared"
)

// SetCode assing the passed code for query
func (query *Query) SetCode(code string) {
	if query == nil {
		panic(errorsapp.ErrNilPointer())
	}

	query.code = code
}

// SetArgs assing the passed args to query
func (query *Query) SetArgs(args ...any) {
	if query == nil {
		panic(errorsapp.ErrNilPointer())
	}

	query.args = args
}

// SetCodeFromSQLFile set the query code equals of code in sql file passed
func (query *Query) SetCodeFromSQLFile(filepath string) {
	if query == nil {
		panic(errorsapp.ErrNilPointer())
	}

	query.code = reposhared.ReadSQLFile(filepath)
}
