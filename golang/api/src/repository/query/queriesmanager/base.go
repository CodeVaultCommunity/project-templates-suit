// Package queriesmanager provides codes to do querys
package queriesmanager

import (
	"mod_name/repository/query/singlequery"
	"mod_name/utils/collections/linkedlist"
)

// Queries type alayas for multiple queries
type Queries = linkedlist.LinkedList[*singlequery.Query]

// QueriesManager struct to dealt with mutiple queries
type QueriesManager struct {
	queries *Queries
}

// New creates a new QueriesManager
func New() *QueriesManager {
	return &QueriesManager{
		queries: linkedlist.New[*singlequery.Query](),
	}
}
