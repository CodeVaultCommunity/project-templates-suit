// Package linkednode provides some nodes implementations
package linkednode

import (
	errorsapp "mod_name/error"
)

// SimpleLinkedNode is a struct to colapse any other type.
// If you pass a struct, you need to specify * (example: SimpleLinkedNode[*MyStruct]).
type SimpleLinkedNode[T any] struct {
	value T
	next  *SimpleLinkedNode[T]
}

// New a constructor for SimpleLinkedNode
func New[T any](value T) *SimpleLinkedNode[T] {
	return &SimpleLinkedNode[T]{
		value: value,
		next:  nil,
	}
}

// GetValue get the value allocated in SimpleLinkedNode.
// If node is empty, it will enter in panic(errorsapp.ErrNilPointer()).
func (node *SimpleLinkedNode[T]) GetValue() T {
	if node == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return node.value
}

/*
SetValue Set the newValue in node and returns the last value.

Example:

node.value = 5;

ret := node.SetValue(7);

ret == 5; node.value == 7.
*/
func (node *SimpleLinkedNode[T]) SetValue(newValue T) T {
	if node == nil {
		panic(errorsapp.ErrNilPointer())
	}

	lastValue := node.value
	node.value = newValue
	return lastValue
}

// GetNext get the next node.
// If have no next value it returns nil.
// If node is nil, it will enter in panic(errorsapp.ErrNilPointer()).
func (node *SimpleLinkedNode[T]) GetNext() *SimpleLinkedNode[T] {
	if node == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return node.next
}

/*
SetNext set the node.next as next argument and will returns the last assigned value.
If node is empty, it will enter in panic(errorsapp.ErrNilPointer()).

Example:

	node := New(5)
	next := New(6)

	last := node.SetNext(next)

Result:

	node.value = 5; node.next = next; last = node.next
*/
func (node *SimpleLinkedNode[T]) SetNext(next *SimpleLinkedNode[T]) *SimpleLinkedNode[T] {
	if node == nil {
		panic(errorsapp.ErrNilPointer())
	}

	last := node.next
	node.next = next
	return last
}

/*
Iterator prepare a Simple Linked Node to be iteraded. This method IS NOT SAFE.

Example:

	node := New(5)
	next := New(6)
	node.SetNext(next)

	for element := range node.Iterator() {
		fmt.Print(element)
	}

Result:

	It will print 5, 6.

ATTENTION:

	To use it, the node can't be nil.
	Additional, if it have some cycle it will loop infitily.
	So, this method IS NOT SAFE.
*/
func (node *SimpleLinkedNode[T]) Iterator() <-chan T {
	if node == nil {
		panic(errorsapp.ErrNilPointer())
	}

	ch := make(chan T)

	go func() {
		for n := node; n != nil; n = n.next {
			ch <- n.value
		}
		close(ch)
	}()

	return ch
}

/*
Map iterate over each node aaplling the operation f and returns a new nodes with new link.

Example:

	node := New(5)
	next := New(6)
	node.SetNext(next)

	new_node_chain := Map(node, func(item int)int(return item*2))

Result:

	new_node_chain -> contains two nodes: first_node.value == 10; second_node.value == 12.
*/
func Map[T, R any](node *SimpleLinkedNode[T], f func(T) R) *SimpleLinkedNode[R] {
	if node == nil || f == nil {
		panic(errorsapp.ErrNilPointer())
	}

	first := New(f(node.value))

	if node.next == nil {
		return first
	}

	actual := first
	for v := range node.next.Iterator() {
		actual.next = New(f(v))
		actual = actual.next
	}

	return first
}

func (node *SimpleLinkedNode[T]) swap(other *SimpleLinkedNode[T]) {
	node.value, other.value = other.value, node.value
}

// Swap gets the node.value and assing it to other.value and vice-versa
func (node *SimpleLinkedNode[T]) Swap(other *SimpleLinkedNode[T]) {
	if node == nil || other == nil {
		panic(errorsapp.ErrNilPointer())
	}

	node.swap(other)
}
