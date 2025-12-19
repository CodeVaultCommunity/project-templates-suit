// Package queriesmanager provides codes to do querys
package queriesmanager

import (
	errorsapp "mod_name/error"
	"mod_name/repository/query/singlequery"
)

// NewFromManyFiles creates a new QueriesManager based on multiple sql files
func NewFromManyFiles(filespath []string, args ...[]any) *QueriesManager {
	if len(filespath) != len(args) {
		panic(errorsapp.ErrMismatchedArgs())
	}

	manager := New()
	for i, file := range filespath {
		manager.queries.AddBack(singlequery.NewFromSQLFile(file, args[i]...))
	}

	return manager
}
