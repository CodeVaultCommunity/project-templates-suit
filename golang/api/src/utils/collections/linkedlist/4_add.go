// Package linkedlist provides linkedlist implementation
package linkedlist

import (
	errorsapp "mod_name/error"
	"mod_name/utils/collections/linkednode"
)

// AddFront adds a element as the first element on list.
// If list is nil, it will enter in panic(errorsapp.ErrNilPointer())
func (list *LinkedList[T]) AddFront(value T) {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	node := linkednode.New(value)

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node.SetNext(list.head)
		list.head = node
	}

	list.size++
}

// AddBack adds a element as the last element on list.
// If list is nil, it will enter in panic(errorsapp.ErrNilPointer())
func (list *LinkedList[T]) AddBack(value T) {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	node := linkednode.New(value)

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.SetNext(node)
		list.tail = node
	}

	list.size++
}

func (list *LinkedList[T]) link(node *linkednode.SimpleLinkedNode[T], value T) {
	newNode := linkednode.New(value)
	newNode.SetNext(node.GetNext())
	node.SetNext(newNode)
	if newNode.GetNext() == nil {
		list.tail = newNode
	}
	list.size++
}

/*
AddAt adds a element at the passed index.

If list is nil, it will enter in panic(errorsapp.ErrNilPointer()).

If the passed index is invalid, it will enter in panic(errorsapp.ErrIndexOutOfBound()).

Example:

	list := NewLinkedList[int]()
	list.AddAt(0, 1) // expected: list.Get(0) == 1
	list.AddAt(1, 2) // expected: list.Get(0) == 1; list.Get(1) == 2
	list.AddAt(1, 1) // expected: list.Get(0) == 1; list.Get(1) == 1; list.Get(2) == 2
	list.AddAt(3, 4) // expected: ErrIndexOutOfBound
*/
func (list *LinkedList[T]) AddAt(index uint, value T) {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	if index == 0 {
		list.AddFront(value)
	} else {
		list.link(list.get(index-1), value)
	}
}

/*
AddAtRev adds a element at the passed index reversed. The index 0 represents the last element on list.

If list is nil, it will enter in panic(errorsapp.ErrNilPointer()).

If the passed index is invalid, it will enter in panic(errorsapp.ErrIndexOutOfBound()).

Example:

	list := NewLinkedList[int]()
	list.AddAtRev(0, 1) // expected: list.Get(0) == 1
	list.AddAtRev(1, 3) // expected: list.Get(0) == 3; list.Get(1) == 1
	list.AddAtRev(1, 2) // expected: list.Get(0) == 3; list.Get(1) == 2; list.Get(2) == 1
	list.AddAtRev(3, 3) // expected: ErrIndexOutOfBound
*/
func (list *LinkedList[T]) AddAtRev(index uint, value T) {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	if index == 0 {
		list.AddBack(value)
	} else {
		if index == list.size {
			list.AddFront(value)
		} else {
			list.link(list.getRev(index), value)
		}
	}
}
