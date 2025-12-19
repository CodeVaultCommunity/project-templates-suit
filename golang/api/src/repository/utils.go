// Package repository - includes utilities functions for fast develop
package repository

import (
	"context"
	"database/sql"
	"mod_name/repository/query/queriesmanager"
	"mod_name/repository/query/singlequery"
	reposhared "mod_name/repository/shared"
	"mod_name/repository/transaction"
	"mod_name/utils/collections/linkedlist"
)

func FastQueryRow(ctx context.Context, sqlfilepath string, args ...any) *sql.Row {
	q := reposhared.ReadSQLFile(sqlfilepath)
	return SQLDB.QueryRowContext(ctx, q, args...)
}

// FastAtomicTransaction execute a query in sqlfile
// If some error occurred it will Rollback and propagate panic
func FastAtomicTransaction(sqlfile string, args ...any) sql.Result {
	tx := transaction.New()
	query := singlequery.NewFromSQLFile(sqlfile, args...)

	tx.Atomic(SQLDB, query)

	return query.Result
}

// FastAtomicTransactions execute multiple queries from sqlfile
// If some error occurred it will Rollback and propagate panic
func FastAtomicTransactions(sqlfiles []string, args ...[]any) *linkedlist.LinkedList[sql.Result] {
	tx := transaction.New()
	query := queriesmanager.NewFromManyFiles(sqlfiles, args...)

	tx.Atomic(SQLDB, query)

	return linkedlist.Map(query.GetQueries(), func(q *singlequery.Query) sql.Result { return q.Result })
}
