package repository

import "database/sql"

type IQuery interface {
	Exec(tx *sql.Tx)
}
