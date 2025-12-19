// Package linkedlist provides linkedlist implementation
package linkedlist

import errorsapp "mod_name/error"

/*
PopFirst removes the first element on List.

If the list is nil, it will enter in panic(errorsapp.ErrNilPointer()).

If the list is empty, it will enter in panic(errorsapp.ErrListIsEmpty()).
*/
func (list *LinkedList[T]) PopFirst() T {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	if list.head == nil {
		panic(errorsapp.ErrListIsEmpty())
	}

	rear := list.head
	list.head = rear.GetNext()
	if list.head == nil {
		list.tail = nil
	}

	list.size--

	return rear.GetValue()
}

/*
PopLast removes the last element on List.

If the list is nil, it will enter in panic(errorsapp.ErrNilPointer()).

If the list is empty, it will enter in panic(errorsapp.ErrListIsEmpty()).
*/
func (list *LinkedList[T]) PopLast() T {
	if list == nil {
		panic(errorsapp.ErrNilPointer())
	}

	if list.head == nil {
		panic(errorsapp.ErrListIsEmpty())
	}

	if list.head == list.tail {
		v := list.head.GetValue()
		list.head = nil
		list.tail = nil
		list.size--
		return v
	}

	tail := list.tail
	preTail := list.head
	for preTail.GetNext() != list.tail {
		preTail = preTail.GetNext()
	}

	list.tail = preTail
	preTail.SetNext(nil)
	list.size--
	return tail.GetValue()
}
