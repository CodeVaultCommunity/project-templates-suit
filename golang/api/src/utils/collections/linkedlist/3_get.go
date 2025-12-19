// Package linkedlist provides linkedlist implementation
package linkedlist

import (
	errorsapp "mod_name/error"
	"mod_name/utils/collections/linkednode"
)

// get private function to get a element at given index.
// If the passed index is invalid, it will enter in panic(errorsapp.ErrIndexOutOfBound())
func (list *LinkedList[T]) get(index uint) *linkednode.SimpleLinkedNode[T] {
	if index > (list.size - 1) {
		panic(errorsapp.ErrIndexOutOfBound())
	}

	if index == 0 {
		return list.head
	}

	if index == (list.size - 1) {
		return list.tail
	}

	node := list.head

	for range index {
		node = node.GetNext()
	}

	return node
}

// getRev private function to get a element at given index reversed.
// The index 0 represents the last element at list.
// If the passed index is invalid, it will enter in panic(errorsapp.ErrIndexOutOfBound())
func (list *LinkedList[T]) getRev(index uint) *linkednode.SimpleLinkedNode[T] {
	if index > list.size {
		panic(errorsapp.ErrIndexOutOfBound())
	}

	if index == 0 {
		return list.tail
	}

	if index == list.size {
		return list.head
	}

	node := list.head

	realIndex := (list.size - index - 1)
	for range realIndex {
		node = node.GetNext()
	}

	return node
}

/*
Get get a element at the given index.

If list is nil, it will enter in panic(errorsapp.ErrNilPointer()).

If the passed index is invalid, it will enter in panic(errorsapp.ErrIndexOutOfBound())
*/
func (list *LinkedList[T]) Get(index uint) T {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return list.get(index).GetValue()
}

/*
GetRev get a element at the given index reversed way.

The index 0 represents the last element on list.

If list is nil, it will enter in panic(errorsapp.ErrNilPointer()).

If the passed index is invalid, it will enter in panic(errorsapp.ErrIndexOutOfBound())
*/
func (list *LinkedList[T]) GetRev(index uint) T {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return list.getRev(index).GetValue()
}

// GetSize returns te current list size.
// If list is nil, it will enter in panic(errorsapp.ErrNilPointer()).
func (list *LinkedList[T]) GetSize() uint {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return list.size
}
