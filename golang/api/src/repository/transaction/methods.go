// Package transaction provides code to handle database transactions
package transaction

import (
	"database/sql"
	"errors"
	errorsapp "mod_name/error"
	"mod_name/repository/query"
	"net/http"
)

// OpenTransaction starts transaction with passed db
func (transaction *Transaction) OpenTransaction(db *sql.DB) {
	if transaction == nil {
		panic(errorsapp.ErrNilPointer())
	}

	if transaction.tx != nil {
		panic(errorsapp.New(http.StatusInternalServerError, errorsapp.DBConnectionFail, "can't connect to db", errors.New("try open a new transaction without close the last")))
	}

	tx, err := db.Begin()
	if err != nil {
		panic(errorsapp.New(http.StatusInternalServerError, errorsapp.DBConnectionFail, "can't start transaction", err))
	}

	transaction.tx = tx
}

// Rollback prepare a roolback for transaction
//
// Example:
//
//	defer tx.Rollback
func (transaction *Transaction) Rollback() {
	if transaction == nil {
		panic(errorsapp.ErrNilPointer())
	}

	if p := recover(); p != nil {
		if transaction.tx != nil {
			err := transaction.tx.Rollback()
			transaction.tx = nil
			if err != nil {
				panic(err)
			}
		}
		panic(p)
	}
}

// Commit saves the transaction in database
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

// Atomic execute query as atomic operation an, if an error occurred, rollback the operation
func (transaction *Transaction) Atomic(db *sql.DB, query query.IQuery) {
	if query == nil {
		panic(errorsapp.ErrNilPointer())
	}
	transaction.OpenTransaction(db)
	defer transaction.Rollback()
	query.Exec(transaction.tx)
	transaction.Commit()
}
