// Package queriesmanager provides codes to do querys
package queriesmanager

import (
	errorsapp "mod_name/error"
	"mod_name/repository/query/singlequery"
)

// AddQuery Adds a query based on code and arguments
func (manager *QueriesManager) AddQuery(code string, args ...any) {
	if manager == nil {
		panic(errorsapp.ErrNilPointer())
	}

	manager.queries.AddBack(singlequery.NewWithCode(code, args...))
}

// AddCodeFromSQLFile adds a query based on sql filepath
func (manager *QueriesManager) AddCodeFromSQLFile(filepath string, args ...any) {
	if manager == nil {
		panic(errorsapp.ErrNilPointer())
	}
	manager.queries.AddBack(singlequery.NewFromSQLFile(filepath, args...))
}

// AddCodeFromSQLFiles add multiple queries based on multiple filespath
func (manager *QueriesManager) AddCodeFromSQLFiles(filespath []string, args ...[]any) {
	if manager == nil {
		panic(errorsapp.ErrNilPointer())
	}

	if len(filespath) != len(args) {
		panic(errorsapp.ErrMismatchedArgs())
	}

	for index, filepath := range filespath {
		manager.queries.AddBack(singlequery.NewFromSQLFile(filepath, args[index]...))
	}
}

// GetQueries returns the queries list
// ATTENTION: you is dealt with pointer here, it could be unsafe
func (manager *QueriesManager) GetQueries() *Queries {
	if manager == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return manager.queries
}
