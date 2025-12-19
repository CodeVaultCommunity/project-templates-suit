// Package singlequery provides codes to do querys
package singlequery

import (
	"database/sql"
	errorsapp "mod_name/error"
)

// Exec implements the IQuery contract
func (query *Query) Exec(tx *sql.Tx) {
	if query == nil || tx == nil {
		panic(errorsapp.ErrNilPointer())
	}

	res, err := tx.Exec(query.code, query.args...)
	if err != nil {
		panic(err)
	}

	query.Result = res
}
