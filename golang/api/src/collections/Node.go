// Package collecions provides some nodes implementations
package collections

import errorsapp "mod_name/error"

type SimpleLinkedNode[T any] struct {
	value *T
	next  *SimpleLinkedNode[T]
}

func NewSLN[T any](value *T) *SimpleLinkedNode[T] {
	return &SimpleLinkedNode[T]{
		value: value,
		next:  nil,
	}
}

func (node *SimpleLinkedNode[T]) GetValue() *T {
	if node == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return node.value
}

func (node *SimpleLinkedNode[T]) SetValue(newValue *T) *T {
	if node == nil {
		panic(errorsapp.ErrNilPointer())
	}

	lastValue := node.value
	node.value = newValue
	return lastValue
}

func (node *SimpleLinkedNode[T]) GetNext() *SimpleLinkedNode[T] {
	if node == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return node.next
}

func (node *SimpleLinkedNode[T]) SetNext(next *SimpleLinkedNode[T]) *SimpleLinkedNode[T] {
	if node == nil {
		panic(errorsapp.ErrNilPointer())
	}

	last := node.next
	node.next = next
	return last
}

func (node *SimpleLinkedNode[T]) IteratorSLN() <-chan *T {
	ch := make(chan *T)

	go func() {
		for n := node; n != nil; n = n.next {
			ch <- n.value
		}
		close(ch)
	}()

	return ch
}

func MapSLN[T, R any](node *SimpleLinkedNode[T], f func(*T) *R) *SimpleLinkedNode[R] {
	if node == nil || f == nil {
		panic(errorsapp.ErrNilPointer())
	}

	first := NewSLN(f(node.value))

	if node.next == nil {
		return first
	}

	actual := first
	for v := range node.next.IteratorSLN() {
		actual.next = NewSLN(f(v))
		actual = actual.next
	}

	return first
}
