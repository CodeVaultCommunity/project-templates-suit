// Package query provides codes to do querys
package query

import (
	errorsapp "mod_name/error"
	"mod_name/repository/shared"
)

func (query *Query) SetCode(code string) {
	if query == nil {
		panic(errorsapp.ErrNilPointer())
	}

	query.code = code
}

func (query *Query) SetArgs(args ...any) {
	query.args = args
}

func (query *Query) SetCodeFromSQLFile(filepath string) {
	if query == nil {
		panic(errorsapp.ErrNilPointer())
	}

	query.code = shared.ReadSQLFile(filepath)
}
