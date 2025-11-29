package transaction

import (
	"database/sql"
	"errors"
	errorsapp "mod_name/error"
	"mod_name/repository"
	"net/http"
)

func (transaction *Transaction) OpenTransaction(db *sql.DB) {
	if transaction == nil {
		panic(errorsapp.ErrNilPointer())
	}

	if transaction.tx == nil {
		panic(errorsapp.New(http.StatusInternalServerError, errorsapp.DBConnectionFail, "can't connect to db", errors.New("try open a new transaction without close the last")))
	}

	tx, err := db.Begin()
	if err != nil {
		panic(errorsapp.New(http.StatusInternalServerError, errorsapp.DBConnectionFail, "can't start transaction", err))
	}

	transaction.tx = tx
}

func (transaction *Transaction) Rollback() {
	if transaction == nil {
		panic(errorsapp.ErrNilPointer())
	}

	if p := recover(); p != nil {
		transaction.tx.Rollback()
		transaction.tx = nil
		panic(p)
	}
}

func (transaction *Transaction) Commit() {
	if transaction == nil {
		panic(errorsapp.ErrNilPointer())
	}

	err := transaction.tx.Commit()

	if err != nil {
		panic(err)
	}

	transaction.tx = nil
}

func (tx *Transaction) Atomic(bd *sql.DB, query repository.IQuery) {
	tx.OpenTransaction(bd)
	defer tx.Rollback()
	query.Exec(tx.tx)
	tx.Commit()
}
