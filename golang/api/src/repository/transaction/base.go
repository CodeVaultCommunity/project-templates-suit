package transaction

import (
	"database/sql"
)

type Transaction struct {
	tx *sql.Tx
}

func New() *Transaction {
	return &Transaction{
		tx: nil,
	}
}
