// Package queriesmanager provides codes to do querys
package queriesmanager

import (
	errorsapp "mod_name/error"
	"mod_name/repository/query"
)

func (manager *QueriesManager) AddQuery(code string, args ...any) {
	if manager == nil {
		panic(errorsapp.ErrNilPointer())
	}

	manager.queries.AddBack(query.NewQueryWithCode(code, args...))
}

func (manager *QueriesManager) AddCodeFromSQLFile(filepath string, args ...any) {
	if manager == nil {
		panic(errorsapp.ErrNilPointer())
	}
	manager.queries.AddBack(query.NewQueryFromSQLFile(filepath, args...))
}

func (manager *QueriesManager) AddCodeFromSQLFiles(filespath []string, args [][]any) {
	if manager == nil {
		panic(errorsapp.ErrNilPointer())
	}

	for index, filepath := range filespath {
		manager.queries.AddBack(query.NewQueryFromSQLFile(filepath, args[index]...))
	}
}

func (manager *QueriesManager) GetQueries() *Queries {
	if manager == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return manager.queries
}
