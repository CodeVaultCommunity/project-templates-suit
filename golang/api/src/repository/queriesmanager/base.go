// Package queriesmanager provides codes to do querys
package queriesmanager

import (
	"mod_name/collections"
	"mod_name/repository/query"
)

type Queries = collections.LinkedList[query.Query]

type QueriesManager struct {
	queries *Queries
}

func NewQueriesManager() *QueriesManager {
	return &QueriesManager{
		queries: collections.NewLinkedList[query.Query](),
	}
}
