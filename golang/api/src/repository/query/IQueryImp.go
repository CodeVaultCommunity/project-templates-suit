// Package query provides codes to do querys
package query

import "database/sql"

func (query *Query) Exec(tx *sql.Tx) {
	_, err := tx.Exec(query.code, query.args...)
	if err != nil {
		panic(err)
	}
}
