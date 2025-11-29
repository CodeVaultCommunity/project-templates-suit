// Package collecions provides some nodes implementations
package collections

import (
	errorsapp "mod_name/error"
)

type LinkedList[T any] struct {
	size uint
	rear *SimpleLinkedNode[T]
	tail *SimpleLinkedNode[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		size: 0,
		rear: nil,
		tail: nil,
	}
}

/*
	=========================================
	============== POP METHODS ==============
	=========================================
*/

func (list *LinkedList[T]) PopFirst() *T {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	if list.rear == nil {
		return nil
	}

	rear := list.rear
	list.rear = rear.next
	if list.rear == nil {
		list.tail = nil
	}

	list.size--

	return rear.value
}

func (list *LinkedList[T]) PopLast() *T {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	if list.rear == nil {
		return nil
	}

	if list.rear == list.tail {
		v := list.rear.value
		list.rear = nil
		list.tail = nil
		list.size--
		return v
	}

	tail := list.tail
	pre_tail := list.rear
	for ; pre_tail.next != tail; pre_tail = pre_tail.next {
	}

	list.tail = pre_tail
	list.size--
	return tail.value
}

/*
=========================================
============== GET METHODS ==============
=========================================
*/
func (list *LinkedList[T]) getAt(index uint) *SimpleLinkedNode[T] {
	if index > (list.size - 1) {
		panic(errorsapp.ErrIndextOutOfBound())
	}

	if index == 0 {
		return list.rear
	}

	if index == (list.size - 1) {
		return list.tail
	}

	node := list.rear

	for range index {
		node = node.next
	}

	return node
}

func (list *LinkedList[T]) getAtRev(index uint) *SimpleLinkedNode[T] {
	if index > (list.size - 1) {
		panic(errorsapp.ErrIndextOutOfBound())
	}

	if index == 0 {
		return list.tail
	}

	if index == (list.size - 1) {
		return list.rear
	}

	node := list.rear

	real_index := (list.size - index - 1)
	for range real_index {
		node = node.next
	}

	return node
}

func (list *LinkedList[T]) Get(index uint) *T {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return list.getAt(index).value
}

func (list *LinkedList[T]) GetRev(index uint) *T {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return list.getAtRev(index).value
}

func (list *LinkedList[T]) GetSize() uint {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return list.size
}

/*
=========================================
============== ADD METHODS ==============
=========================================
*/
func (list *LinkedList[T]) AddFront(value *T) {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	node := NewSLN(value)

	if list.rear == nil {
		list.rear = node
		list.tail = node
	} else {
		node.next = list.rear
		list.rear = node
	}

	list.size++
}

func (list *LinkedList[T]) AddBack(value *T) {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	node := NewSLN(value)

	if list.rear == nil {
		list.rear = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}

	list.size++
}

func (list *LinkedList[T]) link(node *SimpleLinkedNode[T], value *T) {
	new_node := NewSLN(value)
	new_node.next = node.next
	node.next = new_node
	list.size++
}

func (list *LinkedList[T]) AddAt(index uint, value *T) {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	if index == 0 {
		list.AddFront(value)
	} else {
		list.link(list.getAt(index-1), value)
	}
}

func (list *LinkedList[T]) AddAtRev(index uint, value *T) {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	if index == 0 {
		list.AddBack(value)
	} else {
		list.link(list.getAtRev(index), value)
	}
}

/*
==========================================
============ SAFE ADD METHODS ============
==========================================
*/
func (list *LinkedList[T]) SAddFront(value *T) {
	if value == nil {
		panic(errorsapp.ErrNilPointer())
	}
	list.AddFront(value)
}

func (list *LinkedList[T]) SAddBack(value *T) {
	if value == nil {
		panic(errorsapp.ErrNilPointer())
	}
	list.AddBack(value)
}

func (list *LinkedList[T]) SAddAt(index uint, value *T) {
	if list == nil || value == nil {
		panic(errorsapp.ErrNilPointer())
	}
	list.AddAt(index, value)
}

func (list *LinkedList[T]) SAddAtRev(index uint, value *T) {
	if list == nil || value == nil {
		panic(errorsapp.ErrNilPointer())
	}
	list.AddAtRev(index, value)
}

/*
	===========================================
	============ UTILITIES METHODS ============
	===========================================
*/

func (list *LinkedList[T]) IsEmpty() bool {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	return list.rear == nil
}

func (list *LinkedList[T]) IteratorLinkedList() <-chan *T {
	return list.rear.IteratorSLN()
}

func MapLinkedList[T, R any](list *LinkedList[T], f func(*T) *R) *LinkedList[R] {
	if list == nil || f == nil {
		panic(errorsapp.ErrNilPointer())
	}

	new := NewLinkedList[R]()

	for v := range list.IteratorLinkedList() {
		new.AddBack(f(v))
	}

	return new
}

func (list *LinkedList[T]) cleanRear(maintain func(*T) bool) *SimpleLinkedNode[T] {
	var node *SimpleLinkedNode[T]

	for node = list.rear; node != nil; node = list.rear {
		if maintain(node.value) {
			return node
		}

		list.PopFirst()
	}

	return nil
}

func (list *LinkedList[T]) FilterLinkedList(maintain func(*T) bool) {
	if list == nil || maintain == nil {
		panic(errorsapp.ErrNilPointer())
	}

	last_mantained := list.cleanRear(maintain)
	if last_mantained == nil {
		return // early stop
	}

	for actual := last_mantained.next; actual != nil; actual = actual.next {
		if maintain(actual.value) {
			last_mantained = actual
		} else {
			last_mantained.next = actual.next
		}
	}
}

func (list *LinkedList[T]) SomeInLinkedList(condition func(*T) bool) bool {
	if list == nil || condition == nil {
		panic(errorsapp.ErrNilPointer())
	}

	for v := range list.IteratorLinkedList() {
		if condition(v) {
			return true
		}
	}

	return false
}
