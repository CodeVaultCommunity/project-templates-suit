// Package queriesmanager provides codes to do querys
package queriesmanager

import "database/sql"

func (manager *QueriesManager) Exec(tx *sql.Tx) {
	for query := range manager.queries.IteratorLinkedList() {
		query.Exec(tx)
	}
}
