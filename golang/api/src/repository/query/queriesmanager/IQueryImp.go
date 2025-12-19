// Package queriesmanager provides codes to do querys
package queriesmanager

import (
	"database/sql"
	errorsapp "mod_name/error"
)

// Exec Implements the contract IQuery
func (manager *QueriesManager) Exec(tx *sql.Tx) {
	if manager == nil || tx == nil {
		panic(errorsapp.ErrNilPointer())
	}

	for query := range manager.queries.Iter() {
		query.Exec(tx)
	}
}
