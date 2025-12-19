// Package linkedlist provides linkedlist implementation
package linkedlist

import errorsapp "mod_name/error"

// IsEmpty checks if the list is empty.
func (list *LinkedList[T]) IsEmpty() bool {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return list.head == nil
}
