// Package queriesmanager provides codes to do querys
package queriesmanager

import "mod_name/repository/query"

func NewQueryFromMultipleSQLFile(filespath []string) *QueriesManager {
	manager := NewQueriesManager()
	for _, file := range filespath {
		manager.queries.AddBack(query.NewQueryFromSQLFile(file))
	}

	return manager
}
