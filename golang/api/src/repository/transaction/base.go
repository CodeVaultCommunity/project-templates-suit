// Package transaction provides code to handle database transactions
package transaction

import (
	"database/sql"
)

// Transaction struct to manager transaction
type Transaction struct {
	tx *sql.Tx
}

// New creates a new empty transaction
func New() *Transaction {
	return &Transaction{
		tx: nil,
	}
}
