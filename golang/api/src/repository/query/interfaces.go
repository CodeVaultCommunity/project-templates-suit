// Package query provides code for query manager
package query

import "database/sql"

// IQuery contract for queries
type IQuery interface {
	// Exec Execute the Query base on onpened transaction
	Exec(tx *sql.Tx)
}
