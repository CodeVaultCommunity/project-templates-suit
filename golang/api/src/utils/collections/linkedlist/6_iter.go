// Package linkedlist provides linkedlist implementation
package linkedlist

import (
	errorsapp "mod_name/error"
	"mod_name/utils/collections/linkednode"
)

// Iter Prepaer a iterator to Iter over each element of list.
// If the elemnts of list contains some cycle, it will enter in ifinity loop.
func (list *LinkedList[T]) Iter() <-chan T {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return list.head.Iterator()
}

/*
Map function iterate over each list aaplying the function f and returns a new LinkedList with the returns object from f.

Example:

	var list *LinkedList[int] // is a LinkedList with three elements: 1, 2 and 3.
	func duplicate(x int) int // is a function that returns x*2.

	Map(list, duplicate) // in this case, the map function will return a new LinkedList with three elements: 2, 4 and 6.
*/
func Map[T, R any](list *LinkedList[T], f func(T) R) *LinkedList[R] {
	if list == nil || f == nil {
		panic(errorsapp.ErrNilPointer())
	}

	newChain := New[R]()

	for v := range list.Iter() {
		newChain.AddBack(f(v))
	}

	return newChain
}

func (list *LinkedList[T]) cleanRear(maintain func(T) bool) *linkednode.SimpleLinkedNode[T] {
	var node *linkednode.SimpleLinkedNode[T]

	for node = list.head; node != nil; node = list.head {
		if maintain(node.GetValue()) {
			return node
		}

		list.PopFirst()
	}

	return nil
}

/*
Filter gets each element on list and if it not pass on function mantain() it is removed from it.

Example:

	var list *LinkedList[int] // is a linked list with three elements: 1, 2 and 3.
	func isOdd(x int) bool // checks if x is odd number.

	list.Filter(isOdd) // in this case, list will be a LinkedList with elements: 1 and 3.

ATTENTION:

	if list is nil or mantain is nil, it will enter in panic(errorsapp.ErrNilPointer())
*/
func (list *LinkedList[T]) Filter(maintain func(T) bool) {
	if list == nil || maintain == nil {
		panic(errorsapp.ErrNilPointer())
	}

	lastMantained := list.cleanRear(maintain)
	if lastMantained == nil {
		return // early stop
	}

	for actual := lastMantained.GetNext(); actual != nil; actual = actual.GetNext() {
		if maintain(actual.GetValue()) {
			lastMantained.SetNext(actual)
			lastMantained = actual
		} else {
			list.size--
		}
	}

	lastMantained.SetNext(nil)
	list.tail = lastMantained
}

/*
FilterAsNewList returns a new LinkedList with the filtered values from `list`.

Example:

	var list *LinkedList[int] // is a list with three elements: 1, 2 and 3.
	isOdd(x int) bool // checks if x is a odd number.

	list.FilterAsNewList(isOdd) // this method returns a new list with two elements: 1 and 3. The var list still contains the elements 1, 2 and 3.
*/
func (list *LinkedList[T]) FilterAsNewList(maintain func(T) bool) *LinkedList[T] {
	if list == nil || maintain == nil {
		panic(errorsapp.ErrNilPointer())
	}

	listcp := New[T]()

	for element := range list.Iter() {
		listcp.AddBack(element)
	}

	return listcp
}

/*
Some checks if some element returns true in condidtion().

ATTENTION:

	if list is nil or condition is nil, it will enter in panic(errorsapp.ErrNilPointer())
*/
func (list *LinkedList[T]) Some(condition func(T) bool) bool {
	if list == nil || condition == nil {
		panic(errorsapp.ErrNilPointer())
	}

	for v := range list.Iter() {
		if condition(v) {
			return true
		}
	}

	return false
}

/*
All checks if all elements in list returns true in condition().

ATTENTION:

	if list is nil or condition is nil, it will enter in panic(errorsapp.ErrNilPointer()).
*/
func (list *LinkedList[T]) All(condition func(T) bool) bool {
	if list == nil || condition == nil {
		panic(errorsapp.ErrNilPointer())
	}

	for v := range list.Iter() {
		if !condition(v) {
			return false
		}
	}

	return true
}
